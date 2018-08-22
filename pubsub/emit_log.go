package main

import (
	"log"
	"fmt"
	"github.com/streadway/amqp"
	"flag"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	//connect rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@10.40.2.183:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//conn is abstract of the socket
	//create channel for api
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//declareExchange
	err = ch.ExchangeDeclare(
		"logs",
		"fanout",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare an exchange")

	//nameless queue create by customer

	//body receive msg from cmdline
	body := flag.String("body", "Hi..", "message body")
	flag.Parse()
	bodyMsg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(*body),
	}

	//publish message by routekey
	err = ch.Publish(
		"logs",
		"",
		false,
		false,
		bodyMsg,
	)
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Send %s", *body)
}
