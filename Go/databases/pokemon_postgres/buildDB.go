package pokemon_postgres

import (
	d "Go/docker"
	"database/sql"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

func BuildPostgresContainer() {
	image := "postgres"
	user := "postgres"
	pass := "password"
	dab := "pokemon"
	port := "5432"
	porttcp := nat.Port(port + "/tcp")
	cName := "pokemon-postgres"
	cID := ""

	//connect to docker cli
	cli, ctx, err := d.Connect()
	if err != nil {
		println(err.Error())
		return
	}

	// try to get the container if exists already before creating
	getContainer, err := d.GetContainer(cli, ctx, cName)
	if err != nil {
		println(err.Error())
		return
	}

	if getContainer.ID == "" {
		//pull the image if it doesn't exists
		err = d.Pull(cli, ctx, image)
		if err != nil {
			println(err.Error())
		}

		cfg := &container.Config{
			ExposedPorts: nat.PortSet{
				porttcp: struct{}{},
			},
			Image: image,
			Env:   []string{"POSTGRES_PASSWORD=" + pass, "POSTGRES_USER=" + user, "POSTGRES_DB=" + dab},
		}
		hcfg := &container.HostConfig{
			PortBindings: nat.PortMap{
				porttcp: []nat.PortBinding{
					{
						HostIP:   "0.0.0.0",
						HostPort: port,
					},
				},
			},
		}
		// create the new container
		cID, err = d.Create(cli, ctx, cfg, hcfg, nil, nil, cName)
		if err != nil {
			println(err.Error())
			return
		}
	} else {
		cID = getContainer.ID
	}

	// start the container
	err = d.Start(cli, ctx, cID)
	if err != nil {
		println(err.Error())

	}

	//interrogate the container with pg_isready until container is ready before querying
	result := ""
	for{
		if strings.Contains(result, "accepting connections") {
			break
		}
		exec, err := d.Exec(ctx, cli, cID, []string{"pg_isready"})
		if err != nil {
			println(exec.Stderr())
			println(exec.ExitCode)
		}
		result = exec.Stdout()
		time.Sleep(time.Second)
		println(result)
	}


	connStr := fmt.Sprintf("host=localhost port=%s password=%s user=%s "+
		" dbname=%s sslmode=disable", port, pass, user, dab)
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		println(err.Error())
		return
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dab))
	if err != nil && err.Error() != "EOF" {
		println(err.Error())
		return
	}

	_, err = db.Exec("CREATE TABLE POKEMON")

	print("Passed")


}
