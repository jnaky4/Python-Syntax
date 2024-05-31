package main

import (
	"context"
	"database/sql"
	dd "dockerDev/dev"
	"fmt"
	"github.com/docker/docker/api/types"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// todo remove for user friendly input to template
const (
	ContainerImageVersion = "postgres:16"
	ContainerName         = "tPostgres"
	ContainerPort         = "5432"
	DbHost                = "0.0.0.0" //localhost
	DbName                = "Table"
	DbPassword            = "password"
	DbSsl                 = "disable"
	DbUser                = "postgres"
)

// Example building of a container
func main() {
	//example of building Postgres DB
	ctx := context.Background()

	//Network example
	//need to uncomment postgresBuild.HostConfig.NetworkMode in dd.PostgresBuildTemplate()

	gracefulShutdown, err := dd.NetworkBuilder(ctx, "postgresNetwork", types.NetworkCreate{})
	if err != nil {
		log.Fatalf("network build failed -> %s", err.Error())
	}

	defer gracefulShutdown()
	go func() {
		InterruptSignal(gracefulShutdown) //most dev env are closed by interrupt signal and do not trigger defer
	}()

	//build container example
	postgresBuild := dd.PostgresBuildTemplate()
	_, err = dd.BuildContainer(ctx, postgresBuild)
	if err != nil {
		log.Fatalf("Build failed -> %s", err.Error())
	}

	db, err := ConnectDB(GetDBSource())
	if err != nil {
		log.Fatalf("Connection Failed -> %s", err.Error())
	}
	fmt.Printf("Connection Up! %+v\n", db.Stats())

}

func ConnectDB(dbSource string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		return db, err
	}

	for i := 0; i < 3; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(time.Duration(i) * time.Second)
	}

	if err != nil && err.Error() != "EOF" {
		return db, fmt.Errorf("failed to connect to db error -> %s", err.Error())
	}

	return db, nil

}

func GetDBSource() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		DbHost, ContainerPort,
		DbUser, DbPassword,
		DbName, DbSsl,
	)
}

func InterruptSignal(gracefulShutdown func()) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Printf("interrupt signal received, starting graceful shutdown\n")
		gracefulShutdown()
		os.Exit(0)
	}()
}

//todo request example
//func (cli *Client) sendRequest(ctx context.Context, method, path string, query url.Values, body io.Reader, headers headers) (serverResponse, error) {
//	req, err := cli.buildRequest(method, cli.getAPIPath(ctx, path, query), body, headers)
//	if err != nil {
//		return serverResponse{}, err
//	}
//
//	resp, err := cli.doRequest(ctx, req)
//	switch {
//	case errors.Is(err, context.Canceled):
//		return serverResponse{}, errdefs.Cancelled(err)
//	case errors.Is(err, context.DeadlineExceeded):
//		return serverResponse{}, errdefs.Deadline(err)
//	case err == nil:
//		err = cli.checkResponseErr(resp)
//	}
//	return resp, errdefs.FromStatusCode(err, resp.statusCode)
//}
