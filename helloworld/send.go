package main

import (
	"log"
	"fmt"
	"github.com/streadway/amqp"
)

//helper function check return value for amqp call
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
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

	//must declare a queue before send message
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	//define a amqp msg
	body := "HeyMan, Cool!!"
	pubMsg := amqp.Publishing{
		ContentType: "text/plain",
		Body: []byte(body),
	}

	//use channel to publish a message to queue
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		pubMsg,
	)
	failOnError(err, "Failed to publish a message")

	log.Println("Send message finish")
}
