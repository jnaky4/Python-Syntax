package main

import (
	d "Go/docker"
	"database/sql"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq"
)

func main() {
	image := "postgres"
	user := "postgres"
	pass := "pokemon"
	dab := "pokemon"
	port := "5432"
	porttcp := nat.Port(port + "/tcp")

	cli, ctx, err := d.Connect()
	if err != nil {
		println(err.Error())
		return
	}
	err = d.Pull(cli, ctx, image)
	if err != nil {
		println(err.Error())
	}
	//if d.GetContainer(cli, ctx, "")
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
	cID, err := d.Create(cli, ctx, cfg, hcfg, nil, nil, "")
	if err != nil {
		println(err.Error())
		return
	}
	err = d.Start(cli, ctx, cID)
	if err != nil {
		println(err.Error())
		return
	}

	println(cID)
	connStr := fmt.Sprintf("postgresql://%s:%s@localhost:%s/%s?sslmode=disable", user, pass, port, dab)
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		println(err.Error())
		return
	}

	_, err = db.Exec("CREATE DATABASE Pokemon")
	if err != nil {
		println(err.Error())
		return
	}
	print("Passed")

}
