package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafGeoMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafGeoMatchSetSpec   `json:"spec,omitempty"`
	Status            AwsWafGeoMatchSetStatus `json:"status,omitempty"`
}

type AwsWafGeoMatchSetSpecGeoMatchConstraint struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AwsWafGeoMatchSetSpec struct {
	Name               string                  `json:"name"`
	GeoMatchConstraint []AwsWafGeoMatchSetSpec `json:"geo_match_constraint"`
}

type AwsWafGeoMatchSetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafGeoMatchSetList is a list of AwsWafGeoMatchSets
type AwsWafGeoMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafGeoMatchSet CRD objects
	Items []AwsWafGeoMatchSet `json:"items,omitempty"`
}
