package rabbitmq

import (
    "encoding/json"
    "log"
    "72HW/producer/model"

    "github.com/streadway/amqp"
)

func PublishOrder(order model.Order) error {
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

    body, err := json.Marshal(order)
    if err != nil {
        return err
    }

    err = ch.Publish(
        "orders_exchange",
        order.Status,
        false,
        false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        })
    if err != nil {
        return err
    }

    log.Printf(" [x] Sent %s", body)
    return nil
}
