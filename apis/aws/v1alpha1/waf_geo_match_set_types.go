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

type WafGeoMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafGeoMatchSetSpec   `json:"spec,omitempty"`
	Status            WafGeoMatchSetStatus `json:"status,omitempty"`
}

type WafGeoMatchSetSpecGeoMatchConstraint struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type WafGeoMatchSetSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	GeoMatchConstraint *[]WafGeoMatchSetSpec `json:"geo_match_constraint,omitempty"`
	Name               string                `json:"name"`
}

type WafGeoMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafGeoMatchSetList is a list of WafGeoMatchSets
type WafGeoMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafGeoMatchSet CRD objects
	Items []WafGeoMatchSet `json:"items,omitempty"`
}
