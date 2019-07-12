package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermHdinsightMlServicesCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermHdinsightMlServicesClusterSpec   `json:"spec,omitempty"`
	Status            AzurermHdinsightMlServicesClusterStatus `json:"status,omitempty"`
}

type AzurermHdinsightMlServicesClusterSpecGateway struct {
	Enabled  bool   `json:"enabled"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermHdinsightMlServicesClusterSpecStorageAccount struct {
	StorageAccountKey  string `json:"storage_account_key"`
	StorageContainerId string `json:"storage_container_id"`
	IsDefault          bool   `json:"is_default"`
}

type AzurermHdinsightMlServicesClusterSpecRolesHeadNode struct {
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
}

type AzurermHdinsightMlServicesClusterSpecRolesWorkerNode struct {
	VmSize              string   `json:"vm_size"`
	Username            string   `json:"username"`
	Password            string   `json:"password"`
	SshKeys             []string `json:"ssh_keys"`
	SubnetId            string   `json:"subnet_id"`
	VirtualNetworkId    string   `json:"virtual_network_id"`
	MinInstanceCount    int      `json:"min_instance_count"`
	TargetInstanceCount int      `json:"target_instance_count"`
}

type AzurermHdinsightMlServicesClusterSpecRolesZookeeperNode struct {
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
}

type AzurermHdinsightMlServicesClusterSpecRolesEdgeNode struct {
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
}

type AzurermHdinsightMlServicesClusterSpecRoles struct {
	HeadNode      []AzurermHdinsightMlServicesClusterSpecRoles `json:"head_node"`
	WorkerNode    []AzurermHdinsightMlServicesClusterSpecRoles `json:"worker_node"`
	ZookeeperNode []AzurermHdinsightMlServicesClusterSpecRoles `json:"zookeeper_node"`
	EdgeNode      []AzurermHdinsightMlServicesClusterSpecRoles `json:"edge_node"`
}

type AzurermHdinsightMlServicesClusterSpec struct {
	HttpsEndpoint     string                                  `json:"https_endpoint"`
	ResourceGroupName string                                  `json:"resource_group_name"`
	Location          string                                  `json:"location"`
	Tier              string                                  `json:"tier"`
	Gateway           []AzurermHdinsightMlServicesClusterSpec `json:"gateway"`
	Rstudio           bool                                    `json:"rstudio"`
	Tags              map[string]string                       `json:"tags"`
	Name              string                                  `json:"name"`
	ClusterVersion    string                                  `json:"cluster_version"`
	StorageAccount    []AzurermHdinsightMlServicesClusterSpec `json:"storage_account"`
	Roles             []AzurermHdinsightMlServicesClusterSpec `json:"roles"`
	EdgeSshEndpoint   string                                  `json:"edge_ssh_endpoint"`
	SshEndpoint       string                                  `json:"ssh_endpoint"`
}



type AzurermHdinsightMlServicesClusterStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermHdinsightMlServicesClusterList is a list of AzurermHdinsightMlServicesClusters
type AzurermHdinsightMlServicesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermHdinsightMlServicesCluster CRD objects
	Items []AzurermHdinsightMlServicesCluster `json:"items,omitempty"`
}