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

type AzurermHdinsightStormCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermHdinsightStormClusterSpec   `json:"spec,omitempty"`
	Status            AzurermHdinsightStormClusterStatus `json:"status,omitempty"`
}

type AzurermHdinsightStormClusterSpecGateway struct {
	Password string `json:"password"`
	Enabled  bool   `json:"enabled"`
	Username string `json:"username"`
}

type AzurermHdinsightStormClusterSpecRolesWorkerNode struct {
	MinInstanceCount    int      `json:"min_instance_count"`
	TargetInstanceCount int      `json:"target_instance_count"`
	VmSize              string   `json:"vm_size"`
	Username            string   `json:"username"`
	Password            string   `json:"password"`
	SshKeys             []string `json:"ssh_keys"`
	SubnetId            string   `json:"subnet_id"`
	VirtualNetworkId    string   `json:"virtual_network_id"`
}

type AzurermHdinsightStormClusterSpecRolesZookeeperNode struct {
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
}

type AzurermHdinsightStormClusterSpecRolesHeadNode struct {
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
}

type AzurermHdinsightStormClusterSpecRoles struct {
	WorkerNode    []AzurermHdinsightStormClusterSpecRoles `json:"worker_node"`
	ZookeeperNode []AzurermHdinsightStormClusterSpecRoles `json:"zookeeper_node"`
	HeadNode      []AzurermHdinsightStormClusterSpecRoles `json:"head_node"`
}

type AzurermHdinsightStormClusterSpecComponentVersion struct {
	Storm string `json:"storm"`
}

type AzurermHdinsightStormClusterSpecStorageAccount struct {
	StorageAccountKey  string `json:"storage_account_key"`
	StorageContainerId string `json:"storage_container_id"`
	IsDefault          bool   `json:"is_default"`
}

type AzurermHdinsightStormClusterSpec struct {
	ResourceGroupName string                             `json:"resource_group_name"`
	Location          string                             `json:"location"`
	Tier              string                             `json:"tier"`
	Gateway           []AzurermHdinsightStormClusterSpec `json:"gateway"`
	Roles             []AzurermHdinsightStormClusterSpec `json:"roles"`
	Tags              map[string]string                  `json:"tags"`
	HttpsEndpoint     string                             `json:"https_endpoint"`
	SshEndpoint       string                             `json:"ssh_endpoint"`
	Name              string                             `json:"name"`
	ClusterVersion    string                             `json:"cluster_version"`
	ComponentVersion  []AzurermHdinsightStormClusterSpec `json:"component_version"`
	StorageAccount    []AzurermHdinsightStormClusterSpec `json:"storage_account"`
}



type AzurermHdinsightStormClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermHdinsightStormClusterList is a list of AzurermHdinsightStormClusters
type AzurermHdinsightStormClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermHdinsightStormCluster CRD objects
	Items []AzurermHdinsightStormCluster `json:"items,omitempty"`
}