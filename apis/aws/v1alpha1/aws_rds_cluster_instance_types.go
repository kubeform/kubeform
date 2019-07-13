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
	DbSubnetGroupName           string            `json:"db_subnet_group_name"`
	ApplyImmediately            bool              `json:"apply_immediately"`
	AvailabilityZone            string            `json:"availability_zone"`
	CopyTagsToSnapshot          bool              `json:"copy_tags_to_snapshot"`
	Port                        int               `json:"port"`
	MonitoringRoleArn           string            `json:"monitoring_role_arn"`
	PerformanceInsightsKmsKeyId string            `json:"performance_insights_kms_key_id"`
	DbiResourceId               string            `json:"dbi_resource_id"`
	Arn                         string            `json:"arn"`
	IdentifierPrefix            string            `json:"identifier_prefix"`
	Writer                      bool              `json:"writer"`
	KmsKeyId                    string            `json:"kms_key_id"`
	InstanceClass               string            `json:"instance_class"`
	EngineVersion               string            `json:"engine_version"`
	AutoMinorVersionUpgrade     bool              `json:"auto_minor_version_upgrade"`
	MonitoringInterval          int               `json:"monitoring_interval"`
	Identifier                  string            `json:"identifier"`
	PerformanceInsightsEnabled  bool              `json:"performance_insights_enabled"`
	Endpoint                    string            `json:"endpoint"`
	Engine                      string            `json:"engine"`
	PreferredBackupWindow       string            `json:"preferred_backup_window"`
	PromotionTier               int               `json:"promotion_tier"`
	Tags                        map[string]string `json:"tags"`
	PreferredMaintenanceWindow  string            `json:"preferred_maintenance_window"`
	ClusterIdentifier           string            `json:"cluster_identifier"`
	PubliclyAccessible          bool              `json:"publicly_accessible"`
	DbParameterGroupName        string            `json:"db_parameter_group_name"`
	StorageEncrypted            bool              `json:"storage_encrypted"`
}



type AwsRdsClusterInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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