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

type Route53Zone struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ZoneSpec   `json:"spec,omitempty"`
	Status            Route53ZoneStatus `json:"status,omitempty"`
}

type Route53ZoneSpecVpc struct {
	VpcId string `json:"vpc_id"`
}

type Route53ZoneSpec struct {
	// +optional
	Comment string `json:"comment,omitempty"`
	// +optional
	DelegationSetId string `json:"delegation_set_id,omitempty"`
	// +optional
	ForceDestroy bool   `json:"force_destroy,omitempty"`
	Name         string `json:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Vpc *[]Route53ZoneSpec `json:"vpc,omitempty"`
}

type Route53ZoneStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Route53ZoneList is a list of Route53Zones
type Route53ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Route53Zone CRD objects
	Items []Route53Zone `json:"items,omitempty"`
}
