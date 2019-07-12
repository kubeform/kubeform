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

type AzurermNetworkSecurityRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermNetworkSecurityRuleSpec   `json:"spec,omitempty"`
	Status            AzurermNetworkSecurityRuleStatus `json:"status,omitempty"`
}

type AzurermNetworkSecurityRuleSpec struct {
	NetworkSecurityGroupName               string   `json:"network_security_group_name"`
	SourcePortRange                        string   `json:"source_port_range"`
	DestinationPortRanges                  []string `json:"destination_port_ranges"`
	SourceAddressPrefix                    string   `json:"source_address_prefix"`
	DestinationAddressPrefixes             []string `json:"destination_address_prefixes"`
	DestinationApplicationSecurityGroupIds []string `json:"destination_application_security_group_ids"`
	Direction                              string   `json:"direction"`
	Description                            string   `json:"description"`
	Protocol                               string   `json:"protocol"`
	DestinationPortRange                   string   `json:"destination_port_range"`
	SourceAddressPrefixes                  []string `json:"source_address_prefixes"`
	DestinationAddressPrefix               string   `json:"destination_address_prefix"`
	Priority                               int      `json:"priority"`
	ResourceGroupName                      string   `json:"resource_group_name"`
	Name                                   string   `json:"name"`
	SourcePortRanges                       []string `json:"source_port_ranges"`
	SourceApplicationSecurityGroupIds      []string `json:"source_application_security_group_ids"`
	Access                                 string   `json:"access"`
}

type AzurermNetworkSecurityRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermNetworkSecurityRuleList is a list of AzurermNetworkSecurityRules
type AzurermNetworkSecurityRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNetworkSecurityRule CRD objects
	Items []AzurermNetworkSecurityRule `json:"items,omitempty"`
}
