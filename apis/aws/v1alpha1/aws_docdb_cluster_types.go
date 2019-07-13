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

type AwsDocdbCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDocdbClusterSpec   `json:"spec,omitempty"`
	Status            AwsDocdbClusterStatus `json:"status,omitempty"`
}

type AwsDocdbClusterSpec struct {
	MasterUsername               string            `json:"master_username"`
	PreferredMaintenanceWindow   string            `json:"preferred_maintenance_window"`
	HostedZoneId                 string            `json:"hosted_zone_id"`
	Port                         int               `json:"port"`
	DbClusterParameterGroupName  string            `json:"db_cluster_parameter_group_name"`
	FinalSnapshotIdentifier      string            `json:"final_snapshot_identifier"`
	DbSubnetGroupName            string            `json:"db_subnet_group_name"`
	Endpoint                     string            `json:"endpoint"`
	EngineVersion                string            `json:"engine_version"`
	ClusterIdentifierPrefix      string            `json:"cluster_identifier_prefix"`
	BackupRetentionPeriod        int               `json:"backup_retention_period"`
	VpcSecurityGroupIds          []string          `json:"vpc_security_group_ids"`
	ClusterResourceId            string            `json:"cluster_resource_id"`
	KmsKeyId                     string            `json:"kms_key_id"`
	Engine                       string            `json:"engine"`
	SnapshotIdentifier           string            `json:"snapshot_identifier"`
	ApplyImmediately             bool              `json:"apply_immediately"`
	ReaderEndpoint               string            `json:"reader_endpoint"`
	AvailabilityZones            []string          `json:"availability_zones"`
	ClusterIdentifier            string            `json:"cluster_identifier"`
	ClusterMembers               []string          `json:"cluster_members"`
	StorageEncrypted             bool              `json:"storage_encrypted"`
	SkipFinalSnapshot            bool              `json:"skip_final_snapshot"`
	MasterPassword               string            `json:"master_password"`
	PreferredBackupWindow        string            `json:"preferred_backup_window"`
	Arn                          string            `json:"arn"`
	Tags                         map[string]string `json:"tags"`
	EnabledCloudwatchLogsExports []string          `json:"enabled_cloudwatch_logs_exports"`
}



type AwsDocdbClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDocdbClusterList is a list of AwsDocdbClusters
type AwsDocdbClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDocdbCluster CRD objects
	Items []AwsDocdbCluster `json:"items,omitempty"`
}