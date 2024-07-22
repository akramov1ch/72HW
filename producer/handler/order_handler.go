package handler

import (
    "net/http"
    "72HW/producer/model"
    "72HW/producer/rabbitmq"

    "github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
    var order model.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := rabbitmq.PublishOrder(order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, order)
}

func GetOrder(c *gin.Context) {
	orderID := c.Param("id")
	order := model.Order{ID: orderID, Status: "order.retrieved"}
	if err := rabbitmq.PublishOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func UpdateOrder(c *gin.Context) {
    var order model.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    order.Status = "order.updated"
    if err := rabbitmq.PublishOrder(order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
    orderID := c.Param("id")
    order := model.Order{ID: orderID, Status: "order.deleted"}

    if err := rabbitmq.PublishOrder(order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}
