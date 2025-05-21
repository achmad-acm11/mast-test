package manifestAPK

import (
	"encoding/json"
	"os"
	"os/exec"
)

type XAPKManifest struct {
	//XapkVersion      string   `json:"xapk_version"`
	PackageName      string   `json:"package_name"`
	Name             string   `json:"name"`
	VersionCode      string   `json:"version_code"`
	VersionName      string   `json:"version_name"`
	MinSdkVersion    string   `json:"min_sdk_version"`
	TargetSdkVersion string   `json:"target_sdk_version"`
	Permissions      []string `json:"permissions"`
	SplitConfigs     []string `json:"split_configs"`
	TotalSize        int      `json:"total_size"`
	Icon             string   `json:"icon"`
	SplitApks        []struct {
		File string `json:"file"`
		Id   string `json:"id"`
	} `json:"split_apks"`
}

type ManifestXAPKAndroid struct {
	filepathZIP string
}

func NewManifestXAPKAndroid(filepathZIP string) *ManifestXAPKAndroid {
	return &ManifestXAPKAndroid{
		filepathZIP: filepathZIP,
	}
}

func (m ManifestXAPKAndroid) GetResult(path string) ManifestAPKResult {
	var err error
	cmd, err := renameToZipAndExtract(path, m.filepathZIP, err)
	result := ManifestAPKResult{}
	if err != nil {
		result.Err = err
		return result
	}

	cmd = exec.Command("cat", "manifest.json")
	cmd.Stderr = os.Stderr
	cmd.Dir = m.filepathZIP
	out, err := cmd.Output()
	if err != nil {
		os.RemoveAll(m.filepathZIP)
		result.Err = err
		return result
	}

	manifestJSON := string(out)

	var manifest XAPKManifest
	err = json.Unmarshal([]byte(manifestJSON), &manifest)
	if err != nil {
		os.RemoveAll(m.filepathZIP)
		result.Err = err
		return result
	}

	result.PackageName = manifest.PackageName
	result.Version = manifest.VersionName
	result.AppName = manifest.Name

	os.RemoveAll(m.filepathZIP)
	return result
}
