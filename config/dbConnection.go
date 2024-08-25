package config

import (
	"fmt"
	"go_fiber_boilerplate/Apps/User"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB() *gorm.DB {
	LoadEnv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dhaka",
		POSTGRES_HOST, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Db connection error: ", err)
	}
	AutoMigrate(db)
	return db
}

func AutoMigrate(connection *gorm.DB) {
	err := connection.Debug().AutoMigrate(&User.User{})
	if err != nil {
		log.Println("Database migration error: ", err)
	}
}
