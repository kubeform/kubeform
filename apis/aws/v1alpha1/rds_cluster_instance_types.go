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

type RdsClusterInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RdsClusterInstanceSpec   `json:"spec,omitempty"`
	Status            RdsClusterInstanceStatus `json:"status,omitempty"`
}

type RdsClusterInstanceSpec struct {
	// +optional
	AutoMinorVersionUpgrade bool   `json:"auto_minor_version_upgrade,omitempty"`
	ClusterIdentifier       string `json:"cluster_identifier"`
	// +optional
	CopyTagsToSnapshot bool `json:"copy_tags_to_snapshot,omitempty"`
	// +optional
	Engine        string `json:"engine,omitempty"`
	InstanceClass string `json:"instance_class"`
	// +optional
	MonitoringInterval int `json:"monitoring_interval,omitempty"`
	// +optional
	PromotionTier int `json:"promotion_tier,omitempty"`
	// +optional
	PubliclyAccessible bool `json:"publicly_accessible,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type RdsClusterInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RdsClusterInstanceList is a list of RdsClusterInstances
type RdsClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RdsClusterInstance CRD objects
	Items []RdsClusterInstance `json:"items,omitempty"`
}
