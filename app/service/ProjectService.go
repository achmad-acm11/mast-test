package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"mast-integrator/app/dbo/entity"
	"mast-integrator/app/dbo/repository"
	"mast-integrator/app/dto/request"
	"mast-integrator/app/dto/response"
	"mast-integrator/app/exception"
	"mast-integrator/app/helper"
	"mast-integrator/app/shareVar"
)

type ProjectService interface {
	GetAllData(ctx *gin.Context) []response.ProjectResponse
	GetDetailByIdData(ctx *gin.Context, id int) response.ProjectResponse
	//ScanProjectData(ctx *gin.Context, request request.ProjectScanRequest)
	CreateProjectData(ctx *gin.Context, req request.CreateProjectRequest) response.ProjectResponse
	UpdateProjectData(ctx *gin.Context, request request.UpdateProjectRequest, id int) response.ProjectResponse
	DeleteProjectData(ctx *gin.Context, id int)
}

type ProjectServiceImpl struct {
	repo      repository.ProjectRepository
	validator *validator.Validate
	db        *gorm.DB
	stdLog    *helper.StandartLog
}

func NewProjectService(repo repository.ProjectRepository, validate *validator.Validate, db *gorm.DB) *ProjectServiceImpl {
	return &ProjectServiceImpl{
		repo:      repo,
		validator: validate,
		db:        db,
		stdLog:    helper.NewStandardLog(shareVar.Project, shareVar.Service),
	}
}

func (p ProjectServiceImpl) GetAllData(ctx *gin.Context) []response.ProjectResponse {
	projects := p.repo.GetAll(ctx, p.db)

	return response.NewProjectResponseBuilder().List(projects).ListResult()
}

func (p ProjectServiceImpl) GetDetailByIdData(ctx *gin.Context, id int) response.ProjectResponse {
	project := p.getProjectWithException(ctx, id)

	return response.NewProjectResponseBuilder().Default(project).Result()
}

func (p ProjectServiceImpl) CreateProjectData(ctx *gin.Context, req request.CreateProjectRequest) response.ProjectResponse {
	p.stdLog.NameFunc = "CreateProjectData"
	p.stdLog.StartFunction(req)

	err := p.validator.Struct(req)
	helper.ErrorHandlerValidator(err)

	tx := p.db.Begin()
	defer helper.CommitOrRollback(tx)

	project := p.repo.Create(ctx, tx, entity.Project{
		Key:         fmt.Sprintf("1_%s_%d", helper.ToSnakeCase(req.Name), helper.Generate4DigitCode()),
		Name:        req.Name,
		Description: req.Description,
		OsType:      req.Os_type,
		PackageName: req.Package_name,
	})

	p.stdLog.NameFunc = "CreateProjectData"
	p.stdLog.EndFunction(project)

	return response.NewProjectResponseBuilder().Default(project).Result()
}

func (p ProjectServiceImpl) UpdateProjectData(ctx *gin.Context, request request.UpdateProjectRequest, id int) response.ProjectResponse {
	err := p.validator.Struct(request)
	helper.ErrorHandlerValidator(err)

	tx := p.db.Begin()
	defer helper.CommitOrRollback(tx)

	project := p.getProjectWithException(ctx, id)

	project.Name = request.Name
	project.Description = request.Description
	project.PackageName = request.Package_name

	projectNew := p.repo.Update(ctx, tx, project)

	return response.NewProjectResponseBuilder().Default(projectNew).Result()
}

func (p ProjectServiceImpl) DeleteProjectData(ctx *gin.Context, id int) {
	project := p.getProjectWithException(ctx, id)

	tx := p.db.Begin()
	defer helper.CommitOrRollback(tx)
	p.repo.DeleteOne(ctx, tx, project)
}

func (p ProjectServiceImpl) getProjectWithException(ctx *gin.Context, id int) entity.Project {
	project := p.repo.GetOneById(ctx, p.db, id)
	if project.Id == 0 {
		panic(exception.NewNotFoundError(errors.New(shareVar.PROJECT_NOT_FOUND).Error()))
	}

	return project
}
