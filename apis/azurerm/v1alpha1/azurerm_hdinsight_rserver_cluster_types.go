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

type AzurermHdinsightRserverCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermHdinsightRserverClusterSpec   `json:"spec,omitempty"`
	Status            AzurermHdinsightRserverClusterStatus `json:"status,omitempty"`
}

type AzurermHdinsightRserverClusterSpecGateway struct {
	Enabled  bool   `json:"enabled"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermHdinsightRserverClusterSpecRolesHeadNode struct {
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
}

type AzurermHdinsightRserverClusterSpecRolesWorkerNode struct {
	Password            string   `json:"password"`
	SshKeys             []string `json:"ssh_keys"`
	SubnetId            string   `json:"subnet_id"`
	VirtualNetworkId    string   `json:"virtual_network_id"`
	MinInstanceCount    int      `json:"min_instance_count"`
	TargetInstanceCount int      `json:"target_instance_count"`
	VmSize              string   `json:"vm_size"`
	Username            string   `json:"username"`
}

type AzurermHdinsightRserverClusterSpecRolesZookeeperNode struct {
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
}

type AzurermHdinsightRserverClusterSpecRolesEdgeNode struct {
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
}

type AzurermHdinsightRserverClusterSpecRoles struct {
	HeadNode      []AzurermHdinsightRserverClusterSpecRoles `json:"head_node"`
	WorkerNode    []AzurermHdinsightRserverClusterSpecRoles `json:"worker_node"`
	ZookeeperNode []AzurermHdinsightRserverClusterSpecRoles `json:"zookeeper_node"`
	EdgeNode      []AzurermHdinsightRserverClusterSpecRoles `json:"edge_node"`
}

type AzurermHdinsightRserverClusterSpecStorageAccount struct {
	StorageAccountKey  string `json:"storage_account_key"`
	StorageContainerId string `json:"storage_container_id"`
	IsDefault          bool   `json:"is_default"`
}

type AzurermHdinsightRserverClusterSpec struct {
	SshEndpoint       string                               `json:"ssh_endpoint"`
	Name              string                               `json:"name"`
	ClusterVersion    string                               `json:"cluster_version"`
	Gateway           []AzurermHdinsightRserverClusterSpec `json:"gateway"`
	Roles             []AzurermHdinsightRserverClusterSpec `json:"roles"`
	Tags              map[string]string                    `json:"tags"`
	EdgeSshEndpoint   string                               `json:"edge_ssh_endpoint"`
	HttpsEndpoint     string                               `json:"https_endpoint"`
	ResourceGroupName string                               `json:"resource_group_name"`
	Location          string                               `json:"location"`
	Tier              string                               `json:"tier"`
	Rstudio           bool                                 `json:"rstudio"`
	StorageAccount    []AzurermHdinsightRserverClusterSpec `json:"storage_account"`
}

type AzurermHdinsightRserverClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermHdinsightRserverClusterList is a list of AzurermHdinsightRserverClusters
type AzurermHdinsightRserverClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermHdinsightRserverCluster CRD objects
	Items []AzurermHdinsightRserverCluster `json:"items,omitempty"`
}
