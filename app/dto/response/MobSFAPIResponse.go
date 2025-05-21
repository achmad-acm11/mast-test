package response

import "time"

type UploadMobSFResponse struct {
	FileName string `json:"file_name"`
	Hash     string `json:"hash"`
	ScanType string `json:"scan_type"`
}

type ScanStaticResultResponse struct {
	APK StaticAnalysisResultRaw
	IOS IOSStaticAnalysisResultRaw
}

type StaticAnalysisResultRaw struct {
	Title               string   `json:"title"`
	Version             string   `json:"version"`
	FileName            string   `json:"file_name"`
	AppName             string   `json:"app_name"`
	AppType             string   `json:"app_type"`
	Size                string   `json:"size"`
	Md5                 string   `json:"md5"`
	Sha1                string   `json:"sha1"`
	Sha256              string   `json:"sha256"`
	PackageName         string   `json:"package_name"`
	MainActivity        string   `json:"main_activity"`
	ExportedActivities  []string `json:"exported_activities"`
	BrowsableActivities map[string]struct {
		Schemes      []string `json:"schemes"`
		MimeTypes    []string `json:"mime_types"`
		Hosts        []string `json:"hosts"`
		Ports        []string `json:"ports"`
		Paths        []string `json:"paths"`
		PathPrefixs  []string `json:"path_prefixs"`
		PathPatterns []string `json:"path_patterns"`
		Browsable    bool     `json:"browsable"`
	} `json:"browsable_activities"`

	Activities          []string `json:"activities"`
	Receivers           []string `json:"receivers"`
	Providers           []string `json:"providers"`
	Services            []string `json:"services"`
	Libraries           []string `json:"libraries"`
	TargetSdk           string   `json:"target_sdk"`
	MaxSdk              string   `json:"max_sdk"`
	MinSdk              string   `json:"min_sdk"`
	VersionName         string   `json:"version_name"`
	VersionCode         string   `json:"version_code"`
	IconHidden          bool     `json:"icon_hidden"`
	IconFound           bool     `json:"icon_found"`
	CertificateAnalysis struct {
		CertificateInfo     string     `json:"certificate_info"`
		CertificateFindings [][]string `json:"certificate_findings"`
	} `json:"certificate_analysis"`
	Permissions map[string]struct {
		Status      string `json:"status"`
		Info        string `json:"info"`
		Description string `json:"description"`
	} `json:"permissions"`
	ManifestAnalysis []struct {
		Rule        string   `json:"rule"`
		Title       string   `json:"title"`
		Severity    string   `json:"severity"`
		Description string   `json:"description"`
		Name        string   `json:"name"`
		Component   []string `json:"component"`
	} `json:"manifest_analysis"`
	NetworkSecurity []struct {
		Scope       []string `json:"scope"`
		Description string   `json:"description"`
		Severity    string   `json:"severity"`
	} `json:"network_security"`
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
	} `json:"binary_analysis"`
	FileAnalysis []struct {
		Finding string   `json:"finding"`
		Files   []string `json:"files"`
	} `json:"file_analysis"`
	AndroidApi map[string]struct {
		Files    map[string]string `json:"files"`
		Metadata struct {
			Description string `json:"description"`
			Severity    string `json:"severity"`
		} `json:"metadata"`
	} `json:"android_api"`
	CodeAnalysis map[string]struct {
		Files    map[string]string `json:"files"`
		Metadata struct {
			Cvss        float64 `json:"cvss"`
			Cwe         string  `json:"cwe"`
			OwaspMobile string  `json:"owasp-mobile"`
			Masvs       string  `json:"masvs"`
			Ref         string  `json:"ref"`
			Description string  `json:"description"`
			Severity    string  `json:"severity"`
		} `json:"metadata"`
	} `json:"code_analysis"`
	NiapAnalysis map[string]struct {
		Choice      string `json:"choice"`
		Description string `json:"description"`
		Class       string `json:"class"`
	} `json:"niap_analysis"`
	Urls []struct {
		Urls []string `json:"urls"`
		Path string   `json:"path"`
	} `json:"urls"`
	Domains map[string]struct {
		Bad         string `json:"bad"`
		Geolocation struct {
			Ip           string `json:"ip"`
			CountryShort string `json:"country_short"`
			CountryLong  string `json:"country_long"`
			Region       string `json:"region"`
			City         string `json:"city"`
			Latitude     string `json:"latitude"`
			Longitude    string `json:"longitude"`
		} `json:"geolocation"`
	} `json:"domains"`
	Emails []struct {
		Emails []string `json:"emails"`
		Path   string   `json:"path"`
	} `json:"emails"`
	Strings      []string `json:"strings"`
	FirebaseUrls []struct {
		Url  string `json:"url"`
		Open bool   `json:"open"`
	} `json:"firebase_urls"`
	Files         []string `json:"files"`
	ExportedCount struct {
		ExportedActivities int `json:"exported_activities"`
		ExportedServices   int `json:"exported_services"`
		ExportedReceivers  int `json:"exported_receivers"`
		ExportedProviders  int `json:"exported_providers"`
	} `json:"exported_count"`
	Apkid map[string]struct {
		AntiVm    []string `json:"anti_vm"`
		AntiDebug []string `json:"anti_debug"`
		Compiler  []string `json:"compiler"`
	} `json:"apkid"`
	Quark    interface{} `json:"quark"`
	Trackers struct {
		DetectedTrackers int `json:"detected_trackers"`
		TotalTrackers    int `json:"total_trackers"`
		Trackers         []struct {
			Name       string `json:"name"`
			Categories string `json:"categories"`
			Url        string `json:"url"`
		} `json:"trackers"`
	} `json:"trackers"`
	PlaystoreDetails struct {
		Title                    string      `json:"title"`
		Description              string      `json:"description"`
		Summary                  string      `json:"summary"`
		Installs                 string      `json:"installs"`
		MinInstalls              int         `json:"minInstalls"`
		RealInstalls             int         `json:"realInstalls"`
		Score                    interface{} `json:"score"`
		Ratings                  interface{} `json:"ratings"`
		Reviews                  interface{} `json:"reviews"`
		Histogram                []int       `json:"histogram"`
		Price                    int         `json:"price"`
		Free                     bool        `json:"free"`
		Currency                 string      `json:"currency"`
		Sale                     bool        `json:"sale"`
		SaleTime                 interface{} `json:"saleTime"`
		OriginalPrice            interface{} `json:"originalPrice"`
		SaleText                 interface{} `json:"saleText"`
		OffersIAP                bool        `json:"offersIAP"`
		InAppProductPrice        interface{} `json:"inAppProductPrice"`
		Developer                string      `json:"developer"`
		DeveloperId              string      `json:"developerId"`
		DeveloperEmail           string      `json:"developerEmail"`
		DeveloperWebsite         string      `json:"developerWebsite"`
		DeveloperAddress         interface{} `json:"developerAddress"`
		PrivacyPolicy            string      `json:"privacyPolicy"`
		Genre                    string      `json:"genre"`
		GenreId                  string      `json:"genreId"`
		Icon                     string      `json:"icon"`
		HeaderImage              string      `json:"headerImage"`
		Screenshots              []string    `json:"screenshots"`
		Video                    interface{} `json:"video"`
		VideoImage               interface{} `json:"videoImage"`
		ContentRating            string      `json:"contentRating"`
		ContentRatingDescription interface{} `json:"contentRatingDescription"`
		AdSupported              bool        `json:"adSupported"`
		ContainsAds              bool        `json:"containsAds"`
		Released                 interface{} `json:"released"`
		Updated                  int         `json:"updated"`
		Version                  string      `json:"version"`
		RecentChanges            string      `json:"recentChanges"`
		RecentChangesHTML        string      `json:"recentChangesHTML"`
		AppId                    string      `json:"appId"`
		Url                      string      `json:"url"`
		Error                    bool        `json:"error"`
		AndroidVersionText       string      `json:"androidVersionText"`
	} `json:"playstore_details"`
	Secrets interface{} `json:"secrets"`
	Appsec  struct {
		High []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Section     string `json:"section"`
		} `json:"high"`
		Warning []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Section     string `json:"section"`
		} `json:"warning"`
		Info []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Section     string `json:"section"`
		} `json:"info"`
		Secure []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Section     string `json:"section"`
		} `json:"secure"`
		Hotspot []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Section     string `json:"section"`
		} `json:"hotspot"`
		TotalTrackers int    `json:"total_trackers"`
		Trackers      int    `json:"trackers"`
		SecurityScore int    `json:"security_score"`
		AppName       string `json:"app_name"`
		FileName      string `json:"file_name"`
		Hash          string `json:"hash"`
		VersionName   string `json:"version_name"`
	} `json:"appsec"`
	AverageCvss         interface{} `json:"average_cvss"`
	DynamicAnalysisDone bool        `json:"dynamic_analysis_done"`
	VirusTotal          interface{} `json:"virus_total"`
}

type IOSStaticAnalysisResultRaw struct {
	Version        string `json:"version"`
	Title          string `json:"title"`
	FileName       string `json:"file_name"`
	AppName        string `json:"app_name"`
	AppType        string `json:"app_type"`
	Size           string `json:"size"`
	Md5            string `json:"md5"`
	Sha1           string `json:"sha1"`
	Sha256         string `json:"sha256"`
	Build          string `json:"build"`
	AppVersion     string `json:"app_version"`
	SdkName        string `json:"sdk_name"`
	Platform       string `json:"platform"`
	MinOsVersion   string `json:"min_os_version"`
	BundleId       string `json:"bundle_id"`
	BundleUrlTypes []struct {
		CFBundleTypeRole   string   `json:"CFBundleTypeRole"`
		CFBundleURLName    string   `json:"CFBundleURLName"`
		CFBundleURLSchemes []string `json:"CFBundleURLSchemes"`
	} `json:"bundle_url_types"`
	BundleSupportedPlatforms []string `json:"bundle_supported_platforms"`
	IconFound                bool     `json:"icon_found"`
	InfoPlist                string   `json:"info_plist"`
	BinaryInfo               struct {
		Endian  string `json:"endian"`
		Bit     string `json:"bit"`
		Arch    string `json:"arch"`
		Subarch string `json:"subarch"`
	} `json:"binary_info"`
	Permissions map[string]struct {
		Status      string `json:"status"`
		Info        string `json:"info"`
		Description string `json:"description"`
	} `json:"permissions"`
	AtsAnalysis    []interface{} `json:"ats_analysis"`
	BinaryAnalysis map[string]struct {
		DetailedDesc string  `json:"detailed_desc"`
		Severity     string  `json:"severity"`
		Cvss         float64 `json:"cvss"`
		Cwe          string  `json:"cwe"`
		OwaspMobile  string  `json:"owasp-mobile"`
		Masvs        string  `json:"masvs"`
	} `json:"binary_analysis"`
	MachoAnalysis struct {
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
	} `json:"macho_analysis"`
	IosApi map[string]struct {
		Files    map[string]string `json:"files"`
		Metadata struct {
			Description string `json:"description"`
			Severity    string `json:"severity"`
		} `json:"metadata"`
	} `json:"ios_api"`
	CodeAnalysis map[string]struct {
		Files    map[string]string `json:"files"`
		Metadata struct {
			Cvss        float64 `json:"cvss"`
			Cwe         string  `json:"cwe"`
			OwaspMobile string  `json:"owasp-mobile"`
			Masvs       string  `json:"masvs"`
			Ref         string  `json:"ref"`
			Description string  `json:"description"`
			Severity    string  `json:"severity"`
		} `json:"metadata"`
	} `json:"code_analysis"`
	FileAnalysis []struct {
		Issue string `json:"issue"`
		Files []struct {
			FilePath string `json:"file_path"`
			Type     string `json:"type"`
			Hash     string `json:"hash"`
		} `json:"files"`
	} `json:"file_analysis"`
	Libraries []string `json:"libraries"`
	Files     []string `json:"files"`
	Urls      []struct {
		Urls []string `json:"urls"`
		Path string   `json:"path"`
	} `json:"urls"`
	Domains map[string]struct {
		Bad         string `json:"bad"`
		Geolocation struct {
			Ip           string `json:"ip"`
			CountryShort string `json:"country_short"`
			CountryLong  string `json:"country_long"`
			Region       string `json:"region"`
			City         string `json:"city"`
			Latitude     string `json:"latitude"`
			Longitude    string `json:"longitude"`
		} `json:"geolocation"`
	} `json:"domains"`
	Emails []struct {
		Emails []string `json:"emails"`
		Path   string   `json:"path"`
	} `json:"emails"`
	Strings      []string `json:"strings"`
	FirebaseUrls []struct {
		Url  string `json:"url"`
		Open bool   `json:"open"`
	} `json:"firebase_urls"`
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
		Error            bool     `json:"error"`
	} `json:"appstore_details"`
	Secrets  []interface{} `json:"secrets"`
	Trackers struct {
		DetectedTrackers int `json:"detected_trackers"`
		TotalTrackers    int `json:"total_trackers"`
		Trackers         []struct {
			Name       string `json:"name"`
			Categories string `json:"categories"`
			Url        string `json:"url"`
		} `json:"trackers"`
	} `json:"trackers"`
	Appsec struct {
		High []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Section     string `json:"section"`
		} `json:"high"`
		Warning []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Section     string `json:"section"`
		} `json:"warning"`
		Info []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Section     string `json:"section"`
		} `json:"info"`
		Secure []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Section     string `json:"section"`
		} `json:"secure"`
		Hotspot []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Section     string `json:"section"`
		} `json:"hotspot"`
		TotalTrackers int    `json:"total_trackers"`
		Trackers      int    `json:"trackers"`
		SecurityScore int    `json:"security_score"`
		AppName       string `json:"app_name"`
		FileName      string `json:"file_name"`
		Hash          string `json:"hash"`
		VersionName   string `json:"version_name"`
	} `json:"appsec"`
	AverageCvss interface{} `json:"average_cvss"`
	VirusTotal  interface{} `json:"virus_total"`
	BaseUrl     string      `json:"base_url"`
	DwdDir      string      `json:"dwd_dir"`
	HostOs      string      `json:"host_os"`
	Timestamp   time.Time   `json:"timestamp"`
}
