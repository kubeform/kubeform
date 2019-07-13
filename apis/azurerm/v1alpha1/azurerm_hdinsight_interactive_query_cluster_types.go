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

type AzurermHdinsightInteractiveQueryCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermHdinsightInteractiveQueryClusterSpec   `json:"spec,omitempty"`
	Status            AzurermHdinsightInteractiveQueryClusterStatus `json:"status,omitempty"`
}

type AzurermHdinsightInteractiveQueryClusterSpecComponentVersion struct {
	InteractiveHive string `json:"interactive_hive"`
}

type AzurermHdinsightInteractiveQueryClusterSpecGateway struct {
	Enabled  bool   `json:"enabled"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermHdinsightInteractiveQueryClusterSpecStorageAccount struct {
	StorageContainerId string `json:"storage_container_id"`
	IsDefault          bool   `json:"is_default"`
	StorageAccountKey  string `json:"storage_account_key"`
}

type AzurermHdinsightInteractiveQueryClusterSpecRolesHeadNode struct {
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
}

type AzurermHdinsightInteractiveQueryClusterSpecRolesWorkerNode struct {
	SshKeys             []string `json:"ssh_keys"`
	SubnetId            string   `json:"subnet_id"`
	VirtualNetworkId    string   `json:"virtual_network_id"`
	MinInstanceCount    int      `json:"min_instance_count"`
	TargetInstanceCount int      `json:"target_instance_count"`
	VmSize              string   `json:"vm_size"`
	Username            string   `json:"username"`
	Password            string   `json:"password"`
}

type AzurermHdinsightInteractiveQueryClusterSpecRolesZookeeperNode struct {
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
}

type AzurermHdinsightInteractiveQueryClusterSpecRoles struct {
	HeadNode      []AzurermHdinsightInteractiveQueryClusterSpecRoles `json:"head_node"`
	WorkerNode    []AzurermHdinsightInteractiveQueryClusterSpecRoles `json:"worker_node"`
	ZookeeperNode []AzurermHdinsightInteractiveQueryClusterSpecRoles `json:"zookeeper_node"`
}

type AzurermHdinsightInteractiveQueryClusterSpec struct {
	ResourceGroupName string                                        `json:"resource_group_name"`
	Location          string                                        `json:"location"`
	ClusterVersion    string                                        `json:"cluster_version"`
	ComponentVersion  []AzurermHdinsightInteractiveQueryClusterSpec `json:"component_version"`
	Tags              map[string]string                             `json:"tags"`
	HttpsEndpoint     string                                        `json:"https_endpoint"`
	SshEndpoint       string                                        `json:"ssh_endpoint"`
	Name              string                                        `json:"name"`
	Tier              string                                        `json:"tier"`
	Gateway           []AzurermHdinsightInteractiveQueryClusterSpec `json:"gateway"`
	StorageAccount    []AzurermHdinsightInteractiveQueryClusterSpec `json:"storage_account"`
	Roles             []AzurermHdinsightInteractiveQueryClusterSpec `json:"roles"`
}



type AzurermHdinsightInteractiveQueryClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermHdinsightInteractiveQueryClusterList is a list of AzurermHdinsightInteractiveQueryClusters
type AzurermHdinsightInteractiveQueryClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermHdinsightInteractiveQueryCluster CRD objects
	Items []AzurermHdinsightInteractiveQueryCluster `json:"items,omitempty"`
}