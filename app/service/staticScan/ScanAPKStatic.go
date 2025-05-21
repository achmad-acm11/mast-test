package staticScan

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mast-integrator/app/dbo/entity"
	"mast-integrator/app/dto/response"
	"os"
	"time"
)

type ScanAPKStatic struct {
	abstract StaticScanAbstract
}

func NewScanAPKStatic(abstract StaticScanAbstract) *ScanAPKStatic {
	return &ScanAPKStatic{
		abstract: abstract,
	}
}

func (s *ScanAPKStatic) ScanProcess(ctx *gin.Context, param StaticScanProcessParam) {
	//fmt.Printf("%+v\n", param)

	apkVersion := param.ApkVersion
	scan := param.Scan

	defer func() {
		if r := recover(); r != nil {
			s.abstract.processError(ctx, apkVersion, r.(error).Error())
			return
		}
	}()
	//fmt.Printf("Sini")
	resultRaw := s.abstract.mobsfApi.ScanStatic(apkVersion.Filename, apkVersion.HashPath, apkVersion.Extension)

	//fmt.Printf("%+v\n", resultRaw)

	folderPath := "apk_icon/"
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.Mkdir(folderPath, 0755)
	}
	filePathResponse := s.abstract.mobsfApi.DownloadIcon(apkVersion.Id, apkVersion.HashPath)

	fmt.Printf("%+v\n", filePathResponse)

	var dataResult entity.ResultData
	StaticAnalysisResultRaw := resultRaw.APK

	StaticAnalysisResult := RawResultToModelStaticAnalysisResult(StaticAnalysisResultRaw)
	StaticAnalysisResult.ScanVersion = scan.ScanVersion
	StaticAnalysisResult.ScanId = scan.Id
	StaticAnalysisResult.TypeResult = "android"
	dataResult.APK = StaticAnalysisResult

	//fmt.Printf("%+v\n", StaticAnalysisResult)
	//global.Repositories.StaticAnalysisResultRepository.Create(DataResult, 1)

	//update scan status
	apkVersion.StaticScanStatus = 3
	apkVersion.StaticStatusMessage = ""
	//READ!!! remove this after dynamic already implemented
	apkVersion.DynamicScanStatus = 3
	apkVersion.DynamicStatusMessage = ""
	//end of READ!!!

	//update apk version
	apkVersion.MainActivity = StaticAnalysisResultRaw.MainActivity
	apkVersion.TargetSdk = StaticAnalysisResultRaw.TargetSdk
	apkVersion.MinSdk = StaticAnalysisResultRaw.MinSdk
	apkVersion.Size = StaticAnalysisResultRaw.Size
	apkVersion.FileMD5 = StaticAnalysisResultRaw.Md5
	apkVersion.FileSHA5 = StaticAnalysisResultRaw.Sha1
	apkVersion.FileSHA256 = StaticAnalysisResultRaw.Sha256
	apkVersion.IconPath = filePathResponse

	//update scan
	scan.SecurityScore = StaticAnalysisResultRaw.Appsec.SecurityScore
	scan.ActivitiesCount = len(StaticAnalysisResultRaw.Activities)
	scan.ServicesCount = len(StaticAnalysisResultRaw.Services)
	scan.RecieversCount = len(StaticAnalysisResultRaw.Receivers)
	scan.ProvidersCount = len(StaticAnalysisResultRaw.Providers)
	scan.ExportedActivityCount = StaticAnalysisResultRaw.ExportedCount.ExportedActivities
	scan.ExportedServiceCount = StaticAnalysisResultRaw.ExportedCount.ExportedServices
	scan.ExportedReceiverCount = StaticAnalysisResultRaw.ExportedCount.ExportedReceivers
	scan.ExportedProviderCount = StaticAnalysisResultRaw.ExportedCount.ExportedProviders
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
	scan.PrivacyRiskCount = StaticAnalysisResultRaw.Appsec.Trackers
	scan.HighSeverityDistributionCount = len(StaticAnalysisResultRaw.Appsec.High)
	scan.WarningSeverityDistributionCount = len(StaticAnalysisResultRaw.Appsec.Warning)
	scan.InfoSeverityDistributionCount = len(StaticAnalysisResultRaw.Appsec.Info)
	scan.SecureSeverityDistributionCount = len(StaticAnalysisResultRaw.Appsec.Secure)
	scan.HotspotSeverityDistributionCount = len(StaticAnalysisResultRaw.Appsec.Hotspot)
	scan.FinishedAt = time.Now()

	s.abstract.repoApkVersion.Update(ctx, s.abstract.db, apkVersion)
	s.abstract.repo.Update(ctx, s.abstract.db, scan)
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

func RawResultToModelStaticAnalysisResult(rawResult response.StaticAnalysisResultRaw) entity.StaticAnalysisResult {

	var Permissions []entity.Permission
	for permission, data := range rawResult.Permissions {
		Permissions = append(Permissions, entity.Permission{
			Permission:  permission,
			Status:      data.Status,
			Info:        data.Info,
			Description: data.Description,
		})
	}

	var AndroidApi []entity.AndroidApi
	for api, data := range rawResult.AndroidApi {
		var Files []entity.Files
		for filename, value := range data.Files {
			Files = append(Files, entity.Files{
				FileName: filename,
				Value:    value,
			})
		}
		AndroidApi = append(AndroidApi, entity.AndroidApi{
			ApiType: api,
			Files:   Files,
			Metadata: struct {
				Description string `json:"description" bson:"description"`
				Severity    string `json:"severity" bson:"severity"`
			}(data.Metadata),
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

	var NiapAnalysis []entity.NiapAnalysis
	for niapType, data := range rawResult.NiapAnalysis {
		NiapAnalysis = append(NiapAnalysis, entity.NiapAnalysis{
			NiapAnalysisType: niapType,
			Choice:           data.Choice,
			Description:      data.Description,
			Class:            data.Class,
		})
	}

	var BrowsableActivities []entity.BrowsableActivities
	for browsableActivitiesType, data := range rawResult.BrowsableActivities {
		BrowsableActivities = append(BrowsableActivities, entity.BrowsableActivities{
			BrowsableActivitiesType: browsableActivitiesType,
			Schemes:                 data.Schemes,
			MimeTypes:               data.MimeTypes,
			Hosts:                   data.Hosts,
			Ports:                   data.Ports,
			Paths:                   data.Paths,
			PathPrefixs:             data.PathPrefixs,
			PathPatterns:            data.PathPatterns,
			Browsable:               data.Browsable,
		})
	}

	var APKId []entity.APKId
	for apkIdType, data := range rawResult.Apkid {
		APKId = append(APKId, entity.APKId{
			ApkIdType: apkIdType,
			AntiVm:    data.AntiVm,
			AntiDebug: data.AntiDebug,
			Compiler:  data.Compiler,
		})
	}

	var firebaseUrls []entity.FirebaseUrl
	for _, data := range rawResult.FirebaseUrls {
		firebaseUrls = append(firebaseUrls, entity.FirebaseUrl{
			Url:  data.Url,
			Open: data.Open,
		})
	}

	StaticAnalysisResult := entity.StaticAnalysisResult{
		ExportedActivities:  rawResult.ExportedActivities,
		BrowsableActivities: BrowsableActivities,
		Activities:          rawResult.Activities,
		Receivers:           rawResult.Receivers,
		Providers:           rawResult.Providers,
		Services:            rawResult.Services,
		Libraries:           rawResult.Libraries,
		Permissions:         Permissions,
		ManifestAnalysis:    rawResult.ManifestAnalysis,
		NetworkSecurity: []struct {
			Scope       []string `json:"scope" bson:"scope"`
			Description string   `json:"description" bson:"description"`
			Severity    string   `json:"severity" bson:"severity"`
		}(rawResult.NetworkSecurity),
		BinaryAnalysis: rawResult.BinaryAnalysis,
		FileAnalysis:   rawResult.FileAnalysis,
		AndroidApi:     AndroidApi,
		CodeAnalysis:   CodeAnalysis,
		NiapAnalysis:   NiapAnalysis,
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
		Files:        rawResult.Files,
		Apkid:        APKId,
		Quark:        rawResult.Quark,
		Trackers: struct {
			DetectedTrackers int `json:"detected_trackers" bson:"detected_trackers"`
			TotalTrackers    int `json:"total_trackers" bson:"total_trackers"`
			Trackers         []struct {
				Name       string `json:"name" bson:"name"`
				Categories string `json:"categories" bson:"categories"`
				Url        string `json:"url" bson:"url"`
			} `json:"trackers" bson:"trackers"`
		}(rawResult.Trackers),
		PlaystoreDetails: struct {
			Title                    string      `json:"title" bson:"title"`
			Description              string      `json:"description" bson:"description"`
			Summary                  string      `json:"summary" bson:"summary"`
			Installs                 string      `json:"installs" bson:"installs"`
			MinInstalls              int         `json:"min_installs" bson:"min_installs"`
			RealInstalls             int         `json:"real_installs" bson:"real_installs"`
			Score                    interface{} `json:"score" bson:"score"`
			Ratings                  interface{} `json:"ratings" bson:"ratings"`
			Reviews                  interface{} `json:"reviews" bson:"reviews"`
			Histogram                []int       `json:"histogram" bson:"histogram"`
			Price                    int         `json:"price" bson:"price"`
			Free                     bool        `json:"free" bson:"free"`
			Currency                 string      `json:"currency" bson:"currency"`
			Sale                     bool        `json:"sale" bson:"sale"`
			SaleTime                 interface{} `json:"sale_time" bson:"sale_time"`
			OriginalPrice            interface{} `json:"original_price" bson:"original_price"`
			SaleText                 interface{} `json:"sale_text" bson:"sale_text"`
			OffersIAP                bool        `json:"offers_iap" bson:"offers_iap"`
			InAppProductPrice        interface{} `json:"in_app_product_price" bson:"in_app_product_price"`
			Developer                string      `json:"developer" bson:"developer"`
			DeveloperId              string      `json:"developer_id" bson:"developer_id"`
			DeveloperEmail           string      `json:"developer_email" bson:"developer_email"`
			DeveloperWebsite         string      `json:"developer_website" bson:"developer_website"`
			DeveloperAddress         interface{} `json:"developer_address" bson:"developer_address"`
			PrivacyPolicy            string      `json:"privacy_policy" bson:"privacy_policy"`
			Genre                    string      `json:"genre"`
			GenreId                  string      `json:"genre_id" bson:"genre_id"`
			Icon                     string      `json:"icon"`
			HeaderImage              string      `json:"header_image" bson:"header_image"`
			Screenshots              []string    `json:"screenshots"`
			Video                    interface{} `json:"video"`
			VideoImage               interface{} `json:"video_image" bson:"video_image"`
			ContentRating            string      `json:"content_rating" bson:"content_rating"`
			ContentRatingDescription interface{} `json:"content_rating_description" bson:"content_rating_description"`
			AdSupported              bool        `json:"ad_supported" bson:"ad_supported"`
			ContainsAds              bool        `json:"contains_ads" bson:"contains_ads"`
			Released                 interface{} `json:"released"`
			Updated                  int         `json:"updated"`
			Version                  string      `json:"version"`
			RecentChanges            string      `json:"recent_changes" bson:"recent_changes"`
			RecentChangesHTML        string      `json:"recent_changes_html" bson:"recent_changes_html"`
			AppId                    string      `json:"app_id" bson:"app_id"`
			Url                      string      `json:"url"`
			Error                    bool        `json:"error"`
			AndroidVersionText       string      `json:"android_version_text" bson:"android_version_text"`
		}(rawResult.PlaystoreDetails),
		Secrets: rawResult.Secrets,
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
		AverageCvss:         rawResult.AverageCvss,
		DynamicAnalysisDone: rawResult.DynamicAnalysisDone,
		VirusTotal:          rawResult.VirusTotal,
	}
	return StaticAnalysisResult

}
