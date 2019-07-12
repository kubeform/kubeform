package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermSubnetNetworkSecurityGroupAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSubnetNetworkSecurityGroupAssociationSpec   `json:"spec,omitempty"`
	Status            AzurermSubnetNetworkSecurityGroupAssociationStatus `json:"status,omitempty"`
}

type AzurermSubnetNetworkSecurityGroupAssociationSpec struct {
	SubnetId               string `json:"subnet_id"`
	NetworkSecurityGroupId string `json:"network_security_group_id"`
}

type AzurermSubnetNetworkSecurityGroupAssociationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermSubnetNetworkSecurityGroupAssociationList is a list of AzurermSubnetNetworkSecurityGroupAssociations
type AzurermSubnetNetworkSecurityGroupAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSubnetNetworkSecurityGroupAssociation CRD objects
	Items []AzurermSubnetNetworkSecurityGroupAssociation `json:"items,omitempty"`
}
