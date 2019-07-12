package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRouteTableAssociationSpec   `json:"spec,omitempty"`
	Status            AwsRouteTableAssociationStatus `json:"status,omitempty"`
}

type AwsRouteTableAssociationSpec struct {
	SubnetId     string `json:"subnet_id"`
	RouteTableId string `json:"route_table_id"`
}



type AwsRouteTableAssociationStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRouteTableAssociationList is a list of AwsRouteTableAssociations
type AwsRouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRouteTableAssociation CRD objects
	Items []AwsRouteTableAssociation `json:"items,omitempty"`
}