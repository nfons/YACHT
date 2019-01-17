package main

import (
	"github.com/nfons/yacht/cmd"
	"github.com/spf13/viper"
)

var env string
var kfile string

func dirStruct(config AppConfigProperties) {
	return
}

func main() {
	// filename := fmt.Sprintf("%s.conf", env)

	// viper logic
	viper.SetConfigName(env)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	// viperErr := viper.ReadInConfig()
	// if viperErr != nil {
	// 	log.Fatal(viperErr)
	// }
	//
	// config := AppConfigProperties{}
	// strConf := viper.AllSettings()
	// for key, val := range strConf {
	// 	strKey := fmt.Sprintf("%v", key)
	// 	strVal := fmt.Sprintf("%v", val)
	// 	config[strKey] = strVal
	// }

	cmd.Execute()
}
