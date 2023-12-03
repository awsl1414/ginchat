package main

import (
	"fmt"
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "gorm.io/driver/sqlite"
)

// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }

func main() {
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db, _ := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	// db := utils.DB
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// 迁移 schema
	// db.AutoMigrate(&Product{})
	db.AutoMigrate(&models.UserBasic{})

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})
	user := &models.UserBasic{}
	user.Name = "小明"
	db.Create(user)

	// Read
	// var product Product
	// db.First(&product, 1)                 // 根据整型主键查找
	// db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	fmt.Println(db.First(&user, 1))

	// Update - 将 product 的 price 更新为 200
	// db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	db.Model(user).Update("password", "123456")

	// Delete - 删除 product
	// db.Delete(&product, 1)
}
