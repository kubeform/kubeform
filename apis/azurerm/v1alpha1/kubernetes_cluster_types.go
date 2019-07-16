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

type KubernetesCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesClusterSpec   `json:"spec,omitempty"`
	Status            KubernetesClusterStatus `json:"status,omitempty"`
}

type KubernetesClusterSpecAgentPoolProfile struct {
	// +optional
	AvailabilityZones []string `json:"availability_zones,omitempty"`
	// +optional
	Count int `json:"count,omitempty"`
	// +optional
	EnableAutoScaling bool `json:"enable_auto_scaling,omitempty"`
	// +optional
	MaxCount int `json:"max_count,omitempty"`
	// +optional
	MinCount int    `json:"min_count,omitempty"`
	Name     string `json:"name"`
	// +optional
	NodeTaints []string `json:"node_taints,omitempty"`
	// +optional
	OsType string `json:"os_type,omitempty"`
	// +optional
	Type   string `json:"type,omitempty"`
	VmSize string `json:"vm_size"`
	// +optional
	VnetSubnetId string `json:"vnet_subnet_id,omitempty"`
}

type KubernetesClusterSpecLinuxProfileSshKey struct {
	KeyData string `json:"key_data"`
}

type KubernetesClusterSpecLinuxProfile struct {
	AdminUsername string `json:"admin_username"`
	// +kubebuilder:validation:MaxItems=1
	SshKey []KubernetesClusterSpecLinuxProfile `json:"ssh_key"`
}

type KubernetesClusterSpecServicePrincipal struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type KubernetesClusterSpec struct {
	AgentPoolProfile []KubernetesClusterSpec `json:"agent_pool_profile"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ApiServerAuthorizedIpRanges []string `json:"api_server_authorized_ip_ranges,omitempty"`
	DnsPrefix                   string   `json:"dns_prefix"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LinuxProfile      *[]KubernetesClusterSpec `json:"linux_profile,omitempty"`
	Location          string                   `json:"location"`
	Name              string                   `json:"name"`
	ResourceGroupName string                   `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	ServicePrincipal []KubernetesClusterSpec `json:"service_principal"`
}

type KubernetesClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KubernetesClusterList is a list of KubernetesClusters
type KubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubernetesCluster CRD objects
	Items []KubernetesCluster `json:"items,omitempty"`
}
