package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsMainRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsMainRouteTableAssociationSpec   `json:"spec,omitempty"`
	Status            AwsMainRouteTableAssociationStatus `json:"status,omitempty"`
}

type AwsMainRouteTableAssociationSpec struct {
	VpcId                string `json:"vpc_id"`
	RouteTableId         string `json:"route_table_id"`
	OriginalRouteTableId string `json:"original_route_table_id"`
}



type AwsMainRouteTableAssociationStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsMainRouteTableAssociationList is a list of AwsMainRouteTableAssociations
type AwsMainRouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsMainRouteTableAssociation CRD objects
	Items []AwsMainRouteTableAssociation `json:"items,omitempty"`
}