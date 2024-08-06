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
	//defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Couldn't open channel: %s", err)
	}
	//defer ch.Close()
	return ch, nil
}

func PublishMessage(ch *amqp.Channel, message []byte) error {
	queueName := "id_task"
	_, err := ch.QueueDeclare(
		queueName, // имя очереди
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("The queue could not be declared: %s", err)
	}

	err = ch.Publish(
		"",        // exchange
		queueName, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
	return err
}

