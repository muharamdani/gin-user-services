package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muharamdani/gin-user-services/routers"
	"github.com/muharamdani/gin-user-services/utils"
)

func main() {
	mode := utils.Env("SET_MODE")
	port := utils.Env("PORT")
	port = ":" + port

	// Set mode to debug|test|release based on the value of the environment variable
	gin.SetMode(mode)

	// Call one line for all router here
	app := gin.Default()
	routers.Main(app)

	// Listen and Server in 0.0.0.0:8080
	err := app.Run(port)
	if err != nil {
		return
	}
}
