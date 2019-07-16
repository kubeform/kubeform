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

type ContainerService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerServiceSpec   `json:"spec,omitempty"`
	Status            ContainerServiceStatus `json:"status,omitempty"`
}

type ContainerServiceSpecAgentPoolProfile struct {
	// +optional
	Count     int    `json:"count,omitempty"`
	DnsPrefix string `json:"dns_prefix"`
	Name      string `json:"name"`
	VmSize    string `json:"vm_size"`
}

type ContainerServiceSpecDiagnosticsProfile struct {
	Enabled bool `json:"enabled"`
}

type ContainerServiceSpecLinuxProfileSshKey struct {
	KeyData string `json:"key_data"`
}

type ContainerServiceSpecLinuxProfile struct {
	AdminUsername string `json:"admin_username"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	SshKey []ContainerServiceSpecLinuxProfile `json:"ssh_key"`
}

type ContainerServiceSpecMasterProfile struct {
	// +optional
	Count     int    `json:"count,omitempty"`
	DnsPrefix string `json:"dns_prefix"`
}

type ContainerServiceSpecServicePrincipal struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type ContainerServiceSpec struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	AgentPoolProfile []ContainerServiceSpec `json:"agent_pool_profile"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	DiagnosticsProfile []ContainerServiceSpec `json:"diagnostics_profile"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	LinuxProfile []ContainerServiceSpec `json:"linux_profile"`
	Location     string                 `json:"location"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	MasterProfile         []ContainerServiceSpec `json:"master_profile"`
	Name                  string                 `json:"name"`
	OrchestrationPlatform string                 `json:"orchestration_platform"`
	ResourceGroupName     string                 `json:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	ServicePrincipal *[]ContainerServiceSpec `json:"service_principal,omitempty"`
}

type ContainerServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ContainerServiceList is a list of ContainerServices
type ContainerServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ContainerService CRD objects
	Items []ContainerService `json:"items,omitempty"`
}
