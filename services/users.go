package services

import (
	"github.com/gin-gonic/gin"
	"github.com/muharamdani/gin-user-services/config"
	"github.com/muharamdani/gin-user-services/models"
	"strconv"
)

func GetUser(c *gin.Context) (interface{}, error) {
	var users []models.Users
	perPageStr := c.DefaultQuery("per_page", "10")

	// convert perPage to int
	perPageInt, err := strconv.Atoi(perPageStr)
	if err != nil {
		panic(err)
		return nil, err
	}
	config.DB.Limit(perPageInt).Find(&users)
	return users, nil
}

func CreateUser(c *gin.Context) (interface{}, error) {
	var user models.Users
	err := c.ShouldBindJSON(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//func UpdateUser(c *gin.Context) {
//	dt := make(map[string]string)
//	user := c.Params.ByName("name")
//	value, ok := dt[user]
//	if ok {
//		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
//	} else {
//		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
//	}
//}
//
//func DeleteUser(c *gin.Context) {
//	dt := make(map[string]string)
//	user := c.Params.ByName("name")
//	value, ok := dt[user]
//	if ok {
//		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
//	} else {
//		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
//	}
//}
