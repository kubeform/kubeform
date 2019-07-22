package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	Type  string `json:"type" tf:"type"`
	Value string `json:"value" tf:"value"`
}

type WafGeoMatchSetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	GeoMatchConstraint []WafGeoMatchSetSpecGeoMatchConstraint `json:"geoMatchConstraint,omitempty" tf:"geo_match_constraint,omitempty"`
	Name               string                                 `json:"name" tf:"name"`
}

type WafGeoMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
