package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRdsClusterInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRdsClusterInstanceSpec   `json:"spec,omitempty"`
	Status            AwsRdsClusterInstanceStatus `json:"status,omitempty"`
}

type AwsRdsClusterInstanceSpec struct {
	DbSubnetGroupName           string            `json:"db_subnet_group_name"`
	PromotionTier               int               `json:"promotion_tier"`
	PerformanceInsightsEnabled  bool              `json:"performance_insights_enabled"`
	CopyTagsToSnapshot          bool              `json:"copy_tags_to_snapshot"`
	AvailabilityZone            string            `json:"availability_zone"`
	PerformanceInsightsKmsKeyId string            `json:"performance_insights_kms_key_id"`
	Tags                        map[string]string `json:"tags"`
	Arn                         string            `json:"arn"`
	Identifier                  string            `json:"identifier"`
	Endpoint                    string            `json:"endpoint"`
	AutoMinorVersionUpgrade     bool              `json:"auto_minor_version_upgrade"`
	MonitoringInterval          int               `json:"monitoring_interval"`
	IdentifierPrefix            string            `json:"identifier_prefix"`
	InstanceClass               string            `json:"instance_class"`
	ApplyImmediately            bool              `json:"apply_immediately"`
	PreferredMaintenanceWindow  string            `json:"preferred_maintenance_window"`
	ClusterIdentifier           string            `json:"cluster_identifier"`
	Port                        int               `json:"port"`
	PubliclyAccessible          bool              `json:"publicly_accessible"`
	DbiResourceId               string            `json:"dbi_resource_id"`
	MonitoringRoleArn           string            `json:"monitoring_role_arn"`
	DbParameterGroupName        string            `json:"db_parameter_group_name"`
	PreferredBackupWindow       string            `json:"preferred_backup_window"`
	Writer                      bool              `json:"writer"`
	Engine                      string            `json:"engine"`
	EngineVersion               string            `json:"engine_version"`
	KmsKeyId                    string            `json:"kms_key_id"`
	StorageEncrypted            bool              `json:"storage_encrypted"`
}

type AwsRdsClusterInstanceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRdsClusterInstanceList is a list of AwsRdsClusterInstances
type AwsRdsClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRdsClusterInstance CRD objects
	Items []AwsRdsClusterInstance `json:"items,omitempty"`
}
