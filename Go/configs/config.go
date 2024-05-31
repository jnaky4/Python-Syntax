package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"strings"
)

type Config struct{
	FileName string
}

func (c Config) LoadConfig() {
	//c.LoadDefaultConfig()
	c.LoadFileConfig()
	c.ShowStartup()
}

func (c Config) LoadFileConfig() {
	cwd, _ := os.Getwd()
	fPath := path.Join(cwd, "configs", c.FileName+".yml")

	//If local config File doesnt exist, assume Prod env
	if _, err := os.Stat(fPath); os.IsNotExist(err) {
		viper.SetConfigName("prod")
		viper.AddConfigPath("./configs")
	} else {
		viper.SetConfigName(c.FileName)
		viper.AddConfigPath("./configs")
	}

	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("No config file found at %s or ConfigName incorrect: %s\n", fPath, c.FileName)
		println(err.Error())
	}
}

func (c Config) LoadDefaultConfig() {
	viper.SetDefault(LOGLEVEL, "error")
	viper.SetDefault(IMAGE_VERSION, "latest")
	viper.SetDefault(GVERSION, "latest")
	viper.SetDefault(DISPLAYCONFIG, false)
	viper.SetDefault(GVERSION, "latest")
}

func (c Config) ShowStartup() {
	if viper.GetBool(DISPLAYCONFIG){
		fmt.Printf("Current Configuration:\n")
		b, _ := yaml.Marshal(viper.AllSettings())
		fmt.Printf(string(b))
	}
}

func LoadConfig(configName string) {
	LoadDefaultConfig()
	LoadFileConfig(configName)

	if configName == "pokemon_postgres"{
		GetDBSource()
		CheckRequiredConfig()
	}

	//LoadEnvironmentConfig()
	//LoadFlagConfig()

	//ShowStartup(viper.GetBool(DISPLAYCONFIG))
}

func LoadFileConfig(configName string) {
	cwd, _ := os.Getwd()
	fPath := path.Join(cwd, "configs", configName+".yml")

	//If local config File doesnt exist, assume Prod env
	if _, err := os.Stat(fPath); os.IsNotExist(err) {
		viper.SetConfigName("prod")
		viper.AddConfigPath("./configs")
	} else {
		viper.SetConfigName(configName)
		viper.AddConfigPath("./configs")
	}

	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("No config file found at %s or ConfigName incorrect: %s\n", fPath, configName)
		println(err.Error())
	}
}

func LoadDefaultConfig() {
	viper.SetDefault(LOGLEVEL, "error")
	viper.SetDefault(IMAGE_VERSION, "latest")
	viper.SetDefault(DISPLAYCONFIG, true)
}

func ShowStartup(display bool) {
	if display{
		fmt.Printf("Current Configuration:\n")
		b, _ := yaml.Marshal(viper.AllSettings())
		fmt.Printf(string(b))
	}
}

//what does this do?
func LoadEnvironmentConfig() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}