package staticScan

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mast-integrator/app/dbo/api"
	"mast-integrator/app/dbo/entity"
	"mast-integrator/app/dbo/repository"
	"mast-integrator/app/helper"
	"mast-integrator/app/shareVar"
)

type StaticScan interface {
	ScanProcess(ctx *gin.Context, param StaticScanProcessParam)
}

type StaticScanProcessParam struct {
	ApkVersion entity.APKVersion
	Scan       entity.Scan
}

type StaticScanAbstract struct {
	repo           repository.ScanRepository
	repoApkVersion repository.APKVersionRepository
	mobsfApi       api.MobSFAPI
	db             *gorm.DB
	stdLog         *helper.StandartLog
}

func NewStaticScanAbstract(repo repository.ScanRepository, repoApkVersion repository.APKVersionRepository, mobsfApi api.MobSFAPI, db *gorm.DB) *StaticScanAbstract {
	return &StaticScanAbstract{
		repo:           repo,
		repoApkVersion: repoApkVersion,
		mobsfApi:       mobsfApi,
		db:             db,
		stdLog:         helper.NewStandardLog(shareVar.Scan, shareVar.Service),
	}
}

func GetStaticScan(osType string, abstract *StaticScanAbstract) StaticScan {
	if osType == "android" {
		return NewScanAPKStatic(*abstract)
	} else {
		return NewScanIOSStatic()
	}
}

func (s StaticScanAbstract) processError(ctx *gin.Context, apkVersion entity.APKVersion, statusMessage string) {
	apkVersion.StaticScanStatus = 2
	apkVersion.StaticStatusMessage = statusMessage
	s.repoApkVersion.Update(ctx, s.db, apkVersion)

	customFields := logrus.Fields{
		"message":        "Scan Static Analyzer APK Version failed",
		"project_id":     apkVersion.ProjectId,
		"apk_version_id": apkVersion.Id,
		"package_name":   apkVersion.PackageName,
		"version_name":   apkVersion.VersionName,
		"error":          statusMessage,
	}
	s.stdLog.ErrorFunction(customFields)
}
