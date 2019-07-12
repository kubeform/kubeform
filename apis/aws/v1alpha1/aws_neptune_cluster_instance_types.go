package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNeptuneClusterInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNeptuneClusterInstanceSpec   `json:"spec,omitempty"`
	Status            AwsNeptuneClusterInstanceStatus `json:"status,omitempty"`
}

type AwsNeptuneClusterInstanceSpec struct {
	AvailabilityZone           string            `json:"availability_zone"`
	ClusterIdentifier          string            `json:"cluster_identifier"`
	DbiResourceId              string            `json:"dbi_resource_id"`
	Endpoint                   string            `json:"endpoint"`
	Engine                     string            `json:"engine"`
	NeptuneParameterGroupName  string            `json:"neptune_parameter_group_name"`
	Port                       int               `json:"port"`
	Address                    string            `json:"address"`
	Tags                       map[string]string `json:"tags"`
	PreferredMaintenanceWindow string            `json:"preferred_maintenance_window"`
	IdentifierPrefix           string            `json:"identifier_prefix"`
	KmsKeyArn                  string            `json:"kms_key_arn"`
	StorageEncrypted           bool              `json:"storage_encrypted"`
	Arn                        string            `json:"arn"`
	AutoMinorVersionUpgrade    bool              `json:"auto_minor_version_upgrade"`
	EngineVersion              string            `json:"engine_version"`
	InstanceClass              string            `json:"instance_class"`
	PreferredBackupWindow      string            `json:"preferred_backup_window"`
	PromotionTier              int               `json:"promotion_tier"`
	Writer                     bool              `json:"writer"`
	ApplyImmediately           bool              `json:"apply_immediately"`
	NeptuneSubnetGroupName     string            `json:"neptune_subnet_group_name"`
	PubliclyAccessible         bool              `json:"publicly_accessible"`
	Identifier                 string            `json:"identifier"`
}



type AwsNeptuneClusterInstanceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNeptuneClusterInstanceList is a list of AwsNeptuneClusterInstances
type AwsNeptuneClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNeptuneClusterInstance CRD objects
	Items []AwsNeptuneClusterInstance `json:"items,omitempty"`
}