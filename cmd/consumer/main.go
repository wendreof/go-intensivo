package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wendreof/gointensivo/pkg/rabbitmq"
)

//Read the msgs from a RabbitMQ queue and print it!
func main() {
	ch, err := rabbitmq.OpenChannel()

	if err != nil{
		panic(err)
	}

	defer ch.Close()

	out := make(chan amqp.Delivery) //channel
	go rabbitmq.Consume(ch, out) 		//Thread 2

	for msg := range out{
		println(string(msg.Body)) 		//Thread 1
	}
}