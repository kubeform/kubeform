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

type DxGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxGatewaySpec   `json:"spec,omitempty"`
	Status            DxGatewayStatus `json:"status,omitempty"`
}

type DxGatewaySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AmazonSideAsn string `json:"amazonSideAsn" tf:"amazon_side_asn"`
	Name          string `json:"name" tf:"name"`
	// +optional
	OwnerAccountID string `json:"ownerAccountID,omitempty" tf:"owner_account_id,omitempty"`
}

type DxGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DxGatewaySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxGatewayList is a list of DxGateways
type DxGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxGateway CRD objects
	Items []DxGateway `json:"items,omitempty"`
}
