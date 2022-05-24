package migrations

import (
	"github.com/DuduNeves/Estoque_v1/database/entity"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(entity.Products{})
	db.AutoMigrate(entity.User{})
}
