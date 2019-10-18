package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type DataprocCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataprocClusterSpec   `json:"spec,omitempty"`
	Status            DataprocClusterStatus `json:"status,omitempty"`
}

type DataprocClusterSpecClusterConfigGceClusterConfig struct {
	// +optional
	InternalIPOnly bool `json:"internalIPOnly,omitempty" tf:"internal_ip_only,omitempty"`
	// +optional
	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata,omitempty"`
	// +optional
	Network string `json:"network,omitempty" tf:"network,omitempty"`
	// +optional
	ServiceAccount string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`
	// +optional
	ServiceAccountScopes []string `json:"serviceAccountScopes,omitempty" tf:"service_account_scopes,omitempty"`
	// +optional
	Subnetwork string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Zone string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type DataprocClusterSpecClusterConfigInitializationAction struct {
	Script string `json:"script" tf:"script"`
	// +optional
	TimeoutSec int `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`
}

type DataprocClusterSpecClusterConfigMasterConfigDiskConfig struct {
	// +optional
	BootDiskSizeGb int `json:"bootDiskSizeGb,omitempty" tf:"boot_disk_size_gb,omitempty"`
	// +optional
	BootDiskType string `json:"bootDiskType,omitempty" tf:"boot_disk_type,omitempty"`
	// +optional
	NumLocalSsds int `json:"numLocalSsds,omitempty" tf:"num_local_ssds,omitempty"`
}

type DataprocClusterSpecClusterConfigMasterConfig struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiskConfig []DataprocClusterSpecClusterConfigMasterConfigDiskConfig `json:"diskConfig,omitempty" tf:"disk_config,omitempty"`
	// +optional
	InstanceNames []string `json:"instanceNames,omitempty" tf:"instance_names,omitempty"`
	// +optional
	MachineType string `json:"machineType,omitempty" tf:"machine_type,omitempty"`
	// +optional
	NumInstances int `json:"numInstances,omitempty" tf:"num_instances,omitempty"`
}

type DataprocClusterSpecClusterConfigPreemptibleWorkerConfigDiskConfig struct {
	// +optional
	BootDiskSizeGb int `json:"bootDiskSizeGb,omitempty" tf:"boot_disk_size_gb,omitempty"`
}

type DataprocClusterSpecClusterConfigPreemptibleWorkerConfig struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiskConfig []DataprocClusterSpecClusterConfigPreemptibleWorkerConfigDiskConfig `json:"diskConfig,omitempty" tf:"disk_config,omitempty"`
	// +optional
	InstanceNames []string `json:"instanceNames,omitempty" tf:"instance_names,omitempty"`
	// +optional
	NumInstances int `json:"numInstances,omitempty" tf:"num_instances,omitempty"`
}

type DataprocClusterSpecClusterConfigSoftwareConfig struct {
	// +optional
	ImageVersion string `json:"imageVersion,omitempty" tf:"image_version,omitempty"`
	// +optional
	OverrideProperties map[string]string `json:"overrideProperties,omitempty" tf:"override_properties,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
}

type DataprocClusterSpecClusterConfigWorkerConfigDiskConfig struct {
	// +optional
	BootDiskSizeGb int `json:"bootDiskSizeGb,omitempty" tf:"boot_disk_size_gb,omitempty"`
	// +optional
	BootDiskType string `json:"bootDiskType,omitempty" tf:"boot_disk_type,omitempty"`
	// +optional
	NumLocalSsds int `json:"numLocalSsds,omitempty" tf:"num_local_ssds,omitempty"`
}

type DataprocClusterSpecClusterConfigWorkerConfig struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiskConfig []DataprocClusterSpecClusterConfigWorkerConfigDiskConfig `json:"diskConfig,omitempty" tf:"disk_config,omitempty"`
	// +optional
	InstanceNames []string `json:"instanceNames,omitempty" tf:"instance_names,omitempty"`
	// +optional
	MachineType string `json:"machineType,omitempty" tf:"machine_type,omitempty"`
	// +optional
	NumInstances int `json:"numInstances,omitempty" tf:"num_instances,omitempty"`
}

type DataprocClusterSpecClusterConfig struct {
	// +optional
	Bucket string `json:"bucket,omitempty" tf:"bucket,omitempty"`
	// +optional
	// Deprecated
	DeleteAutogenBucket bool `json:"deleteAutogenBucket,omitempty" tf:"delete_autogen_bucket,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	GceClusterConfig []DataprocClusterSpecClusterConfigGceClusterConfig `json:"gceClusterConfig,omitempty" tf:"gce_cluster_config,omitempty"`
	// +optional
	InitializationAction []DataprocClusterSpecClusterConfigInitializationAction `json:"initializationAction,omitempty" tf:"initialization_action,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MasterConfig []DataprocClusterSpecClusterConfigMasterConfig `json:"masterConfig,omitempty" tf:"master_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PreemptibleWorkerConfig []DataprocClusterSpecClusterConfigPreemptibleWorkerConfig `json:"preemptibleWorkerConfig,omitempty" tf:"preemptible_worker_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SoftwareConfig []DataprocClusterSpecClusterConfigSoftwareConfig `json:"softwareConfig,omitempty" tf:"software_config,omitempty"`
	// +optional
	StagingBucket string `json:"stagingBucket,omitempty" tf:"staging_bucket,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	WorkerConfig []DataprocClusterSpecClusterConfigWorkerConfig `json:"workerConfig,omitempty" tf:"worker_config,omitempty"`
}

type DataprocClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	ClusterConfig []DataprocClusterSpecClusterConfig `json:"clusterConfig,omitempty" tf:"cluster_config,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	Name   string            `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
}

type DataprocClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DataprocClusterSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataprocClusterList is a list of DataprocClusters
type DataprocClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataprocCluster CRD objects
	Items []DataprocCluster `json:"items,omitempty"`
}
