package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	// fmt.Println(data)
	c.JSON(200, gin.H{
		"msg": data,
	})
}
