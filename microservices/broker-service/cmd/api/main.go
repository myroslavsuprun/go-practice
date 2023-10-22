package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const webPort = "80"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {
	c, err := connect()
	if err != nil {
		log.Panic(err)
	}
	defer c.Close()

	app := Config{
		Rabbit: c,
	}

	log.Printf("Starting broker service on port %s\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err != nil {
			log.Println("Failed to connect to RabbitMQ")
			counts++
		} else {
			log.Println("Connected to RabbitMQ")
			connection = c
			break
		}

		if counts > 5 {
			return nil, fmt.Errorf("Failed to connect to RabbitMQ")
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		fmt.Println("Waiting ", backOff, " before trying again")
		time.Sleep(backOff)
	}

	return connection, nil
}
