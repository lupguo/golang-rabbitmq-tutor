package gomqtool

import (
	"fmt"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"strings"
)

//config
type appConfig struct {
	AppName  string
	Version  string
	Category string
	File     string
	AmqpUrl  string
}

var Config = appConfig{
	AppName:  "rabbitmq tour by golang",
	Version:  "0.0.1",
	Category: "tour",
	File:     "./config.json",
	AmqpUrl:  "amqp://guest:guest@127.0.0.1:5672",
}

type RabbitmqConfig *struct {
	user string
	pass string
	ip   string
	port int
}

//helper function check return value for amqp call
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

//cmdline get config
func init() {
	//start
	log.Printf(`[app:"%v" version:"%s"]`, Config.AppName, Config.Version)
	log.Println("Common init...")

	//get config file from cmd by pflag
	pflag.StringVarP(&Config.File, "config_file", "c", Config.File, "app config path")
	pflag.Parse()
	log.Printf("[%s]: %s", "config path", Config.File)

	//get amqpurl from config file
	viper.SetConfigFile(Config.File)
	err := viper.ReadInConfig()
	if err != nil {
		FailOnError(err, "Viper cannot read in config")
	}
	Config.AmqpUrl = fmt.Sprintf("amqp://%s:%s@%s:%d",
		viper.GetString("rabbitmq.user"),
		viper.GetString("rabbitmq.pass"),
		viper.GetString("rabbitmq.address"),
		viper.GetInt("rabbitmq.port"),
	)
}

//get config by micro/go-config
func goConfig() {
	//new config
	//conf := config.NewConfig()

	fmt.Println(Config.File)

	//load file
	err := config.Load(file.NewSource(
		file.WithPath(Config.File),
	))
	if err != nil {
		log.Fatalf("%s: %s", "Config load error", err)
	}

	//read config
	conf := config.Map()
	fmt.Println(conf)

	//Scan the config into a struct
	type Host struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	}

	//method 1
	type HostConfig struct {
		Hosts map[string]Host `json:"hosts"`
	}
	var hconf HostConfig
	config.Scan(&hconf)
	fmt.Println(hconf.Hosts["rabbitmq"].Address, hconf.Hosts["rabbitmq"].Port)

	//method 2
	var host Host
	config.Get("hosts", "mysql").Scan(&host)
	fmt.Println(host.Address, host.Port)

	//method 3
	cacheAddress := config.Get("hosts", "cache", "address").String("localhost")
	cachePort := config.Get("hosts", "cache", "port").Int(3000)
	fmt.Println(cacheAddress, cachePort)

	//method 4
	port := config.Get("datastore", "metric", "port").Int(1000)
	fmt.Println(port)

	//method 5
	key := "datastore.metric.port"
	port = config.Get(strings.Split(key, ".")...).Int(3306)
	fmt.Println(port)
}
