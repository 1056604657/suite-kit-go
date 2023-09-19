package cmd

import (
	_ "embed"
	"runtime"
)

var (
	toolName  = "A Go-based command-line tool for Suite operations, also known as `suitectl`."
	buildDate string
	goVersion = runtime.Version()
	repo      = "https://github.houston.softwaregrp.net/SMA-RnD/suite-kit-go"
)

type SmaxConfigJson struct {
	AllowWorkerOnMaster bool `json:"allowWorkerOnMaster"`
	MasterNodes         []struct {
		Hostname          string `json:"hostname"`
		User              string `json:"user"`
		Password          string `json:"password"`
		SkipWarning       bool   `json:"skipWarning"`
		PrivateKey        string `json:"privateKey"`
		Type              string `json:"type"`
		ThinpoolDevice    string `json:"thinpoolDevice"`
		DeviceType        string `json:"deviceType"`
		FlannelIface      string `json:"flannelIface"`
		SkipResourceCheck bool   `json:"skipResourceCheck"`
	} `json:"masterNodes"`
	WorkerNodes []struct {
		Hostname          string `json:"hostname"`
		User              string `json:"user"`
		Password          string `json:"password"`
		SkipWarning       bool   `json:"skipWarning"`
		PrivateKey        string `json:"privateKey"`
		Type              string `json:"type"`
		ThinpoolDevice    string `json:"thinpoolDevice"`
		DeviceType        string `json:"deviceType"`
		FlannelIface      string `json:"flannelIface"`
		SkipResourceCheck bool   `json:"skipResourceCheck"`
	} `json:"workerNodes"`
	LicenseAgreement struct {
		Eula     bool `json:"eula"`
		CallHome bool `json:"callHome"`
	} `json:"licenseAgreement"`
	Connection struct {
		ExternalHostname string `json:"externalHostname"`
		Port             string `json:"port"`
	} `json:"connection"`
	Volumes []struct {
		Type string `json:"type"`
		Name string `json:"name"`
		Host string `json:"host"`
		Path string `json:"path"`
	} `json:"volumes"`
	Database struct {
		Type  string `json:"type"`
		Param struct {
			HighAvailability *bool   `json:"highAvailability,omitempty"`
			DbHost           *string `json:"dbHost,omitempty"`
			DbPort           *string `json:"dbPort,omitempty"`
			DbName           *string `json:"dbName,omitempty"`
			DbUser           *string `json:"dbUser,omitempty"`
			DbPassword       *string `json:"dbPassword,omitempty"`
			DbAutoCreate     *bool   `json:"dbAutoCreate,omitempty"`
		} `json:"param"`
	} `json:"database"`
	Capabilities struct {
		Version             string `json:"version"`
		Suite               string `json:"suite"`
		InstallSize         string `json:"installSize"`
		Edition             string `json:"edition"`
		CapabilitySelection []struct {
			Name string `json:"name"`
		} `json:"capabilitySelection"`
		Configuration []struct {
			ItomSuiteSize        string `json:"itom_suite_size"`
			ItomSuiteInstallType string `json:"itom_suite_install_type"`
			ItomSuiteMode        string `json:"itom_suite_mode"`
			ActivatedServices    []struct {
				Name             string `json:"name"`
				RegistryURL      string `json:"registry_url"`
				ControllerImgTag string `json:"controller_img_tag"`
			} `json:"activated_services"`
			ItomSuiteAutoscaling     bool   `json:"itom_suite_autoscaling"`
			DbownerPassword          string `json:"dbowner_password"`
			IntegrationadminPassword string `json:"integrationadmin_password"`
			InternalDbaownerPassword string `json:"internal_dbaowner_password"`
			Locale                   string `json:"locale"`
			Database                 []struct {
				ProductName      *string `json:"product_name,omitempty"`
				Internal         *bool   `json:"internal,omitempty"`
				DbEngine         *string `json:"db_engine,omitempty"`
				DbServer         *string `json:"db_server,omitempty"`
				DbPort           *string `json:"db_port,omitempty"`
				DbInst           *string `json:"db_inst,omitempty"`
				DbLogin          *string `json:"db_login,omitempty"`
				DbPassword       *string `json:"db_password,omitempty"`
				DbSsl            *bool   `json:"db_ssl,omitempty"`
				DbPgSslMode      *string `json:"db_pg_ssl_mode,omitempty"`
				DbSslCertPath    *string `json:"db_ssl_cert_path,omitempty"`
				DbSslCertContent *string `json:"db_ssl_cert_content,omitempty"`
			} `json:"database"`
			VerticaDatabase []struct {
				ProductName      *string `json:"product_name,omitempty"`
				DbEngine         *string `json:"db_engine,omitempty"`
				DbServer         *string `json:"db_server,omitempty"`
				DbPort           *string `json:"db_port,omitempty"`
				DbInst           *string `json:"db_inst,omitempty"`
				DbLogin          *string `json:"db_login,omitempty"`
				DbPassword       *string `json:"db_password,omitempty"`
				DbSsl            *bool   `json:"db_ssl,omitempty"`
				DbSslCertPath    *string `json:"db_ssl_cert_path,omitempty"`
				DbSslCertContent *string `json:"db_ssl_cert_content,omitempty"`
			} `json:"vertica_database"`
			SysadminPassword string `json:"sysadmin_password"`
			BoadminPassword  string `json:"boadmin_password"`
		} `json:"configuration"`
	} `json:"capabilities"`
}

type FeatureSet struct {
	ID            string   `json:"id"`
	ItsmaServices []string `json:"itsmaServices"`
}

var requireDBServices = map[string]string{
	"itom-cmdb": "cmdb",
	"itom-sam":  "sam",
	"itom-dnd":  "dnd",
	"itom-cmp":  "cmp",
	"itom-cgro": "cgro",
}

var (
	//go:embed template/config.json
	smaxJson string
	//go:embed template/install.properties
	installProperties string

	currentMaster string

	smaxJsonOutput         SmaxConfigJson
	featuresetsJson        []FeatureSet
	smaxVersion            string
	majorVersion           string
	cdfPass                string
	cdfVersion             string
	workers                string
	masters                string
	virtualFqdn            string
	artifactoryBase        string
	nfsServer              string
	dbServer               string
	dbAdminUsername        string
	dbAdminPass            string
	dbSslCert              string
	verticaServer          string
	verticaAdminUsername   string
	verticaAdminPass       string
	verticaSslCert         string
	cdfapiServerDbUser     string
	cdfapiServerDbPassword string
	capabilities           string
	services               string
	nodeRootPass           string
	v                      string
	masterAsWorker         bool
	listResources          bool
	approveAll             bool
	enableVxlan            bool
)

const baseDir = "/home/admin"
