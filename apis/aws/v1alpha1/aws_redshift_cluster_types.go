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
	BucketName  string `json:"bucket_name"`
	S3KeyPrefix string `json:"s3_key_prefix"`
	Enable      bool   `json:"enable"`
}

type AwsRedshiftClusterSpecSnapshotCopy struct {
	RetentionPeriod   int    `json:"retention_period"`
	GrantName         string `json:"grant_name"`
	DestinationRegion string `json:"destination_region"`
}

type AwsRedshiftClusterSpec struct {
	NumberOfNodes                    int                      `json:"number_of_nodes"`
	EnhancedVpcRouting               bool                     `json:"enhanced_vpc_routing"`
	S3KeyPrefix                      string                   `json:"s3_key_prefix"`
	OwnerAccount                     string                   `json:"owner_account"`
	Arn                              string                   `json:"arn"`
	ClusterType                      string                   `json:"cluster_type"`
	ClusterVersion                   string                   `json:"cluster_version"`
	AllowVersionUpgrade              bool                     `json:"allow_version_upgrade"`
	ElasticIp                        string                   `json:"elastic_ip"`
	ClusterRevisionNumber            string                   `json:"cluster_revision_number"`
	Logging                          []AwsRedshiftClusterSpec `json:"logging"`
	VpcSecurityGroupIds              []string                 `json:"vpc_security_group_ids"`
	Encrypted                        bool                     `json:"encrypted"`
	Tags                             map[string]string        `json:"tags"`
	SnapshotCopy                     []AwsRedshiftClusterSpec `json:"snapshot_copy"`
	EnableLogging                    bool                     `json:"enable_logging"`
	DatabaseName                     string                   `json:"database_name"`
	MasterUsername                   string                   `json:"master_username"`
	PreferredMaintenanceWindow       string                   `json:"preferred_maintenance_window"`
	Port                             int                      `json:"port"`
	FinalSnapshotIdentifier          string                   `json:"final_snapshot_identifier"`
	Endpoint                         string                   `json:"endpoint"`
	ClusterPublicKey                 string                   `json:"cluster_public_key"`
	SkipFinalSnapshot                bool                     `json:"skip_final_snapshot"`
	DnsName                          string                   `json:"dns_name"`
	SnapshotIdentifier               string                   `json:"snapshot_identifier"`
	SnapshotClusterIdentifier        string                   `json:"snapshot_cluster_identifier"`
	ClusterIdentifier                string                   `json:"cluster_identifier"`
	NodeType                         string                   `json:"node_type"`
	MasterPassword                   string                   `json:"master_password"`
	KmsKeyId                         string                   `json:"kms_key_id"`
	ClusterSecurityGroups            []string                 `json:"cluster_security_groups"`
	ClusterSubnetGroupName           string                   `json:"cluster_subnet_group_name"`
	AvailabilityZone                 string                   `json:"availability_zone"`
	AutomatedSnapshotRetentionPeriod int                      `json:"automated_snapshot_retention_period"`
	ClusterParameterGroupName        string                   `json:"cluster_parameter_group_name"`
	PubliclyAccessible               bool                     `json:"publicly_accessible"`
	IamRoles                         []string                 `json:"iam_roles"`
	BucketName                       string                   `json:"bucket_name"`
}

type AwsRedshiftClusterStatus struct {
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
