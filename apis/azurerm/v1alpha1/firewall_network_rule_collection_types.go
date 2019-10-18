package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type FirewallNetworkRuleCollection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallNetworkRuleCollectionSpec   `json:"spec,omitempty"`
	Status            FirewallNetworkRuleCollectionStatus `json:"status,omitempty"`
}

type FirewallNetworkRuleCollectionSpecRule struct {
	// +optional
	Description          string   `json:"description,omitempty" tf:"description,omitempty"`
	DestinationAddresses []string `json:"destinationAddresses" tf:"destination_addresses"`
	DestinationPorts     []string `json:"destinationPorts" tf:"destination_ports"`
	Name                 string   `json:"name" tf:"name"`
	Protocols            []string `json:"protocols" tf:"protocols"`
	SourceAddresses      []string `json:"sourceAddresses" tf:"source_addresses"`
}

type FirewallNetworkRuleCollectionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Action            string `json:"action" tf:"action"`
	AzureFirewallName string `json:"azureFirewallName" tf:"azure_firewall_name"`
	Name              string `json:"name" tf:"name"`
	Priority          int    `json:"priority" tf:"priority"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MinItems=1
	Rule []FirewallNetworkRuleCollectionSpecRule `json:"rule" tf:"rule"`
}

type FirewallNetworkRuleCollectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *FirewallNetworkRuleCollectionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
