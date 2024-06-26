package routes

import (
	"github.com/Aftershock123/ecommerce/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()


	r.GET("/get/products", controllers.GetProducts)
	r.POST("/post/products", controllers.CreateProducts)

	r.POST("/post/orders", controllers.CreateOrder)
	r.GET("/get/orders", controllers.GetOrders)

	return r
}
