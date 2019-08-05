package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ContainerNodePool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerNodePoolSpec   `json:"spec,omitempty"`
	Status            ContainerNodePoolStatus `json:"status,omitempty"`
}

type ContainerNodePoolSpecAutoscaling struct {
	MaxNodeCount int `json:"maxNodeCount" tf:"max_node_count"`
	MinNodeCount int `json:"minNodeCount" tf:"min_node_count"`
}

type ContainerNodePoolSpecManagement struct {
	// +optional
	AutoRepair bool `json:"autoRepair,omitempty" tf:"auto_repair,omitempty"`
	// +optional
	AutoUpgrade bool `json:"autoUpgrade,omitempty" tf:"auto_upgrade,omitempty"`
}

type ContainerNodePoolSpecNodeConfigGuestAccelerator struct {
	Count int    `json:"count" tf:"count"`
	Type  string `json:"type" tf:"type"`
}

type ContainerNodePoolSpecNodeConfigTaint struct {
	Effect string `json:"effect" tf:"effect"`
	Key    string `json:"key" tf:"key"`
	Value  string `json:"value" tf:"value"`
}

type ContainerNodePoolSpecNodeConfigWorkloadMetadataConfig struct {
	NodeMetadata string `json:"nodeMetadata" tf:"node_metadata"`
}

type ContainerNodePoolSpecNodeConfig struct {
	// +optional
	DiskSizeGb int `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`
	// +optional
	DiskType string `json:"diskType,omitempty" tf:"disk_type,omitempty"`
	// +optional
	GuestAccelerator []ContainerNodePoolSpecNodeConfigGuestAccelerator `json:"guestAccelerator,omitempty" tf:"guest_accelerator,omitempty"`
	// +optional
	ImageType string `json:"imageType,omitempty" tf:"image_type,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	LocalSsdCount int `json:"localSsdCount,omitempty" tf:"local_ssd_count,omitempty"`
	// +optional
	MachineType string `json:"machineType,omitempty" tf:"machine_type,omitempty"`
	// +optional
	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata,omitempty"`
	// +optional
	MinCPUPlatform string `json:"minCPUPlatform,omitempty" tf:"min_cpu_platform,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
	// +optional
	Preemptible bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`
	// +optional
	ServiceAccount string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// Deprecated
	Taint []ContainerNodePoolSpecNodeConfigTaint `json:"taint,omitempty" tf:"taint,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	WorkloadMetadataConfig []ContainerNodePoolSpecNodeConfigWorkloadMetadataConfig `json:"workloadMetadataConfig,omitempty" tf:"workload_metadata_config,omitempty"`
}

type ContainerNodePoolSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	Autoscaling []ContainerNodePoolSpecAutoscaling `json:"autoscaling,omitempty" tf:"autoscaling,omitempty"`
	Cluster     string                             `json:"cluster" tf:"cluster"`
	// +optional
	InitialNodeCount int `json:"initialNodeCount,omitempty" tf:"initial_node_count,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Management []ContainerNodePoolSpecManagement `json:"management,omitempty" tf:"management,omitempty"`
	// +optional
	// Deprecated
	MaxPodsPerNode int `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	// Deprecated
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NodeConfig []ContainerNodePoolSpecNodeConfig `json:"nodeConfig,omitempty" tf:"node_config,omitempty"`
	// +optional
	NodeCount int `json:"nodeCount,omitempty" tf:"node_count,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
	// +optional
	Zone string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ContainerNodePoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ContainerNodePoolSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ContainerNodePoolList is a list of ContainerNodePools
type ContainerNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ContainerNodePool CRD objects
	Items []ContainerNodePool `json:"items,omitempty"`
}
