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

type AzurermHdinsightMlServicesCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermHdinsightMlServicesClusterSpec   `json:"spec,omitempty"`
	Status            AzurermHdinsightMlServicesClusterStatus `json:"status,omitempty"`
}

type AzurermHdinsightMlServicesClusterSpecStorageAccount struct {
	StorageAccountKey  string `json:"storage_account_key"`
	StorageContainerId string `json:"storage_container_id"`
	IsDefault          bool   `json:"is_default"`
}

type AzurermHdinsightMlServicesClusterSpecRolesWorkerNode struct {
	Username            string   `json:"username"`
	Password            string   `json:"password"`
	SshKeys             []string `json:"ssh_keys"`
	SubnetId            string   `json:"subnet_id"`
	VirtualNetworkId    string   `json:"virtual_network_id"`
	MinInstanceCount    int      `json:"min_instance_count"`
	TargetInstanceCount int      `json:"target_instance_count"`
	VmSize              string   `json:"vm_size"`
}

type AzurermHdinsightMlServicesClusterSpecRolesZookeeperNode struct {
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
}

type AzurermHdinsightMlServicesClusterSpecRolesEdgeNode struct {
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
}

type AzurermHdinsightMlServicesClusterSpecRolesHeadNode struct {
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
}

type AzurermHdinsightMlServicesClusterSpecRoles struct {
	WorkerNode    []AzurermHdinsightMlServicesClusterSpecRoles `json:"worker_node"`
	ZookeeperNode []AzurermHdinsightMlServicesClusterSpecRoles `json:"zookeeper_node"`
	EdgeNode      []AzurermHdinsightMlServicesClusterSpecRoles `json:"edge_node"`
	HeadNode      []AzurermHdinsightMlServicesClusterSpecRoles `json:"head_node"`
}

type AzurermHdinsightMlServicesClusterSpecGateway struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Enabled  bool   `json:"enabled"`
}

type AzurermHdinsightMlServicesClusterSpec struct {
	StorageAccount    []AzurermHdinsightMlServicesClusterSpec `json:"storage_account"`
	Roles             []AzurermHdinsightMlServicesClusterSpec `json:"roles"`
	SshEndpoint       string                                  `json:"ssh_endpoint"`
	Location          string                                  `json:"location"`
	ClusterVersion    string                                  `json:"cluster_version"`
	Tier              string                                  `json:"tier"`
	Rstudio           bool                                    `json:"rstudio"`
	EdgeSshEndpoint   string                                  `json:"edge_ssh_endpoint"`
	HttpsEndpoint     string                                  `json:"https_endpoint"`
	Name              string                                  `json:"name"`
	ResourceGroupName string                                  `json:"resource_group_name"`
	Gateway           []AzurermHdinsightMlServicesClusterSpec `json:"gateway"`
	Tags              map[string]string                       `json:"tags"`
}

type AzurermHdinsightMlServicesClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermHdinsightMlServicesClusterList is a list of AzurermHdinsightMlServicesClusters
type AzurermHdinsightMlServicesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermHdinsightMlServicesCluster CRD objects
	Items []AzurermHdinsightMlServicesCluster `json:"items,omitempty"`
}
