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

type AwsDxGatewayAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxGatewayAssociationSpec   `json:"spec,omitempty"`
	Status            AwsDxGatewayAssociationStatus `json:"status,omitempty"`
}

type AwsDxGatewayAssociationSpec struct {
	AllowedPrefixes                 []string `json:"allowed_prefixes"`
	AssociatedGatewayId             string   `json:"associated_gateway_id"`
	AssociatedGatewayOwnerAccountId string   `json:"associated_gateway_owner_account_id"`
	ProposalId                      string   `json:"proposal_id"`
	VpnGatewayId                    string   `json:"vpn_gateway_id"`
	AssociatedGatewayType           string   `json:"associated_gateway_type"`
	DxGatewayAssociationId          string   `json:"dx_gateway_association_id"`
	DxGatewayId                     string   `json:"dx_gateway_id"`
	DxGatewayOwnerAccountId         string   `json:"dx_gateway_owner_account_id"`
}

type AwsDxGatewayAssociationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDxGatewayAssociationList is a list of AwsDxGatewayAssociations
type AwsDxGatewayAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxGatewayAssociation CRD objects
	Items []AwsDxGatewayAssociation `json:"items,omitempty"`
}
