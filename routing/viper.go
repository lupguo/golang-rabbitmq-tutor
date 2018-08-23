package main

import (
	"gomqtool"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	//init
	gomqtool.Init()

	//viper config manage
	viper.SetConfigFile(gomqtool.TourConfig.Path)
	err := viper.ReadInConfig()
	if err != nil {
		gomqtool.FailOnError(err, "Viper cannot read in config")
	}

	//get config info
	fmt.Println(fmt.Sprintf("amqp://%s:%s@%s:%s", viper.Get("rabbitmq")))
}