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

type GoogleDataprocCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleDataprocClusterSpec   `json:"spec,omitempty"`
	Status            GoogleDataprocClusterStatus `json:"status,omitempty"`
}

type GoogleDataprocClusterSpecClusterConfigWorkerConfigDiskConfig struct {
	NumLocalSsds   int    `json:"num_local_ssds"`
	BootDiskSizeGb int    `json:"boot_disk_size_gb"`
	BootDiskType   string `json:"boot_disk_type"`
}

type GoogleDataprocClusterSpecClusterConfigWorkerConfig struct {
	NumInstances  int                                                  `json:"num_instances"`
	MachineType   string                                               `json:"machine_type"`
	DiskConfig    []GoogleDataprocClusterSpecClusterConfigWorkerConfig `json:"disk_config"`
	InstanceNames []string                                             `json:"instance_names"`
}

type GoogleDataprocClusterSpecClusterConfigPreemptibleWorkerConfigDiskConfig struct {
	BootDiskSizeGb int `json:"boot_disk_size_gb"`
}

type GoogleDataprocClusterSpecClusterConfigPreemptibleWorkerConfig struct {
	NumInstances  int                                                             `json:"num_instances"`
	DiskConfig    []GoogleDataprocClusterSpecClusterConfigPreemptibleWorkerConfig `json:"disk_config"`
	InstanceNames []string                                                        `json:"instance_names"`
}

type GoogleDataprocClusterSpecClusterConfigSoftwareConfig struct {
	ImageVersion       string            `json:"image_version"`
	OverrideProperties map[string]string `json:"override_properties"`
	Properties         map[string]string `json:"properties"`
}

type GoogleDataprocClusterSpecClusterConfigInitializationAction struct {
	Script     string `json:"script"`
	TimeoutSec int    `json:"timeout_sec"`
}

type GoogleDataprocClusterSpecClusterConfigGceClusterConfig struct {
	Metadata             map[string]string `json:"metadata"`
	Zone                 string            `json:"zone"`
	Network              string            `json:"network"`
	Subnetwork           string            `json:"subnetwork"`
	Tags                 []string          `json:"tags"`
	ServiceAccount       string            `json:"service_account"`
	ServiceAccountScopes []string          `json:"service_account_scopes"`
	InternalIpOnly       bool              `json:"internal_ip_only"`
}

type GoogleDataprocClusterSpecClusterConfigMasterConfigDiskConfig struct {
	NumLocalSsds   int    `json:"num_local_ssds"`
	BootDiskSizeGb int    `json:"boot_disk_size_gb"`
	BootDiskType   string `json:"boot_disk_type"`
}

type GoogleDataprocClusterSpecClusterConfigMasterConfig struct {
	NumInstances  int                                                  `json:"num_instances"`
	MachineType   string                                               `json:"machine_type"`
	DiskConfig    []GoogleDataprocClusterSpecClusterConfigMasterConfig `json:"disk_config"`
	InstanceNames []string                                             `json:"instance_names"`
}

type GoogleDataprocClusterSpecClusterConfig struct {
	DeleteAutogenBucket     bool                                     `json:"delete_autogen_bucket"`
	StagingBucket           string                                   `json:"staging_bucket"`
	WorkerConfig            []GoogleDataprocClusterSpecClusterConfig `json:"worker_config"`
	PreemptibleWorkerConfig []GoogleDataprocClusterSpecClusterConfig `json:"preemptible_worker_config"`
	SoftwareConfig          []GoogleDataprocClusterSpecClusterConfig `json:"software_config"`
	InitializationAction    []GoogleDataprocClusterSpecClusterConfig `json:"initialization_action"`
	Bucket                  string                                   `json:"bucket"`
	GceClusterConfig        []GoogleDataprocClusterSpecClusterConfig `json:"gce_cluster_config"`
	MasterConfig            []GoogleDataprocClusterSpecClusterConfig `json:"master_config"`
}

type GoogleDataprocClusterSpec struct {
	Name          string                      `json:"name"`
	Project       string                      `json:"project"`
	Region        string                      `json:"region"`
	Labels        map[string]string           `json:"labels"`
	ClusterConfig []GoogleDataprocClusterSpec `json:"cluster_config"`
}



type GoogleDataprocClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleDataprocClusterList is a list of GoogleDataprocClusters
type GoogleDataprocClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleDataprocCluster CRD objects
	Items []GoogleDataprocCluster `json:"items,omitempty"`
}