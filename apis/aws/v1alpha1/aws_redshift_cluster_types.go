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

type AwsRedshiftCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRedshiftClusterSpec   `json:"spec,omitempty"`
	Status            AwsRedshiftClusterStatus `json:"status,omitempty"`
}

type AwsRedshiftClusterSpecLogging struct {
	Enable      bool   `json:"enable"`
	BucketName  string `json:"bucket_name"`
	S3KeyPrefix string `json:"s3_key_prefix"`
}

type AwsRedshiftClusterSpecSnapshotCopy struct {
	GrantName         string `json:"grant_name"`
	DestinationRegion string `json:"destination_region"`
	RetentionPeriod   int    `json:"retention_period"`
}

type AwsRedshiftClusterSpec struct {
	PreferredMaintenanceWindow       string                   `json:"preferred_maintenance_window"`
	Logging                          []AwsRedshiftClusterSpec `json:"logging"`
	EnableLogging                    bool                     `json:"enable_logging"`
	OwnerAccount                     string                   `json:"owner_account"`
	Tags                             map[string]string        `json:"tags"`
	ClusterSecurityGroups            []string                 `json:"cluster_security_groups"`
	VpcSecurityGroupIds              []string                 `json:"vpc_security_group_ids"`
	AllowVersionUpgrade              bool                     `json:"allow_version_upgrade"`
	S3KeyPrefix                      string                   `json:"s3_key_prefix"`
	SnapshotClusterIdentifier        string                   `json:"snapshot_cluster_identifier"`
	IamRoles                         []string                 `json:"iam_roles"`
	BucketName                       string                   `json:"bucket_name"`
	Arn                              string                   `json:"arn"`
	DatabaseName                     string                   `json:"database_name"`
	Endpoint                         string                   `json:"endpoint"`
	DnsName                          string                   `json:"dns_name"`
	SnapshotCopy                     []AwsRedshiftClusterSpec `json:"snapshot_copy"`
	AvailabilityZone                 string                   `json:"availability_zone"`
	Port                             int                      `json:"port"`
	PubliclyAccessible               bool                     `json:"publicly_accessible"`
	KmsKeyId                         string                   `json:"kms_key_id"`
	SkipFinalSnapshot                bool                     `json:"skip_final_snapshot"`
	ClusterPublicKey                 string                   `json:"cluster_public_key"`
	ClusterRevisionNumber            string                   `json:"cluster_revision_number"`
	SnapshotIdentifier               string                   `json:"snapshot_identifier"`
	ClusterSubnetGroupName           string                   `json:"cluster_subnet_group_name"`
	ClusterParameterGroupName        string                   `json:"cluster_parameter_group_name"`
	AutomatedSnapshotRetentionPeriod int                      `json:"automated_snapshot_retention_period"`
	ClusterVersion                   string                   `json:"cluster_version"`
	Encrypted                        bool                     `json:"encrypted"`
	MasterUsername                   string                   `json:"master_username"`
	MasterPassword                   string                   `json:"master_password"`
	ElasticIp                        string                   `json:"elastic_ip"`
	FinalSnapshotIdentifier          string                   `json:"final_snapshot_identifier"`
	ClusterIdentifier                string                   `json:"cluster_identifier"`
	ClusterType                      string                   `json:"cluster_type"`
	NodeType                         string                   `json:"node_type"`
	NumberOfNodes                    int                      `json:"number_of_nodes"`
	EnhancedVpcRouting               bool                     `json:"enhanced_vpc_routing"`
}



type AwsRedshiftClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRedshiftClusterList is a list of AwsRedshiftClusters
type AwsRedshiftClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRedshiftCluster CRD objects
	Items []AwsRedshiftCluster `json:"items,omitempty"`
}