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

type AwsDaxCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDaxClusterSpec   `json:"spec,omitempty"`
	Status            AwsDaxClusterStatus `json:"status,omitempty"`
}

type AwsDaxClusterSpecNodes struct {
	Port             int    `json:"port"`
	AvailabilityZone string `json:"availability_zone"`
	Id               string `json:"id"`
	Address          string `json:"address"`
}

type AwsDaxClusterSpecServerSideEncryption struct {
	Enabled bool `json:"enabled"`
}

type AwsDaxClusterSpec struct {
	Nodes                 []AwsDaxClusterSpec `json:"nodes"`
	ClusterName           string              `json:"cluster_name"`
	IamRoleArn            string              `json:"iam_role_arn"`
	ReplicationFactor     int                 `json:"replication_factor"`
	NotificationTopicArn  string              `json:"notification_topic_arn"`
	MaintenanceWindow     string              `json:"maintenance_window"`
	SubnetGroupName       string              `json:"subnet_group_name"`
	ConfigurationEndpoint string              `json:"configuration_endpoint"`
	AvailabilityZones     []string            `json:"availability_zones"`
	ParameterGroupName    string              `json:"parameter_group_name"`
	SecurityGroupIds      []string            `json:"security_group_ids"`
	ClusterAddress        string              `json:"cluster_address"`
	Arn                   string              `json:"arn"`
	Description           string              `json:"description"`
	Tags                  map[string]string   `json:"tags"`
	Port                  int                 `json:"port"`
	NodeType              string              `json:"node_type"`
	ServerSideEncryption  []AwsDaxClusterSpec `json:"server_side_encryption"`
}



type AwsDaxClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDaxClusterList is a list of AwsDaxClusters
type AwsDaxClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDaxCluster CRD objects
	Items []AwsDaxCluster `json:"items,omitempty"`
}