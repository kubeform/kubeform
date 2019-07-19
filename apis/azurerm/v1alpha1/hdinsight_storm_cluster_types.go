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

type HdinsightStormCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HdinsightStormClusterSpec   `json:"spec,omitempty"`
	Status            HdinsightStormClusterStatus `json:"status,omitempty"`
}

type HdinsightStormClusterSpecComponentVersion struct {
	Storm string `json:"storm" tf:"storm"`
}

type HdinsightStormClusterSpecGateway struct {
	Enabled bool `json:"enabled" tf:"enabled"`
	// Sensitive Data. Provide secret name which contains one value only
	Password core.LocalObjectReference `json:"password" tf:"password"`
	Username string                    `json:"username" tf:"username"`
}

type HdinsightStormClusterSpecRolesHeadNode struct {
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	Password core.LocalObjectReference `json:"password,omitempty" tf:"password,omitempty"`
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

type HdinsightStormClusterSpecRolesWorkerNode struct {
	// +optional
	MinInstanceCount int `json:"minInstanceCount,omitempty" tf:"min_instance_count,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	Password core.LocalObjectReference `json:"password,omitempty" tf:"password,omitempty"`
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

type HdinsightStormClusterSpecRolesZookeeperNode struct {
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	Password core.LocalObjectReference `json:"password,omitempty" tf:"password,omitempty"`
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

type HdinsightStormClusterSpecRoles struct {
	// +kubebuilder:validation:MaxItems=1
	HeadNode []HdinsightStormClusterSpecRolesHeadNode `json:"headNode" tf:"head_node"`
	// +kubebuilder:validation:MaxItems=1
	WorkerNode []HdinsightStormClusterSpecRolesWorkerNode `json:"workerNode" tf:"worker_node"`
	// +kubebuilder:validation:MaxItems=1
	ZookeeperNode []HdinsightStormClusterSpecRolesZookeeperNode `json:"zookeeperNode" tf:"zookeeper_node"`
}

type HdinsightStormClusterSpecStorageAccount struct {
	IsDefault bool `json:"isDefault" tf:"is_default"`
	// Sensitive Data. Provide secret name which contains one value only
	StorageAccountKey  core.LocalObjectReference `json:"storageAccountKey" tf:"storage_account_key"`
	StorageContainerID string                    `json:"storageContainerID" tf:"storage_container_id"`
}

type HdinsightStormClusterSpec struct {
	ClusterVersion string `json:"clusterVersion" tf:"cluster_version"`
	// +kubebuilder:validation:MaxItems=1
	ComponentVersion []HdinsightStormClusterSpecComponentVersion `json:"componentVersion" tf:"component_version"`
	// +kubebuilder:validation:MaxItems=1
	Gateway           []HdinsightStormClusterSpecGateway `json:"gateway" tf:"gateway"`
	Location          string                             `json:"location" tf:"location"`
	Name              string                             `json:"name" tf:"name"`
	ResourceGroupName string                             `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Roles          []HdinsightStormClusterSpecRoles          `json:"roles" tf:"roles"`
	StorageAccount []HdinsightStormClusterSpecStorageAccount `json:"storageAccount" tf:"storage_account"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	Tier        string                    `json:"tier" tf:"tier"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type HdinsightStormClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
