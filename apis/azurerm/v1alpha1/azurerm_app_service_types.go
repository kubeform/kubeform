package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermAppService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAppServiceSpec   `json:"spec,omitempty"`
	Status            AzurermAppServiceStatus `json:"status,omitempty"`
}

type AzurermAppServiceSpecConnectionString struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type AzurermAppServiceSpecAuthSettingsFacebook struct {
	AppId       string   `json:"app_id"`
	AppSecret   string   `json:"app_secret"`
	OauthScopes []string `json:"oauth_scopes"`
}

type AzurermAppServiceSpecAuthSettingsMicrosoft struct {
	ClientId     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	OauthScopes  []string `json:"oauth_scopes"`
}

type AzurermAppServiceSpecAuthSettingsTwitter struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
}

type AzurermAppServiceSpecAuthSettingsActiveDirectory struct {
	ClientId         string   `json:"client_id"`
	ClientSecret     string   `json:"client_secret"`
	AllowedAudiences []string `json:"allowed_audiences"`
}

type AzurermAppServiceSpecAuthSettingsGoogle struct {
	ClientId     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	OauthScopes  []string `json:"oauth_scopes"`
}

type AzurermAppServiceSpecAuthSettings struct {
	Facebook                    []AzurermAppServiceSpecAuthSettings `json:"facebook"`
	Microsoft                   []AzurermAppServiceSpecAuthSettings `json:"microsoft"`
	Twitter                     []AzurermAppServiceSpecAuthSettings `json:"twitter"`
	TokenStoreEnabled           bool                                `json:"token_store_enabled"`
	ActiveDirectory             []AzurermAppServiceSpecAuthSettings `json:"active_directory"`
	RuntimeVersion              string                              `json:"runtime_version"`
	TokenRefreshExtensionHours  float64                             `json:"token_refresh_extension_hours"`
	UnauthenticatedClientAction string                              `json:"unauthenticated_client_action"`
	AdditionalLoginParams       map[string]string                   `json:"additional_login_params"`
	Issuer                      string                              `json:"issuer"`
	DefaultProvider             string                              `json:"default_provider"`
	Google                      []AzurermAppServiceSpecAuthSettings `json:"google"`
	Enabled                     bool                                `json:"enabled"`
	AllowedExternalRedirectUrls []string                            `json:"allowed_external_redirect_urls"`
}

type AzurermAppServiceSpecLogsApplicationLogsAzureBlobStorage struct {
	RetentionInDays int    `json:"retention_in_days"`
	Level           string `json:"level"`
	SasUrl          string `json:"sas_url"`
}

type AzurermAppServiceSpecLogsApplicationLogs struct {
	AzureBlobStorage []AzurermAppServiceSpecLogsApplicationLogs `json:"azure_blob_storage"`
}

type AzurermAppServiceSpecLogs struct {
	ApplicationLogs []AzurermAppServiceSpecLogs `json:"application_logs"`
}

type AzurermAppServiceSpecSiteConfigCors struct {
	AllowedOrigins     []string `json:"allowed_origins"`
	SupportCredentials bool     `json:"support_credentials"`
}

type AzurermAppServiceSpecSiteConfigIpRestriction struct {
	IpAddress  string `json:"ip_address"`
	SubnetMask string `json:"subnet_mask"`
}

type AzurermAppServiceSpecSiteConfig struct {
	JavaContainerVersion   string                            `json:"java_container_version"`
	ManagedPipelineMode    string                            `json:"managed_pipeline_mode"`
	ScmType                string                            `json:"scm_type"`
	Use32BitWorkerProcess  bool                              `json:"use_32_bit_worker_process"`
	WebsocketsEnabled      bool                              `json:"websockets_enabled"`
	WindowsFxVersion       string                            `json:"windows_fx_version"`
	Cors                   []AzurermAppServiceSpecSiteConfig `json:"cors"`
	MinTlsVersion          string                            `json:"min_tls_version"`
	AlwaysOn               bool                              `json:"always_on"`
	Http2Enabled           bool                              `json:"http2_enabled"`
	PhpVersion             string                            `json:"php_version"`
	PythonVersion          string                            `json:"python_version"`
	RemoteDebuggingVersion string                            `json:"remote_debugging_version"`
	FtpsState              string                            `json:"ftps_state"`
	LinuxFxVersion         string                            `json:"linux_fx_version"`
	VirtualNetworkName     string                            `json:"virtual_network_name"`
	RemoteDebuggingEnabled bool                              `json:"remote_debugging_enabled"`
	AppCommandLine         string                            `json:"app_command_line"`
	DefaultDocuments       []string                          `json:"default_documents"`
	DotnetFrameworkVersion string                            `json:"dotnet_framework_version"`
	IpRestriction          []AzurermAppServiceSpecSiteConfig `json:"ip_restriction"`
	JavaVersion            string                            `json:"java_version"`
	JavaContainer          string                            `json:"java_container"`
	LocalMysqlEnabled      bool                              `json:"local_mysql_enabled"`
}

type AzurermAppServiceSpecSiteCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermAppServiceSpecSourceControl struct {
	RepoUrl string `json:"repo_url"`
	Branch  string `json:"branch"`
}

type AzurermAppServiceSpecIdentity struct {
	Type        string `json:"type"`
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
}

type AzurermAppServiceSpec struct {
	PossibleOutboundIpAddresses string                  `json:"possible_outbound_ip_addresses"`
	ClientAffinityEnabled       bool                    `json:"client_affinity_enabled"`
	Enabled                     bool                    `json:"enabled"`
	ConnectionString            []AzurermAppServiceSpec `json:"connection_string"`
	AuthSettings                []AzurermAppServiceSpec `json:"auth_settings"`
	Logs                        []AzurermAppServiceSpec `json:"logs"`
	DefaultSiteHostname         string                  `json:"default_site_hostname"`
	OutboundIpAddresses         string                  `json:"outbound_ip_addresses"`
	Location                    string                  `json:"location"`
	AppServicePlanId            string                  `json:"app_service_plan_id"`
	SiteConfig                  []AzurermAppServiceSpec `json:"site_config"`
	Name                        string                  `json:"name"`
	AppSettings                 map[string]string       `json:"app_settings"`
	Tags                        map[string]string       `json:"tags"`
	ClientCertEnabled           bool                    `json:"client_cert_enabled"`
	SiteCredential              []AzurermAppServiceSpec `json:"site_credential"`
	SourceControl               []AzurermAppServiceSpec `json:"source_control"`
	Identity                    []AzurermAppServiceSpec `json:"identity"`
	ResourceGroupName           string                  `json:"resource_group_name"`
	HttpsOnly                   bool                    `json:"https_only"`
}

type AzurermAppServiceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermAppServiceList is a list of AzurermAppServices
type AzurermAppServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAppService CRD objects
	Items []AzurermAppService `json:"items,omitempty"`
}
