package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalIpset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafregionalIpsetSpec   `json:"spec,omitempty"`
	Status            AwsWafregionalIpsetStatus `json:"status,omitempty"`
}

type AwsWafregionalIpsetSpecIpSetDescriptor struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AwsWafregionalIpsetSpec struct {
	Name            string                    `json:"name"`
	Arn             string                    `json:"arn"`
	IpSetDescriptor []AwsWafregionalIpsetSpec `json:"ip_set_descriptor"`
}

type AwsWafregionalIpsetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalIpsetList is a list of AwsWafregionalIpsets
type AwsWafregionalIpsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafregionalIpset CRD objects
	Items []AwsWafregionalIpset `json:"items,omitempty"`
}
