package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"io/ioutil"
	"mast-integrator/app/dto/response"
	"mast-integrator/app/exception"
	"mast-integrator/app/helper"
	"mast-integrator/app/shareVar"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type MobSFAPI interface {
	Upload(filepath string) *response.UploadMobSFResponse
	ScanStatic(filename string, hash string, scanType string) *response.ScanStaticResultResponse
	DownloadIcon(apkVersionId int, hashPath string) string
}

type MobSFAPIImpl struct {
	url    string
	key    string
	stdLog *helper.StandartLog
}

func NewMobSFAPI() *MobSFAPIImpl {
	if os.Getenv("APP_ENV") == "" {
		errEnv := godotenv.Load(".env")
		helper.ErrorHandler(errEnv)
	}

	urlMobSFApi := os.Getenv("MOBSF_API_URL")
	MobSFApiKey := os.Getenv("MOBSF_API_KEY")

	return &MobSFAPIImpl{
		url:    urlMobSFApi,
		key:    MobSFApiKey,
		stdLog: helper.NewStandardLog(shareVar.MobSFAPI, shareVar.Service),
	}
}

func (m MobSFAPIImpl) Upload(filePath string) *response.UploadMobSFResponse {
	var responseApi *response.UploadMobSFResponse

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	httpReq, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/upload", m.url), body)
	helper.ErrorHandler(err)
	httpReq.Header.Add("content-type", writer.FormDataContentType())
	httpReq.Header.Add("Authorization", m.key)

	client := &http.Client{}
	response, err := client.Do(httpReq)
	helper.ErrorHandler(err)

	bodyByte, err := ioutil.ReadAll(response.Body)
	helper.ErrorHandler(err)

	if response.StatusCode == http.StatusBadRequest {
		m.logResponseNotOK(response, bodyByte)
		return responseApi
	}
	if response.StatusCode == http.StatusInternalServerError {
		m.logResponseNotOK(response, bodyByte)
		return responseApi
	}
	if response.StatusCode == http.StatusUnauthorized {
		m.logResponseNotOK(response, bodyByte)
		return responseApi
	}

	fmt.Printf("%+v\n", string(bodyByte))

	json.Unmarshal(bodyByte, &responseApi)

	// End log
	//m.stdLog.NameFunc = "Create"
	//m.stdLog.EndFunction(nil)

	return responseApi
}

func (m MobSFAPIImpl) ScanStatic(filename string, hash string, scanType string) *response.ScanStaticResultResponse {
	apkResponseRaw := &response.StaticAnalysisResultRaw{}

	fmt.Printf("Scan Type: %v, FileName: %v\n", scanType, filename)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("scan_type", scanType)
	writer.WriteField("file_name", filename)
	writer.WriteField("hash", hash)
	writer.WriteField("re_scan", "1")
	writer.Close()

	httpReq, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/scan", m.url), body)
	helper.ErrorHandler(err)
	httpReq.Header.Add("content-type", writer.FormDataContentType())
	httpReq.Header.Add("Authorization", m.key)

	client := &http.Client{}
	res, err := client.Do(httpReq)
	helper.ErrorHandler(err)

	bodyByte, err := ioutil.ReadAll(res.Body)
	helper.ErrorHandler(err)
	fmt.Printf("Response: %v\n", res.StatusCode)
	m.swicherResponseStatusAndPanic(res, bodyByte)

	fmt.Printf("%+v\n", string(bodyByte))

	json.Unmarshal(bodyByte, &apkResponseRaw)

	return &response.ScanStaticResultResponse{
		APK: *apkResponseRaw,
	}
}

func (m MobSFAPIImpl) DownloadIcon(apkVersionId int, hashPath string) string {
	var filePathResponse string
	httpReq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/download/%s-icon.png", m.url, hashPath), nil)
	helper.ErrorHandler(err)
	httpReq.Header.Add("Authorization", m.key)

	client := &http.Client{}
	response, err := client.Do(httpReq)
	helper.ErrorHandler(err)

	if response.StatusCode == 200 || response.StatusCode == 201 {
		filePathResponse = "apk_icon/" + strconv.Itoa(apkVersionId) + "_" + hashPath + "-icon.png"

		out, err := os.Create(filePathResponse)
		if err != nil {
			return ""
		}
		defer out.Close()

		file, err := io.Copy(out, response.Body)
		if err != nil {
			return ""
		}

		if file == 0 {
			os.Remove(filePathResponse)
			return ""
		} else {
			return filePathResponse
		}
	}
	return ""
}

func (m MobSFAPIImpl) logResponseNotOK(response *http.Response, bodyByte []byte) {
	res := make(map[string]interface{})

	res["status_code"] = response.StatusCode
	res["body"] = string(bodyByte)
	fmt.Printf("response: %s, status code: %s", string(bodyByte), res["status_code"])
	m.stdLog.WarningFunction(res)
}

func (m MobSFAPIImpl) swicherResponseStatusAndPanic(response *http.Response, bodyByte []byte) {
	if response.StatusCode == http.StatusBadRequest {
		m.logResponseNotOK(response, bodyByte)
		panic(exception.NewInternalServerError(errors.New("Bad Request").Error()))
	}
	if response.StatusCode == http.StatusInternalServerError {
		m.logResponseNotOK(response, bodyByte)
		panic(exception.NewInternalServerError(errors.New("Internal Server Error").Error()))
	}
	if response.StatusCode == http.StatusUnauthorized {
		m.logResponseNotOK(response, bodyByte)
		panic(exception.NewInternalServerError(errors.New("Unauthorized").Error()))
	}
}
