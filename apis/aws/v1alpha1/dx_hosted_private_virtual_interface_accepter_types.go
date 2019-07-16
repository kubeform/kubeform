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

type DxHostedPrivateVirtualInterfaceAccepter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxHostedPrivateVirtualInterfaceAccepterSpec   `json:"spec,omitempty"`
	Status            DxHostedPrivateVirtualInterfaceAccepterStatus `json:"status,omitempty"`
}

type DxHostedPrivateVirtualInterfaceAccepterSpec struct {
	// +optional
	DxGatewayId string `json:"dx_gateway_id,omitempty"`
	// +optional
	Tags               map[string]string `json:"tags,omitempty"`
	VirtualInterfaceId string            `json:"virtual_interface_id"`
	// +optional
	VpnGatewayId string `json:"vpn_gateway_id,omitempty"`
}

type DxHostedPrivateVirtualInterfaceAccepterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
