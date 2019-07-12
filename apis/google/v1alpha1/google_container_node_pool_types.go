package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleContainerNodePool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleContainerNodePoolSpec   `json:"spec,omitempty"`
	Status            GoogleContainerNodePoolStatus `json:"status,omitempty"`
}

type GoogleContainerNodePoolSpecManagement struct {
	AutoRepair  bool `json:"auto_repair"`
	AutoUpgrade bool `json:"auto_upgrade"`
}

type GoogleContainerNodePoolSpecNodeConfigTaint struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
}

type GoogleContainerNodePoolSpecNodeConfigWorkloadMetadataConfig struct {
	NodeMetadata string `json:"node_metadata"`
}

type GoogleContainerNodePoolSpecNodeConfigGuestAccelerator struct {
	Count int    `json:"count"`
	Type  string `json:"type"`
}

type GoogleContainerNodePoolSpecNodeConfig struct {
	MinCpuPlatform         string                                  `json:"min_cpu_platform"`
	OauthScopes            []string                                `json:"oauth_scopes"`
	Taint                  []GoogleContainerNodePoolSpecNodeConfig `json:"taint"`
	WorkloadMetadataConfig []GoogleContainerNodePoolSpecNodeConfig `json:"workload_metadata_config"`
	GuestAccelerator       []GoogleContainerNodePoolSpecNodeConfig `json:"guest_accelerator"`
	ImageType              string                                  `json:"image_type"`
	Preemptible            bool                                    `json:"preemptible"`
	Tags                   []string                                `json:"tags"`
	DiskSizeGb             int                                     `json:"disk_size_gb"`
	ServiceAccount         string                                  `json:"service_account"`
	Metadata               map[string]string                       `json:"metadata"`
	DiskType               string                                  `json:"disk_type"`
	Labels                 map[string]string                       `json:"labels"`
	LocalSsdCount          int                                     `json:"local_ssd_count"`
	MachineType            string                                  `json:"machine_type"`
}

type GoogleContainerNodePoolSpecAutoscaling struct {
	MaxNodeCount int `json:"max_node_count"`
	MinNodeCount int `json:"min_node_count"`
}

type GoogleContainerNodePoolSpec struct {
	InitialNodeCount  int                           `json:"initial_node_count"`
	Management        []GoogleContainerNodePoolSpec `json:"management"`
	Cluster           string                        `json:"cluster"`
	Region            string                        `json:"region"`
	Zone              string                        `json:"zone"`
	NamePrefix        string                        `json:"name_prefix"`
	NodeConfig        []GoogleContainerNodePoolSpec `json:"node_config"`
	NodeCount         int                           `json:"node_count"`
	Autoscaling       []GoogleContainerNodePoolSpec `json:"autoscaling"`
	Version           string                        `json:"version"`
	Project           string                        `json:"project"`
	MaxPodsPerNode    int                           `json:"max_pods_per_node"`
	Name              string                        `json:"name"`
	InstanceGroupUrls []string                      `json:"instance_group_urls"`
}



type GoogleContainerNodePoolStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleContainerNodePoolList is a list of GoogleContainerNodePools
type GoogleContainerNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleContainerNodePool CRD objects
	Items []GoogleContainerNodePool `json:"items,omitempty"`
}