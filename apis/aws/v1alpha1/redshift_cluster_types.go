package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type RedshiftCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedshiftClusterSpec   `json:"spec,omitempty"`
	Status            RedshiftClusterStatus `json:"status,omitempty"`
}

type RedshiftClusterSpecLogging struct {
	// +optional
	BucketName string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`
	Enable     bool   `json:"enable" tf:"enable"`
	// +optional
	S3KeyPrefix string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
}

type RedshiftClusterSpecSnapshotCopy struct {
	DestinationRegion string `json:"destinationRegion" tf:"destination_region"`
	// +optional
	GrantName string `json:"grantName,omitempty" tf:"grant_name,omitempty"`
	// +optional
	RetentionPeriod int `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`
}

type RedshiftClusterSpec struct {
	// +optional
	AllowVersionUpgrade bool `json:"allowVersionUpgrade,omitempty" tf:"allow_version_upgrade,omitempty"`
	// +optional
	AutomatedSnapshotRetentionPeriod int `json:"automatedSnapshotRetentionPeriod,omitempty" tf:"automated_snapshot_retention_period,omitempty"`
	// +optional
	AvailabilityZone  string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	ClusterIdentifier string `json:"clusterIdentifier" tf:"cluster_identifier"`
	// +optional
	ClusterParameterGroupName string `json:"clusterParameterGroupName,omitempty" tf:"cluster_parameter_group_name,omitempty"`
	// +optional
	ClusterPublicKey string `json:"clusterPublicKey,omitempty" tf:"cluster_public_key,omitempty"`
	// +optional
	ClusterRevisionNumber string `json:"clusterRevisionNumber,omitempty" tf:"cluster_revision_number,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ClusterSecurityGroups []string `json:"clusterSecurityGroups,omitempty" tf:"cluster_security_groups,omitempty"`
	// +optional
	ClusterSubnetGroupName string `json:"clusterSubnetGroupName,omitempty" tf:"cluster_subnet_group_name,omitempty"`
	// +optional
	ClusterType string `json:"clusterType,omitempty" tf:"cluster_type,omitempty"`
	// +optional
	ClusterVersion string `json:"clusterVersion,omitempty" tf:"cluster_version,omitempty"`
	// +optional
	DatabaseName string `json:"databaseName,omitempty" tf:"database_name,omitempty"`
	// +optional
	ElasticIP string `json:"elasticIP,omitempty" tf:"elastic_ip,omitempty"`
	// +optional
	Encrypted bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`
	// +optional
	Endpoint string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
	// +optional
	EnhancedVpcRouting bool `json:"enhancedVpcRouting,omitempty" tf:"enhanced_vpc_routing,omitempty"`
	// +optional
	FinalSnapshotIdentifier string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	IamRoles []string `json:"iamRoles,omitempty" tf:"iam_roles,omitempty"`
	// +optional
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Logging []RedshiftClusterSpecLogging `json:"logging,omitempty" tf:"logging,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	MasterPassword core.LocalObjectReference `json:"masterPassword,omitempty" tf:"master_password,omitempty"`
	// +optional
	MasterUsername string `json:"masterUsername,omitempty" tf:"master_username,omitempty"`
	NodeType       string `json:"nodeType" tf:"node_type"`
	// +optional
	NumberOfNodes int `json:"numberOfNodes,omitempty" tf:"number_of_nodes,omitempty"`
	// +optional
	OwnerAccount string `json:"ownerAccount,omitempty" tf:"owner_account,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`
	// +optional
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`
	// +optional
	SkipFinalSnapshot bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`
	// +optional
	SnapshotClusterIdentifier string `json:"snapshotClusterIdentifier,omitempty" tf:"snapshot_cluster_identifier,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SnapshotCopy []RedshiftClusterSpecSnapshotCopy `json:"snapshotCopy,omitempty" tf:"snapshot_copy,omitempty"`
	// +optional
	SnapshotIdentifier string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VpcSecurityGroupIDS []string                  `json:"vpcSecurityGroupIDS,omitempty" tf:"vpc_security_group_ids,omitempty"`
	ProviderRef         core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type RedshiftClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RedshiftClusterList is a list of RedshiftClusters
type RedshiftClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RedshiftCluster CRD objects
	Items []RedshiftCluster `json:"items,omitempty"`
}
