package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermNetworkSecurityRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermNetworkSecurityRuleSpec   `json:"spec,omitempty"`
	Status            AzurermNetworkSecurityRuleStatus `json:"status,omitempty"`
}

type AzurermNetworkSecurityRuleSpec struct {
	Name                                   string   `json:"name"`
	ResourceGroupName                      string   `json:"resource_group_name"`
	Description                            string   `json:"description"`
	DestinationApplicationSecurityGroupIds []string `json:"destination_application_security_group_ids"`
	Direction                              string   `json:"direction"`
	DestinationPortRange                   string   `json:"destination_port_range"`
	SourceAddressPrefixes                  []string `json:"source_address_prefixes"`
	DestinationAddressPrefix               string   `json:"destination_address_prefix"`
	SourceAddressPrefix                    string   `json:"source_address_prefix"`
	SourceApplicationSecurityGroupIds      []string `json:"source_application_security_group_ids"`
	Access                                 string   `json:"access"`
	Priority                               int      `json:"priority"`
	Protocol                               string   `json:"protocol"`
	SourcePortRanges                       []string `json:"source_port_ranges"`
	DestinationPortRanges                  []string `json:"destination_port_ranges"`
	NetworkSecurityGroupName               string   `json:"network_security_group_name"`
	SourcePortRange                        string   `json:"source_port_range"`
	DestinationAddressPrefixes             []string `json:"destination_address_prefixes"`
}

type AzurermNetworkSecurityRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermNetworkSecurityRuleList is a list of AzurermNetworkSecurityRules
type AzurermNetworkSecurityRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNetworkSecurityRule CRD objects
	Items []AzurermNetworkSecurityRule `json:"items,omitempty"`
}
