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

type ComputeRouterInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeRouterInterfaceSpec   `json:"spec,omitempty"`
	Status            ComputeRouterInterfaceStatus `json:"status,omitempty"`
}

type ComputeRouterInterfaceSpec struct {
	// +optional
	IpRange string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`
	Name    string `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region      string                    `json:"region,omitempty" tf:"region,omitempty"`
	Router      string                    `json:"router" tf:"router"`
	VpnTunnel   string                    `json:"vpnTunnel" tf:"vpn_tunnel"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeRouterInterfaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeRouterInterfaceList is a list of ComputeRouterInterfaces
type ComputeRouterInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeRouterInterface CRD objects
	Items []ComputeRouterInterface `json:"items,omitempty"`
}
