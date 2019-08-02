package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DocdbCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DocdbClusterSpec   `json:"spec,omitempty"`
	Status            DocdbClusterStatus `json:"status,omitempty"`
}

type DocdbClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	ApplyImmediately bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`
	// +optional
	BackupRetentionPeriod int `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period,omitempty"`
	// +optional
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`
	// +optional
	ClusterIdentifierPrefix string `json:"clusterIdentifierPrefix,omitempty" tf:"cluster_identifier_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ClusterMembers []string `json:"clusterMembers,omitempty" tf:"cluster_members,omitempty"`
	// +optional
	DbClusterParameterGroupName string `json:"dbClusterParameterGroupName,omitempty" tf:"db_cluster_parameter_group_name,omitempty"`
	// +optional
	DbSubnetGroupName string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name,omitempty"`
	// +optional
	EnabledCloudwatchLogsExports []string `json:"enabledCloudwatchLogsExports,omitempty" tf:"enabled_cloudwatch_logs_exports,omitempty"`
	// +optional
	Engine string `json:"engine,omitempty" tf:"engine,omitempty"`
	// +optional
	EngineVersion string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`
	// +optional
	FinalSnapshotIdentifier string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`
	// +optional
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	// +optional
	MasterPassword string `json:"-" sensitive:"true" tf:"master_password,omitempty"`
	// +optional
	MasterUsername string `json:"masterUsername,omitempty" tf:"master_username,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	PreferredBackupWindow string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`
	// +optional
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`
	// +optional
	SkipFinalSnapshot bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`
	// +optional
	SnapshotIdentifier string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`
	// +optional
	StorageEncrypted bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VpcSecurityGroupIDS []string `json:"vpcSecurityGroupIDS,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type DocdbClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *DocdbClusterSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DocdbClusterList is a list of DocdbClusters
type DocdbClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DocdbCluster CRD objects
	Items []DocdbCluster `json:"items,omitempty"`
}
