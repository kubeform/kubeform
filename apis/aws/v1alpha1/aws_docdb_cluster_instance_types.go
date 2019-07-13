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
	Arn                        string            `json:"arn"`
	DbiResourceId              string            `json:"dbi_resource_id"`
	Endpoint                   string            `json:"endpoint"`
	InstanceClass              string            `json:"instance_class"`
	PubliclyAccessible         bool              `json:"publicly_accessible"`
	IdentifierPrefix           string            `json:"identifier_prefix"`
	Tags                       map[string]string `json:"tags"`
	Writer                     bool              `json:"writer"`
	AutoMinorVersionUpgrade    bool              `json:"auto_minor_version_upgrade"`
	ClusterIdentifier          string            `json:"cluster_identifier"`
	Engine                     string            `json:"engine"`
	EngineVersion              string            `json:"engine_version"`
	Identifier                 string            `json:"identifier"`
	Port                       int               `json:"port"`
	PreferredBackupWindow      string            `json:"preferred_backup_window"`
	PreferredMaintenanceWindow string            `json:"preferred_maintenance_window"`
	StorageEncrypted           bool              `json:"storage_encrypted"`
	ApplyImmediately           bool              `json:"apply_immediately"`
	AvailabilityZone           string            `json:"availability_zone"`
	DbSubnetGroupName          string            `json:"db_subnet_group_name"`
	KmsKeyId                   string            `json:"kms_key_id"`
	PromotionTier              int               `json:"promotion_tier"`
}



type AwsDocdbClusterInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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