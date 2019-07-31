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

type ContainerService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerServiceSpec   `json:"spec,omitempty"`
	Status            ContainerServiceStatus `json:"status,omitempty"`
}

type ContainerServiceSpecAgentPoolProfile struct {
	// +optional
	Count     int    `json:"count,omitempty" tf:"count,omitempty"`
	DnsPrefix string `json:"dnsPrefix" tf:"dns_prefix"`
	Name      string `json:"name" tf:"name"`
	VmSize    string `json:"vmSize" tf:"vm_size"`
}

type ContainerServiceSpecDiagnosticsProfile struct {
	Enabled bool `json:"enabled" tf:"enabled"`
}

type ContainerServiceSpecLinuxProfileSshKey struct {
	KeyData string `json:"keyData" tf:"key_data"`
}

type ContainerServiceSpecLinuxProfile struct {
	AdminUsername string `json:"adminUsername" tf:"admin_username"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	SshKey []ContainerServiceSpecLinuxProfileSshKey `json:"sshKey" tf:"ssh_key"`
}

type ContainerServiceSpecMasterProfile struct {
	// +optional
	Count     int    `json:"count,omitempty" tf:"count,omitempty"`
	DnsPrefix string `json:"dnsPrefix" tf:"dns_prefix"`
}

type ContainerServiceSpecServicePrincipal struct {
	ClientID     string `json:"clientID" tf:"client_id"`
	ClientSecret string `json:"-" sensitive:"true" tf:"client_secret"`
}

type ContainerServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	AgentPoolProfile []ContainerServiceSpecAgentPoolProfile `json:"agentPoolProfile" tf:"agent_pool_profile"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	DiagnosticsProfile []ContainerServiceSpecDiagnosticsProfile `json:"diagnosticsProfile" tf:"diagnostics_profile"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	LinuxProfile []ContainerServiceSpecLinuxProfile `json:"linuxProfile" tf:"linux_profile"`
	Location     string                             `json:"location" tf:"location"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	MasterProfile         []ContainerServiceSpecMasterProfile `json:"masterProfile" tf:"master_profile"`
	Name                  string                              `json:"name" tf:"name"`
	OrchestrationPlatform string                              `json:"orchestrationPlatform" tf:"orchestration_platform"`
	ResourceGroupName     string                              `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	ServicePrincipal []ContainerServiceSpecServicePrincipal `json:"servicePrincipal,omitempty" tf:"service_principal,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ContainerServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
