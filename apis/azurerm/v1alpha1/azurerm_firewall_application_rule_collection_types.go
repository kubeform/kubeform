package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermFirewallApplicationRuleCollection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermFirewallApplicationRuleCollectionSpec   `json:"spec,omitempty"`
	Status            AzurermFirewallApplicationRuleCollectionStatus `json:"status,omitempty"`
}

type AzurermFirewallApplicationRuleCollectionSpecRuleProtocol struct {
	Type string `json:"type"`
	Port int    `json:"port"`
}

type AzurermFirewallApplicationRuleCollectionSpecRule struct {
	Description     string                                             `json:"description"`
	FqdnTags        []string                                           `json:"fqdn_tags"`
	TargetFqdns     []string                                           `json:"target_fqdns"`
	Protocol        []AzurermFirewallApplicationRuleCollectionSpecRule `json:"protocol"`
	Name            string                                             `json:"name"`
	SourceAddresses []string                                           `json:"source_addresses"`
}

type AzurermFirewallApplicationRuleCollectionSpec struct {
	ResourceGroupName string                                         `json:"resource_group_name"`
	Priority          int                                            `json:"priority"`
	Action            string                                         `json:"action"`
	Rule              []AzurermFirewallApplicationRuleCollectionSpec `json:"rule"`
	Name              string                                         `json:"name"`
	AzureFirewallName string                                         `json:"azure_firewall_name"`
}



type AzurermFirewallApplicationRuleCollectionStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermFirewallApplicationRuleCollectionList is a list of AzurermFirewallApplicationRuleCollections
type AzurermFirewallApplicationRuleCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermFirewallApplicationRuleCollection CRD objects
	Items []AzurermFirewallApplicationRuleCollection `json:"items,omitempty"`
}