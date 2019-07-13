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

type AzurermFunctionApp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermFunctionAppSpec   `json:"spec,omitempty"`
	Status            AzurermFunctionAppStatus `json:"status,omitempty"`
}

type AzurermFunctionAppSpecConnectionString struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type AzurermFunctionAppSpecIdentity struct {
	Type        string `json:"type"`
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
}

type AzurermFunctionAppSpecSiteConfig struct {
	AlwaysOn              bool   `json:"always_on"`
	Use32BitWorkerProcess bool   `json:"use_32_bit_worker_process"`
	WebsocketsEnabled     bool   `json:"websockets_enabled"`
	LinuxFxVersion        string `json:"linux_fx_version"`
}

type AzurermFunctionAppSpecSiteCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermFunctionAppSpec struct {
	Version                     string                   `json:"version"`
	AppSettings                 map[string]string        `json:"app_settings"`
	ConnectionString            []AzurermFunctionAppSpec `json:"connection_string"`
	Kind                        string                   `json:"kind"`
	AppServicePlanId            string                   `json:"app_service_plan_id"`
	Tags                        map[string]string        `json:"tags"`
	OutboundIpAddresses         string                   `json:"outbound_ip_addresses"`
	PossibleOutboundIpAddresses string                   `json:"possible_outbound_ip_addresses"`
	Name                        string                   `json:"name"`
	Identity                    []AzurermFunctionAppSpec `json:"identity"`
	StorageConnectionString     string                   `json:"storage_connection_string"`
	EnableBuiltinLogging        bool                     `json:"enable_builtin_logging"`
	Location                    string                   `json:"location"`
	Enabled                     bool                     `json:"enabled"`
	ClientAffinityEnabled       bool                     `json:"client_affinity_enabled"`
	HttpsOnly                   bool                     `json:"https_only"`
	SiteConfig                  []AzurermFunctionAppSpec `json:"site_config"`
	SiteCredential              []AzurermFunctionAppSpec `json:"site_credential"`
	ResourceGroupName           string                   `json:"resource_group_name"`
	DefaultHostname             string                   `json:"default_hostname"`
}



type AzurermFunctionAppStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermFunctionAppList is a list of AzurermFunctionApps
type AzurermFunctionAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermFunctionApp CRD objects
	Items []AzurermFunctionApp `json:"items,omitempty"`
}