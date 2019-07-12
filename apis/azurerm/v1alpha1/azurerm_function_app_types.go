package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermFunctionApp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermFunctionAppSpec   `json:"spec,omitempty"`
	Status            AzurermFunctionAppStatus `json:"status,omitempty"`
}

type AzurermFunctionAppSpecConnectionString struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AzurermFunctionAppSpecIdentity struct {
	Type        string `json:"type"`
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
}

type AzurermFunctionAppSpecSiteCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermFunctionAppSpecSiteConfig struct {
	AlwaysOn              bool   `json:"always_on"`
	Use32BitWorkerProcess bool   `json:"use_32_bit_worker_process"`
	WebsocketsEnabled     bool   `json:"websockets_enabled"`
	LinuxFxVersion        string `json:"linux_fx_version"`
}

type AzurermFunctionAppSpec struct {
	OutboundIpAddresses         string                   `json:"outbound_ip_addresses"`
	Name                        string                   `json:"name"`
	Kind                        string                   `json:"kind"`
	StorageConnectionString     string                   `json:"storage_connection_string"`
	EnableBuiltinLogging        bool                     `json:"enable_builtin_logging"`
	ConnectionString            []AzurermFunctionAppSpec `json:"connection_string"`
	Identity                    []AzurermFunctionAppSpec `json:"identity"`
	Tags                        map[string]string        `json:"tags"`
	PossibleOutboundIpAddresses string                   `json:"possible_outbound_ip_addresses"`
	HttpsOnly                   bool                     `json:"https_only"`
	SiteCredential              []AzurermFunctionAppSpec `json:"site_credential"`
	ResourceGroupName           string                   `json:"resource_group_name"`
	AppSettings                 map[string]string        `json:"app_settings"`
	ClientAffinityEnabled       bool                     `json:"client_affinity_enabled"`
	Location                    string                   `json:"location"`
	AppServicePlanId            string                   `json:"app_service_plan_id"`
	Enabled                     bool                     `json:"enabled"`
	Version                     string                   `json:"version"`
	DefaultHostname             string                   `json:"default_hostname"`
	SiteConfig                  []AzurermFunctionAppSpec `json:"site_config"`
}

type AzurermFunctionAppStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermFunctionAppList is a list of AzurermFunctionApps
type AzurermFunctionAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermFunctionApp CRD objects
	Items []AzurermFunctionApp `json:"items,omitempty"`
}
