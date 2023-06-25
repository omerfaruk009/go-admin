package database

import (
	"fmt"
	"go-admin2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Dns = "host=localhost user=omer password=model-f80 dbname=go-admin port=5432 sslmode=disable TimeZone=Asia/Shanghai"
var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(postgres.Open(Dns), &gorm.Config{})
	if err != nil {
		panic("database bağlanmadı")
	}

	DB = database
	err = database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, models.Product{}, &models.Order{}, &models.OrderItem{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
