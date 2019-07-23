package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type FirewallNATRuleCollection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallNATRuleCollectionSpec   `json:"spec,omitempty"`
	Status            FirewallNATRuleCollectionStatus `json:"status,omitempty"`
}

type FirewallNATRuleCollectionSpecRule struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	DestinationAddresses []string `json:"destinationAddresses" tf:"destination_addresses"`
	// +kubebuilder:validation:UniqueItems=true
	DestinationPorts []string `json:"destinationPorts" tf:"destination_ports"`
	Name             string   `json:"name" tf:"name"`
	// +kubebuilder:validation:UniqueItems=true
	Protocols []string `json:"protocols" tf:"protocols"`
	// +kubebuilder:validation:UniqueItems=true
	SourceAddresses   []string `json:"sourceAddresses" tf:"source_addresses"`
	TranslatedAddress string   `json:"translatedAddress" tf:"translated_address"`
	TranslatedPort    string   `json:"translatedPort" tf:"translated_port"`
}

type FirewallNATRuleCollectionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Action            string `json:"action" tf:"action"`
	AzureFirewallName string `json:"azureFirewallName" tf:"azure_firewall_name"`
	Name              string `json:"name" tf:"name"`
	Priority          int    `json:"priority" tf:"priority"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Rule []FirewallNATRuleCollectionSpecRule `json:"rule" tf:"rule"`
}

type FirewallNATRuleCollectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FirewallNATRuleCollectionList is a list of FirewallNATRuleCollections
type FirewallNATRuleCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FirewallNATRuleCollection CRD objects
	Items []FirewallNATRuleCollection `json:"items,omitempty"`
}
