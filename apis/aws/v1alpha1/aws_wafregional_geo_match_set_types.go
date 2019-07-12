package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalGeoMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafregionalGeoMatchSetSpec   `json:"spec,omitempty"`
	Status            AwsWafregionalGeoMatchSetStatus `json:"status,omitempty"`
}

type AwsWafregionalGeoMatchSetSpecGeoMatchConstraint struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AwsWafregionalGeoMatchSetSpec struct {
	Name               string                          `json:"name"`
	GeoMatchConstraint []AwsWafregionalGeoMatchSetSpec `json:"geo_match_constraint"`
}



type AwsWafregionalGeoMatchSetStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalGeoMatchSetList is a list of AwsWafregionalGeoMatchSets
type AwsWafregionalGeoMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafregionalGeoMatchSet CRD objects
	Items []AwsWafregionalGeoMatchSet `json:"items,omitempty"`
}