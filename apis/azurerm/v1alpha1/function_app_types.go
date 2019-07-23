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

type FunctionApp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionAppSpec   `json:"spec,omitempty"`
	Status            FunctionAppStatus `json:"status,omitempty"`
}

type FunctionAppSpecConnectionString struct {
	Name string `json:"name" tf:"name"`
	Type string `json:"type" tf:"type"`
}

type FunctionAppSpecIdentity struct {
	Type string `json:"type" tf:"type"`
}

type FunctionAppSpecSiteConfig struct {
	// +optional
	AlwaysOn bool `json:"alwaysOn,omitempty" tf:"always_on,omitempty"`
	// +optional
	LinuxFxVersion string `json:"linuxFxVersion,omitempty" tf:"linux_fx_version,omitempty"`
	// +optional
	Use32BitWorkerProcess bool `json:"use32BitWorkerProcess,omitempty" tf:"use_32_bit_worker_process,omitempty"`
	// +optional
	WebsocketsEnabled bool `json:"websocketsEnabled,omitempty" tf:"websockets_enabled,omitempty"`
}

type FunctionAppSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	AppServicePlanID string `json:"appServicePlanID" tf:"app_service_plan_id"`
	// +optional
	AppSettings map[string]string `json:"appSettings,omitempty" tf:"app_settings,omitempty"`
	// +optional
	ClientAffinityEnabled bool `json:"clientAffinityEnabled,omitempty" tf:"client_affinity_enabled,omitempty"`
	// +optional
	ConnectionString []FunctionAppSpecConnectionString `json:"connectionString,omitempty" tf:"connection_string,omitempty"`
	// +optional
	EnableBuiltinLogging bool `json:"enableBuiltinLogging,omitempty" tf:"enable_builtin_logging,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	HttpsOnly bool `json:"httpsOnly,omitempty" tf:"https_only,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity          []FunctionAppSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	Location          string                    `json:"location" tf:"location"`
	Name              string                    `json:"name" tf:"name"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SiteConfig []FunctionAppSpecSiteConfig `json:"siteConfig,omitempty" tf:"site_config,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
}

type FunctionAppStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FunctionAppList is a list of FunctionApps
type FunctionAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FunctionApp CRD objects
	Items []FunctionApp `json:"items,omitempty"`
}
