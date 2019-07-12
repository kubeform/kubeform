package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNeptuneCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNeptuneClusterSpec   `json:"spec,omitempty"`
	Status            AwsNeptuneClusterStatus `json:"status,omitempty"`
}

type AwsNeptuneClusterSpec struct {
	IamRoles                         []string          `json:"iam_roles"`
	KmsKeyArn                        string            `json:"kms_key_arn"`
	PreferredMaintenanceWindow       string            `json:"preferred_maintenance_window"`
	BackupRetentionPeriod            int               `json:"backup_retention_period"`
	ClusterResourceId                string            `json:"cluster_resource_id"`
	ApplyImmediately                 bool              `json:"apply_immediately"`
	AvailabilityZones                []string          `json:"availability_zones"`
	VpcSecurityGroupIds              []string          `json:"vpc_security_group_ids"`
	Arn                              string            `json:"arn"`
	Port                             int               `json:"port"`
	NeptuneClusterParameterGroupName string            `json:"neptune_cluster_parameter_group_name"`
	SkipFinalSnapshot                bool              `json:"skip_final_snapshot"`
	SnapshotIdentifier               string            `json:"snapshot_identifier"`
	FinalSnapshotIdentifier          string            `json:"final_snapshot_identifier"`
	NeptuneSubnetGroupName           string            `json:"neptune_subnet_group_name"`
	HostedZoneId                     string            `json:"hosted_zone_id"`
	StorageEncrypted                 bool              `json:"storage_encrypted"`
	Tags                             map[string]string `json:"tags"`
	ClusterIdentifier                string            `json:"cluster_identifier"`
	EngineVersion                    string            `json:"engine_version"`
	PreferredBackupWindow            string            `json:"preferred_backup_window"`
	Endpoint                         string            `json:"endpoint"`
	Engine                           string            `json:"engine"`
	IamDatabaseAuthenticationEnabled bool              `json:"iam_database_authentication_enabled"`
	ReaderEndpoint                   string            `json:"reader_endpoint"`
	ReplicationSourceIdentifier      string            `json:"replication_source_identifier"`
	ClusterIdentifierPrefix          string            `json:"cluster_identifier_prefix"`
	ClusterMembers                   []string          `json:"cluster_members"`
}



type AwsNeptuneClusterStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNeptuneClusterList is a list of AwsNeptuneClusters
type AwsNeptuneClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNeptuneCluster CRD objects
	Items []AwsNeptuneCluster `json:"items,omitempty"`
}