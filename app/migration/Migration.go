package migration

import (
	"gorm.io/gorm"
	"mast-integrator/app/dbo/entity"
)

func DoMigration(db *gorm.DB) {
	db.AutoMigrate(&entity.Project{})
	db.AutoMigrate(&entity.APKVersion{})
	db.AutoMigrate(&entity.Scan{})
}
