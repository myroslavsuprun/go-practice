package event

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	con *amqp.Connection
}

func (e *Emitter) setup() error {
	ch, err := e.con.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	return declareExchange(ch)
}

func (e *Emitter) Push(event, severity string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch, err := e.con.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	log.Println("Pushing event to RabbitMQ channel")
	return ch.PublishWithContext(
		ctx,
		"logs_topic",
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)
}

func NewEventEmitter(con *amqp.Connection) (emitter Emitter, err error) {
	emitter = Emitter{
		con: con,
	}
	err = emitter.setup()

	return
}
