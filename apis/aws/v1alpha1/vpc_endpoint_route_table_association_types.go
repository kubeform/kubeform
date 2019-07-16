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

type VpcEndpointRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcEndpointRouteTableAssociationSpec   `json:"spec,omitempty"`
	Status            VpcEndpointRouteTableAssociationStatus `json:"status,omitempty"`
}

type VpcEndpointRouteTableAssociationSpec struct {
	RouteTableId  string `json:"route_table_id"`
	VpcEndpointId string `json:"vpc_endpoint_id"`
}

type VpcEndpointRouteTableAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
