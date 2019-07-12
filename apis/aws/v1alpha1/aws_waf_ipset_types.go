package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafIpset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafIpsetSpec   `json:"spec,omitempty"`
	Status            AwsWafIpsetStatus `json:"status,omitempty"`
}

type AwsWafIpsetSpecIpSetDescriptors struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AwsWafIpsetSpec struct {
	Name             string            `json:"name"`
	Arn              string            `json:"arn"`
	IpSetDescriptors []AwsWafIpsetSpec `json:"ip_set_descriptors"`
}



type AwsWafIpsetStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafIpsetList is a list of AwsWafIpsets
type AwsWafIpsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafIpset CRD objects
	Items []AwsWafIpset `json:"items,omitempty"`
}