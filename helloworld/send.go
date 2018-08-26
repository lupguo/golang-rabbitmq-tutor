package main

import (
	"github.com/streadway/amqp"
	"github.com/tkstorm/golang-rabbitmq-tutor/gomqtool"
	"log"
)

var config = gomqtool.Config

func main() {
	//connect rabbitmq
	conn, err := amqp.Dial(config.AmqpUrl)
	gomqtool.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//conn is abstract of the socket
	//create channel for api
	ch, err := conn.Channel()
	gomqtool.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	//must declare a queue before send message
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	gomqtool.FailOnError(err, "Failed to declare a queue")

	//define a amqp msg
	body := "HeyMan, Cool!!"
	pubMsg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	}

	//use channel to publish a message to queue
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		pubMsg,
	)
	gomqtool.FailOnError(err, "Failed to publish a message")

	log.Println("Send message finish")
}
