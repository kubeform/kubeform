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
	HostedZoneId                 string            `json:"hosted_zone_id"`
	Port                         int               `json:"port"`
	PreferredBackupWindow        string            `json:"preferred_backup_window"`
	AvailabilityZones            []string          `json:"availability_zones"`
	DbClusterParameterGroupName  string            `json:"db_cluster_parameter_group_name"`
	SkipFinalSnapshot            bool              `json:"skip_final_snapshot"`
	ApplyImmediately             bool              `json:"apply_immediately"`
	KmsKeyId                     string            `json:"kms_key_id"`
	Tags                         map[string]string `json:"tags"`
	ClusterIdentifier            string            `json:"cluster_identifier"`
	ClusterMembers               []string          `json:"cluster_members"`
	DbSubnetGroupName            string            `json:"db_subnet_group_name"`
	Endpoint                     string            `json:"endpoint"`
	Engine                       string            `json:"engine"`
	MasterUsername               string            `json:"master_username"`
	EnabledCloudwatchLogsExports []string          `json:"enabled_cloudwatch_logs_exports"`
	FinalSnapshotIdentifier      string            `json:"final_snapshot_identifier"`
	PreferredMaintenanceWindow   string            `json:"preferred_maintenance_window"`
	BackupRetentionPeriod        int               `json:"backup_retention_period"`
	ClusterResourceId            string            `json:"cluster_resource_id"`
	ClusterIdentifierPrefix      string            `json:"cluster_identifier_prefix"`
	SnapshotIdentifier           string            `json:"snapshot_identifier"`
	EngineVersion                string            `json:"engine_version"`
	StorageEncrypted             bool              `json:"storage_encrypted"`
	MasterPassword               string            `json:"master_password"`
	VpcSecurityGroupIds          []string          `json:"vpc_security_group_ids"`
	Arn                          string            `json:"arn"`
	ReaderEndpoint               string            `json:"reader_endpoint"`
}

type AwsDocdbClusterStatus struct {
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
