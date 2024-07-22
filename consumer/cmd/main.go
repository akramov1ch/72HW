package main

import (
    "log"
    "72HW/consumer/rabbitmq"
)

func main() {
    if err := rabbitmq.ConsumeOrders(); err != nil {
        log.Fatalf("could not consume orders: %v", err)
    }
}
