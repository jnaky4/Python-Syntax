package metrics

import (
	"Go/docker"
	"context"
)

func BuildGrafana(ctx context.Context)  {
	build := GrafanaContainerBuild()
	docker.BuildContainer(ctx, build)
}

func GrafanaContainerBuild() docker.ContainerBuild {
	return docker.ContainerBuild{
		ImgName:       "grafana/grafana",
		Version:       "latest",
		ContainerName: "pokemon_grafana",
		Port:          "3000",
		EnvVars: []string{
			//fmt.Sprintf("POSTGRES_PASSWORD=%s", viper.GetString(cfg.DBPASSWORD)),
		},
	}
}



