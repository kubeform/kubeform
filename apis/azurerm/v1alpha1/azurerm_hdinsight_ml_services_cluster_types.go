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

type AzurermHdinsightMlServicesClusterSpecRolesHeadNode struct {
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
}

type AzurermHdinsightMlServicesClusterSpecRolesWorkerNode struct {
	MinInstanceCount    int      `json:"min_instance_count"`
	TargetInstanceCount int      `json:"target_instance_count"`
	VmSize              string   `json:"vm_size"`
	Username            string   `json:"username"`
	Password            string   `json:"password"`
	SshKeys             []string `json:"ssh_keys"`
	SubnetId            string   `json:"subnet_id"`
	VirtualNetworkId    string   `json:"virtual_network_id"`
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
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
}

type AzurermHdinsightMlServicesClusterSpecRoles struct {
	HeadNode      []AzurermHdinsightMlServicesClusterSpecRoles `json:"head_node"`
	WorkerNode    []AzurermHdinsightMlServicesClusterSpecRoles `json:"worker_node"`
	ZookeeperNode []AzurermHdinsightMlServicesClusterSpecRoles `json:"zookeeper_node"`
	EdgeNode      []AzurermHdinsightMlServicesClusterSpecRoles `json:"edge_node"`
}

type AzurermHdinsightMlServicesClusterSpecStorageAccount struct {
	IsDefault          bool   `json:"is_default"`
	StorageAccountKey  string `json:"storage_account_key"`
	StorageContainerId string `json:"storage_container_id"`
}

type AzurermHdinsightMlServicesClusterSpecGateway struct {
	Enabled  bool   `json:"enabled"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermHdinsightMlServicesClusterSpec struct {
	HttpsEndpoint     string                                  `json:"https_endpoint"`
	Name              string                                  `json:"name"`
	Location          string                                  `json:"location"`
	Rstudio           bool                                    `json:"rstudio"`
	Roles             []AzurermHdinsightMlServicesClusterSpec `json:"roles"`
	StorageAccount    []AzurermHdinsightMlServicesClusterSpec `json:"storage_account"`
	Tags              map[string]string                       `json:"tags"`
	EdgeSshEndpoint   string                                  `json:"edge_ssh_endpoint"`
	SshEndpoint       string                                  `json:"ssh_endpoint"`
	ResourceGroupName string                                  `json:"resource_group_name"`
	ClusterVersion    string                                  `json:"cluster_version"`
	Tier              string                                  `json:"tier"`
	Gateway           []AzurermHdinsightMlServicesClusterSpec `json:"gateway"`
}



type AzurermHdinsightMlServicesClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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