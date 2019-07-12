package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermNetworkInterfaceApplicationSecurityGroupAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermNetworkInterfaceApplicationSecurityGroupAssociationSpec   `json:"spec,omitempty"`
	Status            AzurermNetworkInterfaceApplicationSecurityGroupAssociationStatus `json:"status,omitempty"`
}

type AzurermNetworkInterfaceApplicationSecurityGroupAssociationSpec struct {
	NetworkInterfaceId         string `json:"network_interface_id"`
	IpConfigurationName        string `json:"ip_configuration_name"`
	ApplicationSecurityGroupId string `json:"application_security_group_id"`
}



type AzurermNetworkInterfaceApplicationSecurityGroupAssociationStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermNetworkInterfaceApplicationSecurityGroupAssociationList is a list of AzurermNetworkInterfaceApplicationSecurityGroupAssociations
type AzurermNetworkInterfaceApplicationSecurityGroupAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNetworkInterfaceApplicationSecurityGroupAssociation CRD objects
	Items []AzurermNetworkInterfaceApplicationSecurityGroupAssociation `json:"items,omitempty"`
}