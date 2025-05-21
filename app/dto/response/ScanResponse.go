package response

import "mast-integrator/app/dbo/entity"

type ScanResponse struct {
	Id           int    `json:"id"`
	APKVersionId int    `json:"apk_version_id"`
	ScanVersion  int    `json:"scan_version"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type ScanResponseBuilder struct {
	singleData ScanResponse
	listData   []ScanResponse
}

func NewScanResponseBuilder() *ScanResponseBuilder {
	return &ScanResponseBuilder{}
}

func (builder *ScanResponseBuilder) Default(scan entity.Scan) *ScanResponseBuilder {
	builder.singleData = mapScan(scan)
	return builder
}

func (builder *ScanResponseBuilder) List(scans []entity.Scan) *ScanResponseBuilder {
	builder.listData = mapListScan(scans)
	return builder
}

func (builder *ScanResponseBuilder) Result() ScanResponse {
	return builder.singleData
}

func (builder *ScanResponseBuilder) ListResult() []ScanResponse {
	return builder.listData
}

func mapScan(scan entity.Scan) ScanResponse {
	response := ScanResponse{
		Id:           scan.Id,
		APKVersionId: scan.APKVersionId,
		CreatedAt:    scan.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    scan.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return response
}

func mapListScan(scanList []entity.Scan) []ScanResponse {
	responses := []ScanResponse{}
	for _, scan := range scanList {
		responses = append(responses, mapScan(scan))
	}
	return responses
}
