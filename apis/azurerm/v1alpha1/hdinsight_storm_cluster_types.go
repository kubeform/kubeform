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

type HdinsightStormCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HdinsightStormClusterSpec   `json:"spec,omitempty"`
	Status            HdinsightStormClusterStatus `json:"status,omitempty"`
}

type HdinsightStormClusterSpecComponentVersion struct {
	Storm string `json:"storm"`
}

type HdinsightStormClusterSpecGateway struct {
	Enabled  bool   `json:"enabled"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type HdinsightStormClusterSpecRolesHeadNode struct {
	// +optional
	Password string `json:"password,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"ssh_keys,omitempty"`
	// +optional
	SubnetId string `json:"subnet_id,omitempty"`
	Username string `json:"username"`
	// +optional
	VirtualNetworkId string `json:"virtual_network_id,omitempty"`
	VmSize           string `json:"vm_size"`
}

type HdinsightStormClusterSpecRolesWorkerNode struct {
	// +optional
	MinInstanceCount int `json:"min_instance_count,omitempty"`
	// +optional
	Password string `json:"password,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"ssh_keys,omitempty"`
	// +optional
	SubnetId            string `json:"subnet_id,omitempty"`
	TargetInstanceCount int    `json:"target_instance_count"`
	Username            string `json:"username"`
	// +optional
	VirtualNetworkId string `json:"virtual_network_id,omitempty"`
	VmSize           string `json:"vm_size"`
}

type HdinsightStormClusterSpecRolesZookeeperNode struct {
	// +optional
	Password string `json:"password,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"ssh_keys,omitempty"`
	// +optional
	SubnetId string `json:"subnet_id,omitempty"`
	Username string `json:"username"`
	// +optional
	VirtualNetworkId string `json:"virtual_network_id,omitempty"`
	VmSize           string `json:"vm_size"`
}

type HdinsightStormClusterSpecRoles struct {
	// +kubebuilder:validation:MaxItems=1
	HeadNode []HdinsightStormClusterSpecRoles `json:"head_node"`
	// +kubebuilder:validation:MaxItems=1
	WorkerNode []HdinsightStormClusterSpecRoles `json:"worker_node"`
	// +kubebuilder:validation:MaxItems=1
	ZookeeperNode []HdinsightStormClusterSpecRoles `json:"zookeeper_node"`
}

type HdinsightStormClusterSpecStorageAccount struct {
	IsDefault          bool   `json:"is_default"`
	StorageAccountKey  string `json:"storage_account_key"`
	StorageContainerId string `json:"storage_container_id"`
}

type HdinsightStormClusterSpec struct {
	ClusterVersion string `json:"cluster_version"`
	// +kubebuilder:validation:MaxItems=1
	ComponentVersion []HdinsightStormClusterSpec `json:"component_version"`
	// +kubebuilder:validation:MaxItems=1
	Gateway           []HdinsightStormClusterSpec `json:"gateway"`
	Location          string                      `json:"location"`
	Name              string                      `json:"name"`
	ResourceGroupName string                      `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Roles          []HdinsightStormClusterSpec `json:"roles"`
	StorageAccount []HdinsightStormClusterSpec `json:"storage_account"`
	Tier           string                      `json:"tier"`
}

type HdinsightStormClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// HdinsightStormClusterList is a list of HdinsightStormClusters
type HdinsightStormClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of HdinsightStormCluster CRD objects
	Items []HdinsightStormCluster `json:"items,omitempty"`
}
