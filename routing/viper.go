package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gomqtool"
)

type rabbitmqConfig struct {
	user    string
	pass    string
	address string
	port    int
}

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
	fmt.Printf("%T \n%v", viper.Get("rabbitmq"), viper.Get("rabbitmq"))
	fmt.Println(fmt.Sprintf("amqp://%v:%v@%v:%v",
		viper.Get("rabbitmq.user"),
		viper.Get("rabbitmq.pass"),
		viper.Get("rabbitmq.address"),
		viper.Get("rabbitmq.port"),
	))
}
