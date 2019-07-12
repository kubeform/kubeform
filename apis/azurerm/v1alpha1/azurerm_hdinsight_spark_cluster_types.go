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

type AzurermHdinsightSparkCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermHdinsightSparkClusterSpec   `json:"spec,omitempty"`
	Status            AzurermHdinsightSparkClusterStatus `json:"status,omitempty"`
}

type AzurermHdinsightSparkClusterSpecComponentVersion struct {
	Spark string `json:"spark"`
}

type AzurermHdinsightSparkClusterSpecGateway struct {
	Enabled  bool   `json:"enabled"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermHdinsightSparkClusterSpecStorageAccount struct {
	StorageContainerId string `json:"storage_container_id"`
	IsDefault          bool   `json:"is_default"`
	StorageAccountKey  string `json:"storage_account_key"`
}

type AzurermHdinsightSparkClusterSpecRolesHeadNode struct {
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
}

type AzurermHdinsightSparkClusterSpecRolesWorkerNode struct {
	VmSize              string   `json:"vm_size"`
	Username            string   `json:"username"`
	Password            string   `json:"password"`
	SshKeys             []string `json:"ssh_keys"`
	SubnetId            string   `json:"subnet_id"`
	VirtualNetworkId    string   `json:"virtual_network_id"`
	MinInstanceCount    int      `json:"min_instance_count"`
	TargetInstanceCount int      `json:"target_instance_count"`
}

type AzurermHdinsightSparkClusterSpecRolesZookeeperNode struct {
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
}

type AzurermHdinsightSparkClusterSpecRoles struct {
	HeadNode      []AzurermHdinsightSparkClusterSpecRoles `json:"head_node"`
	WorkerNode    []AzurermHdinsightSparkClusterSpecRoles `json:"worker_node"`
	ZookeeperNode []AzurermHdinsightSparkClusterSpecRoles `json:"zookeeper_node"`
}

type AzurermHdinsightSparkClusterSpec struct {
	ResourceGroupName string                             `json:"resource_group_name"`
	Location          string                             `json:"location"`
	ClusterVersion    string                             `json:"cluster_version"`
	Tags              map[string]string                  `json:"tags"`
	HttpsEndpoint     string                             `json:"https_endpoint"`
	Name              string                             `json:"name"`
	ComponentVersion  []AzurermHdinsightSparkClusterSpec `json:"component_version"`
	Gateway           []AzurermHdinsightSparkClusterSpec `json:"gateway"`
	StorageAccount    []AzurermHdinsightSparkClusterSpec `json:"storage_account"`
	Roles             []AzurermHdinsightSparkClusterSpec `json:"roles"`
	SshEndpoint       string                             `json:"ssh_endpoint"`
	Tier              string                             `json:"tier"`
}

type AzurermHdinsightSparkClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermHdinsightSparkClusterList is a list of AzurermHdinsightSparkClusters
type AzurermHdinsightSparkClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermHdinsightSparkCluster CRD objects
	Items []AzurermHdinsightSparkCluster `json:"items,omitempty"`
}
