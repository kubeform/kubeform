package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxGatewayAssociationProposal struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxGatewayAssociationProposalSpec   `json:"spec,omitempty"`
	Status            AwsDxGatewayAssociationProposalStatus `json:"status,omitempty"`
}

type AwsDxGatewayAssociationProposalSpec struct {
	AllowedPrefixes         []string `json:"allowed_prefixes"`
	DxGatewayId             string   `json:"dx_gateway_id"`
	DxGatewayOwnerAccountId string   `json:"dx_gateway_owner_account_id"`
	VpnGatewayId            string   `json:"vpn_gateway_id"`
}



type AwsDxGatewayAssociationProposalStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDxGatewayAssociationProposalList is a list of AwsDxGatewayAssociationProposals
type AwsDxGatewayAssociationProposalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxGatewayAssociationProposal CRD objects
	Items []AwsDxGatewayAssociationProposal `json:"items,omitempty"`
}