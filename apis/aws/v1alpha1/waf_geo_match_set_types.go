package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

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

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	GeoMatchConstraint []WafGeoMatchSetSpecGeoMatchConstraint `json:"geoMatchConstraint,omitempty" tf:"geo_match_constraint,omitempty"`
	Name               string                                 `json:"name" tf:"name"`
}

type WafGeoMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *WafGeoMatchSetSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
