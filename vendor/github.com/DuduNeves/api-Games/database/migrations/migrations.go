package migrations

import (
	"github.com/DuduNeves/api-Games/models"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(models.Game{})
}
