package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var cf *Configuration

type Configuration struct {
	EnvironmentPrefix string `mapstructure:"environment_prefix"`
	ServerPort        int    `mapstructure:"server_port"`
	DbConnection      string `mapstructure:"db_connection"`
}

func GetConfig() *Configuration {
	return cf
}

//InitFromFile init config file
func InitFromFile(path string) *Configuration {

	if path == "" {
		viper.AddConfigPath("config")
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
	} else {
		viper.SetConfigFile(path)
	}

	basePath, _ := os.Getwd()

	viper.AutomaticEnv()
	viper.SetEnvPrefix("")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		log.Printf("Config file not found: %v", err)
		panic(err)
	}
	viper.Set("base_path", basePath)
	if err := viper.Unmarshal(&cf); err != nil {
		log.Printf("covert to struct: %v", err)
		log.Fatal(err)
	}
	if path == "" {
		//fmt.sre("File used  %s %+v", viper.ConfigFileUsed(), cf)
		fmt.Printf("File config used  %s\n \n", viper.ConfigFileUsed())
		dataPrinf, _ := json.Marshal(cf)
		fmt.Printf("Config:  %s\n \n", string(dataPrinf))
	}
	return cf

}
