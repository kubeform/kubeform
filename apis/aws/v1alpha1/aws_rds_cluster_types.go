package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRdsCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRdsClusterSpec   `json:"spec,omitempty"`
	Status            AwsRdsClusterStatus `json:"status,omitempty"`
}

type AwsRdsClusterSpecS3Import struct {
	BucketName          string `json:"bucket_name"`
	BucketPrefix        string `json:"bucket_prefix"`
	IngestionRole       string `json:"ingestion_role"`
	SourceEngine        string `json:"source_engine"`
	SourceEngineVersion string `json:"source_engine_version"`
}

type AwsRdsClusterSpecScalingConfiguration struct {
	SecondsUntilAutoPause int  `json:"seconds_until_auto_pause"`
	AutoPause             bool `json:"auto_pause"`
	MaxCapacity           int  `json:"max_capacity"`
	MinCapacity           int  `json:"min_capacity"`
}

type AwsRdsClusterSpec struct {
	DbSubnetGroupName                string              `json:"db_subnet_group_name"`
	DeletionProtection               bool                `json:"deletion_protection"`
	Engine                           string              `json:"engine"`
	MasterUsername                   string              `json:"master_username"`
	KmsKeyId                         string              `json:"kms_key_id"`
	AvailabilityZones                []string            `json:"availability_zones"`
	ClusterIdentifierPrefix          string              `json:"cluster_identifier_prefix"`
	Port                             int                 `json:"port"`
	VpcSecurityGroupIds              []string            `json:"vpc_security_group_ids"`
	PreferredMaintenanceWindow       string              `json:"preferred_maintenance_window"`
	SourceRegion                     string              `json:"source_region"`
	ClusterMembers                   []string            `json:"cluster_members"`
	EngineVersion                    string              `json:"engine_version"`
	S3Import                         []AwsRdsClusterSpec `json:"s3_import"`
	SkipFinalSnapshot                bool                `json:"skip_final_snapshot"`
	MasterPassword                   string              `json:"master_password"`
	IamDatabaseAuthenticationEnabled bool                `json:"iam_database_authentication_enabled"`
	EnabledCloudwatchLogsExports     []string            `json:"enabled_cloudwatch_logs_exports"`
	Tags                             map[string]string   `json:"tags"`
	Endpoint                         string              `json:"endpoint"`
	ScalingConfiguration             []AwsRdsClusterSpec `json:"scaling_configuration"`
	StorageEncrypted                 bool                `json:"storage_encrypted"`
	DatabaseName                     string              `json:"database_name"`
	GlobalClusterIdentifier          string              `json:"global_cluster_identifier"`
	DbClusterParameterGroupName      string              `json:"db_cluster_parameter_group_name"`
	SnapshotIdentifier               string              `json:"snapshot_identifier"`
	PreferredBackupWindow            string              `json:"preferred_backup_window"`
	ReplicationSourceIdentifier      string              `json:"replication_source_identifier"`
	IamRoles                         []string            `json:"iam_roles"`
	ClusterIdentifier                string              `json:"cluster_identifier"`
	CopyTagsToSnapshot               bool                `json:"copy_tags_to_snapshot"`
	BackupRetentionPeriod            int                 `json:"backup_retention_period"`
	ClusterResourceId                string              `json:"cluster_resource_id"`
	BacktrackWindow                  int                 `json:"backtrack_window"`
	ApplyImmediately                 bool                `json:"apply_immediately"`
	EngineMode                       string              `json:"engine_mode"`
	Arn                              string              `json:"arn"`
	HostedZoneId                     string              `json:"hosted_zone_id"`
	ReaderEndpoint                   string              `json:"reader_endpoint"`
	FinalSnapshotIdentifier          string              `json:"final_snapshot_identifier"`
}



type AwsRdsClusterStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRdsClusterList is a list of AwsRdsClusters
type AwsRdsClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRdsCluster CRD objects
	Items []AwsRdsCluster `json:"items,omitempty"`
}