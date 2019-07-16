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

type WafregionalIpset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalIpsetSpec   `json:"spec,omitempty"`
	Status            WafregionalIpsetStatus `json:"status,omitempty"`
}

type WafregionalIpsetSpecIpSetDescriptor struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type WafregionalIpsetSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	IpSetDescriptor *[]WafregionalIpsetSpec `json:"ip_set_descriptor,omitempty"`
	Name            string                  `json:"name"`
}

type WafregionalIpsetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafregionalIpsetList is a list of WafregionalIpsets
type WafregionalIpsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafregionalIpset CRD objects
	Items []WafregionalIpset `json:"items,omitempty"`
}
