package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"mast-integrator/app/controller"
	"mast-integrator/app/dbo/repository"
	"mast-integrator/app/service"
)

func ProjectRoute(router *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) *gin.RouterGroup {
	projectRepository := repository.NewProjectRepository()

	projectService := service.NewProjectService(projectRepository, validate, db)

	projectController := controller.NewProjectController(projectService)

	router.GET("projects", projectController.GetAllHandler)
	router.GET("project/:id", projectController.GetDetailByIdHandler)
	router.POST("project", projectController.CreateProjectHandler)
	router.PUT("project/:id", projectController.UpdateProjectHandler)
	router.DELETE("project/:id", projectController.DeleteProjectHandler)

	return router
}
