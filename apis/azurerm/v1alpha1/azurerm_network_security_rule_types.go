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
	SourceAddressPrefix                    string   `json:"source_address_prefix"`
	DestinationAddressPrefix               string   `json:"destination_address_prefix"`
	SourceAddressPrefixes                  []string `json:"source_address_prefixes"`
	DestinationAddressPrefixes             []string `json:"destination_address_prefixes"`
	Priority                               int      `json:"priority"`
	SourceApplicationSecurityGroupIds      []string `json:"source_application_security_group_ids"`
	DestinationApplicationSecurityGroupIds []string `json:"destination_application_security_group_ids"`
	Access                                 string   `json:"access"`
	Direction                              string   `json:"direction"`
	ResourceGroupName                      string   `json:"resource_group_name"`
	SourcePortRange                        string   `json:"source_port_range"`
	SourcePortRanges                       []string `json:"source_port_ranges"`
	DestinationPortRange                   string   `json:"destination_port_range"`
	DestinationPortRanges                  []string `json:"destination_port_ranges"`
	Name                                   string   `json:"name"`
	NetworkSecurityGroupName               string   `json:"network_security_group_name"`
	Description                            string   `json:"description"`
	Protocol                               string   `json:"protocol"`
}



type AzurermNetworkSecurityRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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