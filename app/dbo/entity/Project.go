package entity

import (
	"gorm.io/gorm"
	"time"
)

type Project struct {
	Id                 int            `gorm:"column:id;type:int;primaryKey;autoIncrement;not null"`
	Key                string         `gorm:"column:key;type:varchar(255);not null"`
	Name               string         `gorm:"column:name;type:varchar(255)"`
	Description        string         `gorm:"column:description;type:varchar(255)"`
	OsType             string         `gorm:"column:os_type;type:varchar(255)"`
	PackageName        string         `gorm:"column:package_name;type:varchar(255)"`
	CurrentScanVersion int            `gorm:"column:current_scan_version;type:int;default:0"`
	CreatedAt          time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;->"`
	UpdatedAt          time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;default:null;->"`
}

func (Project) TableName() string {
	return "projects"
}
