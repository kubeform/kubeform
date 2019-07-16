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

type RdsCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RdsClusterSpec   `json:"spec,omitempty"`
	Status            RdsClusterStatus `json:"status,omitempty"`
}

type RdsClusterSpecS3Import struct {
	BucketName string `json:"bucket_name"`
	// +optional
	BucketPrefix        string `json:"bucket_prefix,omitempty"`
	IngestionRole       string `json:"ingestion_role"`
	SourceEngine        string `json:"source_engine"`
	SourceEngineVersion string `json:"source_engine_version"`
}

type RdsClusterSpecScalingConfiguration struct {
	// +optional
	AutoPause bool `json:"auto_pause,omitempty"`
	// +optional
	MaxCapacity int `json:"max_capacity,omitempty"`
	// +optional
	MinCapacity int `json:"min_capacity,omitempty"`
	// +optional
	SecondsUntilAutoPause int `json:"seconds_until_auto_pause,omitempty"`
}

type RdsClusterSpec struct {
	// +optional
	BacktrackWindow int `json:"backtrack_window,omitempty"`
	// +optional
	BackupRetentionPeriod int `json:"backup_retention_period,omitempty"`
	// +optional
	CopyTagsToSnapshot bool `json:"copy_tags_to_snapshot,omitempty"`
	// +optional
	DeletionProtection bool `json:"deletion_protection,omitempty"`
	// +optional
	EnabledCloudwatchLogsExports []string `json:"enabled_cloudwatch_logs_exports,omitempty"`
	// +optional
	Engine string `json:"engine,omitempty"`
	// +optional
	EngineMode string `json:"engine_mode,omitempty"`
	// +optional
	FinalSnapshotIdentifier string `json:"final_snapshot_identifier,omitempty"`
	// +optional
	GlobalClusterIdentifier string `json:"global_cluster_identifier,omitempty"`
	// +optional
	IamDatabaseAuthenticationEnabled bool `json:"iam_database_authentication_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	IamRoles []string `json:"iam_roles,omitempty"`
	// +optional
	MasterPassword string `json:"master_password,omitempty"`
	// +optional
	ReplicationSourceIdentifier string `json:"replication_source_identifier,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3Import *[]RdsClusterSpec `json:"s3_import,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ScalingConfiguration *[]RdsClusterSpec `json:"scaling_configuration,omitempty"`
	// +optional
	SkipFinalSnapshot bool `json:"skip_final_snapshot,omitempty"`
	// +optional
	SnapshotIdentifier string `json:"snapshot_identifier,omitempty"`
	// +optional
	SourceRegion string `json:"source_region,omitempty"`
	// +optional
	StorageEncrypted bool `json:"storage_encrypted,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type RdsClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
