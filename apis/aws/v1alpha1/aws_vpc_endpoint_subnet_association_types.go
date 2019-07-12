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

type AwsVpcEndpointSubnetAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpcEndpointSubnetAssociationSpec   `json:"spec,omitempty"`
	Status            AwsVpcEndpointSubnetAssociationStatus `json:"status,omitempty"`
}

type AwsVpcEndpointSubnetAssociationSpec struct {
	VpcEndpointId string `json:"vpc_endpoint_id"`
	SubnetId      string `json:"subnet_id"`
}

type AwsVpcEndpointSubnetAssociationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsVpcEndpointSubnetAssociationList is a list of AwsVpcEndpointSubnetAssociations
type AwsVpcEndpointSubnetAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpcEndpointSubnetAssociation CRD objects
	Items []AwsVpcEndpointSubnetAssociation `json:"items,omitempty"`
}
