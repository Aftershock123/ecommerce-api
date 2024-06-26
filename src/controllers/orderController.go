package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/Aftershock123/ecommerce/src/models"
    "github.com/Aftershock123/ecommerce/src/database"
   
)

func CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    database.DB.Create(&order)
    c.JSON(201, order)
}

func GetOrders(c *gin.Context) {
    var orders []models.Order
    database.DB.Find(&orders)
    c.JSON(200, orders)
}