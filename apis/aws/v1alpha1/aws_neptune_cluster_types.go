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
	AvailabilityZones                []string          `json:"availability_zones"`
	ClusterIdentifier                string            `json:"cluster_identifier"`
	ClusterIdentifierPrefix          string            `json:"cluster_identifier_prefix"`
	ReaderEndpoint                   string            `json:"reader_endpoint"`
	StorageEncrypted                 bool              `json:"storage_encrypted"`
	SkipFinalSnapshot                bool              `json:"skip_final_snapshot"`
	SnapshotIdentifier               string            `json:"snapshot_identifier"`
	ReplicationSourceIdentifier      string            `json:"replication_source_identifier"`
	BackupRetentionPeriod            int               `json:"backup_retention_period"`
	HostedZoneId                     string            `json:"hosted_zone_id"`
	PreferredMaintenanceWindow       string            `json:"preferred_maintenance_window"`
	VpcSecurityGroupIds              []string          `json:"vpc_security_group_ids"`
	ClusterResourceId                string            `json:"cluster_resource_id"`
	IamRoles                         []string          `json:"iam_roles"`
	Endpoint                         string            `json:"endpoint"`
	Engine                           string            `json:"engine"`
	PreferredBackupWindow            string            `json:"preferred_backup_window"`
	Tags                             map[string]string `json:"tags"`
	EngineVersion                    string            `json:"engine_version"`
	KmsKeyArn                        string            `json:"kms_key_arn"`
	ApplyImmediately                 bool              `json:"apply_immediately"`
	Arn                              string            `json:"arn"`
	IamDatabaseAuthenticationEnabled bool              `json:"iam_database_authentication_enabled"`
	NeptuneSubnetGroupName           string            `json:"neptune_subnet_group_name"`
	NeptuneClusterParameterGroupName string            `json:"neptune_cluster_parameter_group_name"`
	Port                             int               `json:"port"`
	ClusterMembers                   []string          `json:"cluster_members"`
	FinalSnapshotIdentifier          string            `json:"final_snapshot_identifier"`
}



type AwsNeptuneClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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