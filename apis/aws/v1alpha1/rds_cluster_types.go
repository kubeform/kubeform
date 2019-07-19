package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type RdsCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RdsClusterSpec   `json:"spec,omitempty"`
	Status            RdsClusterStatus `json:"status,omitempty"`
}

type RdsClusterSpecS3Import struct {
	BucketName string `json:"bucketName" tf:"bucket_name"`
	// +optional
	BucketPrefix        string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`
	IngestionRole       string `json:"ingestionRole" tf:"ingestion_role"`
	SourceEngine        string `json:"sourceEngine" tf:"source_engine"`
	SourceEngineVersion string `json:"sourceEngineVersion" tf:"source_engine_version"`
}

type RdsClusterSpecScalingConfiguration struct {
	// +optional
	AutoPause bool `json:"autoPause,omitempty" tf:"auto_pause,omitempty"`
	// +optional
	MaxCapacity int `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`
	// +optional
	MinCapacity int `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`
	// +optional
	SecondsUntilAutoPause int `json:"secondsUntilAutoPause,omitempty" tf:"seconds_until_auto_pause,omitempty"`
}

type RdsClusterSpec struct {
	// +optional
	ApplyImmediately bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`
	// +optional
	BacktrackWindow int `json:"backtrackWindow,omitempty" tf:"backtrack_window,omitempty"`
	// +optional
	BackupRetentionPeriod int `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period,omitempty"`
	// +optional
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`
	// +optional
	ClusterIdentifierPrefix string `json:"clusterIdentifierPrefix,omitempty" tf:"cluster_identifier_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ClusterMembers []string `json:"clusterMembers,omitempty" tf:"cluster_members,omitempty"`
	// +optional
	CopyTagsToSnapshot bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot,omitempty"`
	// +optional
	DatabaseName string `json:"databaseName,omitempty" tf:"database_name,omitempty"`
	// +optional
	DbClusterParameterGroupName string `json:"dbClusterParameterGroupName,omitempty" tf:"db_cluster_parameter_group_name,omitempty"`
	// +optional
	DbSubnetGroupName string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name,omitempty"`
	// +optional
	DeletionProtection bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`
	// +optional
	EnabledCloudwatchLogsExports []string `json:"enabledCloudwatchLogsExports,omitempty" tf:"enabled_cloudwatch_logs_exports,omitempty"`
	// +optional
	Engine string `json:"engine,omitempty" tf:"engine,omitempty"`
	// +optional
	EngineMode string `json:"engineMode,omitempty" tf:"engine_mode,omitempty"`
	// +optional
	EngineVersion string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`
	// +optional
	FinalSnapshotIdentifier string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`
	// +optional
	GlobalClusterIdentifier string `json:"globalClusterIdentifier,omitempty" tf:"global_cluster_identifier,omitempty"`
	// +optional
	IamDatabaseAuthenticationEnabled bool `json:"iamDatabaseAuthenticationEnabled,omitempty" tf:"iam_database_authentication_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	IamRoles []string `json:"iamRoles,omitempty" tf:"iam_roles,omitempty"`
	// +optional
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	MasterPassword core.LocalObjectReference `json:"masterPassword,omitempty" tf:"master_password,omitempty"`
	// +optional
	MasterUsername string `json:"masterUsername,omitempty" tf:"master_username,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	PreferredBackupWindow string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`
	// +optional
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`
	// +optional
	ReplicationSourceIdentifier string `json:"replicationSourceIdentifier,omitempty" tf:"replication_source_identifier,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3Import []RdsClusterSpecS3Import `json:"s3Import,omitempty" tf:"s3_import,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ScalingConfiguration []RdsClusterSpecScalingConfiguration `json:"scalingConfiguration,omitempty" tf:"scaling_configuration,omitempty"`
	// +optional
	SkipFinalSnapshot bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`
	// +optional
	SnapshotIdentifier string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`
	// +optional
	SourceRegion string `json:"sourceRegion,omitempty" tf:"source_region,omitempty"`
	// +optional
	StorageEncrypted bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VpcSecurityGroupIDS []string                  `json:"vpcSecurityGroupIDS,omitempty" tf:"vpc_security_group_ids,omitempty"`
	ProviderRef         core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type RdsClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RdsClusterList is a list of RdsClusters
type RdsClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RdsCluster CRD objects
	Items []RdsCluster `json:"items,omitempty"`
}
