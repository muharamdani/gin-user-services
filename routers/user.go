package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muharamdani/gin-user-services/controllers"
)

func User(router *gin.Engine) {
	router.GET("/users", controllers.GetUser)
	router.POST("/users", controllers.CreateUser)
	//router.PATCH("/user/:name", controllers.UpdateUser)
	//router.DELETE("/user/:name", controllers.DeleteUser)
}
