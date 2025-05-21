package entity

import (
	"gorm.io/gorm"
	"time"
)

type Scan struct {
	Id           int `gorm:"column:id;type:int;primaryKey;autoIncrement;not null"`
	APKVersionId int `gorm:"column:apk_version_id;type:int"`
	ScanVersion  int `gorm:"column:scan_version;type:int"`

	/* Static Information */
	//security Score information
	SecurityScore                    int    `gorm:"column:security_score;type:int"`
	RiskRatingGrade                  string `gorm:"column:risk_rating_grade;type:varchar(255)"`
	PrivacyRiskCount                 int    `gorm:"column:privacy_risk_count;type:int"`
	HighSeverityDistributionCount    int    `gorm:"column:high_severity_distribution_count;type:int"`
	WarningSeverityDistributionCount int    `gorm:"column:warning_severity_distribution_count;type:int"`
	InfoSeverityDistributionCount    int    `gorm:"column:info_severity_distribution_count;type:int"`
	SecureSeverityDistributionCount  int    `gorm:"column:secure_severity_distribution_count;type:int"`
	HotspotSeverityDistributionCount int    `gorm:"column:hotspot_severity_distribution_count;type:int"`

	//android
	ActivitiesCount       int `gorm:"column:activities_count;type:int"`
	ServicesCount         int `gorm:"column:services_count;type:int"`
	RecieversCount        int `gorm:"column:recievers_count;type:int"`
	ProvidersCount        int `gorm:"column:providers_count;type:int"`
	ExportedActivityCount int `gorm:"column:exported_activity_count;type:int"`
	ExportedServiceCount  int `gorm:"column:exported_service_count;type:int"`
	ExportedReceiverCount int `gorm:"column:exported_receiver_count;type:int"`
	ExportedProviderCount int `gorm:"column:exported_provider_count;type:int"`

	/* end static information */
	FinishedAt time.Time `gorm:"column:finished_at;type:timestamp"`

	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;->"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;default:null;->"`
}
