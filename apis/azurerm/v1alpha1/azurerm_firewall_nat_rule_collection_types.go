package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermFirewallNatRuleCollection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermFirewallNatRuleCollectionSpec   `json:"spec,omitempty"`
	Status            AzurermFirewallNatRuleCollectionStatus `json:"status,omitempty"`
}

type AzurermFirewallNatRuleCollectionSpecRule struct {
	Description          string   `json:"description"`
	TranslatedAddress    string   `json:"translated_address"`
	TranslatedPort       string   `json:"translated_port"`
	SourceAddresses      []string `json:"source_addresses"`
	DestinationAddresses []string `json:"destination_addresses"`
	DestinationPorts     []string `json:"destination_ports"`
	Protocols            []string `json:"protocols"`
	Name                 string   `json:"name"`
}

type AzurermFirewallNatRuleCollectionSpec struct {
	Priority          int                                    `json:"priority"`
	Action            string                                 `json:"action"`
	Rule              []AzurermFirewallNatRuleCollectionSpec `json:"rule"`
	Name              string                                 `json:"name"`
	AzureFirewallName string                                 `json:"azure_firewall_name"`
	ResourceGroupName string                                 `json:"resource_group_name"`
}

type AzurermFirewallNatRuleCollectionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermFirewallNatRuleCollectionList is a list of AzurermFirewallNatRuleCollections
type AzurermFirewallNatRuleCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermFirewallNatRuleCollection CRD objects
	Items []AzurermFirewallNatRuleCollection `json:"items,omitempty"`
}
