package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermHdinsightHadoopCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermHdinsightHadoopClusterSpec   `json:"spec,omitempty"`
	Status            AzurermHdinsightHadoopClusterStatus `json:"status,omitempty"`
}

type AzurermHdinsightHadoopClusterSpecRolesHeadNode struct {
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
}

type AzurermHdinsightHadoopClusterSpecRolesWorkerNode struct {
	VirtualNetworkId    string   `json:"virtual_network_id"`
	MinInstanceCount    int      `json:"min_instance_count"`
	TargetInstanceCount int      `json:"target_instance_count"`
	VmSize              string   `json:"vm_size"`
	Username            string   `json:"username"`
	Password            string   `json:"password"`
	SshKeys             []string `json:"ssh_keys"`
	SubnetId            string   `json:"subnet_id"`
}

type AzurermHdinsightHadoopClusterSpecRolesZookeeperNode struct {
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
}

type AzurermHdinsightHadoopClusterSpecRoles struct {
	HeadNode      []AzurermHdinsightHadoopClusterSpecRoles `json:"head_node"`
	WorkerNode    []AzurermHdinsightHadoopClusterSpecRoles `json:"worker_node"`
	ZookeeperNode []AzurermHdinsightHadoopClusterSpecRoles `json:"zookeeper_node"`
}

type AzurermHdinsightHadoopClusterSpecComponentVersion struct {
	Hadoop string `json:"hadoop"`
}

type AzurermHdinsightHadoopClusterSpecGateway struct {
	Enabled  bool   `json:"enabled"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermHdinsightHadoopClusterSpecStorageAccount struct {
	StorageContainerId string `json:"storage_container_id"`
	IsDefault          bool   `json:"is_default"`
	StorageAccountKey  string `json:"storage_account_key"`
}

type AzurermHdinsightHadoopClusterSpec struct {
	HttpsEndpoint     string                              `json:"https_endpoint"`
	Name              string                              `json:"name"`
	ResourceGroupName string                              `json:"resource_group_name"`
	Location          string                              `json:"location"`
	ClusterVersion    string                              `json:"cluster_version"`
	Roles             []AzurermHdinsightHadoopClusterSpec `json:"roles"`
	SshEndpoint       string                              `json:"ssh_endpoint"`
	Tier              string                              `json:"tier"`
	ComponentVersion  []AzurermHdinsightHadoopClusterSpec `json:"component_version"`
	Gateway           []AzurermHdinsightHadoopClusterSpec `json:"gateway"`
	StorageAccount    []AzurermHdinsightHadoopClusterSpec `json:"storage_account"`
	Tags              map[string]string                   `json:"tags"`
}



type AzurermHdinsightHadoopClusterStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermHdinsightHadoopClusterList is a list of AzurermHdinsightHadoopClusters
type AzurermHdinsightHadoopClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermHdinsightHadoopCluster CRD objects
	Items []AzurermHdinsightHadoopCluster `json:"items,omitempty"`
}