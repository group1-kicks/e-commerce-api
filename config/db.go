package config

import (
	cartMdl "e-commerce-api/feature/cart/data"
	orderItemModel "e-commerce-api/feature/order/data"
	orderModel "e-commerce-api/feature/order/data"
	productModel "e-commerce-api/feature/product/data"
	"e-commerce-api/feature/users/data"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(ac AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ac.DBUser, ac.DBPass, ac.DBHost, ac.DBPort, ac.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection error : ", err.Error())
		return nil
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&data.User{})
	db.AutoMigrate(&productModel.Product{})
	db.AutoMigrate(&orderModel.Order{})
	db.AutoMigrate(&orderItemModel.OrderItem{})
	db.AutoMigrate(&cartMdl.Cart{})
}
