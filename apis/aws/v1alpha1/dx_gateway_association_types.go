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

type DxGatewayAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxGatewayAssociationSpec   `json:"spec,omitempty"`
	Status            DxGatewayAssociationStatus `json:"status,omitempty"`
}

type DxGatewayAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AllowedPrefixes []string `json:"allowedPrefixes,omitempty" tf:"allowed_prefixes,omitempty"`
	// +optional
	DxGatewayAssociationID string `json:"dxGatewayAssociationID,omitempty" tf:"dx_gateway_association_id,omitempty"`
	DxGatewayID            string `json:"dxGatewayID" tf:"dx_gateway_id"`
	VpnGatewayID           string `json:"vpnGatewayID" tf:"vpn_gateway_id"`
}

type DxGatewayAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DxGatewayAssociationSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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