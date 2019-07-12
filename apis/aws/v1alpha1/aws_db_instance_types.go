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
	IngestionRole       string `json:"ingestion_role"`
	SourceEngine        string `json:"source_engine"`
	SourceEngineVersion string `json:"source_engine_version"`
	BucketName          string `json:"bucket_name"`
	BucketPrefix        string `json:"bucket_prefix"`
}

type AwsDbInstanceSpec struct {
	StorageType                        string              `json:"storage_type"`
	Iops                               int                 `json:"iops"`
	MaxAllocatedStorage                int                 `json:"max_allocated_storage"`
	DomainIamRoleName                  string              `json:"domain_iam_role_name"`
	PerformanceInsightsRetentionPeriod int                 `json:"performance_insights_retention_period"`
	FinalSnapshotIdentifier            string              `json:"final_snapshot_identifier"`
	MonitoringRoleArn                  string              `json:"monitoring_role_arn"`
	PubliclyAccessible                 bool                `json:"publicly_accessible"`
	Address                            string              `json:"address"`
	Replicas                           []string            `json:"replicas"`
	Arn                                string              `json:"arn"`
	Username                           string              `json:"username"`
	S3Import                           []AwsDbInstanceSpec `json:"s3_import"`
	Endpoint                           string              `json:"endpoint"`
	Password                           string              `json:"password"`
	Engine                             string              `json:"engine"`
	EngineVersion                      string              `json:"engine_version"`
	StorageEncrypted                   bool                `json:"storage_encrypted"`
	Identifier                         string              `json:"identifier"`
	MaintenanceWindow                  string              `json:"maintenance_window"`
	Timezone                           string              `json:"timezone"`
	EnabledCloudwatchLogsExports       []string            `json:"enabled_cloudwatch_logs_exports"`
	Name                               string              `json:"name"`
	CharacterSetName                   string              `json:"character_set_name"`
	SecurityGroupNames                 []string            `json:"security_group_names"`
	PerformanceInsightsEnabled         bool                `json:"performance_insights_enabled"`
	IamDatabaseAuthenticationEnabled   bool                `json:"iam_database_authentication_enabled"`
	ResourceId                         string              `json:"resource_id"`
	PerformanceInsightsKmsKeyId        string              `json:"performance_insights_kms_key_id"`
	AllocatedStorage                   int                 `json:"allocated_storage"`
	AvailabilityZone                   string              `json:"availability_zone"`
	Port                               int                 `json:"port"`
	AllowMajorVersionUpgrade           bool                `json:"allow_major_version_upgrade"`
	IdentifierPrefix                   string              `json:"identifier_prefix"`
	SkipFinalSnapshot                  bool                `json:"skip_final_snapshot"`
	ApplyImmediately                   bool                `json:"apply_immediately"`
	KmsKeyId                           string              `json:"kms_key_id"`
	Domain                             string              `json:"domain"`
	InstanceClass                      string              `json:"instance_class"`
	VpcSecurityGroupIds                []string            `json:"vpc_security_group_ids"`
	HostedZoneId                       string              `json:"hosted_zone_id"`
	Status                             string              `json:"status"`
	BackupRetentionPeriod              int                 `json:"backup_retention_period"`
	MultiAz                            bool                `json:"multi_az"`
	SnapshotIdentifier                 string              `json:"snapshot_identifier"`
	DbSubnetGroupName                  string              `json:"db_subnet_group_name"`
	ParameterGroupName                 string              `json:"parameter_group_name"`
	MonitoringInterval                 int                 `json:"monitoring_interval"`
	Tags                               map[string]string   `json:"tags"`
	DeletionProtection                 bool                `json:"deletion_protection"`
	BackupWindow                       string              `json:"backup_window"`
	LicenseModel                       string              `json:"license_model"`
	CopyTagsToSnapshot                 bool                `json:"copy_tags_to_snapshot"`
	AutoMinorVersionUpgrade            bool                `json:"auto_minor_version_upgrade"`
	OptionGroupName                    string              `json:"option_group_name"`
	CaCertIdentifier                   string              `json:"ca_cert_identifier"`
	ReplicateSourceDb                  string              `json:"replicate_source_db"`
}

type AwsDbInstanceStatus struct {
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
