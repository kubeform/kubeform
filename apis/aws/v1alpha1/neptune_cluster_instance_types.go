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

type NeptuneClusterInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NeptuneClusterInstanceSpec   `json:"spec,omitempty"`
	Status            NeptuneClusterInstanceStatus `json:"status,omitempty"`
}

type NeptuneClusterInstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	ApplyImmediately bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`
	// +optional
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`
	// +optional
	AvailabilityZone  string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	ClusterIdentifier string `json:"clusterIdentifier" tf:"cluster_identifier"`
	// +optional
	Engine string `json:"engine,omitempty" tf:"engine,omitempty"`
	// +optional
	EngineVersion string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`
	// +optional
	Identifier string `json:"identifier,omitempty" tf:"identifier,omitempty"`
	// +optional
	IdentifierPrefix string `json:"identifierPrefix,omitempty" tf:"identifier_prefix,omitempty"`
	InstanceClass    string `json:"instanceClass" tf:"instance_class"`
	// +optional
	NeptuneParameterGroupName string `json:"neptuneParameterGroupName,omitempty" tf:"neptune_parameter_group_name,omitempty"`
	// +optional
	NeptuneSubnetGroupName string `json:"neptuneSubnetGroupName,omitempty" tf:"neptune_subnet_group_name,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	PreferredBackupWindow string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`
	// +optional
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`
	// +optional
	PromotionTier int `json:"promotionTier,omitempty" tf:"promotion_tier,omitempty"`
	// +optional
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NeptuneClusterInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NeptuneClusterInstanceSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NeptuneClusterInstanceList is a list of NeptuneClusterInstances
type NeptuneClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NeptuneClusterInstance CRD objects
	Items []NeptuneClusterInstance `json:"items,omitempty"`
}
