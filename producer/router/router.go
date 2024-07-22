package router

import (
    "github.com/gin-gonic/gin"
    "72HW/producer/handler"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/orders", handler.CreateOrder)
    r.GET("/orders/:id", handler.GetOrder)
    r.PUT("/orders/:id", handler.UpdateOrder)
    r.DELETE("/orders/:id", handler.DeleteOrder)
    return r
}
