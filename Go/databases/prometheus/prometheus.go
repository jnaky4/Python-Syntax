package prometheus

import (
	"Go/docker"
	"context"
)

func BuildPrometheus(ctx context.Context) {
	docker.BuildContainer(ctx, NodeExporterContainerBuild())
	docker.BuildContainer(ctx, PrometheusContainerBuild())
}

func PrometheusContainerBuild() docker.ContainerBuild{
	return docker.ContainerBuild{
		ImgName:       "prom/prometheus",
		Version:       "latest",
		ContainerName: "prometheus",
		Port:          "9090",
		EnvVars: []string{
		},
		Volumes: []string{
			"/Users/Z004X7X/Documents/github/syntax/Go/configs/prometheus.yml:/etc/prometheus/prometheus.yml",
			"/Users/Z004X7X/Documents/prom_persist:/prometheus", //persistent storage location
		},
	}

	// todo restart: unless-stopped
	// todo add network bridges to containerBuilder

}

func NodeExporterContainerBuild() docker.ContainerBuild{
	return docker.ContainerBuild{
		ImgName:       "prom/node-exporter",
		Version:       "latest",
		ContainerName: "node-exporter",
		Port:          "9100",
		EnvVars: []string{
			//fmt.Sprintf("POSTGRES_PASSWORD=%s", viper.GetString(cfg.DBPASSWORD)),
		},
	}
}