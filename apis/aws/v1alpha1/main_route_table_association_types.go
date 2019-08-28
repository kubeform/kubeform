package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type MainRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MainRouteTableAssociationSpec   `json:"spec,omitempty"`
	Status            MainRouteTableAssociationStatus `json:"status,omitempty"`
}

type MainRouteTableAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	OriginalRouteTableID string `json:"originalRouteTableID,omitempty" tf:"original_route_table_id,omitempty"`
	RouteTableID         string `json:"routeTableID" tf:"route_table_id"`
	VpcID                string `json:"vpcID" tf:"vpc_id"`
}



type MainRouteTableAssociationStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *MainRouteTableAssociationSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MainRouteTableAssociationList is a list of MainRouteTableAssociations
type MainRouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MainRouteTableAssociation CRD objects
	Items []MainRouteTableAssociation `json:"items,omitempty"`
}