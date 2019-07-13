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

type AzurermHdinsightKafkaCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermHdinsightKafkaClusterSpec   `json:"spec,omitempty"`
	Status            AzurermHdinsightKafkaClusterStatus `json:"status,omitempty"`
}

type AzurermHdinsightKafkaClusterSpecComponentVersion struct {
	Kafka string `json:"kafka"`
}

type AzurermHdinsightKafkaClusterSpecGateway struct {
	Enabled  bool   `json:"enabled"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermHdinsightKafkaClusterSpecStorageAccount struct {
	StorageAccountKey  string `json:"storage_account_key"`
	StorageContainerId string `json:"storage_container_id"`
	IsDefault          bool   `json:"is_default"`
}

type AzurermHdinsightKafkaClusterSpecRolesHeadNode struct {
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
}

type AzurermHdinsightKafkaClusterSpecRolesWorkerNode struct {
	Password             string   `json:"password"`
	MinInstanceCount     int      `json:"min_instance_count"`
	VmSize               string   `json:"vm_size"`
	Username             string   `json:"username"`
	SshKeys              []string `json:"ssh_keys"`
	SubnetId             string   `json:"subnet_id"`
	VirtualNetworkId     string   `json:"virtual_network_id"`
	TargetInstanceCount  int      `json:"target_instance_count"`
	NumberOfDisksPerNode int      `json:"number_of_disks_per_node"`
}

type AzurermHdinsightKafkaClusterSpecRolesZookeeperNode struct {
	Password         string   `json:"password"`
	SshKeys          []string `json:"ssh_keys"`
	SubnetId         string   `json:"subnet_id"`
	VirtualNetworkId string   `json:"virtual_network_id"`
	VmSize           string   `json:"vm_size"`
	Username         string   `json:"username"`
}

type AzurermHdinsightKafkaClusterSpecRoles struct {
	HeadNode      []AzurermHdinsightKafkaClusterSpecRoles `json:"head_node"`
	WorkerNode    []AzurermHdinsightKafkaClusterSpecRoles `json:"worker_node"`
	ZookeeperNode []AzurermHdinsightKafkaClusterSpecRoles `json:"zookeeper_node"`
}

type AzurermHdinsightKafkaClusterSpec struct {
	Name              string                             `json:"name"`
	ResourceGroupName string                             `json:"resource_group_name"`
	ComponentVersion  []AzurermHdinsightKafkaClusterSpec `json:"component_version"`
	Tags              map[string]string                  `json:"tags"`
	SshEndpoint       string                             `json:"ssh_endpoint"`
	Location          string                             `json:"location"`
	ClusterVersion    string                             `json:"cluster_version"`
	Tier              string                             `json:"tier"`
	Gateway           []AzurermHdinsightKafkaClusterSpec `json:"gateway"`
	StorageAccount    []AzurermHdinsightKafkaClusterSpec `json:"storage_account"`
	Roles             []AzurermHdinsightKafkaClusterSpec `json:"roles"`
	HttpsEndpoint     string                             `json:"https_endpoint"`
}



type AzurermHdinsightKafkaClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermHdinsightKafkaClusterList is a list of AzurermHdinsightKafkaClusters
type AzurermHdinsightKafkaClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermHdinsightKafkaCluster CRD objects
	Items []AzurermHdinsightKafkaCluster `json:"items,omitempty"`
}