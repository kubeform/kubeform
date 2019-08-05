package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type VpnGatewayAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpnGatewayAttachmentSpec   `json:"spec,omitempty"`
	Status            VpnGatewayAttachmentStatus `json:"status,omitempty"`
}

type VpnGatewayAttachmentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	VpcID        string `json:"vpcID" tf:"vpc_id"`
	VpnGatewayID string `json:"vpnGatewayID" tf:"vpn_gateway_id"`
}

type VpnGatewayAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VpnGatewayAttachmentSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpnGatewayAttachmentList is a list of VpnGatewayAttachments
type VpnGatewayAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpnGatewayAttachment CRD objects
	Items []VpnGatewayAttachment `json:"items,omitempty"`
}
