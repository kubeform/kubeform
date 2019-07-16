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

type DxGatewayAssociationProposal struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxGatewayAssociationProposalSpec   `json:"spec,omitempty"`
	Status            DxGatewayAssociationProposalStatus `json:"status,omitempty"`
}

type DxGatewayAssociationProposalSpec struct {
	// +optional
	AssociatedGatewayId     string `json:"associated_gateway_id,omitempty"`
	DxGatewayId             string `json:"dx_gateway_id"`
	DxGatewayOwnerAccountId string `json:"dx_gateway_owner_account_id"`
	// +optional
	// Deprecated
	VpnGatewayId string `json:"vpn_gateway_id,omitempty"`
}

type DxGatewayAssociationProposalStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
