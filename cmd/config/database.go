package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go-grpc"))
	if err != nil {
		log.Fatal("Database connection failed", err.Error())
	}

	return db
}
