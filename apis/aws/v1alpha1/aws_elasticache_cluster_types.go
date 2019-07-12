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

type AwsElasticacheCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElasticacheClusterSpec   `json:"spec,omitempty"`
	Status            AwsElasticacheClusterStatus `json:"status,omitempty"`
}

type AwsElasticacheClusterSpecCacheNodes struct {
	AvailabilityZone string `json:"availability_zone"`
	Id               string `json:"id"`
	Address          string `json:"address"`
	Port             int    `json:"port"`
}

type AwsElasticacheClusterSpec struct {
	EngineVersion              string                      `json:"engine_version"`
	SecurityGroupIds           []string                    `json:"security_group_ids"`
	SnapshotArns               []string                    `json:"snapshot_arns"`
	SnapshotName               string                      `json:"snapshot_name"`
	AvailabilityZones          []string                    `json:"availability_zones"`
	AzMode                     string                      `json:"az_mode"`
	CacheNodes                 []AwsElasticacheClusterSpec `json:"cache_nodes"`
	Engine                     string                      `json:"engine"`
	Tags                       map[string]string           `json:"tags"`
	ParameterGroupName         string                      `json:"parameter_group_name"`
	PreferredAvailabilityZones []string                    `json:"preferred_availability_zones"`
	SecurityGroupNames         []string                    `json:"security_group_names"`
	SubnetGroupName            string                      `json:"subnet_group_name"`
	ApplyImmediately           bool                        `json:"apply_immediately"`
	AvailabilityZone           string                      `json:"availability_zone"`
	ConfigurationEndpoint      string                      `json:"configuration_endpoint"`
	NotificationTopicArn       string                      `json:"notification_topic_arn"`
	MaintenanceWindow          string                      `json:"maintenance_window"`
	NumCacheNodes              int                         `json:"num_cache_nodes"`
	Port                       int                         `json:"port"`
	SnapshotWindow             string                      `json:"snapshot_window"`
	SnapshotRetentionLimit     int                         `json:"snapshot_retention_limit"`
	ClusterAddress             string                      `json:"cluster_address"`
	ClusterId                  string                      `json:"cluster_id"`
	NodeType                   string                      `json:"node_type"`
	ReplicationGroupId         string                      `json:"replication_group_id"`
}

type AwsElasticacheClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsElasticacheClusterList is a list of AwsElasticacheClusters
type AwsElasticacheClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElasticacheCluster CRD objects
	Items []AwsElasticacheCluster `json:"items,omitempty"`
}
