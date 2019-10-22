package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type DxGatewayAssociationProposal struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxGatewayAssociationProposalSpec   `json:"spec,omitempty"`
	Status            DxGatewayAssociationProposalStatus `json:"status,omitempty"`
}

type DxGatewayAssociationProposalSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllowedPrefixes         []string `json:"allowedPrefixes,omitempty" tf:"allowed_prefixes,omitempty"`
	DxGatewayID             string   `json:"dxGatewayID" tf:"dx_gateway_id"`
	DxGatewayOwnerAccountID string   `json:"dxGatewayOwnerAccountID" tf:"dx_gateway_owner_account_id"`
	VpnGatewayID            string   `json:"vpnGatewayID" tf:"vpn_gateway_id"`
}

type DxGatewayAssociationProposalStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DxGatewayAssociationProposalSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
