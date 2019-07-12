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

type AwsRdsClusterInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRdsClusterInstanceSpec   `json:"spec,omitempty"`
	Status            AwsRdsClusterInstanceStatus `json:"status,omitempty"`
}

type AwsRdsClusterInstanceSpec struct {
	Endpoint                    string            `json:"endpoint"`
	Engine                      string            `json:"engine"`
	DbParameterGroupName        string            `json:"db_parameter_group_name"`
	StorageEncrypted            bool              `json:"storage_encrypted"`
	DbiResourceId               string            `json:"dbi_resource_id"`
	MonitoringRoleArn           string            `json:"monitoring_role_arn"`
	Identifier                  string            `json:"identifier"`
	AutoMinorVersionUpgrade     bool              `json:"auto_minor_version_upgrade"`
	PerformanceInsightsKmsKeyId string            `json:"performance_insights_kms_key_id"`
	CopyTagsToSnapshot          bool              `json:"copy_tags_to_snapshot"`
	IdentifierPrefix            string            `json:"identifier_prefix"`
	DbSubnetGroupName           string            `json:"db_subnet_group_name"`
	Port                        int               `json:"port"`
	PreferredBackupWindow       string            `json:"preferred_backup_window"`
	MonitoringInterval          int               `json:"monitoring_interval"`
	PubliclyAccessible          bool              `json:"publicly_accessible"`
	InstanceClass               string            `json:"instance_class"`
	AvailabilityZone            string            `json:"availability_zone"`
	PreferredMaintenanceWindow  string            `json:"preferred_maintenance_window"`
	Tags                        map[string]string `json:"tags"`
	Writer                      bool              `json:"writer"`
	ApplyImmediately            bool              `json:"apply_immediately"`
	PromotionTier               int               `json:"promotion_tier"`
	ClusterIdentifier           string            `json:"cluster_identifier"`
	KmsKeyId                    string            `json:"kms_key_id"`
	Arn                         string            `json:"arn"`
	EngineVersion               string            `json:"engine_version"`
	PerformanceInsightsEnabled  bool              `json:"performance_insights_enabled"`
}

type AwsRdsClusterInstanceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRdsClusterInstanceList is a list of AwsRdsClusterInstances
type AwsRdsClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRdsClusterInstance CRD objects
	Items []AwsRdsClusterInstance `json:"items,omitempty"`
}
