package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermHdinsightHbaseCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermHdinsightHbaseClusterSpec   `json:"spec,omitempty"`
	Status            AzurermHdinsightHbaseClusterStatus `json:"status,omitempty"`
}

type AzurermHdinsightHbaseClusterSpecRolesHeadNode struct {
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
}

type AzurermHdinsightHbaseClusterSpecRolesWorkerNode struct {
	TargetInstanceCount int      `json:"target_instance_count"`
	VmSize              string   `json:"vm_size"`
	Username            string   `json:"username"`
	Password            string   `json:"password"`
	SshKeys             []string `json:"ssh_keys"`
	SubnetId            string   `json:"subnet_id"`
	VirtualNetworkId    string   `json:"virtual_network_id"`
	MinInstanceCount    int      `json:"min_instance_count"`
}

type AzurermHdinsightHbaseClusterSpecRolesZookeeperNode struct {
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
}

type AzurermHdinsightHbaseClusterSpecRoles struct {
	HeadNode      []AzurermHdinsightHbaseClusterSpecRoles `json:"head_node"`
	WorkerNode    []AzurermHdinsightHbaseClusterSpecRoles `json:"worker_node"`
	ZookeeperNode []AzurermHdinsightHbaseClusterSpecRoles `json:"zookeeper_node"`
}

type AzurermHdinsightHbaseClusterSpecComponentVersion struct {
	Hbase string `json:"hbase"`
}

type AzurermHdinsightHbaseClusterSpecGateway struct {
	Enabled  bool   `json:"enabled"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermHdinsightHbaseClusterSpecStorageAccount struct {
	StorageAccountKey  string `json:"storage_account_key"`
	StorageContainerId string `json:"storage_container_id"`
	IsDefault          bool   `json:"is_default"`
}

type AzurermHdinsightHbaseClusterSpec struct {
	Roles             []AzurermHdinsightHbaseClusterSpec `json:"roles"`
	Tags              map[string]string                  `json:"tags"`
	Name              string                             `json:"name"`
	Location          string                             `json:"location"`
	ClusterVersion    string                             `json:"cluster_version"`
	ComponentVersion  []AzurermHdinsightHbaseClusterSpec `json:"component_version"`
	Gateway           []AzurermHdinsightHbaseClusterSpec `json:"gateway"`
	StorageAccount    []AzurermHdinsightHbaseClusterSpec `json:"storage_account"`
	ResourceGroupName string                             `json:"resource_group_name"`
	Tier              string                             `json:"tier"`
	HttpsEndpoint     string                             `json:"https_endpoint"`
	SshEndpoint       string                             `json:"ssh_endpoint"`
}

type AzurermHdinsightHbaseClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermHdinsightHbaseClusterList is a list of AzurermHdinsightHbaseClusters
type AzurermHdinsightHbaseClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermHdinsightHbaseCluster CRD objects
	Items []AzurermHdinsightHbaseCluster `json:"items,omitempty"`
}
