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

type AzurermHdinsightRserverClusterSpecStorageAccount struct {
	StorageContainerId string `json:"storage_container_id"`
	IsDefault          bool   `json:"is_default"`
	StorageAccountKey  string `json:"storage_account_key"`
}

type AzurermHdinsightRserverClusterSpecRolesZookeeperNode struct {
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
}

type AzurermHdinsightRserverClusterSpecRolesEdgeNode struct {
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
}

type AzurermHdinsightRserverClusterSpecRolesHeadNode struct {
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
}

type AzurermHdinsightRserverClusterSpecRolesWorkerNode struct {
	VmSize              string   `json:"vm_size"`
	Username            string   `json:"username"`
	Password            string   `json:"password"`
	SshKeys             []string `json:"ssh_keys"`
	SubnetId            string   `json:"subnet_id"`
	VirtualNetworkId    string   `json:"virtual_network_id"`
	MinInstanceCount    int      `json:"min_instance_count"`
	TargetInstanceCount int      `json:"target_instance_count"`
}

type AzurermHdinsightRserverClusterSpecRoles struct {
	ZookeeperNode []AzurermHdinsightRserverClusterSpecRoles `json:"zookeeper_node"`
	EdgeNode      []AzurermHdinsightRserverClusterSpecRoles `json:"edge_node"`
	HeadNode      []AzurermHdinsightRserverClusterSpecRoles `json:"head_node"`
	WorkerNode    []AzurermHdinsightRserverClusterSpecRoles `json:"worker_node"`
}

type AzurermHdinsightRserverClusterSpec struct {
	ClusterVersion    string                               `json:"cluster_version"`
	Gateway           []AzurermHdinsightRserverClusterSpec `json:"gateway"`
	StorageAccount    []AzurermHdinsightRserverClusterSpec `json:"storage_account"`
	SshEndpoint       string                               `json:"ssh_endpoint"`
	ResourceGroupName string                               `json:"resource_group_name"`
	Location          string                               `json:"location"`
	Tier              string                               `json:"tier"`
	Rstudio           bool                                 `json:"rstudio"`
	Roles             []AzurermHdinsightRserverClusterSpec `json:"roles"`
	Tags              map[string]string                    `json:"tags"`
	EdgeSshEndpoint   string                               `json:"edge_ssh_endpoint"`
	HttpsEndpoint     string                               `json:"https_endpoint"`
	Name              string                               `json:"name"`
}



type AzurermHdinsightRserverClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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