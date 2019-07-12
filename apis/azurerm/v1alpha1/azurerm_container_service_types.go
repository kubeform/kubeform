package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermContainerService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermContainerServiceSpec   `json:"spec,omitempty"`
	Status            AzurermContainerServiceStatus `json:"status,omitempty"`
}

type AzurermContainerServiceSpecLinuxProfileSshKey struct {
	KeyData string `json:"key_data"`
}

type AzurermContainerServiceSpecLinuxProfile struct {
	AdminUsername string                                    `json:"admin_username"`
	SshKey        []AzurermContainerServiceSpecLinuxProfile `json:"ssh_key"`
}

type AzurermContainerServiceSpecAgentPoolProfile struct {
	VmSize    string `json:"vm_size"`
	Name      string `json:"name"`
	Count     int    `json:"count"`
	DnsPrefix string `json:"dns_prefix"`
	Fqdn      string `json:"fqdn"`
}

type AzurermContainerServiceSpecServicePrincipal struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AzurermContainerServiceSpecMasterProfile struct {
	DnsPrefix string `json:"dns_prefix"`
	Fqdn      string `json:"fqdn"`
	Count     int    `json:"count"`
}

type AzurermContainerServiceSpecDiagnosticsProfile struct {
	Enabled    bool   `json:"enabled"`
	StorageUri string `json:"storage_uri"`
}

type AzurermContainerServiceSpec struct {
	LinuxProfile          []AzurermContainerServiceSpec `json:"linux_profile"`
	AgentPoolProfile      []AzurermContainerServiceSpec `json:"agent_pool_profile"`
	ServicePrincipal      []AzurermContainerServiceSpec `json:"service_principal"`
	Tags                  map[string]string             `json:"tags"`
	Name                  string                        `json:"name"`
	Location              string                        `json:"location"`
	ResourceGroupName     string                        `json:"resource_group_name"`
	OrchestrationPlatform string                        `json:"orchestration_platform"`
	MasterProfile         []AzurermContainerServiceSpec `json:"master_profile"`
	DiagnosticsProfile    []AzurermContainerServiceSpec `json:"diagnostics_profile"`
}



type AzurermContainerServiceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermContainerServiceList is a list of AzurermContainerServices
type AzurermContainerServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermContainerService CRD objects
	Items []AzurermContainerService `json:"items,omitempty"`
}