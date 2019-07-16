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

type WafregionalGeoMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalGeoMatchSetSpec   `json:"spec,omitempty"`
	Status            WafregionalGeoMatchSetStatus `json:"status,omitempty"`
}

type WafregionalGeoMatchSetSpecGeoMatchConstraint struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type WafregionalGeoMatchSetSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	GeoMatchConstraint *[]WafregionalGeoMatchSetSpec `json:"geo_match_constraint,omitempty"`
	Name               string                        `json:"name"`
}

type WafregionalGeoMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafregionalGeoMatchSetList is a list of WafregionalGeoMatchSets
type WafregionalGeoMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafregionalGeoMatchSet CRD objects
	Items []WafregionalGeoMatchSet `json:"items,omitempty"`
}
