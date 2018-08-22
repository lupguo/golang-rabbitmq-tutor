package main

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
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
	//make sure the queue exist when the consumer working
	//declare a queue is idempotent-
	//it will only be created if it doesn't exist already
	q, err := ch.QueueDeclare(
		"new_task",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	//setting Qos value for lbs
	err = ch.Qos(
		2,
		0,
		false,
	)
	failOnError(err, "Failed to set Qos")

	//register a consumer for queue by channel
	msgs, err := ch.Consume(
		q.Name,
		"",
		false, //autoAck set false, must through d.Ack(false) acknowledge
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	//read the message from a channel in a goroutine
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s\n", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			log.Printf("Work need %vs", dotCount)
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")

			//manual ack
			d.Ack(false)
		}
	}()
	log.Printf(" [*] Waiting for message. To exit press CTRL+C")
	<-forever
}
