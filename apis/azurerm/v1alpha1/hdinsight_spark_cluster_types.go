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

type HdinsightSparkCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HdinsightSparkClusterSpec   `json:"spec,omitempty"`
	Status            HdinsightSparkClusterStatus `json:"status,omitempty"`
}

type HdinsightSparkClusterSpecComponentVersion struct {
	Spark string `json:"spark" tf:"spark"`
}

type HdinsightSparkClusterSpecGateway struct {
	Enabled  bool   `json:"enabled" tf:"enabled"`
	Username string `json:"username" tf:"username"`
}

type HdinsightSparkClusterSpecRolesHeadNode struct {
	// +optional
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	Username string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id,omitempty"`
	VmSize           string `json:"vmSize" tf:"vm_size"`
}

type HdinsightSparkClusterSpecRolesWorkerNode struct {
	// +optional
	MinInstanceCount int `json:"minInstanceCount,omitempty" tf:"min_instance_count,omitempty"`
	// +optional
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	SubnetID            string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	TargetInstanceCount int    `json:"targetInstanceCount" tf:"target_instance_count"`
	Username            string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id,omitempty"`
	VmSize           string `json:"vmSize" tf:"vm_size"`
}

type HdinsightSparkClusterSpecRolesZookeeperNode struct {
	// +optional
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	Username string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id,omitempty"`
	VmSize           string `json:"vmSize" tf:"vm_size"`
}

type HdinsightSparkClusterSpecRoles struct {
	// +kubebuilder:validation:MaxItems=1
	HeadNode []HdinsightSparkClusterSpecRolesHeadNode `json:"headNode" tf:"head_node"`
	// +kubebuilder:validation:MaxItems=1
	WorkerNode []HdinsightSparkClusterSpecRolesWorkerNode `json:"workerNode" tf:"worker_node"`
	// +kubebuilder:validation:MaxItems=1
	ZookeeperNode []HdinsightSparkClusterSpecRolesZookeeperNode `json:"zookeeperNode" tf:"zookeeper_node"`
}

type HdinsightSparkClusterSpecStorageAccount struct {
	IsDefault          bool   `json:"isDefault" tf:"is_default"`
	StorageContainerID string `json:"storageContainerID" tf:"storage_container_id"`
}

type HdinsightSparkClusterSpec struct {
	Secret         *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`
	ProviderRef    core.LocalObjectReference  `json:"providerRef" tf:"-"`
	ClusterVersion string                     `json:"clusterVersion" tf:"cluster_version"`
	// +kubebuilder:validation:MaxItems=1
	ComponentVersion []HdinsightSparkClusterSpecComponentVersion `json:"componentVersion" tf:"component_version"`
	// +kubebuilder:validation:MaxItems=1
	Gateway           []HdinsightSparkClusterSpecGateway `json:"gateway" tf:"gateway"`
	Location          string                             `json:"location" tf:"location"`
	Name              string                             `json:"name" tf:"name"`
	ResourceGroupName string                             `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Roles          []HdinsightSparkClusterSpecRoles          `json:"roles" tf:"roles"`
	StorageAccount []HdinsightSparkClusterSpecStorageAccount `json:"storageAccount" tf:"storage_account"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Tier string            `json:"tier" tf:"tier"`
}

type HdinsightSparkClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// HdinsightSparkClusterList is a list of HdinsightSparkClusters
type HdinsightSparkClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of HdinsightSparkCluster CRD objects
	Items []HdinsightSparkCluster `json:"items,omitempty"`
}
