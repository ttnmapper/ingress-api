package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"os"
	"ttnmapper-ingress-api/types"
)
import "github.com/go-chi/chi"

var publish_channel = make(chan types.TtnMapperUplinkMessage, 100)

type Configuration struct {
	AmqpHost     string
	AmqpPort     string
	AmqpUser     string
	AmqpPassword string
}

var myConfiguration = Configuration{
	AmqpHost:     "localhost",
	AmqpPort:     "5672",
	AmqpUser:     "user",
	AmqpPassword: "password",
}

func main() {

	file, err := os.Open("conf.json")
	if err != nil {
		log.Print(err.Error())
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&myConfiguration)
	if err != nil {
		log.Println("json error:", err)
	}
	log.Printf("[Configuration]\n%+v\n", myConfiguration) // output: [UserA, UserB]

	router := Routes()

	log.Print("[Routes]")
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	// Start the thread that process the queue
	go publish_from_channel()

	// Start the http endpoint
	log.Fatal(http.ListenAndServe(":8080", router))
}

func publish_from_channel() {
	conn, err := amqp.Dial("amqp://" + myConfiguration.AmqpUser + ":" + myConfiguration.AmqpPassword + "@" + myConfiguration.AmqpHost + ":" + myConfiguration.AmqpPort + "/")
	//if err != nil {
	//	log.Print("Error connecting to RabbitMQ")
	//	return
	//}
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"new_packets", // name
		"fanout",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)

	//var message map[string]interface{}
	for {
		message := <-publish_channel
		log.Printf("Publishing message")

		data, err := json.Marshal(message)
		if err != nil {
			fmt.Printf("marshal failed: %s", err)
			continue
		}

		err = ch.Publish(
			"new_packets", // exchange
			"",            // routing key
			false,         // mandatory
			false,         // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(data),
			})
		failOnError(err, "Failed to publish a message")
	}
}
