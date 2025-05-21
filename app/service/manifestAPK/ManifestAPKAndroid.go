package manifestAPK

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

type ManifestAPKAndroid struct {
	filepathZIP string
}

func NewManifestAPKAndroid(filepathZIP string) *ManifestAPKAndroid {
	return &ManifestAPKAndroid{
		filepathZIP: filepathZIP,
	}
}

func (m ManifestAPKAndroid) GetResult(path string) ManifestAPKResult {
	result := ManifestAPKResult{}
	cmd := exec.Command("aapt", "dump", "badging", path)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		result.Err = err
		return result
	}

	isError := m.checkPackageName(out, &result)
	if isError {
		result.Err = errors.New("package name not found in apk")
		return result
	}

	isError = m.checkVersionName(out, &result)
	if isError {
		result.Err = errors.New("version not found in apk")
		return result
	}

	isError = m.checkAppName(out, &result)
	if isError {
		result.Err = errors.New("app name label not found in apk")
		return result
	}

	return result
}

func (m ManifestAPKAndroid) checkPackageName(out []byte, result *ManifestAPKResult) bool {
	if strings.Contains(string(out), "package: name=") {
		result.PackageName = strings.Split(strings.Split(string(out), "package: name=")[1], " ")[0]
		result.PackageName = strings.Replace(result.PackageName, "'", "", -1)
		result.PackageName = strings.Split(result.PackageName, "\n")[0]
		return false
	} else {
		return true
	}
}

func (m ManifestAPKAndroid) checkVersionName(out []byte, result *ManifestAPKResult) bool {
	if strings.Contains(string(out), "versionName=") {
		result.Version = strings.Split(strings.Split(string(out), "versionName=")[1], " ")[0]
		result.Version = strings.Replace(result.Version, "'", "", -1)
		result.Version = strings.Split(result.Version, "\n")[0]
		return false
	} else {
		return true
	}
}

func (m ManifestAPKAndroid) checkAppName(out []byte, result *ManifestAPKResult) bool {
	if strings.Contains(string(out), "application: label=") {
		result.AppName = strings.Split(strings.Split(string(out), "application: label=")[1], " ")[0]
		result.AppName = strings.Replace(result.AppName, "'", "", -1)
		result.AppName = strings.Split(result.AppName, "\n")[0]
		return false
	} else {
		return true
	}
}
