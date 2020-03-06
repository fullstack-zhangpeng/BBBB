package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	loadConfig()
}

func loadConfig() {
	viper.SetConfigFile("./config/config.json")
	err := viper.ReadInConfig()
	//Handle errors reading the config file
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func writeConfig() {
	viper.SetConfigFile("./config/config.json") //文件名
	err := viper.WriteConfig()                  //写入文件
	if err != nil {                             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
