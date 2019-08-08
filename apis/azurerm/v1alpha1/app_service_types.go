package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AppService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServiceSpec   `json:"spec,omitempty"`
	Status            AppServiceStatus `json:"status,omitempty"`
}

type AppServiceSpecAuthSettingsActiveDirectory struct {
	// +optional
	AllowedAudiences []string `json:"allowedAudiences,omitempty" tf:"allowed_audiences,omitempty"`
	ClientID         string   `json:"clientID" tf:"client_id"`
	// +optional
	ClientSecret string `json:"-" sensitive:"true" tf:"client_secret,omitempty"`
}

type AppServiceSpecAuthSettingsFacebook struct {
	AppID     string `json:"appID" tf:"app_id"`
	AppSecret string `json:"-" sensitive:"true" tf:"app_secret"`
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
}

type AppServiceSpecAuthSettingsGoogle struct {
	ClientID     string `json:"clientID" tf:"client_id"`
	ClientSecret string `json:"-" sensitive:"true" tf:"client_secret"`
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
}

type AppServiceSpecAuthSettingsMicrosoft struct {
	ClientID     string `json:"clientID" tf:"client_id"`
	ClientSecret string `json:"-" sensitive:"true" tf:"client_secret"`
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
}

type AppServiceSpecAuthSettingsTwitter struct {
	ConsumerKey    string `json:"consumerKey" tf:"consumer_key"`
	ConsumerSecret string `json:"-" sensitive:"true" tf:"consumer_secret"`
}

type AppServiceSpecAuthSettings struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ActiveDirectory []AppServiceSpecAuthSettingsActiveDirectory `json:"activeDirectory,omitempty" tf:"active_directory,omitempty"`
	// +optional
	AdditionalLoginParams map[string]string `json:"additionalLoginParams,omitempty" tf:"additional_login_params,omitempty"`
	// +optional
	AllowedExternalRedirectUrls []string `json:"allowedExternalRedirectUrls,omitempty" tf:"allowed_external_redirect_urls,omitempty"`
	// +optional
	DefaultProvider string `json:"defaultProvider,omitempty" tf:"default_provider,omitempty"`
	Enabled         bool   `json:"enabled" tf:"enabled"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Facebook []AppServiceSpecAuthSettingsFacebook `json:"facebook,omitempty" tf:"facebook,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Google []AppServiceSpecAuthSettingsGoogle `json:"google,omitempty" tf:"google,omitempty"`
	// +optional
	Issuer string `json:"issuer,omitempty" tf:"issuer,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Microsoft []AppServiceSpecAuthSettingsMicrosoft `json:"microsoft,omitempty" tf:"microsoft,omitempty"`
	// +optional
	RuntimeVersion string `json:"runtimeVersion,omitempty" tf:"runtime_version,omitempty"`
	// +optional
	TokenRefreshExtensionHours json.Number `json:"tokenRefreshExtensionHours,omitempty" tf:"token_refresh_extension_hours,omitempty"`
	// +optional
	TokenStoreEnabled bool `json:"tokenStoreEnabled,omitempty" tf:"token_store_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Twitter []AppServiceSpecAuthSettingsTwitter `json:"twitter,omitempty" tf:"twitter,omitempty"`
	// +optional
	UnauthenticatedClientAction string `json:"unauthenticatedClientAction,omitempty" tf:"unauthenticated_client_action,omitempty"`
}

type AppServiceSpecConnectionString struct {
	Name  string `json:"name" tf:"name"`
	Type  string `json:"type" tf:"type"`
	Value string `json:"-" sensitive:"true" tf:"value"`
}

type AppServiceSpecIdentity struct {
	// +optional
	PrincipalID string `json:"principalID,omitempty" tf:"principal_id,omitempty"`
	// +optional
	TenantID string `json:"tenantID,omitempty" tf:"tenant_id,omitempty"`
	Type     string `json:"type" tf:"type"`
}

type AppServiceSpecLogsApplicationLogsAzureBlobStorage struct {
	Level           string `json:"level" tf:"level"`
	RetentionInDays int    `json:"retentionInDays" tf:"retention_in_days"`
	SasURL          string `json:"-" sensitive:"true" tf:"sas_url"`
}

type AppServiceSpecLogsApplicationLogs struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AzureBlobStorage []AppServiceSpecLogsApplicationLogsAzureBlobStorage `json:"azureBlobStorage,omitempty" tf:"azure_blob_storage,omitempty"`
}

type AppServiceSpecLogs struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ApplicationLogs []AppServiceSpecLogsApplicationLogs `json:"applicationLogs,omitempty" tf:"application_logs,omitempty"`
}

type AppServiceSpecSiteConfigCors struct {
	// +kubebuilder:validation:UniqueItems=true
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
	// +optional
	SupportCredentials bool `json:"supportCredentials,omitempty" tf:"support_credentials,omitempty"`
}

type AppServiceSpecSiteConfigIpRestriction struct {
	IpAddress string `json:"ipAddress" tf:"ip_address"`
	// +optional
	SubnetMask string `json:"subnetMask,omitempty" tf:"subnet_mask,omitempty"`
}

type AppServiceSpecSiteConfig struct {
	// +optional
	AlwaysOn bool `json:"alwaysOn,omitempty" tf:"always_on,omitempty"`
	// +optional
	AppCommandLine string `json:"appCommandLine,omitempty" tf:"app_command_line,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Cors []AppServiceSpecSiteConfigCors `json:"cors,omitempty" tf:"cors,omitempty"`
	// +optional
	DefaultDocuments []string `json:"defaultDocuments,omitempty" tf:"default_documents,omitempty"`
	// +optional
	DotnetFrameworkVersion string `json:"dotnetFrameworkVersion,omitempty" tf:"dotnet_framework_version,omitempty"`
	// +optional
	FtpsState string `json:"ftpsState,omitempty" tf:"ftps_state,omitempty"`
	// +optional
	Http2Enabled bool `json:"http2Enabled,omitempty" tf:"http2_enabled,omitempty"`
	// +optional
	IpRestriction []AppServiceSpecSiteConfigIpRestriction `json:"ipRestriction,omitempty" tf:"ip_restriction,omitempty"`
	// +optional
	JavaContainer string `json:"javaContainer,omitempty" tf:"java_container,omitempty"`
	// +optional
	JavaContainerVersion string `json:"javaContainerVersion,omitempty" tf:"java_container_version,omitempty"`
	// +optional
	JavaVersion string `json:"javaVersion,omitempty" tf:"java_version,omitempty"`
	// +optional
	LinuxFxVersion string `json:"linuxFxVersion,omitempty" tf:"linux_fx_version,omitempty"`
	// +optional
	LocalMysqlEnabled bool `json:"localMysqlEnabled,omitempty" tf:"local_mysql_enabled,omitempty"`
	// +optional
	ManagedPipelineMode string `json:"managedPipelineMode,omitempty" tf:"managed_pipeline_mode,omitempty"`
	// +optional
	MinTlsVersion string `json:"minTlsVersion,omitempty" tf:"min_tls_version,omitempty"`
	// +optional
	PhpVersion string `json:"phpVersion,omitempty" tf:"php_version,omitempty"`
	// +optional
	PythonVersion string `json:"pythonVersion,omitempty" tf:"python_version,omitempty"`
	// +optional
	RemoteDebuggingEnabled bool `json:"remoteDebuggingEnabled,omitempty" tf:"remote_debugging_enabled,omitempty"`
	// +optional
	RemoteDebuggingVersion string `json:"remoteDebuggingVersion,omitempty" tf:"remote_debugging_version,omitempty"`
	// +optional
	ScmType string `json:"scmType,omitempty" tf:"scm_type,omitempty"`
	// +optional
	Use32BitWorkerProcess bool `json:"use32BitWorkerProcess,omitempty" tf:"use_32_bit_worker_process,omitempty"`
	// +optional
	VirtualNetworkName string `json:"virtualNetworkName,omitempty" tf:"virtual_network_name,omitempty"`
	// +optional
	WebsocketsEnabled bool `json:"websocketsEnabled,omitempty" tf:"websockets_enabled,omitempty"`
	// +optional
	WindowsFxVersion string `json:"windowsFxVersion,omitempty" tf:"windows_fx_version,omitempty"`
}

type AppServiceSpecSiteCredential struct {
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	// +optional
	Username string `json:"username,omitempty" tf:"username,omitempty"`
}

type AppServiceSpecSourceControl struct {
	// +optional
	Branch string `json:"branch,omitempty" tf:"branch,omitempty"`
	// +optional
	RepoURL string `json:"repoURL,omitempty" tf:"repo_url,omitempty"`
}

type AppServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	AppServicePlanID string `json:"appServicePlanID" tf:"app_service_plan_id"`
	// +optional
	AppSettings map[string]string `json:"appSettings,omitempty" tf:"app_settings,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthSettings []AppServiceSpecAuthSettings `json:"authSettings,omitempty" tf:"auth_settings,omitempty"`
	// +optional
	ClientAffinityEnabled bool `json:"clientAffinityEnabled,omitempty" tf:"client_affinity_enabled,omitempty"`
	// +optional
	ClientCertEnabled bool `json:"clientCertEnabled,omitempty" tf:"client_cert_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ConnectionString []AppServiceSpecConnectionString `json:"connectionString,omitempty" tf:"connection_string,omitempty"`
	// +optional
	DefaultSiteHostname string `json:"defaultSiteHostname,omitempty" tf:"default_site_hostname,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	HttpsOnly bool `json:"httpsOnly,omitempty" tf:"https_only,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity []AppServiceSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	Location string                   `json:"location" tf:"location"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Logs []AppServiceSpecLogs `json:"logs,omitempty" tf:"logs,omitempty"`
	Name string               `json:"name" tf:"name"`
	// +optional
	OutboundIPAddresses string `json:"outboundIPAddresses,omitempty" tf:"outbound_ip_addresses,omitempty"`
	// +optional
	PossibleOutboundIPAddresses string `json:"possibleOutboundIPAddresses,omitempty" tf:"possible_outbound_ip_addresses,omitempty"`
	ResourceGroupName           string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SiteConfig []AppServiceSpecSiteConfig `json:"siteConfig,omitempty" tf:"site_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SiteCredential []AppServiceSpecSiteCredential `json:"siteCredential,omitempty" tf:"site_credential,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceControl []AppServiceSpecSourceControl `json:"sourceControl,omitempty" tf:"source_control,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AppServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppServiceSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppServiceList is a list of AppServices
type AppServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppService CRD objects
	Items []AppService `json:"items,omitempty"`
}
