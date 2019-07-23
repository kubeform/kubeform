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

type WafIpset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafIpsetSpec   `json:"spec,omitempty"`
	Status            WafIpsetStatus `json:"status,omitempty"`
}

type WafIpsetSpecIpSetDescriptors struct {
	Type  string `json:"type" tf:"type"`
	Value string `json:"value" tf:"value"`
}

type WafIpsetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true
	IpSetDescriptors []WafIpsetSpecIpSetDescriptors `json:"ipSetDescriptors,omitempty" tf:"ip_set_descriptors,omitempty"`
	Name             string                         `json:"name" tf:"name"`
}

type WafIpsetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafIpsetList is a list of WafIpsets
type WafIpsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafIpset CRD objects
	Items []WafIpset `json:"items,omitempty"`
}
