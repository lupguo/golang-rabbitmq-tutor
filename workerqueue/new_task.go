package main

import (
	"flag"
	"github.com/streadway/amqp"
	"github.com/tkstorm/golang-rabbitmq-tutor/gomqtool"
	"log"
)

var Config = gomqtool.Config

func main() {
	//connect rabbitmq
	conn, err := amqp.Dial(Config.AmqpUrl)
	gomqtool.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//conn is abstract of the socket
	//create channel for api
	ch, err := conn.Channel()
	gomqtool.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	//must declare a queue before send message
	q, err := ch.QueueDeclare(
		"new_task",
		true,
		false,
		false,
		false,
		nil,
	)
	gomqtool.FailOnError(err, "Failed to declare a queue")

	//body receive msg from cmdline
	body := flag.String("body", "Hi.....", "message body")
	flag.Parse()

	//msg
	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(*body),
	}

	//done := make(chan bool)
	//for i := 0; i < 20; i++ {

	for {
		work(ch, q, msg, 1)
	}
	//<-done

	log.Printf("[x] Send Msg Task Ok.")
}

func work(ch *amqp.Channel, q amqp.Queue, msg amqp.Publishing, i int) {
	//for k := 0; k < 20000; k++ {

	err := ch.Publish(
		"",
		q.Name,
		false,
		false,
		msg,
	)
	if err != nil {
		log.Printf("%v: %s\n", err, "Failed to publish a message")
	} else {
		log.Printf("send msg: %s\n", msg.Body)
	}

	//for j := 0; j < 3; j++ {
	//	log.Printf("run[%d][%d] ", i, j)
	//	time.Sleep(1 * time.Second)
	//}
	log.Printf("run[%d] goroutines finished. \n", i)
	//done <- true
	//}
}
