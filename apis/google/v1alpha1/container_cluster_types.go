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

type ContainerCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerClusterSpec   `json:"spec,omitempty"`
	Status            ContainerClusterStatus `json:"status,omitempty"`
}

type ContainerClusterSpecIpAllocationPolicy struct {
	// +optional
	CreateSubnetwork bool `json:"create_subnetwork,omitempty"`
	// +optional
	SubnetworkName string `json:"subnetwork_name,omitempty"`
}

type ContainerClusterSpecMaintenancePolicyDailyMaintenanceWindow struct {
	StartTime string `json:"start_time"`
}

type ContainerClusterSpecMaintenancePolicy struct {
	// +kubebuilder:validation:MaxItems=1
	DailyMaintenanceWindow []ContainerClusterSpecMaintenancePolicy `json:"daily_maintenance_window"`
}

type ContainerClusterSpecMasterAuthorizedNetworksConfig struct{}

type ContainerClusterSpecPodSecurityPolicyConfig struct {
	Enabled bool `json:"enabled"`
}

type ContainerClusterSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// Deprecated
	EnableBinaryAuthorization bool `json:"enable_binary_authorization,omitempty"`
	// +optional
	EnableKubernetesAlpha bool `json:"enable_kubernetes_alpha,omitempty"`
	// +optional
	EnableLegacyAbac bool `json:"enable_legacy_abac,omitempty"`
	// +optional
	// Deprecated
	EnableTpu bool `json:"enable_tpu,omitempty"`
	// +optional
	InitialNodeCount int `json:"initial_node_count,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	IpAllocationPolicy *[]ContainerClusterSpec `json:"ip_allocation_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MaintenancePolicy *[]ContainerClusterSpec `json:"maintenance_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MasterAuthorizedNetworksConfig *[]ContainerClusterSpec `json:"master_authorized_networks_config,omitempty"`
	// +optional
	MinMasterVersion string `json:"min_master_version,omitempty"`
	Name             string `json:"name"`
	// +optional
	Network string `json:"network,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	PodSecurityPolicyConfig *[]ContainerClusterSpec `json:"pod_security_policy_config,omitempty"`
	// +optional
	RemoveDefaultNodePool bool `json:"remove_default_node_pool,omitempty"`
	// +optional
	ResourceLabels map[string]string `json:"resource_labels,omitempty"`
}

type ContainerClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ContainerClusterList is a list of ContainerClusters
type ContainerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ContainerCluster CRD objects
	Items []ContainerCluster `json:"items,omitempty"`
}
