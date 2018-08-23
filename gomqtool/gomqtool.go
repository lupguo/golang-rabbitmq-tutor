package gomqtool

import (
	"fmt"
	"log"
	"flag"
)

////variable
//var ConfigFile *string
//
////const
//const ConstConfigDir  = `I:\Terry.Rod\goLang\rabbitmqGo\src\config.json`
//const ConstVarWebName = "RabbitMq Tour"
//
//type defined
type Config struct {
	Path string
}
var TourConfig Config

//var AppConfig = Config{
//	`I:\Terry.Rod\goLang\rabbitmqGo\src\config.json`,
//}

//helper function check return value for amqp call
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

//get config path from cmdline
func Init() {
	ConfigFile := flag.String("config_file", ".", "set tour app config path")
	flag.Parse()

	TourConfig.Path = *ConfigFile

	log.Printf("[%s]: %s", "config path", *ConfigFile)
}
