package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	AvailabilityZone           string                      `json:"availability_zone"`
	ConfigurationEndpoint      string                      `json:"configuration_endpoint"`
	NodeType                   string                      `json:"node_type"`
	Port                       int                         `json:"port"`
	SnapshotRetentionLimit     int                         `json:"snapshot_retention_limit"`
	ApplyImmediately           bool                        `json:"apply_immediately"`
	CacheNodes                 []AwsElasticacheClusterSpec `json:"cache_nodes"`
	NotificationTopicArn       string                      `json:"notification_topic_arn"`
	PreferredAvailabilityZones []string                    `json:"preferred_availability_zones"`
	SecurityGroupNames         []string                    `json:"security_group_names"`
	SnapshotArns               []string                    `json:"snapshot_arns"`
	AvailabilityZones          []string                    `json:"availability_zones"`
	ClusterAddress             string                      `json:"cluster_address"`
	ClusterId                  string                      `json:"cluster_id"`
	Engine                     string                      `json:"engine"`
	NumCacheNodes              int                         `json:"num_cache_nodes"`
	ParameterGroupName         string                      `json:"parameter_group_name"`
	SnapshotName               string                      `json:"snapshot_name"`
	AzMode                     string                      `json:"az_mode"`
	MaintenanceWindow          string                      `json:"maintenance_window"`
	ReplicationGroupId         string                      `json:"replication_group_id"`
	SecurityGroupIds           []string                    `json:"security_group_ids"`
	SnapshotWindow             string                      `json:"snapshot_window"`
	SubnetGroupName            string                      `json:"subnet_group_name"`
	Tags                       map[string]string           `json:"tags"`
	EngineVersion              string                      `json:"engine_version"`
}

type AwsElasticacheClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheClusterList is a list of AwsElasticacheClusters
type AwsElasticacheClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElasticacheCluster CRD objects
	Items []AwsElasticacheCluster `json:"items,omitempty"`
}
