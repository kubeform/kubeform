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

type AzurermHdinsightStormClusterSpecComponentVersion struct {
	Storm string `json:"storm"`
}

type AzurermHdinsightStormClusterSpecGateway struct {
	Password string `json:"password"`
	Enabled  bool   `json:"enabled"`
	Username string `json:"username"`
}

type AzurermHdinsightStormClusterSpecStorageAccount struct {
	IsDefault          bool   `json:"is_default"`
	StorageAccountKey  string `json:"storage_account_key"`
	StorageContainerId string `json:"storage_container_id"`
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
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
}

type AzurermHdinsightStormClusterSpecRolesWorkerNode struct {
	SshKeys             []string `json:"ssh_keys"`
	SubnetId            string   `json:"subnet_id"`
	VirtualNetworkId    string   `json:"virtual_network_id"`
	MinInstanceCount    int      `json:"min_instance_count"`
	TargetInstanceCount int      `json:"target_instance_count"`
	VmSize              string   `json:"vm_size"`
	Username            string   `json:"username"`
	Password            string   `json:"password"`
}

type AzurermHdinsightStormClusterSpecRoles struct {
	ZookeeperNode []AzurermHdinsightStormClusterSpecRoles `json:"zookeeper_node"`
	HeadNode      []AzurermHdinsightStormClusterSpecRoles `json:"head_node"`
	WorkerNode    []AzurermHdinsightStormClusterSpecRoles `json:"worker_node"`
}

type AzurermHdinsightStormClusterSpec struct {
	Location          string                             `json:"location"`
	ClusterVersion    string                             `json:"cluster_version"`
	Tier              string                             `json:"tier"`
	ComponentVersion  []AzurermHdinsightStormClusterSpec `json:"component_version"`
	Gateway           []AzurermHdinsightStormClusterSpec `json:"gateway"`
	StorageAccount    []AzurermHdinsightStormClusterSpec `json:"storage_account"`
	Roles             []AzurermHdinsightStormClusterSpec `json:"roles"`
	ResourceGroupName string                             `json:"resource_group_name"`
	HttpsEndpoint     string                             `json:"https_endpoint"`
	Tags              map[string]string                  `json:"tags"`
	SshEndpoint       string                             `json:"ssh_endpoint"`
	Name              string                             `json:"name"`
}

type AzurermHdinsightStormClusterStatus struct {
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
