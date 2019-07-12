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

type AzurermFirewallNetworkRuleCollection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermFirewallNetworkRuleCollectionSpec   `json:"spec,omitempty"`
	Status            AzurermFirewallNetworkRuleCollectionStatus `json:"status,omitempty"`
}

type AzurermFirewallNetworkRuleCollectionSpecRule struct {
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	SourceAddresses      []string `json:"source_addresses"`
	DestinationAddresses []string `json:"destination_addresses"`
	DestinationPorts     []string `json:"destination_ports"`
	Protocols            []string `json:"protocols"`
}

type AzurermFirewallNetworkRuleCollectionSpec struct {
	Name              string                                     `json:"name"`
	AzureFirewallName string                                     `json:"azure_firewall_name"`
	ResourceGroupName string                                     `json:"resource_group_name"`
	Priority          int                                        `json:"priority"`
	Action            string                                     `json:"action"`
	Rule              []AzurermFirewallNetworkRuleCollectionSpec `json:"rule"`
}

type AzurermFirewallNetworkRuleCollectionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermFirewallNetworkRuleCollectionList is a list of AzurermFirewallNetworkRuleCollections
type AzurermFirewallNetworkRuleCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermFirewallNetworkRuleCollection CRD objects
	Items []AzurermFirewallNetworkRuleCollection `json:"items,omitempty"`
}
