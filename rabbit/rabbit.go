package rabbit

import (
	"log"
	"github.com/streadway/amqp"
)

func ConnectRabbitMQ() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		log.Fatalf("Couldn't connect to RabbitMQ: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Couldn't open channel: %s", err)
	}

	return ch, nil
}

func PublishMessage(ch *amqp.Channel, message []byte) error {
	queueName := "id_task"
	_, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("The queue could not be declared: %s", err)
	}

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
	return err
}

