package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AppServiceSlot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServiceSlotSpec   `json:"spec,omitempty"`
	Status            AppServiceSlotStatus `json:"status,omitempty"`
}

type AppServiceSlotSpecConnectionString struct {
	Name string `json:"name" tf:"name"`
	Type string `json:"type" tf:"type"`
	// Sensitive Data. Provide secret name which contains one value only
	Value core.LocalObjectReference `json:"value" tf:"value"`
}

type AppServiceSlotSpecIdentity struct {
	Type string `json:"type" tf:"type"`
}

type AppServiceSlotSpecSiteConfigCors struct {
	// +kubebuilder:validation:UniqueItems=true
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

type AppServiceSlotSpec struct {
	AppServiceName   string `json:"appServiceName" tf:"app_service_name"`
	AppServicePlanID string `json:"appServicePlanID" tf:"app_service_plan_id"`
	// +optional
	AppSettings map[string]string `json:"appSettings,omitempty" tf:"app_settings,omitempty"`
	// +optional
	ClientAffinityEnabled bool `json:"clientAffinityEnabled,omitempty" tf:"client_affinity_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ConnectionString []AppServiceSlotSpecConnectionString `json:"connectionString,omitempty" tf:"connection_string,omitempty"`
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
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AppServiceSlotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
