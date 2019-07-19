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

type DefaultVpc struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultVpcSpec   `json:"spec,omitempty"`
	Status            DefaultVpcStatus `json:"status,omitempty"`
}

type DefaultVpcSpec struct {
	// +optional
	EnableClassiclink bool `json:"enableClassiclink,omitempty" tf:"enable_classiclink,omitempty"`
	// +optional
	EnableClassiclinkDNSSupport bool `json:"enableClassiclinkDNSSupport,omitempty" tf:"enable_classiclink_dns_support,omitempty"`
	// +optional
	EnableDNSHostnames bool `json:"enableDNSHostnames,omitempty" tf:"enable_dns_hostnames,omitempty"`
	// +optional
	EnableDNSSupport bool `json:"enableDNSSupport,omitempty" tf:"enable_dns_support,omitempty"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DefaultVpcStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DefaultVpcList is a list of DefaultVpcs
type DefaultVpcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DefaultVpc CRD objects
	Items []DefaultVpc `json:"items,omitempty"`
}
