package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type VpcEndpointRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcEndpointRouteTableAssociationSpec   `json:"spec,omitempty"`
	Status            VpcEndpointRouteTableAssociationStatus `json:"status,omitempty"`
}

type VpcEndpointRouteTableAssociationSpec struct {
	RouteTableID  string                    `json:"routeTableID" tf:"route_table_id"`
	VpcEndpointID string                    `json:"vpcEndpointID" tf:"vpc_endpoint_id"`
	ProviderRef   core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type VpcEndpointRouteTableAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcEndpointRouteTableAssociationList is a list of VpcEndpointRouteTableAssociations
type VpcEndpointRouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcEndpointRouteTableAssociation CRD objects
	Items []VpcEndpointRouteTableAssociation `json:"items,omitempty"`
}
