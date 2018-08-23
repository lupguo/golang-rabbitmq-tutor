package main

import (
	"fmt"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	"log"
)

func main() {
	//new config
	//conf := config.NewConfig()

	//load file
	err := config.Load(file.NewSource(
		file.WithPath(`I:\Terry.Rod\goLang\rabbitmqGo\config.json`),
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
	port := config.Get("datastore","metric","host","port").Int(3378)
	fmt.Println(port)
}
