package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gomqtool"
	"log"
)

var Config = gomqtool.Config

func init() {
	log.Println("Viper init...")
	log.Println(Config)

}

func main() {
	//get config info
	amqpUrl := fmt.Sprintf("amqp://%s:%s@%s:%d",
		viper.GetString("rabbitmq.user"),
		viper.GetString("rabbitmq.pass"),
		viper.GetString("rabbitmq.address"),
		viper.GetInt("rabbitmq.port"),
	)

	fmt.Println(amqpUrl)

}
