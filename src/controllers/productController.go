package controllers

import (
    "github.com/gin-gonic/gin"
	"github.com/Aftershock123/ecommerce/src/models"
    "github.com/Aftershock123/ecommerce/src/database"

)

func CreateProducts(c *gin.Context) {
    var products []models.Product
    if err := c.ShouldBindJSON(&products); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    database.DB.Create(&products)
    c.JSON(201, products)
}
func GetProducts(c *gin.Context) {
    var products []models.Product
    database.DB.Find(&products)
    c.JSON(200, products)
}
