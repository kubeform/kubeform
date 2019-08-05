package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	Port int    `json:"port,omitempty" tf:"port,omitempty"`
	Type string `json:"type" tf:"type"`
}

type FirewallApplicationRuleCollectionSpecRule struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	FqdnTags []string `json:"fqdnTags,omitempty" tf:"fqdn_tags,omitempty"`
	Name     string   `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	Protocol []FirewallApplicationRuleCollectionSpecRuleProtocol `json:"protocol,omitempty" tf:"protocol,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	SourceAddresses []string `json:"sourceAddresses" tf:"source_addresses"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	TargetFqdns []string `json:"targetFqdns,omitempty" tf:"target_fqdns,omitempty"`
}

type FirewallApplicationRuleCollectionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Action            string `json:"action" tf:"action"`
	AzureFirewallName string `json:"azureFirewallName" tf:"azure_firewall_name"`
	Name              string `json:"name" tf:"name"`
	Priority          int    `json:"priority" tf:"priority"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MinItems=1
	Rule []FirewallApplicationRuleCollectionSpecRule `json:"rule" tf:"rule"`
}

type FirewallApplicationRuleCollectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *FirewallApplicationRuleCollectionSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
