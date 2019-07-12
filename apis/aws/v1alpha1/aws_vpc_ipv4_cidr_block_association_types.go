package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcIpv4CidrBlockAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpcIpv4CidrBlockAssociationSpec   `json:"spec,omitempty"`
	Status            AwsVpcIpv4CidrBlockAssociationStatus `json:"status,omitempty"`
}

type AwsVpcIpv4CidrBlockAssociationSpec struct {
	VpcId     string `json:"vpc_id"`
	CidrBlock string `json:"cidr_block"`
}



type AwsVpcIpv4CidrBlockAssociationStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcIpv4CidrBlockAssociationList is a list of AwsVpcIpv4CidrBlockAssociations
type AwsVpcIpv4CidrBlockAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpcIpv4CidrBlockAssociation CRD objects
	Items []AwsVpcIpv4CidrBlockAssociation `json:"items,omitempty"`
}