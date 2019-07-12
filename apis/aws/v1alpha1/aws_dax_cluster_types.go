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
	AvailabilityZone string `json:"availability_zone"`
	Id               string `json:"id"`
	Address          string `json:"address"`
	Port             int    `json:"port"`
}

type AwsDaxClusterSpecServerSideEncryption struct {
	Enabled bool `json:"enabled"`
}

type AwsDaxClusterSpec struct {
	Port                  int                 `json:"port"`
	Nodes                 []AwsDaxClusterSpec `json:"nodes"`
	IamRoleArn            string              `json:"iam_role_arn"`
	SubnetGroupName       string              `json:"subnet_group_name"`
	ConfigurationEndpoint string              `json:"configuration_endpoint"`
	ReplicationFactor     int                 `json:"replication_factor"`
	AvailabilityZones     []string            `json:"availability_zones"`
	Description           string              `json:"description"`
	SecurityGroupIds      []string            `json:"security_group_ids"`
	ServerSideEncryption  []AwsDaxClusterSpec `json:"server_side_encryption"`
	MaintenanceWindow     string              `json:"maintenance_window"`
	Tags                  map[string]string   `json:"tags"`
	ClusterAddress        string              `json:"cluster_address"`
	Arn                   string              `json:"arn"`
	ClusterName           string              `json:"cluster_name"`
	NodeType              string              `json:"node_type"`
	NotificationTopicArn  string              `json:"notification_topic_arn"`
	ParameterGroupName    string              `json:"parameter_group_name"`
}

type AwsDaxClusterStatus struct {
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
