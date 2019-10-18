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

type NetworkSecurityRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSecurityRuleSpec   `json:"spec,omitempty"`
	Status            NetworkSecurityRuleStatus `json:"status,omitempty"`
}

type NetworkSecurityRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Access string `json:"access" tf:"access"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	DestinationAddressPrefix string `json:"destinationAddressPrefix,omitempty" tf:"destination_address_prefix,omitempty"`
	// +optional
	DestinationAddressPrefixes []string `json:"destinationAddressPrefixes,omitempty" tf:"destination_address_prefixes,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DestinationApplicationSecurityGroupIDS []string `json:"destinationApplicationSecurityGroupIDS,omitempty" tf:"destination_application_security_group_ids,omitempty"`
	// +optional
	DestinationPortRange string `json:"destinationPortRange,omitempty" tf:"destination_port_range,omitempty"`
	// +optional
	DestinationPortRanges    []string `json:"destinationPortRanges,omitempty" tf:"destination_port_ranges,omitempty"`
	Direction                string   `json:"direction" tf:"direction"`
	Name                     string   `json:"name" tf:"name"`
	NetworkSecurityGroupName string   `json:"networkSecurityGroupName" tf:"network_security_group_name"`
	Priority                 int      `json:"priority" tf:"priority"`
	Protocol                 string   `json:"protocol" tf:"protocol"`
	ResourceGroupName        string   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SourceAddressPrefix string `json:"sourceAddressPrefix,omitempty" tf:"source_address_prefix,omitempty"`
	// +optional
	SourceAddressPrefixes []string `json:"sourceAddressPrefixes,omitempty" tf:"source_address_prefixes,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceApplicationSecurityGroupIDS []string `json:"sourceApplicationSecurityGroupIDS,omitempty" tf:"source_application_security_group_ids,omitempty"`
	// +optional
	SourcePortRange string `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`
	// +optional
	SourcePortRanges []string `json:"sourcePortRanges,omitempty" tf:"source_port_ranges,omitempty"`
}

type NetworkSecurityRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NetworkSecurityRuleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkSecurityRuleList is a list of NetworkSecurityRules
type NetworkSecurityRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkSecurityRule CRD objects
	Items []NetworkSecurityRule `json:"items,omitempty"`
}
