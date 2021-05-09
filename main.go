package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/tkanos/gonfig"
	"log"
	"net/http"
	"ttnmapper-ingress-api/types"
	"ttnmapper-ingress-api/utils"
)
import "github.com/go-chi/chi"

var publishChannel = make(chan types.TtnMapperUplinkMessage, 100)

type Configuration struct {
	AmqpHost     string `env:"AMQP_HOST"`
	AmqpPort     string `env:"AMQP_PORT"`
	AmqpUser     string `env:"AMQP_USER"`
	AmqpPassword string `env:"AMQP_PASSWORD"`

	HttpListenAddress string `env:"HTTP_LISTEN_ADDRESS"`
}

var myConfiguration = Configuration{
	AmqpHost:     "localhost",
	AmqpPort:     "5672",
	AmqpUser:     "guest",
	AmqpPassword: "guest",

	HttpListenAddress: ":8080",
}

func main() {

	err := gonfig.GetConf("conf.json", &myConfiguration)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("[Configuration]\n%s\n", utils.PrettyPrint(myConfiguration)) // output: [UserA, UserB]

	router := Routes(publishChannel)

	log.Print("[Routes]")
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	// Start the thread that process the queue
	go publishFromChannel()

	// Start the http endpoint
	log.Fatal(http.ListenAndServe(myConfiguration.HttpListenAddress, router))
}

func publishFromChannel() {
	conn, err := amqp.Dial("amqp://" + myConfiguration.AmqpUser + ":" + myConfiguration.AmqpPassword + "@" + myConfiguration.AmqpHost + ":" + myConfiguration.AmqpPort + "/")
	//if err != nil {
	//	log.Print("Error connecting to RabbitMQ")
	//	return
	//}
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
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
		message := <-publishChannel
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
		utils.FailOnError(err, "Failed to publish a message")
	}
}
