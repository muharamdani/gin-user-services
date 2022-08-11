package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muharamdani/gin-user-services/controllers"
)

func User(router *gin.Engine) {
	router.GET("/user/:name", controllers.GetUser)
	router.POST("/user/:name", controllers.CreateUser)
	router.PATCH("/user/:name", controllers.UpdateUser)
	router.DELETE("/user/:name", controllers.DeleteUser)
}
