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
	SecurityGroupNames           []string                             `json:"security_group_names"`
	MaintenanceWindow            string                               `json:"maintenance_window"`
	PrimaryEndpointAddress       string                               `json:"primary_endpoint_address"`
	Port                         int                                  `json:"port"`
	AutoMinorVersionUpgrade      bool                                 `json:"auto_minor_version_upgrade"`
	AvailabilityZones            []string                             `json:"availability_zones"`
	Engine                       string                               `json:"engine"`
	EngineVersion                string                               `json:"engine_version"`
	NodeType                     string                               `json:"node_type"`
	ParameterGroupName           string                               `json:"parameter_group_name"`
	SecurityGroupIds             []string                             `json:"security_group_ids"`
	SubnetGroupName              string                               `json:"subnet_group_name"`
	AutomaticFailoverEnabled     bool                                 `json:"automatic_failover_enabled"`
	AuthToken                    string                               `json:"auth_token"`
	ClusterMode                  []AwsElasticacheReplicationGroupSpec `json:"cluster_mode"`
	TransitEncryptionEnabled     bool                                 `json:"transit_encryption_enabled"`
	AtRestEncryptionEnabled      bool                                 `json:"at_rest_encryption_enabled"`
	NotificationTopicArn         string                               `json:"notification_topic_arn"`
	ReplicationGroupDescription  string                               `json:"replication_group_description"`
	ReplicationGroupId           string                               `json:"replication_group_id"`
	SnapshotWindow               string                               `json:"snapshot_window"`
	Tags                         map[string]string                    `json:"tags"`
	ConfigurationEndpointAddress string                               `json:"configuration_endpoint_address"`
	MemberClusters               []string                             `json:"member_clusters"`
	SnapshotArns                 []string                             `json:"snapshot_arns"`
	SnapshotRetentionLimit       int                                  `json:"snapshot_retention_limit"`
	SnapshotName                 string                               `json:"snapshot_name"`
	ApplyImmediately             bool                                 `json:"apply_immediately"`
}



type AwsElasticacheReplicationGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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