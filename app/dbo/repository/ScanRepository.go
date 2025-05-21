package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mast-integrator/app/dbo/entity"
	"mast-integrator/app/helper"
	"time"
)

type ScanRepository interface {
	Create(ctx *gin.Context, db *gorm.DB, scan entity.Scan) entity.Scan
	Update(ctx *gin.Context, db *gorm.DB, scan entity.Scan) entity.Scan
	GetAllByAPKVersionId(ctx *gin.Context, db *gorm.DB, apkVersionId int) []entity.Scan
	GetLastByAPKVersionId(ctx *gin.Context, db *gorm.DB, apkVersionId int) entity.Scan
	DeleteAllByAPKVersionId(ctx *gin.Context, db *gorm.DB, apkVersionId int)
}

type ScanRepositoryImpl struct {
}

func NewScanRepository() *ScanRepositoryImpl {
	return &ScanRepositoryImpl{}
}

func (s ScanRepositoryImpl) Create(ctx *gin.Context, db *gorm.DB, scan entity.Scan) entity.Scan {
	err := db.WithContext(ctx).Create(&scan).Error
	helper.ErrorHandler(err)

	return scan
}

func (s ScanRepositoryImpl) Update(ctx *gin.Context, db *gorm.DB, scan entity.Scan) entity.Scan {
	scan.UpdatedAt = time.Now()

	err := db.WithContext(ctx).Model(&scan).Updates(scan).Error
	helper.ErrorHandler(err)

	return scan
}

func (s ScanRepositoryImpl) GetAllByAPKVersionId(ctx *gin.Context, db *gorm.DB, apkVersionId int) []entity.Scan {
	scans := []entity.Scan{}

	err := db.WithContext(ctx).
		Where("apk_version_id = ?", apkVersionId).
		Find(&scans).Error
	helper.ErrorHandler(err)

	return scans
}

func (s ScanRepositoryImpl) GetLastByAPKVersionId(ctx *gin.Context, db *gorm.DB, apkVersionId int) entity.Scan {
	scan := entity.Scan{}

	err := db.WithContext(ctx).
		Where("apk_version_id = ?", apkVersionId).
		Order("id desc").Find(&scan).Error
	helper.ErrorHandler(err)

	return scan
}

func (s ScanRepositoryImpl) DeleteAllByAPKVersionId(ctx *gin.Context, db *gorm.DB, apkVersionId int) {
	scans := []entity.Scan{}

	err := db.WithContext(ctx).
		Where("apk_version_id = ?", apkVersionId).
		Delete(&scans).Error
	helper.ErrorHandler(err)
}
