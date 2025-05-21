package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"mast-integrator/app/controller"
	"mast-integrator/app/dbo/api"
	"mast-integrator/app/dbo/repository"
	"mast-integrator/app/service"
)

func ScanRoute(router *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) *gin.RouterGroup {
	scanRepository := repository.NewScanRepository()
	apkVersionRepository := repository.NewAPKVersionRepository()
	mobsfApi := api.NewMobSFAPI()

	scanService := service.NewScanService(scanRepository, apkVersionRepository, mobsfApi, db, validate)

	scanController := controller.NewScanController(scanService)

	router.POST("scan", scanController.TriggerScanHandler)
	router.GET("scans", scanController.ScansHandler)

	return router
}
