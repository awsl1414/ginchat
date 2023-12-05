package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Sussess 200 {string} json{"code", "message"}
// @Router /user/get_user_list [get]
func GetUserList(c *gin.Context) {

	data := models.GetUserList()
	// fmt.Println(data)
	c.JSON(200, gin.H{
		"msg": data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @Param name query string false "用户名"
// @Param password query string false "密码"
// @Param repassword query string false "确认密码"
// @Sussess 200 {string} json{"code", "message"}
// @Router /user/create_user [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(-1, gin.H{
			"msg": "密码不一致",
		})
	}
	user.Password = password
	models.CreateUser(user)
	// fmt.Println(data)
	c.JSON(200, gin.H{
		"msg": "新增用户成功",
	})
}

// DeleteUser
// @Summary 新增用户
// @Tags 用户模块
// @Param name query string false "用户名"
// @Param password query string false "密码"
// @Param repassword query string false "确认密码"
// @Sussess 200 {string} json{"code", "message"}
// @Router /user/create_user [get]
func CreateUse(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(-1, gin.H{
			"msg": "密码不一致",
		})
	}
	user.Password = password
	models.CreateUser(user)
	// fmt.Println(data)
	c.JSON(200, gin.H{
		"msg": "新增用户成功",
	})
}
