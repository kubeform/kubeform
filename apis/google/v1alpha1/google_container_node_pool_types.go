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

type GoogleContainerNodePool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleContainerNodePoolSpec   `json:"spec,omitempty"`
	Status            GoogleContainerNodePoolStatus `json:"status,omitempty"`
}

type GoogleContainerNodePoolSpecNodeConfigTaint struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
}

type GoogleContainerNodePoolSpecNodeConfigGuestAccelerator struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type GoogleContainerNodePoolSpecNodeConfigWorkloadMetadataConfig struct {
	NodeMetadata string `json:"node_metadata"`
}

type GoogleContainerNodePoolSpecNodeConfig struct {
	Taint                  []GoogleContainerNodePoolSpecNodeConfig `json:"taint"`
	Metadata               map[string]string                       `json:"metadata"`
	Tags                   []string                                `json:"tags"`
	GuestAccelerator       []GoogleContainerNodePoolSpecNodeConfig `json:"guest_accelerator"`
	LocalSsdCount          int                                     `json:"local_ssd_count"`
	Labels                 map[string]string                       `json:"labels"`
	MinCpuPlatform         string                                  `json:"min_cpu_platform"`
	Preemptible            bool                                    `json:"preemptible"`
	ServiceAccount         string                                  `json:"service_account"`
	DiskSizeGb             int                                     `json:"disk_size_gb"`
	ImageType              string                                  `json:"image_type"`
	OauthScopes            []string                                `json:"oauth_scopes"`
	WorkloadMetadataConfig []GoogleContainerNodePoolSpecNodeConfig `json:"workload_metadata_config"`
	DiskType               string                                  `json:"disk_type"`
	MachineType            string                                  `json:"machine_type"`
}

type GoogleContainerNodePoolSpecManagement struct {
	AutoRepair  bool `json:"auto_repair"`
	AutoUpgrade bool `json:"auto_upgrade"`
}

type GoogleContainerNodePoolSpecAutoscaling struct {
	MinNodeCount int `json:"min_node_count"`
	MaxNodeCount int `json:"max_node_count"`
}

type GoogleContainerNodePoolSpec struct {
	Region            string                        `json:"region"`
	MaxPodsPerNode    int                           `json:"max_pods_per_node"`
	Name              string                        `json:"name"`
	NodeConfig        []GoogleContainerNodePoolSpec `json:"node_config"`
	InstanceGroupUrls []string                      `json:"instance_group_urls"`
	Management        []GoogleContainerNodePoolSpec `json:"management"`
	Autoscaling       []GoogleContainerNodePoolSpec `json:"autoscaling"`
	Zone              string                        `json:"zone"`
	NamePrefix        string                        `json:"name_prefix"`
	Version           string                        `json:"version"`
	InitialNodeCount  int                           `json:"initial_node_count"`
	NodeCount         int                           `json:"node_count"`
	Project           string                        `json:"project"`
	Cluster           string                        `json:"cluster"`
}



type GoogleContainerNodePoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleContainerNodePoolList is a list of GoogleContainerNodePools
type GoogleContainerNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleContainerNodePool CRD objects
	Items []GoogleContainerNodePool `json:"items,omitempty"`
}