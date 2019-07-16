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

type DocdbClusterInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DocdbClusterInstanceSpec   `json:"spec,omitempty"`
	Status            DocdbClusterInstanceStatus `json:"status,omitempty"`
}

type DocdbClusterInstanceSpec struct {
	// +optional
	AutoMinorVersionUpgrade bool   `json:"auto_minor_version_upgrade,omitempty"`
	ClusterIdentifier       string `json:"cluster_identifier"`
	// +optional
	Engine        string `json:"engine,omitempty"`
	InstanceClass string `json:"instance_class"`
	// +optional
	PromotionTier int `json:"promotion_tier,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DocdbClusterInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
