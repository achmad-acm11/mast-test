package response

import "mast-integrator/app/dbo/entity"

type ProjectResponse struct {
	Id                 int    `json:"id"`
	Key                string `json:"key"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	OsType             string `json:"os_type"`
	PackageName        string `json:"package_name"`
	CurrentScanVersion int    `json:"current_scan_version"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

type ProjectResponseBuilder struct {
	singleData ProjectResponse
	listData   []ProjectResponse
}

func NewProjectResponseBuilder() *ProjectResponseBuilder {
	return &ProjectResponseBuilder{}
}

func (builder *ProjectResponseBuilder) Default(project entity.Project) *ProjectResponseBuilder {
	builder.singleData = mapProject(project)
	return builder
}

func (builder *ProjectResponseBuilder) List(projects []entity.Project) *ProjectResponseBuilder {
	builder.listData = mapListProject(projects)
	return builder
}

func (builder *ProjectResponseBuilder) Result() ProjectResponse {
	return builder.singleData
}

func (builder *ProjectResponseBuilder) ListResult() []ProjectResponse {
	return builder.listData
}

func mapProject(project entity.Project) ProjectResponse {
	response := ProjectResponse{
		Id:                 project.Id,
		Key:                project.Key,
		Name:               project.Name,
		Description:        project.Description,
		OsType:             project.OsType,
		PackageName:        project.PackageName,
		CurrentScanVersion: project.CurrentScanVersion,
		CreatedAt:          project.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:          project.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return response
}

func mapListProject(projectList []entity.Project) []ProjectResponse {
	responses := []ProjectResponse{}
	for _, project := range projectList {
		responses = append(responses, mapProject(project))
	}
	return responses
}
