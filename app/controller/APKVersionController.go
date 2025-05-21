package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mast-integrator/app/dto/request"
	"mast-integrator/app/exception"
	"mast-integrator/app/helper"
	"mast-integrator/app/service"
	"net/http"
	"strconv"
)

type APKVersionController struct {
	service service.APKVersionService
}

func NewAPKVersionController(service service.APKVersionService) *APKVersionController {
	return &APKVersionController{
		service: service,
	}
}

func (a *APKVersionController) UploadAPKHandler(ctx *gin.Context) {
	req := new(request.UploadAPKVersionForm)

	mobileProjectId, _ := strconv.Atoi(ctx.PostForm("mobile_project_id"))
	req.MobileProjectId = mobileProjectId
	req.APKFile, _ = ctx.FormFile("apk_file")
	req.APKVersionId = ctx.PostForm("apk_version_id")
	req.Overwrite = ctx.PostForm("overwrite")

	response := a.service.UploadAPKData(ctx, *req)

	ctx.JSON(http.StatusOK, response)
}

func (a *APKVersionController) GetApkVersionDetailByIdHandler(ctx *gin.Context) {
	apkVersionId, err := strconv.Atoi(ctx.Param("id"))
	helper.ErrorHandlerValidator(err)

	response := a.service.GetAPKVersionData(ctx, apkVersionId)
	ctx.JSON(http.StatusOK, response)
}

func (a *APKVersionController) GetApkVersionsByProjectHandler(ctx *gin.Context) {
	var projectId int
	var err error
	if ctx.Query("project_id") == "" {
		panic(exception.NewBadRequestError(errors.New("query parameter project_id is required").Error()))
		return
	} else {
		projectId, err = strconv.Atoi(ctx.Query("project_id"))
		if err != nil {
			panic(exception.NewBadRequestError(errors.New("query parameter project_id is must be integer").Error()))
			return
		}
	}

	responses := a.service.GetAPKVersionsData(ctx, projectId)
	ctx.JSON(http.StatusOK, responses)
}
