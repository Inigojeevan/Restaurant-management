package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(IncomingRoutes *gin.Engine) {
	IncomingRoutes.GET("/orders", controller.GetOrders())
	IncomingRoutes.GET("/orders/order_id", controller.GetOrder())
	IncomingRoutes.POST("/orders", controller.CreateOrder())
	IncomingRoutes.PATCH("/orders/order_id", controller.UpdateOrder())
}
