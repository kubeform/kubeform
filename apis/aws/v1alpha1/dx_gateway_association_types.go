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

type DxGatewayAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxGatewayAssociationSpec   `json:"spec,omitempty"`
	Status            DxGatewayAssociationStatus `json:"status,omitempty"`
}

type DxGatewayAssociationSpec struct {
	DxGatewayId string `json:"dx_gateway_id"`
	// +optional
	ProposalId string `json:"proposal_id,omitempty"`
	// +optional
	// Deprecated
	VpnGatewayId string `json:"vpn_gateway_id,omitempty"`
}

type DxGatewayAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxGatewayAssociationList is a list of DxGatewayAssociations
type DxGatewayAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxGatewayAssociation CRD objects
	Items []DxGatewayAssociation `json:"items,omitempty"`
}
