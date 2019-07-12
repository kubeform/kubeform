package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcDhcpOptionsAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsVpcDhcpOptionsAssociationSpec   `json:"spec,omitempty"`
	Status            AwsVpcDhcpOptionsAssociationStatus `json:"status,omitempty"`
}

type AwsVpcDhcpOptionsAssociationSpec struct {
	VpcId         string `json:"vpc_id"`
	DhcpOptionsId string `json:"dhcp_options_id"`
}



type AwsVpcDhcpOptionsAssociationStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcDhcpOptionsAssociationList is a list of AwsVpcDhcpOptionsAssociations
type AwsVpcDhcpOptionsAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsVpcDhcpOptionsAssociation CRD objects
	Items []AwsVpcDhcpOptionsAssociation `json:"items,omitempty"`
}