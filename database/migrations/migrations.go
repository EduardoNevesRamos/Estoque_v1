package migrations

import (
	"github.com/DuduNeves/Estoque_v1/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Products{})
}
