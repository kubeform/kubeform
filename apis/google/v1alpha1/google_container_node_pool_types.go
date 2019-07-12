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

type GoogleContainerNodePoolSpecAutoscaling struct {
	MinNodeCount int `json:"min_node_count"`
	MaxNodeCount int `json:"max_node_count"`
}

type GoogleContainerNodePoolSpecManagement struct {
	AutoRepair  bool `json:"auto_repair"`
	AutoUpgrade bool `json:"auto_upgrade"`
}

type GoogleContainerNodePoolSpecNodeConfigWorkloadMetadataConfig struct {
	NodeMetadata string `json:"node_metadata"`
}

type GoogleContainerNodePoolSpecNodeConfigGuestAccelerator struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}

type GoogleContainerNodePoolSpecNodeConfigTaint struct {
	Value  string `json:"value"`
	Effect string `json:"effect"`
	Key    string `json:"key"`
}

type GoogleContainerNodePoolSpecNodeConfig struct {
	DiskType               string                                  `json:"disk_type"`
	Metadata               map[string]string                       `json:"metadata"`
	MinCpuPlatform         string                                  `json:"min_cpu_platform"`
	Tags                   []string                                `json:"tags"`
	Labels                 map[string]string                       `json:"labels"`
	OauthScopes            []string                                `json:"oauth_scopes"`
	ServiceAccount         string                                  `json:"service_account"`
	WorkloadMetadataConfig []GoogleContainerNodePoolSpecNodeConfig `json:"workload_metadata_config"`
	DiskSizeGb             int                                     `json:"disk_size_gb"`
	GuestAccelerator       []GoogleContainerNodePoolSpecNodeConfig `json:"guest_accelerator"`
	ImageType              string                                  `json:"image_type"`
	MachineType            string                                  `json:"machine_type"`
	LocalSsdCount          int                                     `json:"local_ssd_count"`
	Preemptible            bool                                    `json:"preemptible"`
	Taint                  []GoogleContainerNodePoolSpecNodeConfig `json:"taint"`
}

type GoogleContainerNodePoolSpec struct {
	Region            string                        `json:"region"`
	Zone              string                        `json:"zone"`
	InitialNodeCount  int                           `json:"initial_node_count"`
	InstanceGroupUrls []string                      `json:"instance_group_urls"`
	Cluster           string                        `json:"cluster"`
	Name              string                        `json:"name"`
	Autoscaling       []GoogleContainerNodePoolSpec `json:"autoscaling"`
	NamePrefix        string                        `json:"name_prefix"`
	MaxPodsPerNode    int                           `json:"max_pods_per_node"`
	Management        []GoogleContainerNodePoolSpec `json:"management"`
	Project           string                        `json:"project"`
	NodeConfig        []GoogleContainerNodePoolSpec `json:"node_config"`
	NodeCount         int                           `json:"node_count"`
	Version           string                        `json:"version"`
}

type GoogleContainerNodePoolStatus struct {
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
