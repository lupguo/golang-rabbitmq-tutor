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
	//make sure the queue exist when the consumer working
	//declare a queue is idempotent-
	//it will only be created if it doesn't exist already
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	gomqtool.FailOnError(err, "Failed to declare a queue")

	//register a consumer for queue by channel
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	gomqtool.FailOnError(err, "Failed to register a consumer")

	//read the message from a channel in a goroutine
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for message. To exit press CTRL+C")
	<-forever
}
