package response

import "mast-integrator/app/dbo/entity"

type APKVersionResponse struct {
	Id        int `json:"id"`
	ProjectId int `json:"project_id"`

	AppName     string `json:"app_name"`
	AppType     string `json:"app_type"`
	PackageName string `json:"package_name"`
	VersionName string `json:"version_name"`
	HashPath    string `json:"hash_path"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type APKVersionResponseBuilder struct {
	singleData APKVersionResponse
	listData   []APKVersionResponse
}

func NewAPKVersionResponseBuilder() *APKVersionResponseBuilder {
	return &APKVersionResponseBuilder{}
}

func (builder *APKVersionResponseBuilder) Default(apkVersion entity.APKVersion) *APKVersionResponseBuilder {
	builder.singleData = mapApkVersion(apkVersion)
	return builder
}

func (builder *APKVersionResponseBuilder) List(apkVersions []entity.APKVersion) *APKVersionResponseBuilder {
	builder.listData = mapListApkVersion(apkVersions)
	return builder
}

func (builder *APKVersionResponseBuilder) Result() APKVersionResponse {
	return builder.singleData
}

func (builder *APKVersionResponseBuilder) ListResult() []APKVersionResponse {
	return builder.listData
}

func mapApkVersion(apkVersion entity.APKVersion) APKVersionResponse {
	response := APKVersionResponse{
		Id:          apkVersion.Id,
		ProjectId:   apkVersion.ProjectId,
		AppName:     apkVersion.AppName,
		AppType:     apkVersion.AppType,
		PackageName: apkVersion.PackageName,
		VersionName: apkVersion.VersionName,
		HashPath:    apkVersion.HashPath,
		CreatedAt:   apkVersion.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   apkVersion.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return response
}

func mapListApkVersion(apkVersions []entity.APKVersion) []APKVersionResponse {
	responses := []APKVersionResponse{}
	for _, apkVersion := range apkVersions {
		responses = append(responses, mapApkVersion(apkVersion))
	}
	return responses
}
