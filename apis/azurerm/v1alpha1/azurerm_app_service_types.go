package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AzurermAppService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAppServiceSpec   `json:"spec,omitempty"`
	Status            AzurermAppServiceStatus `json:"status,omitempty"`
}

type AzurermAppServiceSpecSiteCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
	ManagedPipelineMode    string                            `json:"managed_pipeline_mode"`
	RemoteDebuggingEnabled bool                              `json:"remote_debugging_enabled"`
	Use32BitWorkerProcess  bool                              `json:"use_32_bit_worker_process"`
	MinTlsVersion          string                            `json:"min_tls_version"`
	Cors                   []AzurermAppServiceSpecSiteConfig `json:"cors"`
	AlwaysOn               bool                              `json:"always_on"`
	DefaultDocuments       []string                          `json:"default_documents"`
	DotnetFrameworkVersion string                            `json:"dotnet_framework_version"`
	RemoteDebuggingVersion string                            `json:"remote_debugging_version"`
	ScmType                string                            `json:"scm_type"`
	FtpsState              string                            `json:"ftps_state"`
	AppCommandLine         string                            `json:"app_command_line"`
	JavaContainer          string                            `json:"java_container"`
	LocalMysqlEnabled      bool                              `json:"local_mysql_enabled"`
	PhpVersion             string                            `json:"php_version"`
	VirtualNetworkName     string                            `json:"virtual_network_name"`
	IpRestriction          []AzurermAppServiceSpecSiteConfig `json:"ip_restriction"`
	JavaVersion            string                            `json:"java_version"`
	JavaContainerVersion   string                            `json:"java_container_version"`
	LinuxFxVersion         string                            `json:"linux_fx_version"`
	WindowsFxVersion       string                            `json:"windows_fx_version"`
	Http2Enabled           bool                              `json:"http2_enabled"`
	PythonVersion          string                            `json:"python_version"`
	WebsocketsEnabled      bool                              `json:"websockets_enabled"`
}

type AzurermAppServiceSpecLogsApplicationLogsAzureBlobStorage struct {
	Level           string `json:"level"`
	SasUrl          string `json:"sas_url"`
	RetentionInDays int    `json:"retention_in_days"`
}

type AzurermAppServiceSpecLogsApplicationLogs struct {
	AzureBlobStorage []AzurermAppServiceSpecLogsApplicationLogs `json:"azure_blob_storage"`
}

type AzurermAppServiceSpecLogs struct {
	ApplicationLogs []AzurermAppServiceSpecLogs `json:"application_logs"`
}

type AzurermAppServiceSpecIdentity struct {
	Type        string `json:"type"`
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
}

type AzurermAppServiceSpecAuthSettingsActiveDirectory struct {
	ClientSecret     string   `json:"client_secret"`
	AllowedAudiences []string `json:"allowed_audiences"`
	ClientId         string   `json:"client_id"`
}

type AzurermAppServiceSpecAuthSettingsFacebook struct {
	AppId       string   `json:"app_id"`
	AppSecret   string   `json:"app_secret"`
	OauthScopes []string `json:"oauth_scopes"`
}

type AzurermAppServiceSpecAuthSettingsGoogle struct {
	OauthScopes  []string `json:"oauth_scopes"`
	ClientId     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
}

type AzurermAppServiceSpecAuthSettingsMicrosoft struct {
	OauthScopes  []string `json:"oauth_scopes"`
	ClientId     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
}

type AzurermAppServiceSpecAuthSettingsTwitter struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
}

type AzurermAppServiceSpecAuthSettings struct {
	TokenStoreEnabled           bool                                `json:"token_store_enabled"`
	UnauthenticatedClientAction string                              `json:"unauthenticated_client_action"`
	Enabled                     bool                                `json:"enabled"`
	RuntimeVersion              string                              `json:"runtime_version"`
	ActiveDirectory             []AzurermAppServiceSpecAuthSettings `json:"active_directory"`
	Facebook                    []AzurermAppServiceSpecAuthSettings `json:"facebook"`
	DefaultProvider             string                              `json:"default_provider"`
	Google                      []AzurermAppServiceSpecAuthSettings `json:"google"`
	AdditionalLoginParams       map[string]string                   `json:"additional_login_params"`
	AllowedExternalRedirectUrls []string                            `json:"allowed_external_redirect_urls"`
	Issuer                      string                              `json:"issuer"`
	TokenRefreshExtensionHours  float64                             `json:"token_refresh_extension_hours"`
	Microsoft                   []AzurermAppServiceSpecAuthSettings `json:"microsoft"`
	Twitter                     []AzurermAppServiceSpecAuthSettings `json:"twitter"`
}

type AzurermAppServiceSpecSourceControl struct {
	RepoUrl string `json:"repo_url"`
	Branch  string `json:"branch"`
}

type AzurermAppServiceSpecConnectionString struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type AzurermAppServiceSpec struct {
	SiteCredential              []AzurermAppServiceSpec `json:"site_credential"`
	DefaultSiteHostname         string                  `json:"default_site_hostname"`
	Location                    string                  `json:"location"`
	SiteConfig                  []AzurermAppServiceSpec `json:"site_config"`
	Logs                        []AzurermAppServiceSpec `json:"logs"`
	HttpsOnly                   bool                    `json:"https_only"`
	OutboundIpAddresses         string                  `json:"outbound_ip_addresses"`
	PossibleOutboundIpAddresses string                  `json:"possible_outbound_ip_addresses"`
	Identity                    []AzurermAppServiceSpec `json:"identity"`
	ResourceGroupName           string                  `json:"resource_group_name"`
	AppServicePlanId            string                  `json:"app_service_plan_id"`
	Enabled                     bool                    `json:"enabled"`
	Name                        string                  `json:"name"`
	AuthSettings                []AzurermAppServiceSpec `json:"auth_settings"`
	Tags                        map[string]string       `json:"tags"`
	SourceControl               []AzurermAppServiceSpec `json:"source_control"`
	ClientAffinityEnabled       bool                    `json:"client_affinity_enabled"`
	ClientCertEnabled           bool                    `json:"client_cert_enabled"`
	AppSettings                 map[string]string       `json:"app_settings"`
	ConnectionString            []AzurermAppServiceSpec `json:"connection_string"`
}

type AzurermAppServiceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermAppServiceList is a list of AzurermAppServices
type AzurermAppServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAppService CRD objects
	Items []AzurermAppService `json:"items,omitempty"`
}
