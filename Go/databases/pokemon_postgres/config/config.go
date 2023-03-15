package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	_ "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

func LoadConfig(){
	//LoadDefaultConfig()
	LoadFileConfig()
	SetLogging()
	//LoadEnvironmentConfig()
	//LoadFlagConfig()
	//CheckRequiredConfig()
	//ShowStartup()
	GetDBContext()
}

func SetLogging(){
	loglevel, err := log.ParseLevel(viper.GetString(LOGLEVEL))
	if err != nil {
		log.Fatalf("Invalid loglevel. %v", err)
	}
	log.SetLevel(loglevel)
}

func LoadFileConfig(){
	if _, err := os.Stat("./config.yml"); os.IsNotExist(err) {
		viper.SetConfigName("config")
		viper.AddConfigPath("/etc")
	} else {
		viper.SetConfigName("ConfigLocal")
	}
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./databases/pokemon_postgres/config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Warnf("No config file found. %v\n", err)
	}

}

func GetDBContext() {
	viper.Set(CONTEXT, fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		viper.GetString(HOST), viper.GetInt(PORT), viper.GetString(USER), viper.GetString(PASSWORD), viper.GetString(DBNAME)))
}