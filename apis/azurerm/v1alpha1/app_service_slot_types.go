package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type AppServiceSlot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServiceSlotSpec   `json:"spec,omitempty"`
	Status            AppServiceSlotStatus `json:"status,omitempty"`
}

type AppServiceSlotSpecAuthSettingsActiveDirectory struct {
	// +optional
	AllowedAudiences []string `json:"allowedAudiences,omitempty" tf:"allowed_audiences,omitempty"`
	ClientID         string   `json:"clientID" tf:"client_id"`
	// +optional
	ClientSecret string `json:"-" sensitive:"true" tf:"client_secret,omitempty"`
}

type AppServiceSlotSpecAuthSettingsFacebook struct {
	AppID     string `json:"appID" tf:"app_id"`
	AppSecret string `json:"-" sensitive:"true" tf:"app_secret"`
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
}

type AppServiceSlotSpecAuthSettingsGoogle struct {
	ClientID     string `json:"clientID" tf:"client_id"`
	ClientSecret string `json:"-" sensitive:"true" tf:"client_secret"`
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
}

type AppServiceSlotSpecAuthSettingsMicrosoft struct {
	ClientID     string `json:"clientID" tf:"client_id"`
	ClientSecret string `json:"-" sensitive:"true" tf:"client_secret"`
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
}

type AppServiceSlotSpecAuthSettingsTwitter struct {
	ConsumerKey    string `json:"consumerKey" tf:"consumer_key"`
	ConsumerSecret string `json:"-" sensitive:"true" tf:"consumer_secret"`
}

type AppServiceSlotSpecAuthSettings struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ActiveDirectory []AppServiceSlotSpecAuthSettingsActiveDirectory `json:"activeDirectory,omitempty" tf:"active_directory,omitempty"`
	// +optional
	AdditionalLoginParams map[string]string `json:"additionalLoginParams,omitempty" tf:"additional_login_params,omitempty"`
	// +optional
	AllowedExternalRedirectUrls []string `json:"allowedExternalRedirectUrls,omitempty" tf:"allowed_external_redirect_urls,omitempty"`
	// +optional
	DefaultProvider string `json:"defaultProvider,omitempty" tf:"default_provider,omitempty"`
	Enabled         bool   `json:"enabled" tf:"enabled"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Facebook []AppServiceSlotSpecAuthSettingsFacebook `json:"facebook,omitempty" tf:"facebook,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Google []AppServiceSlotSpecAuthSettingsGoogle `json:"google,omitempty" tf:"google,omitempty"`
	// +optional
	Issuer string `json:"issuer,omitempty" tf:"issuer,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Microsoft []AppServiceSlotSpecAuthSettingsMicrosoft `json:"microsoft,omitempty" tf:"microsoft,omitempty"`
	// +optional
	RuntimeVersion string `json:"runtimeVersion,omitempty" tf:"runtime_version,omitempty"`
	// +optional
	TokenRefreshExtensionHours json.Number `json:"tokenRefreshExtensionHours,omitempty" tf:"token_refresh_extension_hours,omitempty"`
	// +optional
	TokenStoreEnabled bool `json:"tokenStoreEnabled,omitempty" tf:"token_store_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Twitter []AppServiceSlotSpecAuthSettingsTwitter `json:"twitter,omitempty" tf:"twitter,omitempty"`
	// +optional
	UnauthenticatedClientAction string `json:"unauthenticatedClientAction,omitempty" tf:"unauthenticated_client_action,omitempty"`
}

type AppServiceSlotSpecConnectionString struct {
	Name  string `json:"name" tf:"name"`
	Type  string `json:"type" tf:"type"`
	Value string `json:"-" sensitive:"true" tf:"value"`
}

type AppServiceSlotSpecIdentity struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids,omitempty"`
	// +optional
	PrincipalID string `json:"principalID,omitempty" tf:"principal_id,omitempty"`
	// +optional
	TenantID string `json:"tenantID,omitempty" tf:"tenant_id,omitempty"`
	Type     string `json:"type" tf:"type"`
}

type AppServiceSlotSpecSiteConfigCors struct {
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
	// +optional
	SupportCredentials bool `json:"supportCredentials,omitempty" tf:"support_credentials,omitempty"`
}

type AppServiceSlotSpecSiteConfigIpRestriction struct {
	IpAddress string `json:"ipAddress" tf:"ip_address"`
	// +optional
	SubnetMask string `json:"subnetMask,omitempty" tf:"subnet_mask,omitempty"`
}

type AppServiceSlotSpecSiteConfig struct {
	// +optional
	AlwaysOn bool `json:"alwaysOn,omitempty" tf:"always_on,omitempty"`
	// +optional
	AppCommandLine string `json:"appCommandLine,omitempty" tf:"app_command_line,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Cors []AppServiceSlotSpecSiteConfigCors `json:"cors,omitempty" tf:"cors,omitempty"`
	// +optional
	DefaultDocuments []string `json:"defaultDocuments,omitempty" tf:"default_documents,omitempty"`
	// +optional
	DotnetFrameworkVersion string `json:"dotnetFrameworkVersion,omitempty" tf:"dotnet_framework_version,omitempty"`
	// +optional
	FtpsState string `json:"ftpsState,omitempty" tf:"ftps_state,omitempty"`
	// +optional
	Http2Enabled bool `json:"http2Enabled,omitempty" tf:"http2_enabled,omitempty"`
	// +optional
	IpRestriction []AppServiceSlotSpecSiteConfigIpRestriction `json:"ipRestriction,omitempty" tf:"ip_restriction,omitempty"`
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
	MinTLSVersion string `json:"minTLSVersion,omitempty" tf:"min_tls_version,omitempty"`
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

type AppServiceSlotSpecSiteCredential struct {
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	// +optional
	Username string `json:"username,omitempty" tf:"username,omitempty"`
}

type AppServiceSlotSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	AppServiceName   string `json:"appServiceName" tf:"app_service_name"`
	AppServicePlanID string `json:"appServicePlanID" tf:"app_service_plan_id"`
	// +optional
	AppSettings map[string]string `json:"appSettings,omitempty" tf:"app_settings,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthSettings []AppServiceSlotSpecAuthSettings `json:"authSettings,omitempty" tf:"auth_settings,omitempty"`
	// +optional
	ClientAffinityEnabled bool `json:"clientAffinityEnabled,omitempty" tf:"client_affinity_enabled,omitempty"`
	// +optional
	ConnectionString []AppServiceSlotSpecConnectionString `json:"connectionString,omitempty" tf:"connection_string,omitempty"`
	// +optional
	DefaultSiteHostname string `json:"defaultSiteHostname,omitempty" tf:"default_site_hostname,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	HttpsOnly bool `json:"httpsOnly,omitempty" tf:"https_only,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity          []AppServiceSlotSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	Location          string                       `json:"location" tf:"location"`
	Name              string                       `json:"name" tf:"name"`
	ResourceGroupName string                       `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SiteConfig []AppServiceSlotSpecSiteConfig `json:"siteConfig,omitempty" tf:"site_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SiteCredential []AppServiceSlotSpecSiteCredential `json:"siteCredential,omitempty" tf:"site_credential,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AppServiceSlotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppServiceSlotSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppServiceSlotList is a list of AppServiceSlots
type AppServiceSlotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppServiceSlot CRD objects
	Items []AppServiceSlot `json:"items,omitempty"`
}
