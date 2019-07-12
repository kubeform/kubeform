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

type AwsNeptuneCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNeptuneClusterSpec   `json:"spec,omitempty"`
	Status            AwsNeptuneClusterStatus `json:"status,omitempty"`
}

type AwsNeptuneClusterSpec struct {
	Endpoint                         string            `json:"endpoint"`
	EngineVersion                    string            `json:"engine_version"`
	ClusterMembers                   []string          `json:"cluster_members"`
	FinalSnapshotIdentifier          string            `json:"final_snapshot_identifier"`
	IamDatabaseAuthenticationEnabled bool              `json:"iam_database_authentication_enabled"`
	ReaderEndpoint                   string            `json:"reader_endpoint"`
	KmsKeyArn                        string            `json:"kms_key_arn"`
	NeptuneClusterParameterGroupName string            `json:"neptune_cluster_parameter_group_name"`
	Port                             int               `json:"port"`
	PreferredMaintenanceWindow       string            `json:"preferred_maintenance_window"`
	ClusterIdentifierPrefix          string            `json:"cluster_identifier_prefix"`
	ClusterResourceId                string            `json:"cluster_resource_id"`
	HostedZoneId                     string            `json:"hosted_zone_id"`
	IamRoles                         []string          `json:"iam_roles"`
	SnapshotIdentifier               string            `json:"snapshot_identifier"`
	VpcSecurityGroupIds              []string          `json:"vpc_security_group_ids"`
	Tags                             map[string]string `json:"tags"`
	Arn                              string            `json:"arn"`
	Engine                           string            `json:"engine"`
	PreferredBackupWindow            string            `json:"preferred_backup_window"`
	StorageEncrypted                 bool              `json:"storage_encrypted"`
	NeptuneSubnetGroupName           string            `json:"neptune_subnet_group_name"`
	ReplicationSourceIdentifier      string            `json:"replication_source_identifier"`
	SkipFinalSnapshot                bool              `json:"skip_final_snapshot"`
	ApplyImmediately                 bool              `json:"apply_immediately"`
	AvailabilityZones                []string          `json:"availability_zones"`
	BackupRetentionPeriod            int               `json:"backup_retention_period"`
	ClusterIdentifier                string            `json:"cluster_identifier"`
}

type AwsNeptuneClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsNeptuneClusterList is a list of AwsNeptuneClusters
type AwsNeptuneClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNeptuneCluster CRD objects
	Items []AwsNeptuneCluster `json:"items,omitempty"`
}
