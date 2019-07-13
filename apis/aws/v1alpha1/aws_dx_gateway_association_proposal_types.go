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

type AwsDxGatewayAssociationProposal struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxGatewayAssociationProposalSpec   `json:"spec,omitempty"`
	Status            AwsDxGatewayAssociationProposalStatus `json:"status,omitempty"`
}

type AwsDxGatewayAssociationProposalSpec struct {
	AllowedPrefixes                 []string `json:"allowed_prefixes"`
	AssociatedGatewayId             string   `json:"associated_gateway_id"`
	AssociatedGatewayOwnerAccountId string   `json:"associated_gateway_owner_account_id"`
	AssociatedGatewayType           string   `json:"associated_gateway_type"`
	DxGatewayId                     string   `json:"dx_gateway_id"`
	DxGatewayOwnerAccountId         string   `json:"dx_gateway_owner_account_id"`
	VpnGatewayId                    string   `json:"vpn_gateway_id"`
}



type AwsDxGatewayAssociationProposalStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDxGatewayAssociationProposalList is a list of AwsDxGatewayAssociationProposals
type AwsDxGatewayAssociationProposalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxGatewayAssociationProposal CRD objects
	Items []AwsDxGatewayAssociationProposal `json:"items,omitempty"`
}