package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeNetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeNetworkSpec   `json:"spec,omitempty"`
	Status            ComputeNetworkStatus `json:"status,omitempty"`
}

type ComputeNetworkSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutoCreateSubnetworks bool `json:"autoCreateSubnetworks,omitempty" tf:"auto_create_subnetworks,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	GatewayIpv4 string `json:"gatewayIpv4,omitempty" tf:"gateway_ipv4,omitempty"`
	// +optional
	// Deprecated
	Ipv4Range string `json:"ipv4Range,omitempty" tf:"ipv4_range,omitempty"`
	Name      string `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	RoutingMode string `json:"routingMode,omitempty" tf:"routing_mode,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type ComputeNetworkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeNetworkSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeNetworkList is a list of ComputeNetworks
type ComputeNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeNetwork CRD objects
	Items []ComputeNetwork `json:"items,omitempty"`
}
