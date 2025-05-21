package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"mast-integrator/app/dbo/api"
	"mast-integrator/app/dbo/entity"
	"mast-integrator/app/dbo/repository"
	"mast-integrator/app/dto/request"
	"mast-integrator/app/dto/response"
	"mast-integrator/app/exception"
	"mast-integrator/app/helper"
	"mast-integrator/app/shareVar"
	"time"
)

type APKVersionService interface {
	UploadAPKData(ctx *gin.Context, req request.UploadAPKVersionForm) response.APKVersionResponse
	GetAPKVersionData(ctx *gin.Context, apkVersionId int) response.APKVersionResponse
	GetAPKVersionsData(ctx *gin.Context, projectId int) []response.APKVersionResponse
}

type APKVersionServiceImpl struct {
	repo        repository.APKVersionRepository
	repoProject repository.ProjectRepository
	mobsfApi    api.MobSFAPI
	validator   *validator.Validate
	db          *gorm.DB
	stdLog      *helper.StandartLog
}

func NewAPKVersionService(repo repository.APKVersionRepository, repoProject repository.ProjectRepository, mobsfApi api.MobSFAPI, validate *validator.Validate, db *gorm.DB) *APKVersionServiceImpl {
	return &APKVersionServiceImpl{
		repo:        repo,
		repoProject: repoProject,
		mobsfApi:    mobsfApi,
		validator:   validate,
		db:          db,
		stdLog:      helper.NewStandardLog(shareVar.APKVersion, shareVar.Service),
	}
}

func (a APKVersionServiceImpl) GetAPKVersionData(ctx *gin.Context, apkVersionId int) response.APKVersionResponse {
	apkVersion := a.repo.GetOneById(ctx, a.db, apkVersionId)
	if apkVersion.Id == 0 {
		panic(exception.NewNotFoundError("APK Version Not Found"))
	}
	return response.NewAPKVersionResponseBuilder().Default(apkVersion).Result()
}

func (a APKVersionServiceImpl) GetAPKVersionsData(ctx *gin.Context, projectId int) []response.APKVersionResponse {
	apkVersions := a.repo.GetAllByProjectId(ctx, a.db, projectId)

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		helper.ErrorHandler(err)
	}
	t1 := time.Now().In(loc)

	for _, apkVersion := range apkVersions {
		t2 := apkVersion.UpdatedAt
		diff := t1.Sub(t2).Minutes()
		if apkVersion.StaticScanStatus == 1 {
			if diff > 15 {
				apkVersion.StaticScanStatus = 2
				apkVersion.StaticStatusMessage = "Scan timeout"
				a.repo.Update(ctx, a.db, apkVersion)
			}
		}

		if apkVersion.DynamicScanStatus == 1 {
			if diff > 15 {
				apkVersion.DynamicScanStatus = 2
				apkVersion.DynamicStatusMessage = "Scan timeout"
				a.repo.Update(ctx, a.db, apkVersion)
			}
		}
	}

	return response.NewAPKVersionResponseBuilder().List(apkVersions).ListResult()
}

func (a APKVersionServiceImpl) getProjectWithException(ctx *gin.Context, id int) entity.Project {
	project := a.repoProject.GetOneById(ctx, a.db, id)
	if project.Id == 0 {
		panic(exception.NewNotFoundError(errors.New(shareVar.PROJECT_NOT_FOUND).Error()))
	}
	return project
}
