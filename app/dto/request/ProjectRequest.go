package request

type CreateProjectRequest struct {
	Name         string `validate:"required" json:"name"`
	Description  string `json:"description"`
	Os_type      string `validate:"required" json:"os_type"`
	Package_name string `validate:"required" json:"package_name"`
}

type UpdateProjectRequest struct {
	Name         string `validate:"required" json:"name"`
	Description  string `json:"description"`
	Package_name string `validate:"required" json:"package_name"`
}
