package request

type ScanTriggerRequest struct {
	APKVersionId int    `json:"apk_version_id" validate:"required,numeric,gt=0"`
	ScanType     string `json:"scan_type" validate:"required,oneof=static dynamic all"`
}
