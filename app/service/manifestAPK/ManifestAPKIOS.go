package manifestAPK

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

type ManifestAPKIOS struct {
	filepathZIP string
}

func NewManifestAPKIOS(filepathZIP string) *ManifestAPKIOS {
	return &ManifestAPKIOS{
		filepathZIP: filepathZIP,
	}
}

func (m ManifestAPKIOS) GetResult(path string) ManifestAPKResult {
	var err error
	cmd, err := renameToZipAndExtract(path, m.filepathZIP, err)
	result := ManifestAPKResult{}
	if err != nil {
		result.Err = err
		return result
	}

	cmd = exec.Command("ls")
	cmd.Stderr = os.Stderr
	cmd.Dir = m.filepathZIP + "/Payload"
	out, err := cmd.Output()
	if err != nil {
		os.RemoveAll(m.filepathZIP)
		result.Err = err
		return result
	}

	pathAPP := strings.Split(string(out), " ")[0]
	pathAPP = strings.Replace(pathAPP, " ", "", -1)
	pathAPP = strings.Replace(pathAPP, "\n", "", -1)

	cmd = exec.Command("cat", pathAPP+"/Info.plist")
	cmd.Stderr = os.Stderr
	cmd.Dir = m.filepathZIP + "/Payload"
	out, err = cmd.Output()
	if err != nil {
		os.RemoveAll(m.filepathZIP)
		result.Err = err
		return result
	}

	if !strings.Contains(string(out), "<!DOCTYPE plist PUBLIC") {
		cmd = exec.Command("plistutil", "-i", "Info.plist")
		cmd.Stderr = os.Stderr
		cmd.Dir = m.filepathZIP + "/Payload/" + pathAPP
		out, err = cmd.Output()
		if err != nil {
			os.RemoveAll(m.filepathZIP)
			result.Err = err
			return result
		}
	}
	isError := m.checkPackageName(out, &result)
	if isError {
		result.Err = errors.New("package name not found in apk")
		return result
	}

	isError = m.checkVersionName(out, &result)
	if isError {
		result.Err = errors.New("Version name not found in apk")
		return result
	}

	isError = m.checkAppName(out, &result)
	if isError {
		result.Err = errors.New("App name not found in apk")
		return result
	}

	return result
}

func (m ManifestAPKIOS) checkPackageName(out []byte, result *ManifestAPKResult) bool {
	if strings.Contains(string(out), "CFBundleIdentifier") {
		getString := strings.Split(string(out), "<key>CFBundleIdentifier</key>")
		if len(getString) > 1 {
			getStringEnd := strings.Split(getString[1], "</string>")
			if len(getStringEnd) > 1 {
				result.PackageName = getStringEnd[0]
				result.PackageName = strings.Replace(result.PackageName, "<string>", "", -1)
				result.PackageName = strings.Replace(result.PackageName, "</key>", "", -1)
				result.PackageName = strings.Replace(result.PackageName, "\n", "", -1)
				result.PackageName = strings.Replace(result.PackageName, "\t", "", -1)
				return false
			} else {
				return true
			}
		} else {
			return true
		}
	} else {
		return true
	}
}

func (m ManifestAPKIOS) checkVersionName(out []byte, result *ManifestAPKResult) bool {
	if strings.Contains(string(out), "CFBundleShortVersionString") {
		getString := strings.Split(string(out), "CFBundleShortVersionString")
		if len(getString) > 1 {
			getStringEnd := strings.Split(getString[1], "</string>")
			if len(getStringEnd) > 1 {
				result.Version = getStringEnd[0]
				result.Version = strings.Replace(result.Version, "<string>", "", -1)
				result.Version = strings.Replace(result.Version, "</key>", "", -1)
				result.Version = strings.Replace(result.Version, "\n", "", -1)
				result.Version = strings.Replace(result.Version, "\t", "", -1)
				return false
			} else {
				return true
			}
		} else {
			return true
		}
	} else {
		return true
	}
}

func (m ManifestAPKIOS) checkAppName(out []byte, result *ManifestAPKResult) bool {
	if strings.Contains(string(out), "CFBundleName") {
		getString := strings.Split(string(out), "CFBundleName")
		if len(getString) > 1 {
			getStringEnd := strings.Split(getString[1], "</string>")
			if len(getStringEnd) > 1 {
				result.AppName = getStringEnd[0]
				result.AppName = strings.Replace(result.AppName, "<string>", "", -1)
				result.AppName = strings.Replace(result.AppName, "</key>", "", -1)
				result.AppName = strings.Replace(result.AppName, "\n", "", -1)
				result.AppName = strings.Replace(result.AppName, "\t", "", -1)
				return false
			} else {
				return true
			}
		} else {
			return true
		}
	} else {
		return true
	}
}
