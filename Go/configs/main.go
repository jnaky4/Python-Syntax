package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"strings"
)
const (
	LOGLEVEL = "loglevel"
	TEST_USERNAME = "test.username"
	SHOWCONFIG = "showconfig"
)

func main(){
	LoadConfig()
	fmt.Printf("%s\n", viper.GetString(LOGLEVEL))
}

func LoadConfig(){
	LoadDefaultConfig()
	LoadFileConfig()
	LoadEnvironmentConfig()
	//LoadFlagConfig()
	ShowStartup()
}

func LoadFileConfig(){
	cwd, _ := os.Getwd()
	path := path.Join( cwd, "configs", "config.yml")
	fmt.Printf("%v\n", path)
	if _, err := os.Stat(path); os.IsNotExist(err){
		fmt.Printf("UHOH\n")
		//viper.SetConfigName("prod")
		//viper.AddConfigPath("/etc")
	} else{
		fmt.Printf("it exists\n")
		//viper.SetConfigName("dev")
	}
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil{
		fmt.Printf("No config file found\n")
	}
}

func LoadEnvironmentConfig(){
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func LoadDefaultConfig(){
	viper.SetDefault(LOGLEVEL, "info")
	viper.SetDefault(TEST_USERNAME, "usr")
}

//func LoadFlagConfig(){
//	pflag.Bool(SHOWCONFIG, false, "Display expected run config and exit")
//}

func ShowStartup(){
	log.Infof("Configuration Loaded")
	log.Infof("LogLevel: %v", viper.GetString(LOGLEVEL))
}
