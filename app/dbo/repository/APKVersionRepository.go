package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mast-integrator/app/dbo/entity"
	"mast-integrator/app/helper"
	"time"
)

type APKVersionRepository interface {
	Create(ctx *gin.Context, db *gorm.DB, apkVersion entity.APKVersion) entity.APKVersion
	Update(ctx *gin.Context, db *gorm.DB, apkVersion entity.APKVersion) entity.APKVersion
	GetOneById(ctx *gin.Context, db *gorm.DB, id int) entity.APKVersion
	GetOneByVersionNameAndProjectId(ctx *gin.Context, db *gorm.DB, versionName string, projectId int) entity.APKVersion
	GetOneByProjectId(ctx *gin.Context, db *gorm.DB, projectId int) entity.APKVersion
	GetAllByProjectId(ctx *gin.Context, db *gorm.DB, projectId int) []entity.APKVersion
	DeleteAllByProjectId(ctx *gin.Context, db *gorm.DB, projectId int)
}

type APKVersionRepositoryImpl struct{}

func NewAPKVersionRepository() *APKVersionRepositoryImpl {
	return &APKVersionRepositoryImpl{}
}

func (A APKVersionRepositoryImpl) Create(ctx *gin.Context, db *gorm.DB, apkVersion entity.APKVersion) entity.APKVersion {
	err := db.WithContext(ctx).Create(&apkVersion).Error
	helper.ErrorHandler(err)

	return apkVersion
}

func (A APKVersionRepositoryImpl) Update(ctx *gin.Context, db *gorm.DB, apkVersion entity.APKVersion) entity.APKVersion {
	apkVersion.UpdatedAt = time.Now()

	err := db.WithContext(ctx).Model(&apkVersion).Updates(apkVersion).Error
	helper.ErrorHandler(err)

	return apkVersion
}

func (A APKVersionRepositoryImpl) GetOneById(ctx *gin.Context, db *gorm.DB, id int) entity.APKVersion {
	apkVersion := entity.APKVersion{}

	err := db.WithContext(ctx).Preload("Project").Where("id = ?", id).Find(&apkVersion).Error
	helper.ErrorHandler(err)

	return apkVersion
}

func (A APKVersionRepositoryImpl) GetOneByVersionNameAndProjectId(ctx *gin.Context, db *gorm.DB, versionName string, projectId int) entity.APKVersion {
	apkVersion := entity.APKVersion{}

	err := db.WithContext(ctx).
		Where("version_name = ?", versionName).
		Where("project_id = ?", projectId).
		Find(&apkVersion).Error
	helper.ErrorHandler(err)

	return apkVersion
}

func (A APKVersionRepositoryImpl) GetOneByProjectId(ctx *gin.Context, db *gorm.DB, projectId int) entity.APKVersion {
	apkVersion := entity.APKVersion{}

	err := db.WithContext(ctx).
		Where("project_id = ?", projectId).
		Find(&apkVersion).Error
	helper.ErrorHandler(err)

	return apkVersion
}

func (A APKVersionRepositoryImpl) GetAllByProjectId(ctx *gin.Context, db *gorm.DB, projectId int) []entity.APKVersion {
	apkVersions := []entity.APKVersion{}

	err := db.WithContext(ctx).
		Where("project_id = ?", projectId).
		Find(&apkVersions).Error
	helper.ErrorHandler(err)

	return apkVersions
}

func (A APKVersionRepositoryImpl) DeleteAllByProjectId(ctx *gin.Context, db *gorm.DB, projectId int) {
	apkVersions := []entity.APKVersion{}

	err := db.WithContext(ctx).
		Where("project_id = ?", projectId).
		Delete(&apkVersions).Error
	helper.ErrorHandler(err)
}
