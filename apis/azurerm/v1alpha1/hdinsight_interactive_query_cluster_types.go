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

type HdinsightInteractiveQueryCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HdinsightInteractiveQueryClusterSpec   `json:"spec,omitempty"`
	Status            HdinsightInteractiveQueryClusterStatus `json:"status,omitempty"`
}

type HdinsightInteractiveQueryClusterSpecComponentVersion struct {
	InteractiveHive string `json:"interactive_hive"`
}

type HdinsightInteractiveQueryClusterSpecGateway struct {
	Enabled  bool   `json:"enabled"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type HdinsightInteractiveQueryClusterSpecRolesHeadNode struct {
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

type HdinsightInteractiveQueryClusterSpecRolesWorkerNode struct {
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

type HdinsightInteractiveQueryClusterSpecRolesZookeeperNode struct {
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

type HdinsightInteractiveQueryClusterSpecRoles struct {
	// +kubebuilder:validation:MaxItems=1
	HeadNode []HdinsightInteractiveQueryClusterSpecRoles `json:"head_node"`
	// +kubebuilder:validation:MaxItems=1
	WorkerNode []HdinsightInteractiveQueryClusterSpecRoles `json:"worker_node"`
	// +kubebuilder:validation:MaxItems=1
	ZookeeperNode []HdinsightInteractiveQueryClusterSpecRoles `json:"zookeeper_node"`
}

type HdinsightInteractiveQueryClusterSpecStorageAccount struct {
	IsDefault          bool   `json:"is_default"`
	StorageAccountKey  string `json:"storage_account_key"`
	StorageContainerId string `json:"storage_container_id"`
}

type HdinsightInteractiveQueryClusterSpec struct {
	ClusterVersion string `json:"cluster_version"`
	// +kubebuilder:validation:MaxItems=1
	ComponentVersion []HdinsightInteractiveQueryClusterSpec `json:"component_version"`
	// +kubebuilder:validation:MaxItems=1
	Gateway           []HdinsightInteractiveQueryClusterSpec `json:"gateway"`
	Location          string                                 `json:"location"`
	Name              string                                 `json:"name"`
	ResourceGroupName string                                 `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Roles          []HdinsightInteractiveQueryClusterSpec `json:"roles"`
	StorageAccount []HdinsightInteractiveQueryClusterSpec `json:"storage_account"`
	Tier           string                                 `json:"tier"`
}

type HdinsightInteractiveQueryClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// HdinsightInteractiveQueryClusterList is a list of HdinsightInteractiveQueryClusters
type HdinsightInteractiveQueryClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of HdinsightInteractiveQueryCluster CRD objects
	Items []HdinsightInteractiveQueryCluster `json:"items,omitempty"`
}
