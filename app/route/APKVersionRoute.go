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

func APKVersionRoute(router *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) *gin.RouterGroup {
	repo := repository.NewAPKVersionRepository()
	repoProject := repository.NewProjectRepository()
	mobsfApi := api.NewMobSFAPI()

	apkVersionService := service.NewAPKVersionService(repo, repoProject, mobsfApi, validate, db)

	apkVersionController := controller.NewAPKVersionController(apkVersionService)

	router.POST("apk-version/upload", apkVersionController.UploadAPKHandler)
	router.GET("apk-version/:id", apkVersionController.GetApkVersionDetailByIdHandler)
	router.GET("apk-versions", apkVersionController.GetApkVersionsByProjectHandler)

	return router
}
