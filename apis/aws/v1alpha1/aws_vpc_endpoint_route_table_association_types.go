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

type AwsVpcEndpointRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpcEndpointRouteTableAssociationSpec   `json:"spec,omitempty"`
	Status            AwsVpcEndpointRouteTableAssociationStatus `json:"status,omitempty"`
}

type AwsVpcEndpointRouteTableAssociationSpec struct {
	VpcEndpointId string `json:"vpc_endpoint_id"`
	RouteTableId  string `json:"route_table_id"`
}



type AwsVpcEndpointRouteTableAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsVpcEndpointRouteTableAssociationList is a list of AwsVpcEndpointRouteTableAssociations
type AwsVpcEndpointRouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpcEndpointRouteTableAssociation CRD objects
	Items []AwsVpcEndpointRouteTableAssociation `json:"items,omitempty"`
}