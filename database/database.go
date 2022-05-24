package database

import (
	"log"
	"time"

	"github.com/DuduNeves/Estoque_v1/database/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := "eduardo:123456@tcp(172.18.0.2:3306)/storage?parseTime=true&loc=Local"
	database, err := gorm.Open(mysql.Open(str), &gorm.Config{})

	if err != nil {
		log.Fatal("Error: ", err)
	}

	db = database

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)

}

func GetDatabase() *gorm.DB {
	return db
}
