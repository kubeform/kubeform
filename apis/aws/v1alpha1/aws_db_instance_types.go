package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	Name                             string              `json:"name"`
	Port                             int                 `json:"port"`
	IamDatabaseAuthenticationEnabled bool                `json:"iam_database_authentication_enabled"`
	EnabledCloudwatchLogsExports     []string            `json:"enabled_cloudwatch_logs_exports"`
	DomainIamRoleName                string              `json:"domain_iam_role_name"`
	S3Import                         []AwsDbInstanceSpec `json:"s3_import"`
	HostedZoneId                     string              `json:"hosted_zone_id"`
	ReplicateSourceDb                string              `json:"replicate_source_db"`
	SnapshotIdentifier               string              `json:"snapshot_identifier"`
	Password                         string              `json:"password"`
	Identifier                       string              `json:"identifier"`
	LicenseModel                     string              `json:"license_model"`
	FinalSnapshotIdentifier          string              `json:"final_snapshot_identifier"`
	KmsKeyId                         string              `json:"kms_key_id"`
	Arn                              string              `json:"arn"`
	VpcSecurityGroupIds              []string            `json:"vpc_security_group_ids"`
	AutoMinorVersionUpgrade          bool                `json:"auto_minor_version_upgrade"`
	AllowMajorVersionUpgrade         bool                `json:"allow_major_version_upgrade"`
	Timezone                         string              `json:"timezone"`
	SkipFinalSnapshot                bool                `json:"skip_final_snapshot"`
	ResourceId                       string              `json:"resource_id"`
	StorageEncrypted                 bool                `json:"storage_encrypted"`
	IdentifierPrefix                 string              `json:"identifier_prefix"`
	DbSubnetGroupName                string              `json:"db_subnet_group_name"`
	Status                           string              `json:"status"`
	BackupRetentionPeriod            int                 `json:"backup_retention_period"`
	BackupWindow                     string              `json:"backup_window"`
	MonitoringInterval               int                 `json:"monitoring_interval"`
	Engine                           string              `json:"engine"`
	AvailabilityZone                 string              `json:"availability_zone"`
	Iops                             int                 `json:"iops"`
	CaCertIdentifier                 string              `json:"ca_cert_identifier"`
	CharacterSetName                 string              `json:"character_set_name"`
	ParameterGroupName               string              `json:"parameter_group_name"`
	Domain                           string              `json:"domain"`
	Username                         string              `json:"username"`
	AllocatedStorage                 int                 `json:"allocated_storage"`
	MultiAz                          bool                `json:"multi_az"`
	Endpoint                         string              `json:"endpoint"`
	SecurityGroupNames               []string            `json:"security_group_names"`
	Address                          string              `json:"address"`
	ApplyImmediately                 bool                `json:"apply_immediately"`
	EngineVersion                    string              `json:"engine_version"`
	MaintenanceWindow                string              `json:"maintenance_window"`
	StorageType                      string              `json:"storage_type"`
	CopyTagsToSnapshot               bool                `json:"copy_tags_to_snapshot"`
	Replicas                         []string            `json:"replicas"`
	MonitoringRoleArn                string              `json:"monitoring_role_arn"`
	OptionGroupName                  string              `json:"option_group_name"`
	DeletionProtection               bool                `json:"deletion_protection"`
	InstanceClass                    string              `json:"instance_class"`
	PubliclyAccessible               bool                `json:"publicly_accessible"`
	Tags                             map[string]string   `json:"tags"`
}



type AwsDbInstanceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbInstanceList is a list of AwsDbInstances
type AwsDbInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDbInstance CRD objects
	Items []AwsDbInstance `json:"items,omitempty"`
}