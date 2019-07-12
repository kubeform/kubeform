package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDocdbClusterInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDocdbClusterInstanceSpec   `json:"spec,omitempty"`
	Status            AwsDocdbClusterInstanceStatus `json:"status,omitempty"`
}

type AwsDocdbClusterInstanceSpec struct {
	AutoMinorVersionUpgrade    bool              `json:"auto_minor_version_upgrade"`
	AvailabilityZone           string            `json:"availability_zone"`
	Endpoint                   string            `json:"endpoint"`
	KmsKeyId                   string            `json:"kms_key_id"`
	StorageEncrypted           bool              `json:"storage_encrypted"`
	Arn                        string            `json:"arn"`
	DbiResourceId              string            `json:"dbi_resource_id"`
	EngineVersion              string            `json:"engine_version"`
	IdentifierPrefix           string            `json:"identifier_prefix"`
	Tags                       map[string]string `json:"tags"`
	ClusterIdentifier          string            `json:"cluster_identifier"`
	DbSubnetGroupName          string            `json:"db_subnet_group_name"`
	Engine                     string            `json:"engine"`
	Identifier                 string            `json:"identifier"`
	InstanceClass              string            `json:"instance_class"`
	Port                       int               `json:"port"`
	Writer                     bool              `json:"writer"`
	ApplyImmediately           bool              `json:"apply_immediately"`
	PreferredBackupWindow      string            `json:"preferred_backup_window"`
	PreferredMaintenanceWindow string            `json:"preferred_maintenance_window"`
	PromotionTier              int               `json:"promotion_tier"`
	PubliclyAccessible         bool              `json:"publicly_accessible"`
}



type AwsDocdbClusterInstanceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDocdbClusterInstanceList is a list of AwsDocdbClusterInstances
type AwsDocdbClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDocdbClusterInstance CRD objects
	Items []AwsDocdbClusterInstance `json:"items,omitempty"`
}