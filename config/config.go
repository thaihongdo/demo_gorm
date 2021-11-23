package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

var cf *Configuration

type Configuration struct {
	ServerPort   int
	DbConnection string
}

func GetConfig() *Configuration {
	return cf
}

func InitFromFile(filePathStr string, basePath string) {
	viper.SetConfigFile(filePathStr)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Config file not found: %v", err)
	} else {

		cf = &Configuration{
			ServerPort:   viper.GetInt(".server_port"),
			DbConnection: viper.GetString(".db_connection"),
		}
		log.Println(viper.ConfigFileUsed())
		log.Printf("COnfig %+v", *cf)
	}
}
