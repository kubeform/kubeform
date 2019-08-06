package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	InteractiveHive string `json:"interactiveHive" tf:"interactive_hive"`
}

type HdinsightInteractiveQueryClusterSpecGateway struct {
	Enabled  bool   `json:"enabled" tf:"enabled"`
	Password string `json:"-" sensitive:"true" tf:"password"`
	Username string `json:"username" tf:"username"`
}

type HdinsightInteractiveQueryClusterSpecRolesHeadNode struct {
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
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

type HdinsightInteractiveQueryClusterSpecRolesWorkerNode struct {
	// +optional
	MinInstanceCount int `json:"minInstanceCount,omitempty" tf:"min_instance_count,omitempty"`
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
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

type HdinsightInteractiveQueryClusterSpecRolesZookeeperNode struct {
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
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

type HdinsightInteractiveQueryClusterSpecRoles struct {
	// +kubebuilder:validation:MaxItems=1
	HeadNode []HdinsightInteractiveQueryClusterSpecRolesHeadNode `json:"headNode" tf:"head_node"`
	// +kubebuilder:validation:MaxItems=1
	WorkerNode []HdinsightInteractiveQueryClusterSpecRolesWorkerNode `json:"workerNode" tf:"worker_node"`
	// +kubebuilder:validation:MaxItems=1
	ZookeeperNode []HdinsightInteractiveQueryClusterSpecRolesZookeeperNode `json:"zookeeperNode" tf:"zookeeper_node"`
}

type HdinsightInteractiveQueryClusterSpecStorageAccount struct {
	IsDefault          bool   `json:"isDefault" tf:"is_default"`
	StorageAccountKey  string `json:"-" sensitive:"true" tf:"storage_account_key"`
	StorageContainerID string `json:"storageContainerID" tf:"storage_container_id"`
}

type HdinsightInteractiveQueryClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	KubeFormSecret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	ClusterVersion string `json:"clusterVersion" tf:"cluster_version"`
	// +kubebuilder:validation:MaxItems=1
	ComponentVersion []HdinsightInteractiveQueryClusterSpecComponentVersion `json:"componentVersion" tf:"component_version"`
	// +kubebuilder:validation:MaxItems=1
	Gateway []HdinsightInteractiveQueryClusterSpecGateway `json:"gateway" tf:"gateway"`
	// +optional
	HttpsEndpoint     string `json:"httpsEndpoint,omitempty" tf:"https_endpoint,omitempty"`
	Location          string `json:"location" tf:"location"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Roles []HdinsightInteractiveQueryClusterSpecRoles `json:"roles" tf:"roles"`
	// +optional
	SshEndpoint    string                                               `json:"sshEndpoint,omitempty" tf:"ssh_endpoint,omitempty"`
	StorageAccount []HdinsightInteractiveQueryClusterSpecStorageAccount `json:"storageAccount" tf:"storage_account"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Tier string            `json:"tier" tf:"tier"`
}

type HdinsightInteractiveQueryClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *HdinsightInteractiveQueryClusterSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
