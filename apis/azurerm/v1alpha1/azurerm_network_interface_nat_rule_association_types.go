package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermNetworkInterfaceNatRuleAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermNetworkInterfaceNatRuleAssociationSpec   `json:"spec,omitempty"`
	Status            AzurermNetworkInterfaceNatRuleAssociationStatus `json:"status,omitempty"`
}

type AzurermNetworkInterfaceNatRuleAssociationSpec struct {
	IpConfigurationName string `json:"ip_configuration_name"`
	NatRuleId           string `json:"nat_rule_id"`
	NetworkInterfaceId  string `json:"network_interface_id"`
}



type AzurermNetworkInterfaceNatRuleAssociationStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermNetworkInterfaceNatRuleAssociationList is a list of AzurermNetworkInterfaceNatRuleAssociations
type AzurermNetworkInterfaceNatRuleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNetworkInterfaceNatRuleAssociation CRD objects
	Items []AzurermNetworkInterfaceNatRuleAssociation `json:"items,omitempty"`
}