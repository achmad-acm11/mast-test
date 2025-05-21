package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"mast-integrator/app/dbo/api"
	"mast-integrator/app/dbo/entity"
	"mast-integrator/app/dbo/repository"
	"mast-integrator/app/dto/response"
	"mast-integrator/app/exception"
	"mast-integrator/app/helper"
	"mast-integrator/app/service/staticScan"
	"mast-integrator/app/shareVar"
)

type ScanService interface {
	TriggerScanData(ctx *gin.Context, apkVersionId int, scanType string)
	GetScansData(ctx *gin.Context, apkVersionId int) []response.ScanResponse
}

type ScanServiceImpl struct {
	repo           repository.ScanRepository
	repoApkVersion repository.APKVersionRepository
	mobsfApi       api.MobSFAPI
	validator      *validator.Validate
	db             *gorm.DB
	stdLog         *helper.StandartLog
}

func NewScanService(repo repository.ScanRepository, repoApkVersion repository.APKVersionRepository, mobsfApi api.MobSFAPI, db *gorm.DB, validate *validator.Validate) *ScanServiceImpl {
	return &ScanServiceImpl{
		repo:           repo,
		mobsfApi:       mobsfApi,
		repoApkVersion: repoApkVersion,
		db:             db,
		validator:      validate,
		stdLog:         helper.NewStandardLog(shareVar.Scan, shareVar.Service),
	}
}

func (s ScanServiceImpl) TriggerScanData(ctx *gin.Context, apkVersionId int, scanType string) {
	apkVersion, scan := s.initScanData(ctx, apkVersionId, scanType)
	//fmt.Printf("apkVersion:%v, scan:%v\n", apkVersion, scan)
	abstract := staticScan.NewStaticScanAbstract(s.repo, s.repoApkVersion, s.mobsfApi, s.db)
	if scanType == "static" {
		param := staticScan.StaticScanProcessParam{
			ApkVersion: apkVersion,
			Scan:       scan,
		}
		go staticScan.GetStaticScan(apkVersion.Project.OsType, abstract).ScanProcess(ctx, param)
		//if  == "android" {
		//	go s.AsyncScanAPK(ctx, apkVersion, scan)
		//} else {
		//	go s.AsyncScanIOS(ctx, apkVersion, scan)
		//}
	} else if scanType == "dynamic" {
		//global.Services.DynamicAnalysisService.AsyncScan(APKVersion, Scan)
	} else {
		//if apkVersion.Project.OsType == "android" {
		//	go s.AsyncScanAPK(ctx, apkVersion, scan)
		//} else {
		//	go s.AsyncScanIOS(ctx, apkVersion, scan)
		//}
		//global.Services.DynamicAnalysisService.AsyncScan(APKVersion, Scan)
	}
}

func (s ScanServiceImpl) GetScansData(ctx *gin.Context, apkVersionId int) []response.ScanResponse {
	scans := s.repo.GetAllByAPKVersionId(ctx, s.db, apkVersionId)

	return response.NewScanResponseBuilder().List(scans).ListResult()
}

func (s ScanServiceImpl) initScanData(ctx *gin.Context, apkVersionId int, scanType string) (entity.APKVersion, entity.Scan) {
	var isErrorStatus = false

	apkVersion := s.getApkVersionWithThrow(ctx, apkVersionId)

	if apkVersion.StaticScanStatus == 3 {
		panic(exception.NewBadRequestError(errors.New("apk version already scanned, please reupload apk version for rescan all").Error()))
	}

	if apkVersion.StaticScanStatus == 1 || apkVersion.DynamicScanStatus == 1 {
		panic(exception.NewConflictError(errors.New("apk version on scan process").Error()))
	}

	//if scanType == "static" && apkVersion.StaticScanStatus != 2 {
	//	panic(exception.NewBadRequestError(errors.New("only failed static scan can be re-scan").Error()))
	//}
	//
	//if scanType == "static" && apkVersion.StaticScanStatus == 2 {
	//	apkVersion.StaticScanStatus = 1
	//	isErrorStatus = true
	//}

	//s.repoApkVersion.Update(ctx, s.db, apkVersion)
	//fields := logrus.Fields{
	//	"message":        "Scan APK Version started",
	//	"project_id":     apkVersion.ProjectId,
	//	"apk_version_id": apkVersion.Id,
	//	"package_name":   apkVersion.PackageName,
	//	"version_name":   apkVersion.VersionName,
	//}
	//s.stdLog.InfoFunction(fields)

	currentScanData := s.getLastScanDataByCurrentVersion(ctx, apkVersion.Id, isErrorStatus)

	return apkVersion, currentScanData
}

func (s ScanServiceImpl) getLastScanDataByCurrentVersion(ctx *gin.Context, apkVersionId int, isErrorStatus bool) entity.Scan {
	scanVersion := 1
	lastScan := s.repo.GetLastByAPKVersionId(ctx, s.db, apkVersionId)

	if !isErrorStatus {
		if lastScan.Id != 0 {
			scanVersion = lastScan.ScanVersion + 1
		}
		scan := entity.Scan{
			APKVersionId: apkVersionId,
			ScanVersion:  scanVersion,
		}
		scanNew := s.repo.Create(ctx, s.db, scan)
		return scanNew
	} else {
		if lastScan.Id == 0 {
			scanNew := entity.Scan{
				APKVersionId: apkVersionId,
				ScanVersion:  scanVersion,
			}
			s.repo.Create(ctx, s.db, scanNew)
			return scanNew
		} else {
			return lastScan
		}
	}
}

func (s ScanServiceImpl) getApkVersionWithThrow(ctx *gin.Context, apkVersionId int) entity.APKVersion {
	apkVersion := s.repoApkVersion.GetOneById(ctx, s.db, apkVersionId)
	if apkVersion.Id == 0 {
		panic(exception.NewNotFoundError(errors.New("apk version not found").Error()))
	}
	return apkVersion
}
