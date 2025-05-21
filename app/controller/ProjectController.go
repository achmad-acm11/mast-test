package controller

import (
	"github.com/gin-gonic/gin"
	"mast-integrator/app/dto/request"
	"mast-integrator/app/helper"
	"mast-integrator/app/service"
	"mast-integrator/app/shareVar"
	"net/http"
	"strconv"
)

type ProjectController struct {
	service service.ProjectService
}

func NewProjectController(service service.ProjectService) *ProjectController {
	return &ProjectController{
		service: service,
	}
}

func (p *ProjectController) GetAllHandler(ctx *gin.Context) {
	responses := p.service.GetAllData(ctx)

	ctx.JSON(http.StatusOK, responses)
}

func (p *ProjectController) GetDetailByIdHandler(ctx *gin.Context) {
	projectId, err := strconv.Atoi(ctx.Param("id"))
	helper.ErrorHandlerValidator(err)

	response := p.service.GetDetailByIdData(ctx, projectId)

	ctx.JSON(http.StatusOK, response)
}

func (p *ProjectController) CreateProjectHandler(ctx *gin.Context) {
	var request request.CreateProjectRequest
	err := ctx.ShouldBindJSON(&request)
	helper.ErrorHandler(err)

	response := p.service.CreateProjectData(ctx, request)

	ctx.JSON(http.StatusOK, response)
}

func (p *ProjectController) UpdateProjectHandler(ctx *gin.Context) {
	projectId, err := strconv.Atoi(ctx.Param("id"))
	helper.ErrorHandler(err)

	var request request.UpdateProjectRequest
	err = ctx.ShouldBindJSON(&request)
	helper.ErrorHandler(err)

	response := p.service.UpdateProjectData(ctx, request, projectId)

	ctx.JSON(http.StatusOK, response)
}

func (p *ProjectController) DeleteProjectHandler(ctx *gin.Context) {
	projectId, err := strconv.Atoi(ctx.Param("id"))
	helper.ErrorHandler(err)

	p.service.DeleteProjectData(ctx, projectId)

	ctx.JSON(http.StatusOK, shareVar.PROJECT_DELETED)
}
