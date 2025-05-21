package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"mast-integrator/app/dbo/entity"
	"mast-integrator/app/dto/request"
	"mast-integrator/app/dto/response"
	"mast-integrator/app/exception"
	"mast-integrator/app/helper"
	"mast-integrator/app/service/manifestAPK"
	"mime/multipart"
	"os"
	"slices"
	"strconv"
	"strings"
)

type checkAPKVersionParam struct {
	currentApkVersion entity.APKVersion
	latestApkVersion  entity.APKVersion
	isOverwrite       bool
	apkVersionId      int
	filePath          string
	uploadDir         string
	versionName       string
	packageName       string
	appName           string
	project           entity.Project
	extension         string
}

func (a APKVersionServiceImpl) UploadAPKData(ctx *gin.Context, req request.UploadAPKVersionForm) response.APKVersionResponse {
	apkVersionRequest := validationRequest(req, a.validator)

	filePath := saveAPKFile(ctx, apkVersionRequest.APKFile)

	if apkVersionRequest.Overwrite == true && apkVersionRequest.APKVersionId != 0 {
		panic(exception.NewBadRequestError(errors.New("overwrite and apk version id cannot be used together").Error()))
	}

	project := a.getProjectWithException(ctx, apkVersionRequest.ProjectId)

	curDir, _ := os.Getwd()
	uploadDir := curDir + "/tmp"
	extension := apkVersionRequest.APKFile.Filename[strings.LastIndex(req.APKFile.Filename, ".")+1:]
	checkOsTypeAndExtension(project.OsType, extension, filePath)

	manifestAPKResult := manifestAPK.GetManifestAPK(project.OsType, extension, project.Id).GetResult(filePath)
	errParse := manifestAPKResult.Err
	if errParse != nil {
		deleteFilePathUploadDir(filePath, uploadDir, project, extension)
		throwException(exception.BadRequestError{}, "error get manifest "+project.OsType+","+errParse.Error())
	}
	
	log.Fatal(0)

	var checkApkVersion entity.APKVersion
	checkLatest := a.repo.GetOneByProjectId(ctx, a.db, apkVersionRequest.ProjectId)
	checkApkVersion = a.repo.GetOneByVersionNameAndProjectId(ctx, a.db, manifestAPKResult.Version, apkVersionRequest.ProjectId)

	if checkLatest.Id != 0 {
		param := checkAPKVersionParam{
			currentApkVersion: checkApkVersion,
			latestApkVersion:  checkLatest,
			isOverwrite:       apkVersionRequest.Overwrite,
			apkVersionId:      apkVersionRequest.APKVersionId,
			filePath:          filePath,
			uploadDir:         uploadDir,
			versionName:       manifestAPKResult.Version,
			packageName:       manifestAPKResult.PackageName,
			appName:           manifestAPKResult.AppName,
			project:           project,
			extension:         extension,
		}
		a.checkAPKVersionData(ctx, param)
	}

	mobsfResponse := a.mobsfApi.Upload(filePath)
	if mobsfResponse == nil {
		deleteFilePathUploadDir(filePath, uploadDir, project, extension)
		throwException(exception.InternalServerError{}, "Upload Failed")
	}

	var apkVersion entity.APKVersion
	if checkApkVersion.Id != 0 {
		param := request.OverwriteAPKVersionRequest{
			CurrentAPKVersion: checkApkVersion,
			ProjectId:         project.Id,
			PackageName:       manifestAPKResult.PackageName,
			VersionName:       manifestAPKResult.Version,
			AppName:           manifestAPKResult.AppName,
			FileName:          mobsfResponse.FileName,
			HashPath:          mobsfResponse.Hash,
			ScanType:          mobsfResponse.ScanType,
			Extension:         extension,
		}
		apkVersion = a.OverwriteAPKVersionData(ctx, param)
	} else {
		param := request.CreateAPKVersionRequest{
			ProjectId:   project.Id,
			PackageName: manifestAPKResult.PackageName,
			VersionName: manifestAPKResult.Version,
			AppName:     manifestAPKResult.AppName,
			FileName:    mobsfResponse.FileName,
			HashPath:    mobsfResponse.Hash,
			ScanType:    mobsfResponse.ScanType,
			Extension:   extension,
		}
		apkVersion = a.CreateAPKVersionData(ctx, param)
	}

	if checkLatest.Id == 0 {
		project.PackageName = manifestAPKResult.PackageName
		a.repoProject.Update(ctx, a.db, project)
	}

	deleteFilePathUploadDir(filePath, uploadDir, project, extension)

	apkVersionResponse := response.NewAPKVersionResponseBuilder().Default(apkVersion).Result()
	return apkVersionResponse
}

func (a APKVersionServiceImpl) checkAPKVersionData(ctx *gin.Context, param checkAPKVersionParam) {
	isCurrentApkVersionExist := (param.currentApkVersion.Id != 0)
	isOverwrite := param.isOverwrite
	isStaticScanOnRunning := param.currentApkVersion.StaticScanStatus == 1
	isDynamicScanOnRunning := param.currentApkVersion.DynamicScanStatus == 1
	isApkVersionSelected := param.apkVersionId != 0

	// TODO CurrentApkVersion Exist and want to overwrite but currentApkVersion on progress scanning
	if (isCurrentApkVersionExist && isOverwrite) && (isStaticScanOnRunning || isDynamicScanOnRunning) {
		deleteFilePathUploadDir(param.filePath, param.uploadDir, param.project, param.extension)
		throwException(exception.ConflictError{}, "apk version on progress, please wait until the process is finished")
	}

	// TODO CurrentApkVersion Exist and don't want to overwrite but target apkVersionId is zero
	if (isCurrentApkVersionExist && isOverwrite == false) && isApkVersionSelected == false {
		deleteFilePathUploadDir(param.filePath, param.uploadDir, param.project, param.extension)
		throwException(exception.ConflictError{}, "the version already exist, checklist the overwrite flag to replace the version")
	}

	// TODO CurrentApkVersion Exist and don't want to overwrite but target apkVersionId is not zero
	if (isCurrentApkVersionExist && isOverwrite == false) && isApkVersionSelected {
		checkApkVersion := a.repo.GetOneById(ctx, a.db, param.apkVersionId)
		if checkApkVersion.Id == 0 {
			deleteFilePathUploadDir(param.filePath, param.uploadDir, param.project, param.extension)
			throwException(exception.NotFoundError{}, "apk version not found")
		} else {
			if param.versionName != checkApkVersion.VersionName {
				deleteFilePathUploadDir(param.filePath, param.uploadDir, param.project, param.extension)
				throwException(exception.BadRequestError{}, "the version invalid, current version is "+checkApkVersion.VersionName+" and the version that you upload is "+param.versionName)
			}
			if checkApkVersion.StaticScanStatus == 1 || checkApkVersion.DynamicScanStatus == 1 {
				deleteFilePathUploadDir(param.filePath, param.uploadDir, param.project, param.extension)
				throwException(exception.ConflictError{}, "apk version on progress, please wait until the process is finished")
			}
		}

	}

	// TODO Current PackageName not same with packageName apk upload
	if param.project.PackageName != param.packageName {
		deleteFilePathUploadDir(param.filePath, param.uploadDir, param.project, param.extension)
		throwException(exception.BadRequestError{}, "package name not match, current package name is "+param.project.PackageName+" and package name in uploaded apk is "+param.packageName)
	}

	// TODO Current AppName not same with appname apk upload
	if param.latestApkVersion.AppName != param.appName {
		deleteFilePathUploadDir(param.filePath, param.uploadDir, param.project, param.extension)
		throwException(exception.BadRequestError{}, "app name not match, current app name is "+param.latestApkVersion.AppName+" and app name in uploaded apk is "+param.appName)
	}
}

func (a APKVersionServiceImpl) CreateAPKVersionData(ctx *gin.Context, req request.CreateAPKVersionRequest) entity.APKVersion {
	apkVersion := entity.APKVersion{}

	apkVersion.ProjectId = req.ProjectId
	apkVersion.PackageName = req.PackageName
	apkVersion.VersionName = req.VersionName
	apkVersion.AppName = req.AppName
	apkVersion.Filename = req.FileName
	apkVersion.HashPath = req.HashPath
	apkVersion.ScanType = req.ScanType
	apkVersion.StaticScanStatus = 0
	apkVersion.DynamicScanStatus = 0
	apkVersion.StaticStatusMessage = ""
	apkVersion.DynamicStatusMessage = ""
	apkVersion.Extension = req.Extension

	a.repo.Create(ctx, a.db, apkVersion)

	return apkVersion
}

func (a APKVersionServiceImpl) OverwriteAPKVersionData(ctx *gin.Context, req request.OverwriteAPKVersionRequest) entity.APKVersion {
	currentApkVersion := req.CurrentAPKVersion

	currentApkVersion.ProjectId = req.ProjectId
	currentApkVersion.PackageName = req.PackageName
	currentApkVersion.VersionName = req.VersionName
	currentApkVersion.AppName = req.AppName
	currentApkVersion.Filename = req.FileName
	currentApkVersion.HashPath = req.HashPath
	currentApkVersion.ScanType = req.ScanType
	currentApkVersion.StaticScanStatus = 0
	currentApkVersion.DynamicScanStatus = 0
	currentApkVersion.StaticStatusMessage = ""
	currentApkVersion.DynamicStatusMessage = ""
	currentApkVersion.Extension = req.Extension

	a.repo.Update(ctx, a.db, currentApkVersion)

	return currentApkVersion
}

func validationRequest(req request.UploadAPKVersionForm, validator *validator.Validate) *request.UploadAPKVersionRequest {
	apkVersionParam := request.NewUploadAPKVersionParam(req, validator)
	errMessage, isError := apkVersionParam.ValidateRequest()
	if isError == true {
		panic(exception.NewValidationError(errMessage))
	}
	return apkVersionParam.GetResultRequest()
}

func saveAPKFile(ctx *gin.Context, file *multipart.FileHeader) string {
	fileName := file.Filename
	curDir, _ := os.Getwd()
	uploadDir := curDir + "/tmp"

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, 0755)
	}
	filePath := uploadDir + "/" + fileName
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		helper.ErrorHandler(err)
	}
	return filePath
}

func checkOsTypeAndExtension(ostype string, extension string, filePath string) {
	if ostype == "android" && slices.Contains([]string{"apk", "xapk"}, extension) == false {
		os.Remove(filePath)
		throwException(exception.BadRequestError{}, "file extension must be apk or xapk")
	} else if ostype == "ios" && extension != "ipa" {
		os.Remove(filePath)
		throwException(exception.BadRequestError{}, "file extension must be ipa")
	}
}

func deleteFilePathUploadDir(filePath string, uploadDir string, project entity.Project, extension string) {
	os.Remove(filePath)
	if project.OsType != "android" {
		os.RemoveAll(uploadDir + "/" + strconv.Itoa(project.Id) + "/")
	} else if project.OsType == "android" && extension == "xapk" {
		os.RemoveAll(uploadDir + "/" + strconv.Itoa(project.Id) + "/")
	}
}

func throwException(errType any, errMessage string) {
	switch errType.(type) {
	case exception.NotFoundError:
		panic(exception.NewNotFoundError(errors.New(errMessage).Error()))
	case exception.BadRequestError:
		panic(exception.NewBadRequestError(errors.New(errMessage).Error()))
	case exception.ConflictError:
		panic(exception.NewConflictError(errors.New(errMessage).Error()))
	case exception.InternalServerError:
		panic(exception.NewInternalServerError(errors.New(errMessage).Error()))
	}
}
