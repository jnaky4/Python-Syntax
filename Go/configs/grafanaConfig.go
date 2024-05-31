package config

type GrafanaConfig struct {
	Config
}

func NewGrafanaConfig(configName string) GrafanaConfig{
	return GrafanaConfig{ Config: Config{FileName: configName}}
}

func (g GrafanaConfig) LoadGrafanaConfig(){
	g.LoadConfig()
	g.LoadDefaultConfig()
}
