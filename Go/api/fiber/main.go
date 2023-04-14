package config

import (
	"database/sql"
	"fmt"
	_ "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

func LoadConfig()(*sql.DB,error){
	//LoadDefaultConfig()
	LoadFileConfig("./databases/pokemon_postgres/config","config.yml")
	//SetLogging()
	//LoadEnvironmentConfig()
	//LoadFlagConfig()
	//CheckRequiredConfig()
	//ShowStartup()
	viper.Set(CONTEXT, GetDBContext())
	db, err := sql.Open("postgres", viper.GetString(CONTEXT))
	if err != nil {
		return nil, err
	}
	return db, nil
}

//func SetLogging(){
//	loglevel, err := log.ParseLevel(viper.GetString(LOGLEVEL))
//	if err != nil {
//		log.Fatalf("Invalid loglevel. %v", err)
//	}
//	log.SetLevel(loglevel)
//}

func LoadFileConfig(path string, cname string) error {
	if _, err := os.Stat(cname); os.IsNotExist(err) {
	//if _, err := os.Stat("./config.yml"); os.IsNotExist(err) {
		viper.SetConfigName("config")
		viper.AddConfigPath("/etc")
	} else {
		viper.SetConfigName("ConfigLocal")
	}
	viper.SetConfigType("yaml")
	//viper.AddConfigPath("./databases/pokemon_postgres/config")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		return err
		//log.Warnf("No config file found. %v\n", err)

	}
	return nil
}

func GetDBContext() string{
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		viper.GetString(HOST), viper.GetInt(PORT), viper.GetString(USER), viper.GetString(PASSWORD), viper.GetString(DBNAME))
}
