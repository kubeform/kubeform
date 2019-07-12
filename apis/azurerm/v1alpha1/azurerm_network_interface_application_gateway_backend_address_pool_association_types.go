package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationSpec   `json:"spec,omitempty"`
	Status            AzurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationStatus `json:"status,omitempty"`
}

type AzurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationSpec struct {
	NetworkInterfaceId   string `json:"network_interface_id"`
	IpConfigurationName  string `json:"ip_configuration_name"`
	BackendAddressPoolId string `json:"backend_address_pool_id"`
}

type AzurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList is a list of AzurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociations
type AzurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation CRD objects
	Items []AzurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation `json:"items,omitempty"`
}
