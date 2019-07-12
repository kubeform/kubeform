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

type AzurermSubnetRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSubnetRouteTableAssociationSpec   `json:"spec,omitempty"`
	Status            AzurermSubnetRouteTableAssociationStatus `json:"status,omitempty"`
}

type AzurermSubnetRouteTableAssociationSpec struct {
	RouteTableId string `json:"route_table_id"`
	SubnetId     string `json:"subnet_id"`
}

type AzurermSubnetRouteTableAssociationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermSubnetRouteTableAssociationList is a list of AzurermSubnetRouteTableAssociations
type AzurermSubnetRouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSubnetRouteTableAssociation CRD objects
	Items []AzurermSubnetRouteTableAssociation `json:"items,omitempty"`
}
