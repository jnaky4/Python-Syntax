package main

import (
	cfg "Go/configs"
	"Go/databases/prometheus"
	"Go/metrics"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strings"
)

func main(){
	ctx := context.Background()

	//c := cfg.NewPgConfig("pokemon_postgres")
	//c.LoadDBConfig()
	//
	//pg.BuildDB(ctx)


	//d := docker.Manager{}
	//err := d.Connect(ctx)
	//if err != nil {
	//	fmt.Printf("fatal Connect() error -> %v", err)
	//}
	//defer d.Client.Close()

	//containerName := "pokemon-postgres"
	//
	//// Command to execute inside the Docker container
	//cmd := []string{"ls", "-l"}
	//
	//d.Exec(ctx, containerName, cmd)


	c := cfg.NewGrafanaConfig("grafana")
	c.LoadGrafanaConfig()

	metrics.BuildGrafana(ctx)
	prometheus.BuildPrometheus(ctx)


}

func getContainerID(ctx context.Context, cli *client.Client, containerName string) (string, error) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return "", err
	}

	for _, container := range containers {
		for _, name := range container.Names {
			if strings.Contains(name, containerName) {
				return container.ID, nil
			}
		}
	}

	return "", fmt.Errorf("Container not found: %s", containerName)
}