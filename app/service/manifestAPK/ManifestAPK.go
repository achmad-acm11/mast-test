package manifestAPK

import (
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type ManifestAPKResult struct {
	PackageName string
	Version     string
	AppName     string
	Err         error
}

type ManifestAPK interface {
	GetResult(path string) ManifestAPKResult
}

func GetManifestAPK(osType string, extension string, projectId int) ManifestAPK {
	curDir, _ := os.Getwd()
	uploadDir := curDir + "/tmp"
	filepathZIP := uploadDir + "/" + strconv.Itoa(projectId) + "/"
	os.Mkdir(filepathZIP, 0777)

	if osType == "android" {
		if extension == "xapk" {
			return NewManifestXAPKAndroid(filepathZIP)
		} else {
			return NewManifestAPKAndroid(filepathZIP)
		}
	} else {
		return NewManifestAPKIOS(filepathZIP)
	}
}

func renameToZipAndExtract(path string, filepathZIP string, err error) (*exec.Cmd, error) {
	fileIPA, _ := os.Open(path)
	defer fileIPA.Close()

	filename := strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	fileZIP, _ := os.Create(filepathZIP + filename + ".zip")
	defer fileZIP.Close()
	_, err = io.Copy(fileZIP, fileIPA)

	if err != nil {
		os.RemoveAll(filepathZIP)
		return nil, err
	}

	cmd := exec.Command("unzip", filename+".zip")
	cmd.Stderr = os.Stderr
	cmd.Dir = filepathZIP
	_, err = cmd.Output()
	if err != nil {
		os.RemoveAll(filepathZIP)
		return nil, err
	}
	return cmd, nil
}
