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

type AzurermFunctionAppSpecIdentity struct {
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
	Type        string `json:"type"`
}

type AzurermFunctionAppSpecSiteConfig struct {
	LinuxFxVersion        string `json:"linux_fx_version"`
	AlwaysOn              bool   `json:"always_on"`
	Use32BitWorkerProcess bool   `json:"use_32_bit_worker_process"`
	WebsocketsEnabled     bool   `json:"websockets_enabled"`
}

type AzurermFunctionAppSpecSiteCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermFunctionAppSpecConnectionString struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type AzurermFunctionAppSpec struct {
	Identity                    []AzurermFunctionAppSpec `json:"identity"`
	Tags                        map[string]string        `json:"tags"`
	OutboundIpAddresses         string                   `json:"outbound_ip_addresses"`
	StorageConnectionString     string                   `json:"storage_connection_string"`
	ResourceGroupName           string                   `json:"resource_group_name"`
	Location                    string                   `json:"location"`
	Kind                        string                   `json:"kind"`
	Enabled                     bool                     `json:"enabled"`
	AppSettings                 map[string]string        `json:"app_settings"`
	SiteConfig                  []AzurermFunctionAppSpec `json:"site_config"`
	Name                        string                   `json:"name"`
	DefaultHostname             string                   `json:"default_hostname"`
	PossibleOutboundIpAddresses string                   `json:"possible_outbound_ip_addresses"`
	ClientAffinityEnabled       bool                     `json:"client_affinity_enabled"`
	SiteCredential              []AzurermFunctionAppSpec `json:"site_credential"`
	AppServicePlanId            string                   `json:"app_service_plan_id"`
	EnableBuiltinLogging        bool                     `json:"enable_builtin_logging"`
	ConnectionString            []AzurermFunctionAppSpec `json:"connection_string"`
	HttpsOnly                   bool                     `json:"https_only"`
	Version                     string                   `json:"version"`
}

type AzurermFunctionAppStatus struct {
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
