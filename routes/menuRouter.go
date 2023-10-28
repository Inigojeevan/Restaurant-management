package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(IncomingRoutes *gin.Engine) {
	IncomingRoutes.GET("/menus", controller.GetMenus())
	IncomingRoutes.GET("/menus/menu_id", controller.GetMenu())
	IncomingRoutes.POST("/menus", controller.CreateMenu())
	IncomingRoutes.PATCH("/menus/menu_id", controller.UpdateMenu())
}
