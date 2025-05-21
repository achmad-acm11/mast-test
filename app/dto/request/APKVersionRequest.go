package request

import (
	"github.com/go-playground/validator/v10"
	"mast-integrator/app/dbo/entity"
	"mast-integrator/app/helper"
	"mime/multipart"
	"strconv"
	"strings"
)

type CreateAPKVersionRequest struct {
	ProjectId   int
	PackageName string
	VersionName string
	AppName     string
	FileName    string
	HashPath    string
	ScanType    string
	Extension   string
}

type OverwriteAPKVersionRequest struct {
	CurrentAPKVersion entity.APKVersion
	ProjectId         int
	PackageName       string
	VersionName       string
	AppName           string
	FileName          string
	HashPath          string
	ScanType          string
	Extension         string
}

type UploadAPKVersionForm struct {
	MobileProjectId int
	APKFile         *multipart.FileHeader
	Overwrite       string
	APKVersionId    string
}

type UploadAPKVersionRequest struct {
	ProjectId    int `json:"project_id"`
	APKFile      *multipart.FileHeader
	Overwrite    bool `json:"overwrite"`
	APKVersionId int  `json:"apk_version_id"`
}

type uploadAPKVersionParam struct {
	requestForm   UploadAPKVersionForm
	request       *UploadAPKVersionRequest
	messagesError map[string]string
	isError       bool
	validate      *validator.Validate
}

func NewUploadAPKVersionParam(requestForm UploadAPKVersionForm, validate *validator.Validate) *uploadAPKVersionParam {
	return &uploadAPKVersionParam{
		requestForm:   requestForm,
		request:       &UploadAPKVersionRequest{},
		isError:       false,
		messagesError: make(map[string]string),
		validate:      validate,
	}
}

func (p *uploadAPKVersionParam) validateProjectId(val int) {
	if err := p.validate.Var(val, "required,numeric"); err != nil {
		msg := helper.GetMessageOneErrorValidator(err, "project_id")
		p.messagesError[strings.ToLower("project_id")] = msg
		p.isError = true
		return
	}
	p.request.ProjectId = val
}

func (p *uploadAPKVersionParam) validateAPKFile(val *multipart.FileHeader) {
	if err := p.validate.Var(val, "required"); err != nil {
		msg := helper.GetMessageOneErrorValidator(err, "apk_file")
		p.messagesError[strings.ToLower("apk_file")] = msg
		p.isError = true
		return
	}
	p.request.APKFile = val
}

func (p *uploadAPKVersionParam) validateOverwrite(val string) {
	if val != "" {
		if err := p.validate.Var(val, "bool"); err != nil {
			msg := helper.GetMessageOneErrorValidator(err, "overwrite")
			p.messagesError[strings.ToLower("overwrite")] = msg
			p.isError = true
			return
		}
		overwrite, _ := strconv.ParseBool(val)
		p.request.Overwrite = overwrite
	} else {
		p.request.Overwrite = false
	}
}

func (p *uploadAPKVersionParam) validateAPKVersionId(val string) {
	if val != "" {
		if err := p.validate.Var(val, "required,numeric"); err != nil {
			msg := helper.GetMessageOneErrorValidator(err, "apk_version_id")
			p.messagesError[strings.ToLower("apk_version_id")] = msg
			p.isError = true
			return
		}
		p.request.APKVersionId, _ = strconv.Atoi(val)
	}
}

func (p *uploadAPKVersionParam) ValidateRequest() (map[string]string, bool) {
	p.validateProjectId(p.requestForm.MobileProjectId)

	if p.isError == true {
		return p.messagesError, p.isError
	}

	p.validateAPKFile(p.requestForm.APKFile)

	if p.isError == true {
		return p.messagesError, p.isError
	}

	p.validateOverwrite(p.requestForm.Overwrite)

	return p.messagesError, p.isError
}

func (p *uploadAPKVersionParam) GetResultRequest() *UploadAPKVersionRequest {
	return p.request
}

type UploadAPK struct {
	Rules            map[string]string
	Message          map[string][]string
	RequestUploadAPK UploadAPKVersionRequest
}
