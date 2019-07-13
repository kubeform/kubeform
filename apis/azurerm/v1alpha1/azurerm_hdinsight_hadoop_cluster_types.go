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

type AzurermHdinsightHadoopCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermHdinsightHadoopClusterSpec   `json:"spec,omitempty"`
	Status            AzurermHdinsightHadoopClusterStatus `json:"status,omitempty"`
}

type AzurermHdinsightHadoopClusterSpecStorageAccount struct {
	StorageAccountKey  string `json:"storage_account_key"`
	StorageContainerId string `json:"storage_container_id"`
	IsDefault          bool   `json:"is_default"`
}

type AzurermHdinsightHadoopClusterSpecRolesHeadNode struct {
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
}

type AzurermHdinsightHadoopClusterSpecRolesWorkerNode struct {
	Password            string   `json:"password"`
	SshKeys             []string `json:"ssh_keys"`
	SubnetId            string   `json:"subnet_id"`
	VirtualNetworkId    string   `json:"virtual_network_id"`
	MinInstanceCount    int      `json:"min_instance_count"`
	TargetInstanceCount int      `json:"target_instance_count"`
	VmSize              string   `json:"vm_size"`
	Username            string   `json:"username"`
}

type AzurermHdinsightHadoopClusterSpecRolesZookeeperNode struct {
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
}

type AzurermHdinsightHadoopClusterSpecRoles struct {
	HeadNode      []AzurermHdinsightHadoopClusterSpecRoles `json:"head_node"`
	WorkerNode    []AzurermHdinsightHadoopClusterSpecRoles `json:"worker_node"`
	ZookeeperNode []AzurermHdinsightHadoopClusterSpecRoles `json:"zookeeper_node"`
}

type AzurermHdinsightHadoopClusterSpecGateway struct {
	Enabled  bool   `json:"enabled"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermHdinsightHadoopClusterSpecComponentVersion struct {
	Hadoop string `json:"hadoop"`
}

type AzurermHdinsightHadoopClusterSpec struct {
	StorageAccount    []AzurermHdinsightHadoopClusterSpec `json:"storage_account"`
	Roles             []AzurermHdinsightHadoopClusterSpec `json:"roles"`
	Tags              map[string]string                   `json:"tags"`
	ResourceGroupName string                              `json:"resource_group_name"`
	Location          string                              `json:"location"`
	ClusterVersion    string                              `json:"cluster_version"`
	Tier              string                              `json:"tier"`
	Gateway           []AzurermHdinsightHadoopClusterSpec `json:"gateway"`
	SshEndpoint       string                              `json:"ssh_endpoint"`
	Name              string                              `json:"name"`
	ComponentVersion  []AzurermHdinsightHadoopClusterSpec `json:"component_version"`
	HttpsEndpoint     string                              `json:"https_endpoint"`
}



type AzurermHdinsightHadoopClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermHdinsightHadoopClusterList is a list of AzurermHdinsightHadoopClusters
type AzurermHdinsightHadoopClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermHdinsightHadoopCluster CRD objects
	Items []AzurermHdinsightHadoopCluster `json:"items,omitempty"`
}