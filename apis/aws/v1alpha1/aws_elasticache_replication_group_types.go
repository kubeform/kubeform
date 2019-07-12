package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheReplicationGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElasticacheReplicationGroupSpec   `json:"spec,omitempty"`
	Status            AwsElasticacheReplicationGroupStatus `json:"status,omitempty"`
}

type AwsElasticacheReplicationGroupSpecClusterMode struct {
	ReplicasPerNodeGroup int `json:"replicas_per_node_group"`
	NumNodeGroups        int `json:"num_node_groups"`
}

type AwsElasticacheReplicationGroupSpec struct {
	Port                         int                                  `json:"port"`
	ConfigurationEndpointAddress string                               `json:"configuration_endpoint_address"`
	Engine                       string                               `json:"engine"`
	NumberCacheClusters          int                                  `json:"number_cache_clusters"`
	SnapshotName                 string                               `json:"snapshot_name"`
	TransitEncryptionEnabled     bool                                 `json:"transit_encryption_enabled"`
	ApplyImmediately             bool                                 `json:"apply_immediately"`
	AvailabilityZones            []string                             `json:"availability_zones"`
	EngineVersion                string                               `json:"engine_version"`
	ParameterGroupName           string                               `json:"parameter_group_name"`
	ReplicationGroupId           string                               `json:"replication_group_id"`
	SnapshotWindow               string                               `json:"snapshot_window"`
	SubnetGroupName              string                               `json:"subnet_group_name"`
	AtRestEncryptionEnabled      bool                                 `json:"at_rest_encryption_enabled"`
	AutoMinorVersionUpgrade      bool                                 `json:"auto_minor_version_upgrade"`
	NodeType                     string                               `json:"node_type"`
	NotificationTopicArn         string                               `json:"notification_topic_arn"`
	PrimaryEndpointAddress       string                               `json:"primary_endpoint_address"`
	ReplicationGroupDescription  string                               `json:"replication_group_description"`
	SnapshotRetentionLimit       int                                  `json:"snapshot_retention_limit"`
	AutomaticFailoverEnabled     bool                                 `json:"automatic_failover_enabled"`
	ClusterMode                  []AwsElasticacheReplicationGroupSpec `json:"cluster_mode"`
	MaintenanceWindow            string                               `json:"maintenance_window"`
	SecurityGroupNames           []string                             `json:"security_group_names"`
	Tags                         map[string]string                    `json:"tags"`
	SecurityGroupIds             []string                             `json:"security_group_ids"`
	AuthToken                    string                               `json:"auth_token"`
	MemberClusters               []string                             `json:"member_clusters"`
	SnapshotArns                 []string                             `json:"snapshot_arns"`
}

type AwsElasticacheReplicationGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheReplicationGroupList is a list of AwsElasticacheReplicationGroups
type AwsElasticacheReplicationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElasticacheReplicationGroup CRD objects
	Items []AwsElasticacheReplicationGroup `json:"items,omitempty"`
}
