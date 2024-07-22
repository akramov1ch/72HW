package handler

import (
	"fmt"
	"log"
	"72HW/consumer/database"
	"72HW/consumer/model"
)

func HandleOrder(order model.Order) {
	if err := validateOrder(order); err != nil {
		log.Printf("Order validation failed: %v", err)
		return
	}

	switch order.Status {
	case "order.pending":
		log.Println("Order is pending")
	case "order.completed":
		log.Println("Order is completed")
	case "order.canceled":
		log.Println("Order is canceled")
	case "order.updated":
		log.Println("Order is updated")
	case "order.deleted":
		log.Println("Order is deleted")
	}

	if err := database.SaveOrder(order); err != nil {
		log.Printf("Failed to save order: %v", err)
		return
	}

	log.Printf("Order processed: %v", order)
}

func validateOrder(order model.Order) error {
	if order.ID == "" || order.Status == "" {
		return fmt.Errorf("invalid order data")
	}
	return nil
}
