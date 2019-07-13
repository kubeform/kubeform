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

type AwsWafregionalIpset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafregionalIpsetSpec   `json:"spec,omitempty"`
	Status            AwsWafregionalIpsetStatus `json:"status,omitempty"`
}

type AwsWafregionalIpsetSpecIpSetDescriptor struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type AwsWafregionalIpsetSpec struct {
	Name            string                    `json:"name"`
	Arn             string                    `json:"arn"`
	IpSetDescriptor []AwsWafregionalIpsetSpec `json:"ip_set_descriptor"`
}



type AwsWafregionalIpsetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsWafregionalIpsetList is a list of AwsWafregionalIpsets
type AwsWafregionalIpsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafregionalIpset CRD objects
	Items []AwsWafregionalIpset `json:"items,omitempty"`
}