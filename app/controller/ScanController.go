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

type ScanController struct {
	service service.ScanService
}

func NewScanController(service service.ScanService) *ScanController {
	return &ScanController{
		service: service,
	}
}

func (s *ScanController) TriggerScanHandler(ctx *gin.Context) {
	var request request.ScanTriggerRequest
	err := ctx.ShouldBindJSON(&request)
	helper.ErrorHandler(err)

	s.service.TriggerScanData(ctx, request.APKVersionId, request.ScanType)

	ctx.JSON(http.StatusAccepted, "Success Scan APK Version, please wait until scan is complete")
}

func (s *ScanController) ScansHandler(ctx *gin.Context) {
	var apkVersionId int
	var err error
	if ctx.Query("apk_version_id") == "" {
		panic(exception.NewBadRequestError(errors.New("query parameter apk_version_id is required").Error()))
		return
	} else {
		apkVersionId, err = strconv.Atoi(ctx.Query("apk_version_id"))
		if err != nil {
			panic(exception.NewBadRequestError(errors.New("query parameter apk_version_id is must be integer").Error()))
			return
		}
	}

	responses := s.service.GetScansData(ctx, apkVersionId)
	ctx.JSON(http.StatusOK, responses)
}
