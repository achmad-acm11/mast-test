package entity

import "time"

type ResultData struct {
	APK StaticAnalysisResult `json:"apk"`
	IOS IOSStaticAnalysisResult
}

type StaticAnalysisResult struct {
	TypeResult          string                `json:"type_result" bson:"type_result"`
	ScanId              int                   `json:"scan_id" bson:"scan_id"`
	Scan                *Scan                 `json:"scan,omitempty"`
	ScanVersion         int                   `json:"scan_version" bson:"scan_version"`
	CreatedAt           time.Time             `json:"created_at" bson:"created_at"`
	UpdatedAt           time.Time             `json:"updated_at" bson:"updated_at"`
	ExportedActivities  []string              `json:"exported_activities" bson:"exported_activities"`
	BrowsableActivities []BrowsableActivities `json:"browsable_activities" bson:"browsable_activities"`
	Activities          []string              `json:"activities" bson:"activities"`
	Receivers           []string              `json:"receivers" bson:"receivers"`
	Providers           []string              `json:"providers" bson:"providers"`
	Services            []string              `json:"services" bson:"services"`
	Libraries           []string              `json:"libraries" bson:"libraries"`
	CertificateAnalysis struct {
		CertificateInfo     string     `json:"certificate_info" bson:"certificate_info"`
		CertificateFindings [][]string `json:"certificate_findings" bson:"certificate_findings"`
	} `json:"certificate_analysis" bson:"certificate_analysis"`
	Permissions      []Permission `json:"permissions" bson:"permissions"`
	ManifestAnalysis []struct {
		Rule        string   `json:"rule"`
		Title       string   `json:"title"`
		Severity    string   `json:"severity"`
		Description string   `json:"description"`
		Name        string   `json:"name"`
		Component   []string `json:"component"`
	} `json:"manifest_analysis" bson:"manifest_analysis"`
	NetworkSecurity []struct {
		Scope       []string `json:"scope" bson:"scope"`
		Description string   `json:"description" bson:"description"`
		Severity    string   `json:"severity" bson:"severity"`
	} `json:"network_security" bson:"network_security"`
	BinaryAnalysis []struct {
		Name string `json:"name"`
		Nx   struct {
			IsNx        bool   `json:"is_nx"`
			Severity    string `json:"severity"`
			Description string `json:"description"`
		} `json:"nx"`
		StackCanary struct {
			HasCanary   bool   `json:"has_canary"`
			Severity    string `json:"severity"`
			Description string `json:"description"`
		} `json:"stack_canary"`
		Rpath struct {
			Rpath       interface{} `json:"rpath"`
			Severity    string      `json:"severity"`
			Description string      `json:"description"`
		} `json:"rpath"`
		Runpath struct {
			Runpath     interface{} `json:"runpath"`
			Severity    string      `json:"severity"`
			Description string      `json:"description"`
		} `json:"runpath"`
		Fortify struct {
			IsFortified bool   `json:"is_fortified"`
			Severity    string `json:"severity"`
			Description string `json:"description"`
		} `json:"fortify"`
		Symbol struct {
			IsStripped  bool   `json:"is_stripped"`
			Severity    string `json:"severity"`
			Description string `json:"description"`
		} `json:"symbol"`
	} `json:"binary_analysis" bson:"binary_analysis"`
	FileAnalysis []struct {
		Finding string   `json:"finding"`
		Files   []string `json:"files"`
	} `json:"file_analysis" bson:"file_analysis"`
	AndroidApi   []AndroidApi   `json:"android_api" bson:"android_api"`
	CodeAnalysis []CodeAnalysis `json:"code_analysis" bson:"code_analysis"`
	NiapAnalysis []NiapAnalysis `json:"niap_analysis" bson:"niap_analysis"`
	Urls         []struct {
		Urls []string `json:"urls" bson:"urls"`
		Path string   `json:"path" bson:"path"`
	} `json:"urls" bson:"urls"`
	Domains []Domains `json:"domains" bson:"domains"`
	Emails  []struct {
		Emails []string `json:"emails" bson:"emails"`
		Path   string   `json:"path" bson:"path"`
	} `json:"emails" bson:"emails"`
	Strings      []string      `json:"strings" bson:"strings"`
	FirebaseUrls []FirebaseUrl `json:"firebase_urls"`
	Files        []string      `json:"files" bson:"files"`
	Apkid        []APKId       `json:"apkid" bson:"apkid"`
	Quark        interface{}   `json:"quark" bson:"quark"`
	Trackers     struct {
		DetectedTrackers int `json:"detected_trackers" bson:"detected_trackers"`
		TotalTrackers    int `json:"total_trackers" bson:"total_trackers"`
		Trackers         []struct {
			Name       string `json:"name" bson:"name"`
			Categories string `json:"categories" bson:"categories"`
			Url        string `json:"url" bson:"url"`
		} `json:"trackers" bson:"trackers"`
	} `json:"trackers" bson:"trackers"`
	PlaystoreDetails struct {
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
	} `json:"playstore_details" bson:"playstore_details"`
	Secrets interface{} `json:"secrets" bson:"secrets"`
	Appsec  struct {
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
	} `json:"appsec" bson:"appsec"`
	AverageCvss         interface{} `json:"average_cvss" bson:"average_cvss"`
	DynamicAnalysisDone bool        `json:"dynamic_analysis_done" bson:"dynamic_analysis_done"`
	VirusTotal          interface{} `json:"virus_total" bson:"virus_total"`
}

type IOSStaticAnalysisResult struct {
	TypeResult               string              `json:"type_result" bson:"type_result"`
	ScanId                   int                 `json:"scan_id" bson:"scan_id"`
	Scan                     *Scan               `json:"scan,omitempty"`
	ScanVersion              int                 `json:"scan_version" bson:"scan_version"`
	CreatedAt                time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt                time.Time           `json:"updated_at" bson:"updated_at"`
	BundleUrlTypes           []BundleUrlType     `json:"bundle_url_types"`
	BundleSupportedPlatforms []string            `json:"bundle_supported_platforms" bson:"bundle_supported_platforms"`
	IconFound                bool                `json:"icon_found" bson:"icon_found"`
	InfoPlist                string              `json:"info_plist" bson:"info_plist"`
	Permissions              []Permission        `json:"permissions" bson:"permissions"`
	AtsAnalysis              interface{}         `json:"ats_analysis" bson:"ats_analysis"`
	BinaryAnalysis           []IOSBinaryAnalysis `json:"binary_analysis" bson:"binary_analysis"`
	MachoAnalysis            struct {
		Name string `json:"name"`
		Nx   struct {
			HasNx       bool   `json:"has_nx"`
			Severity    string `json:"severity"`
			Description string `json:"description"`
		} `json:"nx"`
		Pie struct {
			HasPie      bool   `json:"has_pie"`
			Severity    string `json:"severity"`
			Description string `json:"description"`
		} `json:"pie"`
		StackCanary struct {
			HasCanary   bool   `json:"has_canary"`
			Severity    string `json:"severity"`
			Description string `json:"description"`
		} `json:"stack_canary"`
		Arc struct {
			HasArc      bool   `json:"has_arc"`
			Severity    string `json:"severity"`
			Description string `json:"description"`
		} `json:"arc"`
		Rpath struct {
			HasRpath    bool   `json:"has_rpath"`
			Severity    string `json:"severity"`
			Description string `json:"description"`
		} `json:"rpath"`
		CodeSignature struct {
			HasCodeSignature bool   `json:"has_code_signature"`
			Severity         string `json:"severity"`
			Description      string `json:"description"`
		} `json:"code_signature"`
		Encrypted struct {
			IsEncrypted bool   `json:"is_encrypted"`
			Severity    string `json:"severity"`
			Description string `json:"description"`
		} `json:"encrypted"`
		Symbol struct {
			IsStripped  bool   `json:"is_stripped"`
			Severity    string `json:"severity"`
			Description string `json:"description"`
		} `json:"symbol"`
	} `json:"macho_analysis" bson:"macho_analysis"`
	IosApi       []AndroidApi   `json:"ios_api" bson:"ios_api"`
	CodeAnalysis []CodeAnalysis `json:"code_analysis" bson:"code_analysis"`
	FileAnalysis []struct {
		Issue string `json:"issue" bson:"issue"`
		Files []struct {
			FilePath string `json:"file_path" bson:"file_path"`
			Type     string `json:"type" bson:"type"`
			Hash     string `json:"hash" bson:"hash"`
		} `json:"files" bson:"files"`
	} `json:"file_analysis" bson:"file_analysis"`
	Libraries []string `json:"libraries" bson:"libraries"`
	Files     []string `json:"files" bson:"files"`
	Urls      []struct {
		Urls []string `json:"urls" bson:"urls"`
		Path string   `json:"path" bson:"path"`
	} `json:"urls" bson:"urls"`
	Domains []Domains `json:"domains" bson:"domains"`
	Emails  []struct {
		Emails []string `json:"emails" bson:"emails"`
		Path   string   `json:"path" bson:"path"`
	} `json:"emails" bson:"emails"`
	Strings         []string      `json:"strings" bson:"strings"`
	FirebaseUrls    []FirebaseUrl `json:"firebase_urls"`
	AppstoreDetails struct {
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
	} `json:"appstore_details" bson:"appstore_details"`
	Secrets  interface{} `json:"secrets" bson:"secrets"`
	Trackers struct {
		DetectedTrackers int `json:"detected_trackers" bson:"detected_trackers"`
		TotalTrackers    int `json:"total_trackers" bson:"total_trackers"`
		Trackers         []struct {
			Name       string `json:"name" bson:"name"`
			Categories string `json:"categories" bson:"categories"`
			Url        string `json:"url" bson:"url"`
		} `json:"trackers" bson:"trackers"`
	} `json:"trackers" bson:"trackers"`
	Appsec struct {
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
	} `json:"appsec" bson:"appsec"`
	AverageCvss interface{} `json:"average_cvss" bson:"average_cvss"`
	VirusTotal  interface{} `json:"virus_total" bson:"virus_total"`
	BaseUrl     string      `json:"base_url" bson:"base_url"`
	DwdDir      string      `json:"dwd_dir" bson:"dwd_dir"`
	HostOs      string      `json:"host_os" bson:"host_os"`
	Timestamp   time.Time   `json:"timestamp" bson:"timestamp"`
}

type BundleUrlType struct {
	CFBundleTypeRole   string   `json:"CFBundleTypeRole"`
	CFBundleURLName    string   `json:"CFBundleURLName"`
	CFBundleURLSchemes []string `json:"CFBundleURLSchemes"`
}

type IOSBinaryAnalysis struct {
	IosBinaryAnalysisType string  `json:"ios_binary_analysis_type" bson:"ios_binary_analysis_type"`
	DetailedDesc          string  `json:"detailed_desc" bson:"detailed_desc"`
	Severity              string  `json:"severity" bson:"severity"`
	Cvss                  float64 `json:"cvss" bson:"cvss"`
	Cwe                   string  `json:"cwe" bson:"cwe"`
	OwaspMobile           string  `json:"owasp-mobile" bson:"owasp-mobile"`
	Masvs                 string  `json:"masvs" bson:"masvs"`
}

type BrowsableActivities struct {
	BrowsableActivitiesType string   `json:"browsable_activities_type" bson:"browsable_activities_type"`
	Schemes                 []string `json:"schemes" bson:"schemes"`
	MimeTypes               []string `json:"mime_types" bson:"mime_types"`
	Hosts                   []string `json:"hosts" bson:"hosts"`
	Ports                   []string `json:"ports" bson:"ports"`
	Paths                   []string `json:"paths" bson:"paths"`
	PathPrefixs             []string `json:"path_prefixs" bson:"path_prefixs"`
	PathPatterns            []string `json:"path_patterns" bson:"path_patterns"`
	Browsable               bool     `json:"browsable" bson:"browsable"`
}

type Permission struct {
	Permission  string `json:"permission" bson:"permission"`
	Status      string `json:"status" bson:"status"`
	Info        string `json:"info" bson:"info"`
	Description string `json:"description" bson:"description"`
}

type AndroidApi struct {
	ApiType  string  `json:"api_type" bson:"api_type"`
	Files    []Files `json:"files" bson:"files"`
	Metadata struct {
		Description string `json:"description" bson:"description"`
		Severity    string `json:"severity" bson:"severity"`
	} `json:"metadata" bson:"metadata"`
}

type Files struct {
	FileName string `json:"file_name" bson:"file_name"`
	Value    string `json:"value" bson:"value"`
}

type CodeAnalysis struct {
	CodeAnalysisType string  `json:"code_analysis_type" bson:"code_analysis_type"`
	Files            []Files `json:"files" bson:"files"`
	Metadata         struct {
		Cvss        float64 `json:"cvss" bson:"cvss"`
		Cwe         string  `json:"cwe" bson:"cwe"`
		OwaspMobile string  `json:"owasp-mobile" bson:"owasp-mobile"`
		Masvs       string  `json:"masvs" bson:"masvs"`
		Ref         string  `json:"ref" bson:"ref"`
		Description string  `json:"description" bson:"description"`
		Severity    string  `json:"severity" bson:"severity"`
	} `json:"metadata" bson:"metadata"`
}

type NiapAnalysis struct {
	NiapAnalysisType string `json:"niap_analysis_type" bson:"niap_analysis_type"`
	Choice           string `json:"choice" bson:"choice"`
	Description      string `json:"description" bson:"description"`
	Class            string `json:"class" bson:"class"`
}

type Domains struct {
	DomainType  string `json:"domain_type" bson:"domain_type"`
	Bad         string `json:"bad" bson:"bad"`
	Geolocation struct {
		Ip           string `json:"ip" bson:"ip"`
		CountryShort string `json:"country_short" bson:"country_short"`
		CountryLong  string `json:"country_long" bson:"country_long"`
		Region       string `json:"region" bson:"region"`
		City         string `json:"city" bson:"city"`
		Latitude     string `json:"latitude" bson:"latitude"`
		Longitude    string `json:"longitude" bson:"longitude"`
	} `json:"geolocation" bson:"geolocation"`
}

type FirebaseUrl struct {
	Url  string `json:"url"`
	Open bool   `json:"open"`
}

type APKId struct {
	ApkIdType string   `json:"apk_id_type" bson:"apk_id_type"`
	AntiVm    []string `json:"anti_vm"`
	AntiDebug []string `json:"anti_debug"`
	Compiler  []string `json:"compiler"`
}
