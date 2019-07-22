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

type DxHostedPrivateVirtualInterfaceAccepter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxHostedPrivateVirtualInterfaceAccepterSpec   `json:"spec,omitempty"`
	Status            DxHostedPrivateVirtualInterfaceAccepterStatus `json:"status,omitempty"`
}

type DxHostedPrivateVirtualInterfaceAccepterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	DxGatewayID string `json:"dxGatewayID,omitempty" tf:"dx_gateway_id,omitempty"`
	// +optional
	Tags               map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	VirtualInterfaceID string            `json:"virtualInterfaceID" tf:"virtual_interface_id"`
	// +optional
	VpnGatewayID string `json:"vpnGatewayID,omitempty" tf:"vpn_gateway_id,omitempty"`
}

type DxHostedPrivateVirtualInterfaceAccepterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxHostedPrivateVirtualInterfaceAccepterList is a list of DxHostedPrivateVirtualInterfaceAccepters
type DxHostedPrivateVirtualInterfaceAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxHostedPrivateVirtualInterfaceAccepter CRD objects
	Items []DxHostedPrivateVirtualInterfaceAccepter `json:"items,omitempty"`
}
