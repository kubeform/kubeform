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
	NumberCacheClusters          int                                  `json:"number_cache_clusters"`
	ReplicationGroupId           string                               `json:"replication_group_id"`
	AvailabilityZones            []string                             `json:"availability_zones"`
	MemberClusters               []string                             `json:"member_clusters"`
	ReplicationGroupDescription  string                               `json:"replication_group_description"`
	Tags                         map[string]string                    `json:"tags"`
	AutoMinorVersionUpgrade      bool                                 `json:"auto_minor_version_upgrade"`
	AutomaticFailoverEnabled     bool                                 `json:"automatic_failover_enabled"`
	ClusterMode                  []AwsElasticacheReplicationGroupSpec `json:"cluster_mode"`
	SnapshotArns                 []string                             `json:"snapshot_arns"`
	ConfigurationEndpointAddress string                               `json:"configuration_endpoint_address"`
	EngineVersion                string                               `json:"engine_version"`
	SecurityGroupNames           []string                             `json:"security_group_names"`
	SnapshotWindow               string                               `json:"snapshot_window"`
	ApplyImmediately             bool                                 `json:"apply_immediately"`
	ParameterGroupName           string                               `json:"parameter_group_name"`
	PrimaryEndpointAddress       string                               `json:"primary_endpoint_address"`
	AtRestEncryptionEnabled      bool                                 `json:"at_rest_encryption_enabled"`
	MaintenanceWindow            string                               `json:"maintenance_window"`
	Port                         int                                  `json:"port"`
	SnapshotRetentionLimit       int                                  `json:"snapshot_retention_limit"`
	SubnetGroupName              string                               `json:"subnet_group_name"`
	TransitEncryptionEnabled     bool                                 `json:"transit_encryption_enabled"`
	AuthToken                    string                               `json:"auth_token"`
	SecurityGroupIds             []string                             `json:"security_group_ids"`
	Engine                       string                               `json:"engine"`
	NodeType                     string                               `json:"node_type"`
	NotificationTopicArn         string                               `json:"notification_topic_arn"`
	SnapshotName                 string                               `json:"snapshot_name"`
}

type AwsElasticacheReplicationGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsElasticacheReplicationGroupList is a list of AwsElasticacheReplicationGroups
type AwsElasticacheReplicationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElasticacheReplicationGroup CRD objects
	Items []AwsElasticacheReplicationGroup `json:"items,omitempty"`
}
