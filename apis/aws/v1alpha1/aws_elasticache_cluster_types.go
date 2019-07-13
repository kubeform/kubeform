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
	Id               string `json:"id"`
	Address          string `json:"address"`
	Port             int    `json:"port"`
	AvailabilityZone string `json:"availability_zone"`
}

type AwsElasticacheClusterSpec struct {
	ConfigurationEndpoint      string                      `json:"configuration_endpoint"`
	PreferredAvailabilityZones []string                    `json:"preferred_availability_zones"`
	SnapshotRetentionLimit     int                         `json:"snapshot_retention_limit"`
	SnapshotWindow             string                      `json:"snapshot_window"`
	Tags                       map[string]string           `json:"tags"`
	SnapshotArns               []string                    `json:"snapshot_arns"`
	AvailabilityZone           string                      `json:"availability_zone"`
	ClusterId                  string                      `json:"cluster_id"`
	MaintenanceWindow          string                      `json:"maintenance_window"`
	NumCacheNodes              int                         `json:"num_cache_nodes"`
	ReplicationGroupId         string                      `json:"replication_group_id"`
	Port                       int                         `json:"port"`
	SecurityGroupNames         []string                    `json:"security_group_names"`
	SecurityGroupIds           []string                    `json:"security_group_ids"`
	ApplyImmediately           bool                        `json:"apply_immediately"`
	AvailabilityZones          []string                    `json:"availability_zones"`
	CacheNodes                 []AwsElasticacheClusterSpec `json:"cache_nodes"`
	EngineVersion              string                      `json:"engine_version"`
	ParameterGroupName         string                      `json:"parameter_group_name"`
	SnapshotName               string                      `json:"snapshot_name"`
	SubnetGroupName            string                      `json:"subnet_group_name"`
	AzMode                     string                      `json:"az_mode"`
	ClusterAddress             string                      `json:"cluster_address"`
	Engine                     string                      `json:"engine"`
	NodeType                   string                      `json:"node_type"`
	NotificationTopicArn       string                      `json:"notification_topic_arn"`
}



type AwsElasticacheClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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