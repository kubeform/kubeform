package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type DbInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbInstanceSpec   `json:"spec,omitempty"`
	Status            DbInstanceStatus `json:"status,omitempty"`
}

type DbInstanceSpecS3Import struct {
	BucketName string `json:"bucketName" tf:"bucket_name"`
	// +optional
	BucketPrefix        string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`
	IngestionRole       string `json:"ingestionRole" tf:"ingestion_role"`
	SourceEngine        string `json:"sourceEngine" tf:"source_engine"`
	SourceEngineVersion string `json:"sourceEngineVersion" tf:"source_engine_version"`
}

type DbInstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Address string `json:"address,omitempty" tf:"address,omitempty"`
	// +optional
	AllocatedStorage int `json:"allocatedStorage,omitempty" tf:"allocated_storage,omitempty"`
	// +optional
	AllowMajorVersionUpgrade bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade,omitempty"`
	// +optional
	ApplyImmediately bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`
	// +optional
	AvailabilityZone string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	// +optional
	BackupRetentionPeriod int `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period,omitempty"`
	// +optional
	BackupWindow string `json:"backupWindow,omitempty" tf:"backup_window,omitempty"`
	// +optional
	CaCertIdentifier string `json:"caCertIdentifier,omitempty" tf:"ca_cert_identifier,omitempty"`
	// +optional
	CharacterSetName string `json:"characterSetName,omitempty" tf:"character_set_name,omitempty"`
	// +optional
	CopyTagsToSnapshot bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot,omitempty"`
	// +optional
	DbSubnetGroupName string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name,omitempty"`
	// +optional
	DeletionProtection bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`
	// +optional
	Domain string `json:"domain,omitempty" tf:"domain,omitempty"`
	// +optional
	DomainIamRoleName string `json:"domainIamRoleName,omitempty" tf:"domain_iam_role_name,omitempty"`
	// +optional
	EnabledCloudwatchLogsExports []string `json:"enabledCloudwatchLogsExports,omitempty" tf:"enabled_cloudwatch_logs_exports,omitempty"`
	// +optional
	Endpoint string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
	// +optional
	Engine string `json:"engine,omitempty" tf:"engine,omitempty"`
	// +optional
	EngineVersion string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`
	// +optional
	FinalSnapshotIdentifier string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`
	// +optional
	HostedZoneID string `json:"hostedZoneID,omitempty" tf:"hosted_zone_id,omitempty"`
	// +optional
	IamDatabaseAuthenticationEnabled bool `json:"iamDatabaseAuthenticationEnabled,omitempty" tf:"iam_database_authentication_enabled,omitempty"`
	// +optional
	Identifier string `json:"identifier,omitempty" tf:"identifier,omitempty"`
	// +optional
	IdentifierPrefix string `json:"identifierPrefix,omitempty" tf:"identifier_prefix,omitempty"`
	InstanceClass    string `json:"instanceClass" tf:"instance_class"`
	// +optional
	Iops int `json:"iops,omitempty" tf:"iops,omitempty"`
	// +optional
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	// +optional
	LicenseModel string `json:"licenseModel,omitempty" tf:"license_model,omitempty"`
	// +optional
	MaintenanceWindow string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`
	// +optional
	MonitoringInterval int `json:"monitoringInterval,omitempty" tf:"monitoring_interval,omitempty"`
	// +optional
	MonitoringRoleArn string `json:"monitoringRoleArn,omitempty" tf:"monitoring_role_arn,omitempty"`
	// +optional
	MultiAz bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	OptionGroupName string `json:"optionGroupName,omitempty" tf:"option_group_name,omitempty"`
	// +optional
	ParameterGroupName string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`
	// +optional
	Replicas []string `json:"replicas,omitempty" tf:"replicas,omitempty"`
	// +optional
	ReplicateSourceDb string `json:"replicateSourceDb,omitempty" tf:"replicate_source_db,omitempty"`
	// +optional
	ResourceID string `json:"resourceID,omitempty" tf:"resource_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3Import []DbInstanceSpecS3Import `json:"s3Import,omitempty" tf:"s3_import,omitempty"`
	// +optional
	SecurityGroupNames []string `json:"securityGroupNames,omitempty" tf:"security_group_names,omitempty"`
	// +optional
	SkipFinalSnapshot bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`
	// +optional
	SnapshotIdentifier string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	StorageEncrypted bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`
	// +optional
	StorageType string `json:"storageType,omitempty" tf:"storage_type,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Timezone string `json:"timezone,omitempty" tf:"timezone,omitempty"`
	// +optional
	Username string `json:"username,omitempty" tf:"username,omitempty"`
	// +optional
	VpcSecurityGroupIDS []string `json:"vpcSecurityGroupIDS,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type DbInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DbInstanceSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
