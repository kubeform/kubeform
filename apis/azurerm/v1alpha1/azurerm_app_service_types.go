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

type AzurermAppServiceSpecSourceControl struct {
	Branch  string `json:"branch"`
	RepoUrl string `json:"repo_url"`
}

type AzurermAppServiceSpecIdentity struct {
	Type        string `json:"type"`
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
}

type AzurermAppServiceSpecSiteConfigCors struct {
	SupportCredentials bool     `json:"support_credentials"`
	AllowedOrigins     []string `json:"allowed_origins"`
}

type AzurermAppServiceSpecSiteConfigIpRestriction struct {
	IpAddress  string `json:"ip_address"`
	SubnetMask string `json:"subnet_mask"`
}

type AzurermAppServiceSpecSiteConfig struct {
	DefaultDocuments       []string                          `json:"default_documents"`
	JavaVersion            string                            `json:"java_version"`
	JavaContainerVersion   string                            `json:"java_container_version"`
	RemoteDebuggingEnabled bool                              `json:"remote_debugging_enabled"`
	RemoteDebuggingVersion string                            `json:"remote_debugging_version"`
	ScmType                string                            `json:"scm_type"`
	LinuxFxVersion         string                            `json:"linux_fx_version"`
	AppCommandLine         string                            `json:"app_command_line"`
	Use32BitWorkerProcess  bool                              `json:"use_32_bit_worker_process"`
	FtpsState              string                            `json:"ftps_state"`
	VirtualNetworkName     string                            `json:"virtual_network_name"`
	Cors                   []AzurermAppServiceSpecSiteConfig `json:"cors"`
	IpRestriction          []AzurermAppServiceSpecSiteConfig `json:"ip_restriction"`
	Http2Enabled           bool                              `json:"http2_enabled"`
	JavaContainer          string                            `json:"java_container"`
	ManagedPipelineMode    string                            `json:"managed_pipeline_mode"`
	PhpVersion             string                            `json:"php_version"`
	AlwaysOn               bool                              `json:"always_on"`
	LocalMysqlEnabled      bool                              `json:"local_mysql_enabled"`
	PythonVersion          string                            `json:"python_version"`
	WebsocketsEnabled      bool                              `json:"websockets_enabled"`
	WindowsFxVersion       string                            `json:"windows_fx_version"`
	MinTlsVersion          string                            `json:"min_tls_version"`
	DotnetFrameworkVersion string                            `json:"dotnet_framework_version"`
}

type AzurermAppServiceSpecLogsApplicationLogsAzureBlobStorage struct {
	SasUrl          string `json:"sas_url"`
	RetentionInDays int    `json:"retention_in_days"`
	Level           string `json:"level"`
}

type AzurermAppServiceSpecLogsApplicationLogs struct {
	AzureBlobStorage []AzurermAppServiceSpecLogsApplicationLogs `json:"azure_blob_storage"`
}

type AzurermAppServiceSpecLogs struct {
	ApplicationLogs []AzurermAppServiceSpecLogs `json:"application_logs"`
}

type AzurermAppServiceSpecAuthSettingsGoogle struct {
	ClientId     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	OauthScopes  []string `json:"oauth_scopes"`
}

type AzurermAppServiceSpecAuthSettingsFacebook struct {
	AppId       string   `json:"app_id"`
	AppSecret   string   `json:"app_secret"`
	OauthScopes []string `json:"oauth_scopes"`
}

type AzurermAppServiceSpecAuthSettingsTwitter struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
}

type AzurermAppServiceSpecAuthSettingsMicrosoft struct {
	ClientId     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	OauthScopes  []string `json:"oauth_scopes"`
}

type AzurermAppServiceSpecAuthSettingsActiveDirectory struct {
	ClientId         string   `json:"client_id"`
	ClientSecret     string   `json:"client_secret"`
	AllowedAudiences []string `json:"allowed_audiences"`
}

type AzurermAppServiceSpecAuthSettings struct {
	Google                      []AzurermAppServiceSpecAuthSettings `json:"google"`
	AllowedExternalRedirectUrls []string                            `json:"allowed_external_redirect_urls"`
	Facebook                    []AzurermAppServiceSpecAuthSettings `json:"facebook"`
	DefaultProvider             string                              `json:"default_provider"`
	Issuer                      string                              `json:"issuer"`
	RuntimeVersion              string                              `json:"runtime_version"`
	TokenRefreshExtensionHours  float64                             `json:"token_refresh_extension_hours"`
	Twitter                     []AzurermAppServiceSpecAuthSettings `json:"twitter"`
	Enabled                     bool                                `json:"enabled"`
	AdditionalLoginParams       map[string]string                   `json:"additional_login_params"`
	Microsoft                   []AzurermAppServiceSpecAuthSettings `json:"microsoft"`
	TokenStoreEnabled           bool                                `json:"token_store_enabled"`
	UnauthenticatedClientAction string                              `json:"unauthenticated_client_action"`
	ActiveDirectory             []AzurermAppServiceSpecAuthSettings `json:"active_directory"`
}

type AzurermAppServiceSpecConnectionString struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AzurermAppServiceSpecSiteCredential struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type AzurermAppServiceSpec struct {
	Location                    string                  `json:"location"`
	ClientAffinityEnabled       bool                    `json:"client_affinity_enabled"`
	Tags                        map[string]string       `json:"tags"`
	SourceControl               []AzurermAppServiceSpec `json:"source_control"`
	Name                        string                  `json:"name"`
	Identity                    []AzurermAppServiceSpec `json:"identity"`
	AppServicePlanId            string                  `json:"app_service_plan_id"`
	SiteConfig                  []AzurermAppServiceSpec `json:"site_config"`
	Logs                        []AzurermAppServiceSpec `json:"logs"`
	HttpsOnly                   bool                    `json:"https_only"`
	PossibleOutboundIpAddresses string                  `json:"possible_outbound_ip_addresses"`
	ResourceGroupName           string                  `json:"resource_group_name"`
	AuthSettings                []AzurermAppServiceSpec `json:"auth_settings"`
	ClientCertEnabled           bool                    `json:"client_cert_enabled"`
	ConnectionString            []AzurermAppServiceSpec `json:"connection_string"`
	SiteCredential              []AzurermAppServiceSpec `json:"site_credential"`
	OutboundIpAddresses         string                  `json:"outbound_ip_addresses"`
	Enabled                     bool                    `json:"enabled"`
	AppSettings                 map[string]string       `json:"app_settings"`
	DefaultSiteHostname         string                  `json:"default_site_hostname"`
}



type AzurermAppServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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