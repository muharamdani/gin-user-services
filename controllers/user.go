package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/muharamdani/gin-user-services/services"
	"github.com/muharamdani/gin-user-services/utils"
)

func GetUser(c *gin.Context) {
	data, err := services.GetUser(c)
	if err != nil {
		utils.ResponseInternalServerError(c, err.Error())
		return
	}
	message := "Get user data success"
	utils.ResponseOK(c, message, data)
}

func CreateUser(c *gin.Context) {
	data, _ := services.CreateUser(c)
	message := "Create user data success"
	utils.ResponseOK(c, message, data)
	return
}

//
//func UpdateUser(c *gin.Context) {
//	data, err := services.UpdateUser(c)
//	if err != nil {
//		utils.ResponseInternalServerError(c, err.Error())
//		return
//	}
//	message := "Update user data success"
//	utils.ResponseOK(c, message, data)
//}
//
//func DeleteUser(c *gin.Context) {
//	data, err := services.DeleteUser(c)
//	if err != nil {
//		utils.ResponseInternalServerError(c, err.Error())
//		return
//	}
//	message := "Delete user data success"
//	utils.ResponseOK(c, message, data)
//}
