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

type AwsRdsCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRdsClusterSpec   `json:"spec,omitempty"`
	Status            AwsRdsClusterStatus `json:"status,omitempty"`
}

type AwsRdsClusterSpecS3Import struct {
	SourceEngine        string `json:"source_engine"`
	SourceEngineVersion string `json:"source_engine_version"`
	BucketName          string `json:"bucket_name"`
	BucketPrefix        string `json:"bucket_prefix"`
	IngestionRole       string `json:"ingestion_role"`
}

type AwsRdsClusterSpecScalingConfiguration struct {
	MaxCapacity           int  `json:"max_capacity"`
	MinCapacity           int  `json:"min_capacity"`
	SecondsUntilAutoPause int  `json:"seconds_until_auto_pause"`
	AutoPause             bool `json:"auto_pause"`
}

type AwsRdsClusterSpec struct {
	AvailabilityZones                []string            `json:"availability_zones"`
	HostedZoneId                     string              `json:"hosted_zone_id"`
	FinalSnapshotIdentifier          string              `json:"final_snapshot_identifier"`
	Port                             int                 `json:"port"`
	VpcSecurityGroupIds              []string            `json:"vpc_security_group_ids"`
	IamDatabaseAuthenticationEnabled bool                `json:"iam_database_authentication_enabled"`
	ClusterIdentifier                string              `json:"cluster_identifier"`
	MasterPassword                   string              `json:"master_password"`
	Arn                              string              `json:"arn"`
	Endpoint                         string              `json:"endpoint"`
	MasterUsername                   string              `json:"master_username"`
	SnapshotIdentifier               string              `json:"snapshot_identifier"`
	PreferredMaintenanceWindow       string              `json:"preferred_maintenance_window"`
	DeletionProtection               bool                `json:"deletion_protection"`
	S3Import                         []AwsRdsClusterSpec `json:"s3_import"`
	BackupRetentionPeriod            int                 `json:"backup_retention_period"`
	ApplyImmediately                 bool                `json:"apply_immediately"`
	ReplicationSourceIdentifier      string              `json:"replication_source_identifier"`
	BacktrackWindow                  int                 `json:"backtrack_window"`
	CopyTagsToSnapshot               bool                `json:"copy_tags_to_snapshot"`
	DatabaseName                     string              `json:"database_name"`
	DbSubnetGroupName                string              `json:"db_subnet_group_name"`
	ReaderEndpoint                   string              `json:"reader_endpoint"`
	EngineMode                       string              `json:"engine_mode"`
	IamRoles                         []string            `json:"iam_roles"`
	Tags                             map[string]string   `json:"tags"`
	ClusterIdentifierPrefix          string              `json:"cluster_identifier_prefix"`
	DbClusterParameterGroupName      string              `json:"db_cluster_parameter_group_name"`
	Engine                           string              `json:"engine"`
	SkipFinalSnapshot                bool                `json:"skip_final_snapshot"`
	ClusterResourceId                string              `json:"cluster_resource_id"`
	EnabledCloudwatchLogsExports     []string            `json:"enabled_cloudwatch_logs_exports"`
	ClusterMembers                   []string            `json:"cluster_members"`
	GlobalClusterIdentifier          string              `json:"global_cluster_identifier"`
	EngineVersion                    string              `json:"engine_version"`
	StorageEncrypted                 bool                `json:"storage_encrypted"`
	PreferredBackupWindow            string              `json:"preferred_backup_window"`
	SourceRegion                     string              `json:"source_region"`
	ScalingConfiguration             []AwsRdsClusterSpec `json:"scaling_configuration"`
	KmsKeyId                         string              `json:"kms_key_id"`
}

type AwsRdsClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRdsClusterList is a list of AwsRdsClusters
type AwsRdsClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRdsCluster CRD objects
	Items []AwsRdsCluster `json:"items,omitempty"`
}
