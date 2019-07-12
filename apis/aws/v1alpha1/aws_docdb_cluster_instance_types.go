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

type AwsDocdbClusterInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDocdbClusterInstanceSpec   `json:"spec,omitempty"`
	Status            AwsDocdbClusterInstanceStatus `json:"status,omitempty"`
}

type AwsDocdbClusterInstanceSpec struct {
	PreferredBackupWindow      string            `json:"preferred_backup_window"`
	PubliclyAccessible         bool              `json:"publicly_accessible"`
	StorageEncrypted           bool              `json:"storage_encrypted"`
	Arn                        string            `json:"arn"`
	Endpoint                   string            `json:"endpoint"`
	EngineVersion              string            `json:"engine_version"`
	KmsKeyId                   string            `json:"kms_key_id"`
	Identifier                 string            `json:"identifier"`
	IdentifierPrefix           string            `json:"identifier_prefix"`
	InstanceClass              string            `json:"instance_class"`
	Port                       int               `json:"port"`
	ApplyImmediately           bool              `json:"apply_immediately"`
	DbiResourceId              string            `json:"dbi_resource_id"`
	ClusterIdentifier          string            `json:"cluster_identifier"`
	DbSubnetGroupName          string            `json:"db_subnet_group_name"`
	Engine                     string            `json:"engine"`
	PreferredMaintenanceWindow string            `json:"preferred_maintenance_window"`
	PromotionTier              int               `json:"promotion_tier"`
	Tags                       map[string]string `json:"tags"`
	AutoMinorVersionUpgrade    bool              `json:"auto_minor_version_upgrade"`
	AvailabilityZone           string            `json:"availability_zone"`
	Writer                     bool              `json:"writer"`
}

type AwsDocdbClusterInstanceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDocdbClusterInstanceList is a list of AwsDocdbClusterInstances
type AwsDocdbClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDocdbClusterInstance CRD objects
	Items []AwsDocdbClusterInstance `json:"items,omitempty"`
}
