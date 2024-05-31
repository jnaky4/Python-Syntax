package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type PgConfig struct {
	Config
}

func (p PgConfig) LoadDBConfig(){
	p.LoadConfig()
	p.GetDBSource()
	p.CheckRequiredConfig()
}

func (p PgConfig) GetDBSource() {
	GetDBSource()
}

func (p PgConfig) CheckRequiredConfig() {
	CheckRequiredConfig()
}

func NewPgConfig(configName string) PgConfig{
	return PgConfig{ Config: Config{FileName: configName}}
}

func LoadDBConfig(configName string){
	LoadDefaultConfig()
	LoadFileConfig(configName)
	ShowStartup(viper.GetBool(DISPLAYCONFIG))
	GetDBSource()
	CheckRequiredConfig()
}

func CheckRequiredConfig() {
	required := []string{
		DB_NAME,
		DB_HOST,
		DB_USER,
		CONTAINER_PORT,
		DB_PASSWORD,
		DB_SSL,
		DB_SOURCE,
	}
	for _, v := range required {
		if viper.GetString(v) == "" {
			log.Fatalf("Required config '%v' is not defined.\n", v)
		}
	}
}

func GetDBSource() {
	viper.Set(DB_SOURCE, fmt.Sprintf(
		"host=%s port=%d " +
			"user=%s password=%s " +
			"dbname=%s sslmode=%s",
		viper.GetString(DB_HOST), viper.GetInt(CONTAINER_PORT),
		viper.GetString(DB_USER), viper.GetString(DB_PASSWORD),
		viper.GetString(DB_NAME), viper.GetString(DB_SSL)))
}
