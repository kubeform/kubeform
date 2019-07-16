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

type HdinsightHbaseCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HdinsightHbaseClusterSpec   `json:"spec,omitempty"`
	Status            HdinsightHbaseClusterStatus `json:"status,omitempty"`
}

type HdinsightHbaseClusterSpecComponentVersion struct {
	Hbase string `json:"hbase"`
}

type HdinsightHbaseClusterSpecGateway struct {
	Enabled  bool   `json:"enabled"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type HdinsightHbaseClusterSpecRolesHeadNode struct {
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

type HdinsightHbaseClusterSpecRolesWorkerNode struct {
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

type HdinsightHbaseClusterSpecRolesZookeeperNode struct {
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

type HdinsightHbaseClusterSpecRoles struct {
	// +kubebuilder:validation:MaxItems=1
	HeadNode []HdinsightHbaseClusterSpecRoles `json:"head_node"`
	// +kubebuilder:validation:MaxItems=1
	WorkerNode []HdinsightHbaseClusterSpecRoles `json:"worker_node"`
	// +kubebuilder:validation:MaxItems=1
	ZookeeperNode []HdinsightHbaseClusterSpecRoles `json:"zookeeper_node"`
}

type HdinsightHbaseClusterSpecStorageAccount struct {
	IsDefault          bool   `json:"is_default"`
	StorageAccountKey  string `json:"storage_account_key"`
	StorageContainerId string `json:"storage_container_id"`
}

type HdinsightHbaseClusterSpec struct {
	ClusterVersion string `json:"cluster_version"`
	// +kubebuilder:validation:MaxItems=1
	ComponentVersion []HdinsightHbaseClusterSpec `json:"component_version"`
	// +kubebuilder:validation:MaxItems=1
	Gateway           []HdinsightHbaseClusterSpec `json:"gateway"`
	Location          string                      `json:"location"`
	Name              string                      `json:"name"`
	ResourceGroupName string                      `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Roles          []HdinsightHbaseClusterSpec `json:"roles"`
	StorageAccount []HdinsightHbaseClusterSpec `json:"storage_account"`
	Tier           string                      `json:"tier"`
}

type HdinsightHbaseClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// HdinsightHbaseClusterList is a list of HdinsightHbaseClusters
type HdinsightHbaseClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of HdinsightHbaseCluster CRD objects
	Items []HdinsightHbaseCluster `json:"items,omitempty"`
}
