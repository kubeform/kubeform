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

type FirewallApplicationRuleCollection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallApplicationRuleCollectionSpec   `json:"spec,omitempty"`
	Status            FirewallApplicationRuleCollectionStatus `json:"status,omitempty"`
}

type FirewallApplicationRuleCollectionSpecRuleProtocol struct {
	// +optional
	Port int    `json:"port,omitempty"`
	Type string `json:"type"`
}

type FirewallApplicationRuleCollectionSpecRule struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	FqdnTags []string `json:"fqdn_tags,omitempty"`
	Name     string   `json:"name"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	Protocol *[]FirewallApplicationRuleCollectionSpecRule `json:"protocol,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	SourceAddresses []string `json:"source_addresses"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	TargetFqdns []string `json:"target_fqdns,omitempty"`
}

type FirewallApplicationRuleCollectionSpec struct {
	Action            string `json:"action"`
	AzureFirewallName string `json:"azure_firewall_name"`
	Name              string `json:"name"`
	Priority          int    `json:"priority"`
	ResourceGroupName string `json:"resource_group_name"`
	// +kubebuilder:validation:MinItems=1
	Rule []FirewallApplicationRuleCollectionSpec `json:"rule"`
}

type FirewallApplicationRuleCollectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FirewallApplicationRuleCollectionList is a list of FirewallApplicationRuleCollections
type FirewallApplicationRuleCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FirewallApplicationRuleCollection CRD objects
	Items []FirewallApplicationRuleCollection `json:"items,omitempty"`
}
