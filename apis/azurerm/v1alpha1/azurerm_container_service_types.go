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

type AzurermContainerService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermContainerServiceSpec   `json:"spec,omitempty"`
	Status            AzurermContainerServiceStatus `json:"status,omitempty"`
}

type AzurermContainerServiceSpecMasterProfile struct {
	DnsPrefix string `json:"dns_prefix"`
	Fqdn      string `json:"fqdn"`
	Count     int    `json:"count"`
}

type AzurermContainerServiceSpecServicePrincipal struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AzurermContainerServiceSpecDiagnosticsProfile struct {
	Enabled    bool   `json:"enabled"`
	StorageUri string `json:"storage_uri"`
}

type AzurermContainerServiceSpecLinuxProfileSshKey struct {
	KeyData string `json:"key_data"`
}

type AzurermContainerServiceSpecLinuxProfile struct {
	SshKey        []AzurermContainerServiceSpecLinuxProfile `json:"ssh_key"`
	AdminUsername string                                    `json:"admin_username"`
}

type AzurermContainerServiceSpecAgentPoolProfile struct {
	Name      string `json:"name"`
	Count     int    `json:"count"`
	DnsPrefix string `json:"dns_prefix"`
	Fqdn      string `json:"fqdn"`
	VmSize    string `json:"vm_size"`
}

type AzurermContainerServiceSpec struct {
	MasterProfile         []AzurermContainerServiceSpec `json:"master_profile"`
	ServicePrincipal      []AzurermContainerServiceSpec `json:"service_principal"`
	DiagnosticsProfile    []AzurermContainerServiceSpec `json:"diagnostics_profile"`
	Tags                  map[string]string             `json:"tags"`
	Name                  string                        `json:"name"`
	ResourceGroupName     string                        `json:"resource_group_name"`
	OrchestrationPlatform string                        `json:"orchestration_platform"`
	LinuxProfile          []AzurermContainerServiceSpec `json:"linux_profile"`
	AgentPoolProfile      []AzurermContainerServiceSpec `json:"agent_pool_profile"`
	Location              string                        `json:"location"`
}



type AzurermContainerServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermContainerServiceList is a list of AzurermContainerServices
type AzurermContainerServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermContainerService CRD objects
	Items []AzurermContainerService `json:"items,omitempty"`
}