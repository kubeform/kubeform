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
	DbiResourceId              string            `json:"dbi_resource_id"`
	Endpoint                   string            `json:"endpoint"`
	EngineVersion              string            `json:"engine_version"`
	KmsKeyArn                  string            `json:"kms_key_arn"`
	Address                    string            `json:"address"`
	ApplyImmediately           bool              `json:"apply_immediately"`
	Arn                        string            `json:"arn"`
	AvailabilityZone           string            `json:"availability_zone"`
	Identifier                 string            `json:"identifier"`
	Port                       int               `json:"port"`
	PreferredMaintenanceWindow string            `json:"preferred_maintenance_window"`
	NeptuneSubnetGroupName     string            `json:"neptune_subnet_group_name"`
	PromotionTier              int               `json:"promotion_tier"`
	ClusterIdentifier          string            `json:"cluster_identifier"`
	Engine                     string            `json:"engine"`
	InstanceClass              string            `json:"instance_class"`
	NeptuneParameterGroupName  string            `json:"neptune_parameter_group_name"`
	StorageEncrypted           bool              `json:"storage_encrypted"`
	Tags                       map[string]string `json:"tags"`
	Writer                     bool              `json:"writer"`
	AutoMinorVersionUpgrade    bool              `json:"auto_minor_version_upgrade"`
	IdentifierPrefix           string            `json:"identifier_prefix"`
	PreferredBackupWindow      string            `json:"preferred_backup_window"`
	PubliclyAccessible         bool              `json:"publicly_accessible"`
}



type AwsNeptuneClusterInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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