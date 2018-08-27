package main

import (
	"github.com/streadway/amqp"
	"github.com/tkstorm/golang-rabbitmq-tutor/gomqtool"
	"log"
)

var WorkerConfig = gomqtool.Config

func main() {
	//connect rabbitmq
	conn, err := amqp.Dial(WorkerConfig.AmqpUrl)
	gomqtool.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//conn is abstract of the socket
	//create channel for api
	ch, err := conn.Channel()
	gomqtool.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	//setting Qos value of channel
	//err = ch.Qos(
	//	100,
	//	0,
	//	false,
	//)
	//gomqtool.FailOnError(err, "Failed to set Qos")

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
	gomqtool.FailOnError(err, "Failed to declare a queue")

	//read the message from a channel in a goroutine
	count := 0
	forever := make(chan bool)
	//for i := 0; i < 20; i++ {
	consume(ch, q, count)
	//}
	<-forever
	log.Printf(" [*] Waiting for message. To exit press CTRL+C")
}

func consume(ch *amqp.Channel, q amqp.Queue, count int) {
	count++
	log.Printf("start consume[%d]\n", count)

	//register a consumer for queue by channel
	msgs, err := ch.Consume(
		q.Name,
		"",
		true, //autoAck set false, must through d.Ack(false) acknowledge
		false,
		false,
		false,
		nil,
	)
	gomqtool.FailOnError(err, "Failed to register a consumer")

	for d := range msgs {
		log.Printf("Received a message: %s\n", d.Body)
		//dotCount := bytes.Count(d.Body, []byte(".")) / 2
		//log.Printf("Work need %vs", dotCount)
		//t := time.Duration(dotCount)
		//time.Sleep(t * time.Second)
		log.Printf("Done")

		//manual ack
		//d.Ack(false)
	}
}
