package main

import (
	"fmt"
	"github.com/tkstorm/golang-rabbitmq-tutor/gomqtool"
	"log"
)

var Config = gomqtool.Config

func init() {
	log.Println("Viper init...")
	log.Println(Config)
}

func main() {
	//get config info
	fmt.Println(Config.AmqpUrl)

}
