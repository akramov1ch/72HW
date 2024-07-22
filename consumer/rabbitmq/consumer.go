package rabbitmq

import (
	"72HW/consumer/handler"
	"72HW/consumer/model"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

func ConsumeOrders() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"orders_exchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	statuses := []string{"order.pending", "order.completed", "order.canceled", "order.updated", "order.deleted"}

	for _, status := range statuses {
		go consume(ch, status)
	}

	select {}
}

func consume(ch *amqp.Channel, status string) {
	q, err := ch.QueueDeclare(
		status,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	err = ch.QueueBind(
		q.Name,
		status,
		"orders_exchange",
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	for d := range msgs {
		var order model.Order
		if err := json.Unmarshal(d.Body, &order); err != nil {
			log.Printf("Failed to unmarshal order: %v", err)
			continue
		}

		handler.HandleOrder(order)
	}
}
