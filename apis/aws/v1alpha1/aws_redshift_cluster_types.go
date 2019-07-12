package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRedshiftClusterSpec   `json:"spec,omitempty"`
	Status            AwsRedshiftClusterStatus `json:"status,omitempty"`
}

type AwsRedshiftClusterSpecSnapshotCopy struct {
	GrantName         string `json:"grant_name"`
	DestinationRegion string `json:"destination_region"`
	RetentionPeriod   int    `json:"retention_period"`
}

type AwsRedshiftClusterSpecLogging struct {
	Enable      bool   `json:"enable"`
	BucketName  string `json:"bucket_name"`
	S3KeyPrefix string `json:"s3_key_prefix"`
}

type AwsRedshiftClusterSpec struct {
	PreferredMaintenanceWindow       string                   `json:"preferred_maintenance_window"`
	ElasticIp                        string                   `json:"elastic_ip"`
	IamRoles                         []string                 `json:"iam_roles"`
	SnapshotCopy                     []AwsRedshiftClusterSpec `json:"snapshot_copy"`
	SnapshotIdentifier               string                   `json:"snapshot_identifier"`
	Tags                             map[string]string        `json:"tags"`
	ClusterIdentifier                string                   `json:"cluster_identifier"`
	MasterPassword                   string                   `json:"master_password"`
	VpcSecurityGroupIds              []string                 `json:"vpc_security_group_ids"`
	AutomatedSnapshotRetentionPeriod int                      `json:"automated_snapshot_retention_period"`
	ClusterRevisionNumber            string                   `json:"cluster_revision_number"`
	Port                             int                      `json:"port"`
	Endpoint                         string                   `json:"endpoint"`
	ClusterType                      string                   `json:"cluster_type"`
	NumberOfNodes                    int                      `json:"number_of_nodes"`
	EnhancedVpcRouting               bool                     `json:"enhanced_vpc_routing"`
	KmsKeyId                         string                   `json:"kms_key_id"`
	SkipFinalSnapshot                bool                     `json:"skip_final_snapshot"`
	DatabaseName                     string                   `json:"database_name"`
	NodeType                         string                   `json:"node_type"`
	AllowVersionUpgrade              bool                     `json:"allow_version_upgrade"`
	FinalSnapshotIdentifier          string                   `json:"final_snapshot_identifier"`
	EnableLogging                    bool                     `json:"enable_logging"`
	OwnerAccount                     string                   `json:"owner_account"`
	MasterUsername                   string                   `json:"master_username"`
	ClusterSubnetGroupName           string                   `json:"cluster_subnet_group_name"`
	ClusterVersion                   string                   `json:"cluster_version"`
	Logging                          []AwsRedshiftClusterSpec `json:"logging"`
	S3KeyPrefix                      string                   `json:"s3_key_prefix"`
	ClusterSecurityGroups            []string                 `json:"cluster_security_groups"`
	AvailabilityZone                 string                   `json:"availability_zone"`
	ClusterParameterGroupName        string                   `json:"cluster_parameter_group_name"`
	BucketName                       string                   `json:"bucket_name"`
	SnapshotClusterIdentifier        string                   `json:"snapshot_cluster_identifier"`
	PubliclyAccessible               bool                     `json:"publicly_accessible"`
	Encrypted                        bool                     `json:"encrypted"`
	DnsName                          string                   `json:"dns_name"`
	ClusterPublicKey                 string                   `json:"cluster_public_key"`
}

type AwsRedshiftClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftClusterList is a list of AwsRedshiftClusters
type AwsRedshiftClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRedshiftCluster CRD objects
	Items []AwsRedshiftCluster `json:"items,omitempty"`
}
