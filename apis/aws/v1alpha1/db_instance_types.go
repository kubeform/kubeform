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

type DbInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbInstanceSpec   `json:"spec,omitempty"`
	Status            DbInstanceStatus `json:"status,omitempty"`
}

type DbInstanceSpecS3Import struct {
	BucketName string `json:"bucket_name"`
	// +optional
	BucketPrefix        string `json:"bucket_prefix,omitempty"`
	IngestionRole       string `json:"ingestion_role"`
	SourceEngine        string `json:"source_engine"`
	SourceEngineVersion string `json:"source_engine_version"`
}

type DbInstanceSpec struct {
	// +optional
	AllowMajorVersionUpgrade bool `json:"allow_major_version_upgrade,omitempty"`
	// +optional
	AutoMinorVersionUpgrade bool `json:"auto_minor_version_upgrade,omitempty"`
	// +optional
	CopyTagsToSnapshot bool `json:"copy_tags_to_snapshot,omitempty"`
	// +optional
	DeletionProtection bool `json:"deletion_protection,omitempty"`
	// +optional
	Domain string `json:"domain,omitempty"`
	// +optional
	DomainIamRoleName string `json:"domain_iam_role_name,omitempty"`
	// +optional
	EnabledCloudwatchLogsExports []string `json:"enabled_cloudwatch_logs_exports,omitempty"`
	// +optional
	FinalSnapshotIdentifier string `json:"final_snapshot_identifier,omitempty"`
	// +optional
	IamDatabaseAuthenticationEnabled bool   `json:"iam_database_authentication_enabled,omitempty"`
	InstanceClass                    string `json:"instance_class"`
	// +optional
	Iops int `json:"iops,omitempty"`
	// +optional
	MaxAllocatedStorage int `json:"max_allocated_storage,omitempty"`
	// +optional
	MonitoringInterval int `json:"monitoring_interval,omitempty"`
	// +optional
	Password string `json:"password,omitempty"`
	// +optional
	PerformanceInsightsEnabled bool `json:"performance_insights_enabled,omitempty"`
	// +optional
	PubliclyAccessible bool `json:"publicly_accessible,omitempty"`
	// +optional
	ReplicateSourceDb string `json:"replicate_source_db,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3Import *[]DbInstanceSpec `json:"s3_import,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupNames []string `json:"security_group_names,omitempty"`
	// +optional
	SkipFinalSnapshot bool `json:"skip_final_snapshot,omitempty"`
	// +optional
	SnapshotIdentifier string `json:"snapshot_identifier,omitempty"`
	// +optional
	StorageEncrypted bool `json:"storage_encrypted,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DbInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DbInstanceList is a list of DbInstances
type DbInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DbInstance CRD objects
	Items []DbInstance `json:"items,omitempty"`
}
