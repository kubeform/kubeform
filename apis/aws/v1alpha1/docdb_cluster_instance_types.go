package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DocdbClusterInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DocdbClusterInstanceSpec   `json:"spec,omitempty"`
	Status            DocdbClusterInstanceStatus `json:"status,omitempty"`
}

type DocdbClusterInstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	ApplyImmediately bool   `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`
	Arn              string `json:"arn" tf:"arn"`
	// +optional
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`
	// +optional
	AvailabilityZone  string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	ClusterIdentifier string `json:"clusterIdentifier" tf:"cluster_identifier"`
	DbSubnetGroupName string `json:"dbSubnetGroupName" tf:"db_subnet_group_name"`
	DbiResourceID     string `json:"dbiResourceID" tf:"dbi_resource_id"`
	Endpoint          string `json:"endpoint" tf:"endpoint"`
	// +optional
	Engine        string `json:"engine,omitempty" tf:"engine,omitempty"`
	EngineVersion string `json:"engineVersion" tf:"engine_version"`
	// +optional
	Identifier string `json:"identifier,omitempty" tf:"identifier,omitempty"`
	// +optional
	IdentifierPrefix      string `json:"identifierPrefix,omitempty" tf:"identifier_prefix,omitempty"`
	InstanceClass         string `json:"instanceClass" tf:"instance_class"`
	KmsKeyID              string `json:"kmsKeyID" tf:"kms_key_id"`
	Port                  int    `json:"port" tf:"port"`
	PreferredBackupWindow string `json:"preferredBackupWindow" tf:"preferred_backup_window"`
	// +optional
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`
	// +optional
	PromotionTier      int  `json:"promotionTier,omitempty" tf:"promotion_tier,omitempty"`
	PubliclyAccessible bool `json:"publiclyAccessible" tf:"publicly_accessible"`
	StorageEncrypted   bool `json:"storageEncrypted" tf:"storage_encrypted"`
	// +optional
	Tags   map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Writer bool              `json:"writer" tf:"writer"`
}

type DocdbClusterInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DocdbClusterInstanceSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DocdbClusterInstanceList is a list of DocdbClusterInstances
type DocdbClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DocdbClusterInstance CRD objects
	Items []DocdbClusterInstance `json:"items,omitempty"`
}
