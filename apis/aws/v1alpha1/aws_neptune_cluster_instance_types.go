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

type AwsNeptuneClusterInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNeptuneClusterInstanceSpec   `json:"spec,omitempty"`
	Status            AwsNeptuneClusterInstanceStatus `json:"status,omitempty"`
}

type AwsNeptuneClusterInstanceSpec struct {
	NeptuneParameterGroupName  string            `json:"neptune_parameter_group_name"`
	NeptuneSubnetGroupName     string            `json:"neptune_subnet_group_name"`
	Port                       int               `json:"port"`
	PreferredMaintenanceWindow string            `json:"preferred_maintenance_window"`
	Tags                       map[string]string `json:"tags"`
	AvailabilityZone           string            `json:"availability_zone"`
	EngineVersion              string            `json:"engine_version"`
	AutoMinorVersionUpgrade    bool              `json:"auto_minor_version_upgrade"`
	Identifier                 string            `json:"identifier"`
	Address                    string            `json:"address"`
	ApplyImmediately           bool              `json:"apply_immediately"`
	Endpoint                   string            `json:"endpoint"`
	KmsKeyArn                  string            `json:"kms_key_arn"`
	PreferredBackupWindow      string            `json:"preferred_backup_window"`
	PromotionTier              int               `json:"promotion_tier"`
	PubliclyAccessible         bool              `json:"publicly_accessible"`
	StorageEncrypted           bool              `json:"storage_encrypted"`
	Arn                        string            `json:"arn"`
	ClusterIdentifier          string            `json:"cluster_identifier"`
	Writer                     bool              `json:"writer"`
	IdentifierPrefix           string            `json:"identifier_prefix"`
	InstanceClass              string            `json:"instance_class"`
	DbiResourceId              string            `json:"dbi_resource_id"`
	Engine                     string            `json:"engine"`
}

type AwsNeptuneClusterInstanceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsNeptuneClusterInstanceList is a list of AwsNeptuneClusterInstances
type AwsNeptuneClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNeptuneClusterInstance CRD objects
	Items []AwsNeptuneClusterInstance `json:"items,omitempty"`
}
