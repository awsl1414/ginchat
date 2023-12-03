package main

import (
	"ginchat/router"
	"ginchat/utils"
)

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"msg": "pong",
// 		})
// 	})
// 	r.Run()
// }

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	r := router.Router()
	r.Run("localhost:8081")
}
