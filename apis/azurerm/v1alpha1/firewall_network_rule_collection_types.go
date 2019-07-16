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

type FirewallNetworkRuleCollection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallNetworkRuleCollectionSpec   `json:"spec,omitempty"`
	Status            FirewallNetworkRuleCollectionStatus `json:"status,omitempty"`
}

type FirewallNetworkRuleCollectionSpecRule struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	DestinationAddresses []string `json:"destination_addresses"`
	// +kubebuilder:validation:UniqueItems=true
	DestinationPorts []string `json:"destination_ports"`
	Name             string   `json:"name"`
	// +kubebuilder:validation:UniqueItems=true
	Protocols []string `json:"protocols"`
	// +kubebuilder:validation:UniqueItems=true
	SourceAddresses []string `json:"source_addresses"`
}

type FirewallNetworkRuleCollectionSpec struct {
	Action            string `json:"action"`
	AzureFirewallName string `json:"azure_firewall_name"`
	Name              string `json:"name"`
	Priority          int    `json:"priority"`
	ResourceGroupName string `json:"resource_group_name"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Rule []FirewallNetworkRuleCollectionSpec `json:"rule"`
}

type FirewallNetworkRuleCollectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FirewallNetworkRuleCollectionList is a list of FirewallNetworkRuleCollections
type FirewallNetworkRuleCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FirewallNetworkRuleCollection CRD objects
	Items []FirewallNetworkRuleCollection `json:"items,omitempty"`
}
