package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDocdbCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDocdbClusterSpec   `json:"spec,omitempty"`
	Status            AwsDocdbClusterStatus `json:"status,omitempty"`
}

type AwsDocdbClusterSpec struct {
	PreferredMaintenanceWindow   string            `json:"preferred_maintenance_window"`
	KmsKeyId                     string            `json:"kms_key_id"`
	ClusterResourceId            string            `json:"cluster_resource_id"`
	ClusterIdentifier            string            `json:"cluster_identifier"`
	ClusterMembers               []string          `json:"cluster_members"`
	HostedZoneId                 string            `json:"hosted_zone_id"`
	FinalSnapshotIdentifier      string            `json:"final_snapshot_identifier"`
	DbSubnetGroupName            string            `json:"db_subnet_group_name"`
	Engine                       string            `json:"engine"`
	SnapshotIdentifier           string            `json:"snapshot_identifier"`
	Tags                         map[string]string `json:"tags"`
	EngineVersion                string            `json:"engine_version"`
	MasterPassword               string            `json:"master_password"`
	EnabledCloudwatchLogsExports []string          `json:"enabled_cloudwatch_logs_exports"`
	Endpoint                     string            `json:"endpoint"`
	StorageEncrypted             bool              `json:"storage_encrypted"`
	PreferredBackupWindow        string            `json:"preferred_backup_window"`
	ReaderEndpoint               string            `json:"reader_endpoint"`
	SkipFinalSnapshot            bool              `json:"skip_final_snapshot"`
	Port                         int               `json:"port"`
	VpcSecurityGroupIds          []string          `json:"vpc_security_group_ids"`
	Arn                          string            `json:"arn"`
	ClusterIdentifierPrefix      string            `json:"cluster_identifier_prefix"`
	MasterUsername               string            `json:"master_username"`
	ApplyImmediately             bool              `json:"apply_immediately"`
	AvailabilityZones            []string          `json:"availability_zones"`
	DbClusterParameterGroupName  string            `json:"db_cluster_parameter_group_name"`
	BackupRetentionPeriod        int               `json:"backup_retention_period"`
}

type AwsDocdbClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDocdbClusterList is a list of AwsDocdbClusters
type AwsDocdbClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDocdbCluster CRD objects
	Items []AwsDocdbCluster `json:"items,omitempty"`
}
