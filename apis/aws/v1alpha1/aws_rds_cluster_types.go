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

type AwsRdsClusterSpecScalingConfiguration struct {
	AutoPause             bool `json:"auto_pause"`
	MaxCapacity           int  `json:"max_capacity"`
	MinCapacity           int  `json:"min_capacity"`
	SecondsUntilAutoPause int  `json:"seconds_until_auto_pause"`
}

type AwsRdsClusterSpecS3Import struct {
	SourceEngine        string `json:"source_engine"`
	SourceEngineVersion string `json:"source_engine_version"`
	BucketName          string `json:"bucket_name"`
	BucketPrefix        string `json:"bucket_prefix"`
	IngestionRole       string `json:"ingestion_role"`
}

type AwsRdsClusterSpec struct {
	ScalingConfiguration             []AwsRdsClusterSpec `json:"scaling_configuration"`
	BackupRetentionPeriod            int                 `json:"backup_retention_period"`
	BacktrackWindow                  int                 `json:"backtrack_window"`
	ClusterIdentifier                string              `json:"cluster_identifier"`
	ClusterIdentifierPrefix          string              `json:"cluster_identifier_prefix"`
	ClusterMembers                   []string            `json:"cluster_members"`
	DatabaseName                     string              `json:"database_name"`
	DbSubnetGroupName                string              `json:"db_subnet_group_name"`
	ReplicationSourceIdentifier      string              `json:"replication_source_identifier"`
	StorageEncrypted                 bool                `json:"storage_encrypted"`
	FinalSnapshotIdentifier          string              `json:"final_snapshot_identifier"`
	MasterUsername                   string              `json:"master_username"`
	Tags                             map[string]string   `json:"tags"`
	Engine                           string              `json:"engine"`
	ApplyImmediately                 bool                `json:"apply_immediately"`
	VpcSecurityGroupIds              []string            `json:"vpc_security_group_ids"`
	EnabledCloudwatchLogsExports     []string            `json:"enabled_cloudwatch_logs_exports"`
	AvailabilityZones                []string            `json:"availability_zones"`
	GlobalClusterIdentifier          string              `json:"global_cluster_identifier"`
	PreferredMaintenanceWindow       string              `json:"preferred_maintenance_window"`
	KmsKeyId                         string              `json:"kms_key_id"`
	IamDatabaseAuthenticationEnabled bool                `json:"iam_database_authentication_enabled"`
	ClusterResourceId                string              `json:"cluster_resource_id"`
	S3Import                         []AwsRdsClusterSpec `json:"s3_import"`
	SkipFinalSnapshot                bool                `json:"skip_final_snapshot"`
	MasterPassword                   string              `json:"master_password"`
	SourceRegion                     string              `json:"source_region"`
	CopyTagsToSnapshot               bool                `json:"copy_tags_to_snapshot"`
	DeletionProtection               bool                `json:"deletion_protection"`
	ReaderEndpoint                   string              `json:"reader_endpoint"`
	HostedZoneId                     string              `json:"hosted_zone_id"`
	EngineMode                       string              `json:"engine_mode"`
	EngineVersion                    string              `json:"engine_version"`
	Arn                              string              `json:"arn"`
	Endpoint                         string              `json:"endpoint"`
	DbClusterParameterGroupName      string              `json:"db_cluster_parameter_group_name"`
	SnapshotIdentifier               string              `json:"snapshot_identifier"`
	Port                             int                 `json:"port"`
	PreferredBackupWindow            string              `json:"preferred_backup_window"`
	IamRoles                         []string            `json:"iam_roles"`
}



type AwsRdsClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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