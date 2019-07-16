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

type ComputeGlobalAddress struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeGlobalAddressSpec   `json:"spec,omitempty"`
	Status            ComputeGlobalAddressStatus `json:"status,omitempty"`
}

type ComputeGlobalAddressSpec struct {
	// +optional
	AddressType string `json:"address_type,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	IpVersion string `json:"ip_version,omitempty"`
	// +optional
	// Deprecated
	Labels map[string]string `json:"labels,omitempty"`
	Name   string            `json:"name"`
	// +optional
	// Deprecated
	Network string `json:"network,omitempty"`
	// +optional
	// Deprecated
	PrefixLength int `json:"prefix_length,omitempty"`
	// +optional
	// Deprecated
	Purpose string `json:"purpose,omitempty"`
}

type ComputeGlobalAddressStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeGlobalAddressList is a list of ComputeGlobalAddresss
type ComputeGlobalAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeGlobalAddress CRD objects
	Items []ComputeGlobalAddress `json:"items,omitempty"`
}
