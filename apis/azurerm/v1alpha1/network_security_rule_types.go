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

type NetworkSecurityRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSecurityRuleSpec   `json:"spec,omitempty"`
	Status            NetworkSecurityRuleStatus `json:"status,omitempty"`
}

type NetworkSecurityRuleSpec struct {
	Access string `json:"access"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	DestinationAddressPrefix string `json:"destination_address_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DestinationAddressPrefixes []string `json:"destination_address_prefixes,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	DestinationApplicationSecurityGroupIds []string `json:"destination_application_security_group_ids,omitempty"`
	// +optional
	DestinationPortRange string `json:"destination_port_range,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DestinationPortRanges    []string `json:"destination_port_ranges,omitempty"`
	Direction                string   `json:"direction"`
	Name                     string   `json:"name"`
	NetworkSecurityGroupName string   `json:"network_security_group_name"`
	Priority                 int      `json:"priority"`
	Protocol                 string   `json:"protocol"`
	ResourceGroupName        string   `json:"resource_group_name"`
	// +optional
	SourceAddressPrefix string `json:"source_address_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SourceAddressPrefixes []string `json:"source_address_prefixes,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	SourceApplicationSecurityGroupIds []string `json:"source_application_security_group_ids,omitempty"`
	// +optional
	SourcePortRange string `json:"source_port_range,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SourcePortRanges []string `json:"source_port_ranges,omitempty"`
}

type NetworkSecurityRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
