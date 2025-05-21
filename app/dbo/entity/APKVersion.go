package entity

import (
	"gorm.io/gorm"
	"time"
)

type APKVersion struct {
	Id        int `gorm:"column:id;type:int;primaryKey;autoIncrement;not null"`
	ProjectId int `gorm:"column:project_id;type:int;not null"`
	// app information
	AppName     string `gorm:"column:app_name;type:varchar(255)"`
	AppType     string `gorm:"column:app_type;type:varchar(255)"`
	PackageName string `gorm:"column:package_name;type:varchar(255)"`
	VersionName string `gorm:"column:version_name;type:varchar(255)"`
	HashPath    string `gorm:"column:hash_path;type:varchar(255)"`

	//android
	TargetSdk    string `gorm:"column:target_sdk;type:varchar(255)"`
	MaxSdk       string `gorm:"column:max_sdk;type:varchar(255)"`
	MinSdk       string `gorm:"column:min_sdk;type:varchar(255)"`
	VersionCode  string `gorm:"column:version_code;type:varchar(255)"`
	MainActivity string `gorm:"column:main_activity;type:varchar(255)"`

	//ios
	Build             string `gorm:"column:build;type:varchar(255)"`
	SdkName           string `gorm:"column:sdk_name;type:varchar(255)"`
	Platform          string `gorm:"column:platform;type:varchar(255)"`
	MinOsVersion      string `gorm:"column:min_os_version;type:varchar(255)"`
	BinaryInfoEndian  string `gorm:"column:binary_info_endian;type:varchar(255)"`
	BinaryInfoBit     string `gorm:"column:binary_info_bit;type:varchar(255)"`
	BinaryInfoArch    string `gorm:"column:binary_info_arch;type:varchar(255)"`
	BinaryInfoSubarch string `gorm:"column:binary_info_subarch;type:varchar(255)"`

	//file information
	IconPath   string `gorm:"column:icon_path;type:varchar(255)"`
	Filename   string `gorm:"column:file_name;type:varchar(255)"`
	Size       string `gorm:"column:size;type:varchar(255)"`
	FileMD5    string `gorm:"column:file_md_5;type:varchar(255)"`
	FileSHA5   string `gorm:"column:file_sh_5;type:varchar(255)"`
	FileSHA256 string `gorm:"column:file_sh_256;type:varchar(255)"`

	//scan information
	StaticScanStatus     int    `gorm:"column:static_scan_status;type:int"`
	StaticStatusMessage  string `gorm:"column:static_status_message;type:varchar(255)"`
	DynamicScanStatus    int    `gorm:"column:dynamic_scan_status;type:int"`
	DynamicStatusMessage string `gorm:"column:dynamic_status_message;type:varchar(255)"`
	ScanType             string `gorm:"column:scan_type;type:varchar(255)"`
	Extension            string `gorm:"column:extension;type:varchar(255)"`

	Project *Project `gorm:"foreignKey:ProjectId;references:Id"`

	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;->"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;default:null;->"`
}
