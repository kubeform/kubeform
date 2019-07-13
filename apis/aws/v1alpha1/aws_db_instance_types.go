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

type AwsDbInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDbInstanceSpec   `json:"spec,omitempty"`
	Status            AwsDbInstanceStatus `json:"status,omitempty"`
}

type AwsDbInstanceSpecS3Import struct {
	BucketName          string `json:"bucket_name"`
	BucketPrefix        string `json:"bucket_prefix"`
	IngestionRole       string `json:"ingestion_role"`
	SourceEngine        string `json:"source_engine"`
	SourceEngineVersion string `json:"source_engine_version"`
}

type AwsDbInstanceSpec struct {
	MultiAz                            bool                `json:"multi_az"`
	Port                               int                 `json:"port"`
	Address                            string              `json:"address"`
	KmsKeyId                           string              `json:"kms_key_id"`
	IamDatabaseAuthenticationEnabled   bool                `json:"iam_database_authentication_enabled"`
	PerformanceInsightsEnabled         bool                `json:"performance_insights_enabled"`
	Iops                               int                 `json:"iops"`
	SecurityGroupNames                 []string            `json:"security_group_names"`
	S3Import                           []AwsDbInstanceSpec `json:"s3_import"`
	Arn                                string              `json:"arn"`
	VpcSecurityGroupIds                []string            `json:"vpc_security_group_ids"`
	MonitoringInterval                 int                 `json:"monitoring_interval"`
	PerformanceInsightsRetentionPeriod int                 `json:"performance_insights_retention_period"`
	InstanceClass                      string              `json:"instance_class"`
	ParameterGroupName                 string              `json:"parameter_group_name"`
	Status                             string              `json:"status"`
	SnapshotIdentifier                 string              `json:"snapshot_identifier"`
	Tags                               map[string]string   `json:"tags"`
	Password                           string              `json:"password"`
	MaxAllocatedStorage                int                 `json:"max_allocated_storage"`
	Endpoint                           string              `json:"endpoint"`
	HostedZoneId                       string              `json:"hosted_zone_id"`
	Username                           string              `json:"username"`
	StorageType                        string              `json:"storage_type"`
	Identifier                         string              `json:"identifier"`
	MaintenanceWindow                  string              `json:"maintenance_window"`
	ApplyImmediately                   bool                `json:"apply_immediately"`
	MonitoringRoleArn                  string              `json:"monitoring_role_arn"`
	CaCertIdentifier                   string              `json:"ca_cert_identifier"`
	AutoMinorVersionUpgrade            bool                `json:"auto_minor_version_upgrade"`
	Timezone                           string              `json:"timezone"`
	StorageEncrypted                   bool                `json:"storage_encrypted"`
	Domain                             string              `json:"domain"`
	Name                               string              `json:"name"`
	Engine                             string              `json:"engine"`
	IdentifierPrefix                   string              `json:"identifier_prefix"`
	DeletionProtection                 bool                `json:"deletion_protection"`
	EngineVersion                      string              `json:"engine_version"`
	DbSubnetGroupName                  string              `json:"db_subnet_group_name"`
	EnabledCloudwatchLogsExports       []string            `json:"enabled_cloudwatch_logs_exports"`
	AllocatedStorage                   int                 `json:"allocated_storage"`
	OptionGroupName                    string              `json:"option_group_name"`
	LicenseModel                       string              `json:"license_model"`
	SkipFinalSnapshot                  bool                `json:"skip_final_snapshot"`
	ReplicateSourceDb                  string              `json:"replicate_source_db"`
	ResourceId                         string              `json:"resource_id"`
	BackupRetentionPeriod              int                 `json:"backup_retention_period"`
	BackupWindow                       string              `json:"backup_window"`
	PerformanceInsightsKmsKeyId        string              `json:"performance_insights_kms_key_id"`
	AvailabilityZone                   string              `json:"availability_zone"`
	PubliclyAccessible                 bool                `json:"publicly_accessible"`
	FinalSnapshotIdentifier            string              `json:"final_snapshot_identifier"`
	CopyTagsToSnapshot                 bool                `json:"copy_tags_to_snapshot"`
	Replicas                           []string            `json:"replicas"`
	DomainIamRoleName                  string              `json:"domain_iam_role_name"`
	CharacterSetName                   string              `json:"character_set_name"`
	AllowMajorVersionUpgrade           bool                `json:"allow_major_version_upgrade"`
}



type AwsDbInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDbInstanceList is a list of AwsDbInstances
type AwsDbInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDbInstance CRD objects
	Items []AwsDbInstance `json:"items,omitempty"`
}