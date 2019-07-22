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

type DxGatewayAssociationProposal struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxGatewayAssociationProposalSpec   `json:"spec,omitempty"`
	Status            DxGatewayAssociationProposalStatus `json:"status,omitempty"`
}

type DxGatewayAssociationProposalSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AllowedPrefixes []string `json:"allowedPrefixes,omitempty" tf:"allowed_prefixes,omitempty"`
	// +optional
	AssociatedGatewayID     string `json:"associatedGatewayID,omitempty" tf:"associated_gateway_id,omitempty"`
	DxGatewayID             string `json:"dxGatewayID" tf:"dx_gateway_id"`
	DxGatewayOwnerAccountID string `json:"dxGatewayOwnerAccountID" tf:"dx_gateway_owner_account_id"`
	// +optional
	// Deprecated
	VpnGatewayID string `json:"vpnGatewayID,omitempty" tf:"vpn_gateway_id,omitempty"`
}

type DxGatewayAssociationProposalStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxGatewayAssociationProposalList is a list of DxGatewayAssociationProposals
type DxGatewayAssociationProposalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxGatewayAssociationProposal CRD objects
	Items []DxGatewayAssociationProposal `json:"items,omitempty"`
}
