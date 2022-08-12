package routers

import (
	"github.com/gin-gonic/gin"
)

func Main(router *gin.Engine) {
	// Call router here
	User(router)
}
