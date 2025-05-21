package staticScan

import (
	"github.com/gin-gonic/gin"
	"mast-integrator/app/dbo/entity"
	"mast-integrator/app/dto/response"
	"os"
	"time"
)

type ScanIOSStatic struct {
	abstract StaticScanAbstract
}

func NewScanIOSStatic() *ScanIOSStatic {
	return &ScanIOSStatic{}
}

func (s *ScanIOSStatic) ScanProcess(ctx *gin.Context, param StaticScanProcessParam) {
	apkVersion := param.ApkVersion
	scan := param.Scan

	defer func() {
		if r := recover(); r != nil {
			s.abstract.processError(ctx, apkVersion, r.(error).Error())
			return
		}
	}()
	resultRaw := s.abstract.mobsfApi.ScanStatic(apkVersion.Filename, apkVersion.HashPath, "ipa")
	//if err != nil {
	//	processError(apkVersion, global, err.Error())
	//	return nil, nil, nil
	//}

	//download icon
	folderPath := "apk_icon/"
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.Mkdir(folderPath, 0755)
	}
	filePathResponse := s.abstract.mobsfApi.DownloadIcon(apkVersion.Id, apkVersion.HashPath)
	//if Err != nil {
	//	processError(apkVersion, global, Err.Error())
	//	return nil, nil, nil
	//}

	//save scan result
	var dataResult entity.ResultData
	StaticAnalysisResultRaw := resultRaw.IOS
	StaticAnalysisResult := RawResultToModelIOSStaticAnalysisResult(StaticAnalysisResultRaw)
	StaticAnalysisResult.ScanVersion = scan.ScanVersion
	StaticAnalysisResult.ScanId = scan.Id
	StaticAnalysisResult.TypeResult = "ios"
	dataResult.IOS = StaticAnalysisResult
	//global.Repositories.StaticAnalysisResultRepository.Create(DataResult, 1)

	//update scan status
	apkVersion.StaticScanStatus = 3
	apkVersion.StaticStatusMessage = ""
	//READ!!! remove this after dynamic already implemented
	apkVersion.DynamicScanStatus = 3
	apkVersion.DynamicStatusMessage = ""
	//end of READ!!!

	//update apk version
	apkVersion.Build = StaticAnalysisResultRaw.Build
	apkVersion.SdkName = StaticAnalysisResultRaw.SdkName
	apkVersion.Platform = StaticAnalysisResultRaw.Platform
	apkVersion.MinOsVersion = StaticAnalysisResultRaw.MinOsVersion
	apkVersion.BinaryInfoEndian = StaticAnalysisResultRaw.BinaryInfo.Endian
	apkVersion.BinaryInfoBit = StaticAnalysisResultRaw.BinaryInfo.Bit
	apkVersion.BinaryInfoArch = StaticAnalysisResultRaw.BinaryInfo.Arch
	apkVersion.BinaryInfoSubarch = StaticAnalysisResultRaw.BinaryInfo.Subarch

	apkVersion.Size = StaticAnalysisResultRaw.Size
	apkVersion.FileMD5 = StaticAnalysisResultRaw.Md5
	apkVersion.FileSHA5 = StaticAnalysisResultRaw.Sha1
	apkVersion.FileSHA256 = StaticAnalysisResultRaw.Sha256
	apkVersion.IconPath = filePathResponse

	scan.SecurityScore = StaticAnalysisResultRaw.Appsec.SecurityScore
	scan.SecurityScore = StaticAnalysisResultRaw.Appsec.SecurityScore
	if StaticAnalysisResultRaw.Appsec.SecurityScore > 95 {
		scan.RiskRatingGrade = "A"
	} else if StaticAnalysisResultRaw.Appsec.SecurityScore > 80 {
		scan.RiskRatingGrade = "B"
	} else if StaticAnalysisResultRaw.Appsec.SecurityScore > 65 {
		scan.RiskRatingGrade = "C"
	} else if StaticAnalysisResultRaw.Appsec.SecurityScore > 50 {
		scan.RiskRatingGrade = "D"
	} else {
		scan.RiskRatingGrade = "E"
	}
	scan.HighSeverityDistributionCount = len(StaticAnalysisResultRaw.Appsec.High)
	scan.WarningSeverityDistributionCount = len(StaticAnalysisResultRaw.Appsec.Warning)
	scan.InfoSeverityDistributionCount = len(StaticAnalysisResultRaw.Appsec.Info)
	scan.SecureSeverityDistributionCount = len(StaticAnalysisResultRaw.Appsec.Secure)
	scan.HotspotSeverityDistributionCount = len(StaticAnalysisResultRaw.Appsec.Hotspot)
	scan.FinishedAt = time.Now()

	s.abstract.repoApkVersion.Update(ctx, s.abstract.db, apkVersion)
	s.abstract.repo.Update(ctx, s.abstract.db, scan)
	//global.Repositories.StaticAnalysisResultRepository.Create(DataResult, 2)
	//logger.SetLogConsole(logger.LogData{
	//	Message: "Scan Static Analyzer APK Version finished",
	//	CustomFields: logrus.Fields{
	//		"project_id":     apkVersion.MobileProjectId,
	//		"apk_version_id": apkVersion.ID,
	//		"package_name":   apkVersion.PackageName,
	//		"version_name":   apkVersion.VersionName,
	//	},
	//})
}

func RawResultToModelIOSStaticAnalysisResult(rawResult response.IOSStaticAnalysisResultRaw) entity.IOSStaticAnalysisResult {

	var Permissions []entity.Permission
	for permission, data := range rawResult.Permissions {
		Permissions = append(Permissions, entity.Permission{
			Permission:  permission,
			Status:      data.Status,
			Info:        data.Info,
			Description: data.Description,
		})
	}

	var IosAPI []entity.AndroidApi
	for api, data := range rawResult.IosApi {
		var Files []entity.Files
		for filename, value := range data.Files {
			Files = append(Files, entity.Files{
				FileName: filename,
				Value:    value,
			})
		}
		IosAPI = append(IosAPI, entity.AndroidApi{
			ApiType: api,
			Files:   Files,
			Metadata: struct {
				Description string `json:"description" bson:"description"`
				Severity    string `json:"severity" bson:"severity"`
			}(data.Metadata),
		})
	}

	var IOSBinaryAnalysis []entity.IOSBinaryAnalysis
	for IOSBinaryAnalysisType, data := range rawResult.BinaryAnalysis {
		IOSBinaryAnalysis = append(IOSBinaryAnalysis, entity.IOSBinaryAnalysis{
			IosBinaryAnalysisType: IOSBinaryAnalysisType,
			DetailedDesc:          data.DetailedDesc,
			Severity:              data.Severity,
			Cvss:                  float64(data.Cvss),
			Cwe:                   data.Cwe,
			OwaspMobile:           data.OwaspMobile,
			Masvs:                 data.Masvs,
		})

	}

	var CodeAnalysis []entity.CodeAnalysis
	for CodeAnalysisType, data := range rawResult.CodeAnalysis {

		var Files []entity.Files
		for filename, value := range data.Files {
			Files = append(Files, entity.Files{
				FileName: filename,
				Value:    value,
			})
		}
		CodeAnalysis = append(CodeAnalysis, entity.CodeAnalysis{
			CodeAnalysisType: CodeAnalysisType,
			Files:            Files,
			Metadata: struct {
				Cvss        float64 `json:"cvss" bson:"cvss"`
				Cwe         string  `json:"cwe" bson:"cwe"`
				OwaspMobile string  `json:"owasp-mobile" bson:"owasp-mobile"`
				Masvs       string  `json:"masvs" bson:"masvs"`
				Ref         string  `json:"ref" bson:"ref"`
				Description string  `json:"description" bson:"description"`
				Severity    string  `json:"severity" bson:"severity"`
			}(data.Metadata),
		})
	}

	var Domains []entity.Domains
	for domain, data := range rawResult.Domains {
		Domains = append(Domains, entity.Domains{
			DomainType: domain,
			Bad:        data.Bad,
			Geolocation: struct {
				Ip           string `json:"ip" bson:"ip"`
				CountryShort string `json:"country_short" bson:"country_short"`
				CountryLong  string `json:"country_long" bson:"country_long"`
				Region       string `json:"region" bson:"region"`
				City         string `json:"city" bson:"city"`
				Latitude     string `json:"latitude" bson:"latitude"`
				Longitude    string `json:"longitude" bson:"longitude"`
			}(data.Geolocation),
		})
	}

	var firebaseUrls []entity.FirebaseUrl
	for _, data := range rawResult.FirebaseUrls {
		firebaseUrls = append(firebaseUrls, entity.FirebaseUrl{
			Url:  data.Url,
			Open: data.Open,
		})
	}

	//var NiapAnalysis []entity.NiapAnalysis
	//for niapType, data := range rawResult.NiapAnalysis {
	//	NiapAnalysis = append(NiapAnalysis, entity.NiapAnalysis{
	//		NiapAnalysisType: niapType,
	//		Choice:           data.Choice,
	//		Description:      data.Description,
	//		Class:            data.Class,
	//	})
	//}

	var BundleUrlTypes []entity.BundleUrlType
	for _, data := range rawResult.BundleUrlTypes {
		BundleUrlTypes = append(BundleUrlTypes, entity.BundleUrlType{
			CFBundleTypeRole:   data.CFBundleTypeRole,
			CFBundleURLName:    data.CFBundleURLName,
			CFBundleURLSchemes: data.CFBundleURLSchemes,
		})
	}
	StaticAnalysisResult := entity.IOSStaticAnalysisResult{
		BundleUrlTypes:           BundleUrlTypes,
		BundleSupportedPlatforms: rawResult.BundleSupportedPlatforms,
		InfoPlist:                rawResult.InfoPlist,
		Permissions:              Permissions,
		AtsAnalysis:              rawResult.AtsAnalysis,
		BinaryAnalysis:           IOSBinaryAnalysis,
		MachoAnalysis:            rawResult.MachoAnalysis,
		IosApi:                   IosAPI,
		CodeAnalysis:             CodeAnalysis,
		FileAnalysis: []struct {
			Issue string `json:"issue" bson:"issue"`
			Files []struct {
				FilePath string `json:"file_path" bson:"file_path"`
				Type     string `json:"type" bson:"type"`
				Hash     string `json:"hash" bson:"hash"`
			} `json:"files" bson:"files"`
		}(rawResult.FileAnalysis),
		Libraries: rawResult.Libraries,
		Files:     rawResult.Files,
		Urls: []struct {
			Urls []string `json:"urls" bson:"urls"`
			Path string   `json:"path" bson:"path"`
		}(rawResult.Urls),
		Domains: Domains,
		Emails: []struct {
			Emails []string `json:"emails" bson:"emails"`
			Path   string   `json:"path" bson:"path"`
		}(rawResult.Emails),
		Strings:      rawResult.Strings,
		FirebaseUrls: firebaseUrls,
		AppstoreDetails: struct {
			Features         []string `json:"features"`
			Icon             string   `json:"icon"`
			DeveloperId      int      `json:"developer_id"`
			Developer        string   `json:"developer"`
			DeveloperUrl     string   `json:"developer_url"`
			DeveloperWebsite string   `json:"developer_website"`
			SupportedDevices []string `json:"supported_devices"`
			Title            string   `json:"title"`
			AppId            string   `json:"app_id"`
			Category         []string `json:"category"`
			Description      string   `json:"description"`
			Price            float64  `json:"price"`
			ItunesUrl        string   `json:"itunes_url"`
			Score            float64  `json:"score"`
			Error            bool     `json:"error" bson:"error"`
		}(rawResult.AppstoreDetails),
		Secrets: rawResult.Secrets,
		Trackers: struct {
			DetectedTrackers int `json:"detected_trackers" bson:"detected_trackers"`
			TotalTrackers    int `json:"total_trackers" bson:"total_trackers"`
			Trackers         []struct {
				Name       string `json:"name" bson:"name"`
				Categories string `json:"categories" bson:"categories"`
				Url        string `json:"url" bson:"url"`
			} `json:"trackers" bson:"trackers"`
		}(rawResult.Trackers),
		Appsec: struct {
			High []struct {
				Title       string `json:"title" bson:"title"`
				Description string `json:"description" bson:"description"`
				Section     string `json:"section" bson:"section"`
			} `json:"high" bson:"high"`
			Warning []struct {
				Title       string `json:"title" bson:"title"`
				Description string `json:"description" bson:"description"`
				Section     string `json:"section" bson:"section"`
			} `json:"warning" bson:"warning"`
			Info []struct {
				Title       string `json:"title" bson:"title"`
				Description string `json:"description" bson:"description"`
				Section     string `json:"section" bson:"section"`
			} `json:"info" bson:"info"`
			Secure []struct {
				Title       string `json:"title" bson:"title"`
				Description string `json:"description" bson:"description"`
				Section     string `json:"section" bson:"section"`
			} `json:"secure" bson:"secure"`
			Hotspot []struct {
				Title       string `json:"title" bson:"title"`
				Description string `json:"description" bson:"description"`
				Section     string `json:"section" bson:"section"`
			} `json:"hotspot" bson:"hotspot"`
			TotalTrackers int    `json:"total_trackers" bson:"total_trackers"`
			Trackers      int    `json:"trackers" bson:"trackers"`
			SecurityScore int    `json:"security_score" bson:"security_score"`
			AppName       string `json:"app_name" bson:"app_name"`
			FileName      string `json:"file_name" bson:"file_name"`
			Hash          string `json:"hash" bson:"hash"`
			VersionName   string `json:"version_name" bson:"version_name"`
		}(rawResult.Appsec),
		AverageCvss: rawResult.AverageCvss,
		VirusTotal:  rawResult.VirusTotal,
		BaseUrl:     rawResult.BaseUrl,
		DwdDir:      rawResult.DwdDir,
		HostOs:      rawResult.HostOs,
		Timestamp:   rawResult.Timestamp,
	}
	return StaticAnalysisResult

}
