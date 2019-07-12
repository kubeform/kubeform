package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxGatewayAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxGatewayAssociationSpec   `json:"spec,omitempty"`
	Status            AwsDxGatewayAssociationStatus `json:"status,omitempty"`
}

type AwsDxGatewayAssociationSpec struct {
	VpnGatewayId           string   `json:"vpn_gateway_id"`
	DxGatewayAssociationId string   `json:"dx_gateway_association_id"`
	AllowedPrefixes        []string `json:"allowed_prefixes"`
	DxGatewayId            string   `json:"dx_gateway_id"`
}

type AwsDxGatewayAssociationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDxGatewayAssociationList is a list of AwsDxGatewayAssociations
type AwsDxGatewayAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxGatewayAssociation CRD objects
	Items []AwsDxGatewayAssociation `json:"items,omitempty"`
}
