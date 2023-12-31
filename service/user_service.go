package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"math/rand"
	"strconv"

	"github.com/asaskevich/govalidator"
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

// GetUserList
// @Summary 查询用户
// @Tags 用户模块
// @Param name query string false "用户名"
// @Param password query string false "密码"
// @Sussess 200 {string} json{"code", "message"}
// @Router /user/get_user_list [get]
func FindUserByNameAndPwd(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	user := models.FindUserByName(name)
	if user.Identity == "" {
		c.JSON(200, gin.H{
			"msg": "用户名不存在",
		})

	}
	flag := utils.VaildPassword(password, user.Salt, user.Password)
	if !flag {
		c.JSON(200, gin.H{
			"msg": "密码不正确",
		})

	}
	data := models.FindUserByNameAndPwd(name, password)
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

	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(-1, gin.H{
			"msg": "用户名已注册",
		})
		return
	}
	if password != repassword {
		c.JSON(-1, gin.H{
			"msg": "密码不一致",
		})
		return
	}
	salt := fmt.Sprintf("%0.6d", rand.Int31())
	// user.Password = password
	user.Password = utils.MakePassword(password, salt)
	models.CreateUser(user)
	// fmt.Println(data)
	c.JSON(200, gin.H{
		"msg": "新增用户成功",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Sussess 200 {string} json{"code", "message"}
// @Router /user/delete_user [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"msg": "删除用户成功",
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Sussess 200 {string} json{"code", "message"}
// @Router /user/update_user [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"msg": "修改用户成功",
		})
	}
}
