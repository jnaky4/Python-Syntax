package dev

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"os"
	"path"
	"strings"
)

const (
	postgresImage         = "postgres:16"
	postgresContainerName = "tPostgres"
	postgresContainerPort = "5432"
	DbHost                = "0.0.0.0" //localhost
	postgresDbName        = "Table"
	postgresPassword      = "password"
	DbSsl                 = "disable"
	postgresUser          = "postgres"
)

//todo how should programmer pass in data to manage templates?

func PostgresBuildTemplate() ContainerBuild {
	postgresBuild := NewContainerBuilder()
	postgresBuild.ContainerName = postgresContainerName
	postgresBuild.Config.Image = postgresImage
	postgresBuild.Config.Env = []string{
		fmt.Sprintf("POSTGRES_PASSWORD=%s", postgresPassword),
		fmt.Sprintf("POSTGRES_USER=%s", postgresUser),
		fmt.Sprintf("POSTGRES_DB=%s", postgresDbName),
	}

	postgresBuild.Config.Hostname = fmt.Sprintf("%s-hostnameexample", postgresBuild.Config.Image)
	postgresBuild.HostConfig.NetworkMode = "postgresNetwork"

	postgresBuild.Config.ExposedPorts["5432/tcp"] = struct{}{} // Define ports to be exposed (has to be same as hostconfig.portbindings.newport)
	postgresBuild.HostConfig.PortBindings = nat.PortMap{
		"5432/tcp": []nat.PortBinding{
			{HostIP: "", HostPort: postgresContainerPort},
		},
	}
	return postgresBuild
}

func RedisBuildTemplate() ContainerBuild {
	redisBuild := NewContainerBuilder()
	redisBuild.ContainerName = "ssRedis"
	redisBuild.Config.Env = []string{}
	redisBuild.Config.Image = "redis:latest"
	redisBuild.Config.ExposedPorts["6379/tcp"] = struct{}{}
	redisBuild.HostConfig.PortBindings = nat.PortMap{
		"6379/tcp": []nat.PortBinding{
			{HostIP: "", HostPort: "6379"},
		},
	}
	return redisBuild
}

func MinioBuildTemplate() ContainerBuild {

	minioBuild := NewContainerBuilder()
	minioBuild.ContainerName = "ssMinio"
	minioBuild.Config.Image = "minio/minio:latest"
	minioBuild.Config.Env = []string{
		fmt.Sprintf("MINIO_ROOT_USER=%s", "minio"),
		fmt.Sprintf("MINIO_ROOT_PASSWORD=%s", "minio123"), //Must be 8 chars or will fail
	}
	minioBuild.Config.Cmd = []string{"server", "/data"}
	minioBuild.Config.ExposedPorts["9000/tcp"] = struct{}{}
	minioBuild.HostConfig.PortBindings = nat.PortMap{
		"9000/tcp": []nat.PortBinding{
			{HostIP: "", HostPort: "9000"},
		},
	}
	minioBuild.Config.Volumes = map[string]struct {
	}{
		"/minio/data": {
			//Driver: "local",
			//Labels: []string{},
			//Mountpoint: "/var/lib/docker/volumes/my-vol/_data",
			//Name: "my-vol",
			//Options: []string{},
			//Scope: "local",
		},
	}
	return minioBuild
}

func StateStoreMinioBuildTemplate() ContainerBuild {

	minioBuild := NewContainerBuilder()

	minioBuild.ContainerName = "ssMinio"

	// Container Config settings
	minioBuild.Config.Image = "minio/minio:latest"
	minioBuild.Config.Env = []string{
		fmt.Sprintf("MINIO_ROOT_USER=%s", "toss_access_key"),
		fmt.Sprintf("MINIO_ROOT_PASSWORD=%s", "toss_test_secret_key"), //Must be 8 chars or will fail
		//fmt.Sprintf("OBJECTSTORE_BUCKET=%s", "abucket"),
	}
	minioBuild.Config.Hostname = "minio"
	minioBuild.Config.Cmd = strings.Split("server /data --console-address :9001 --address :9000", " ")

	//Container Configs for these must have for Host Configs as well
	minioBuild.Config.Volumes = map[string]struct{}{
		"/data":                      {},
		"/etc/nginx/htpasswd":        {},
		"/etc/nginx/ssl/cert.crt":    {},
		"/etc/nginx/ssl/private.key": {},
	}
	minioBuild.Config.ExposedPorts["9000/tcp"] = struct{}{}
	minioBuild.Config.ExposedPorts["9001/tcp"] = struct{}{}

	//Host Config Settings
	minioBuild.HostConfig.NetworkMode = "statestoretoss_shared"
	minioBuild.HostConfig.PortBindings = nat.PortMap{
		"9000/tcp": []nat.PortBinding{
			{HostIP: "", HostPort: "9000"},
		},
		"9001/tcp": []nat.PortBinding{
			{HostIP: "", HostPort: "9001"},
		},
	}

	getwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("error getting cwd -> %s", err.Error())
		return ContainerBuild{}
	}

	minioBuild.HostConfig.Binds = []string{
		path.Join(getwd, "data:/data:rw"),
		path.Join(getwd, "htpasswd:/etc/nginx/htpasswd:rw,z"),
		path.Join(getwd, "cert.crt:/etc/nginx/ssl/cert.crt:rw,z"),
		path.Join(getwd, "private.key:/etc/nginx/ssl/private.key:rw,z"),
	}

	minioBuild.NetworkConfig.EndpointsConfig = map[string]*network.EndpointSettings{
		"statestoretoss_shared": {
			//DriverOpts: map[string]string{
			//	"driver": "bridge",
			//},
			Aliases: []string{"toss_minio"},
		},
	}

	return minioBuild
}

// NetworkBuilder creates a network and returns a network cleanup function
func NetworkBuilder(ctx context.Context, networkName string, network types.NetworkCreate) (func(), error) {
	//network must be built/cleaned up prior to building a container that depends on network

	d, err := BuildNetwork(ctx, networkName, network)
	if err != nil {
		return nil, fmt.Errorf("failed to build network -> %s", err.Error())
	}

	return func() {
		_ = d.GracefulShutdown(ctx, networkName)
	}, nil
}
