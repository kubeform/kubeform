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

type HdinsightMlServicesCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HdinsightMlServicesClusterSpec   `json:"spec,omitempty"`
	Status            HdinsightMlServicesClusterStatus `json:"status,omitempty"`
}

type HdinsightMlServicesClusterSpecGateway struct {
	Enabled  bool   `json:"enabled"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type HdinsightMlServicesClusterSpecRolesEdgeNode struct {
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

type HdinsightMlServicesClusterSpecRolesHeadNode struct {
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

type HdinsightMlServicesClusterSpecRolesWorkerNode struct {
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

type HdinsightMlServicesClusterSpecRolesZookeeperNode struct {
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

type HdinsightMlServicesClusterSpecRoles struct {
	// +kubebuilder:validation:MaxItems=1
	EdgeNode []HdinsightMlServicesClusterSpecRoles `json:"edge_node"`
	// +kubebuilder:validation:MaxItems=1
	HeadNode []HdinsightMlServicesClusterSpecRoles `json:"head_node"`
	// +kubebuilder:validation:MaxItems=1
	WorkerNode []HdinsightMlServicesClusterSpecRoles `json:"worker_node"`
	// +kubebuilder:validation:MaxItems=1
	ZookeeperNode []HdinsightMlServicesClusterSpecRoles `json:"zookeeper_node"`
}

type HdinsightMlServicesClusterSpecStorageAccount struct {
	IsDefault          bool   `json:"is_default"`
	StorageAccountKey  string `json:"storage_account_key"`
	StorageContainerId string `json:"storage_container_id"`
}

type HdinsightMlServicesClusterSpec struct {
	ClusterVersion string `json:"cluster_version"`
	// +kubebuilder:validation:MaxItems=1
	Gateway           []HdinsightMlServicesClusterSpec `json:"gateway"`
	Location          string                           `json:"location"`
	Name              string                           `json:"name"`
	ResourceGroupName string                           `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Roles          []HdinsightMlServicesClusterSpec `json:"roles"`
	Rstudio        bool                             `json:"rstudio"`
	StorageAccount []HdinsightMlServicesClusterSpec `json:"storage_account"`
	Tier           string                           `json:"tier"`
}

type HdinsightMlServicesClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// HdinsightMlServicesClusterList is a list of HdinsightMlServicesClusters
type HdinsightMlServicesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of HdinsightMlServicesCluster CRD objects
	Items []HdinsightMlServicesCluster `json:"items,omitempty"`
}
