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

type NetworkSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSecurityGroupSpec   `json:"spec,omitempty"`
	Status            NetworkSecurityGroupStatus `json:"status,omitempty"`
}

type NetworkSecurityGroupSpecSecurityRule struct {
	Access string `json:"access" tf:"access"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	DestinationAddressPrefix string `json:"destinationAddressPrefix,omitempty" tf:"destination_address_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DestinationAddressPrefixes []string `json:"destinationAddressPrefixes,omitempty" tf:"destination_address_prefixes,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DestinationApplicationSecurityGroupIDS []string `json:"destinationApplicationSecurityGroupIDS,omitempty" tf:"destination_application_security_group_ids,omitempty"`
	// +optional
	DestinationPortRange string `json:"destinationPortRange,omitempty" tf:"destination_port_range,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DestinationPortRanges []string `json:"destinationPortRanges,omitempty" tf:"destination_port_ranges,omitempty"`
	Direction             string   `json:"direction" tf:"direction"`
	Name                  string   `json:"name" tf:"name"`
	Priority              int      `json:"priority" tf:"priority"`
	Protocol              string   `json:"protocol" tf:"protocol"`
	// +optional
	SourceAddressPrefix string `json:"sourceAddressPrefix,omitempty" tf:"source_address_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SourceAddressPrefixes []string `json:"sourceAddressPrefixes,omitempty" tf:"source_address_prefixes,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SourceApplicationSecurityGroupIDS []string `json:"sourceApplicationSecurityGroupIDS,omitempty" tf:"source_application_security_group_ids,omitempty"`
	// +optional
	SourcePortRange string `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SourcePortRanges []string `json:"sourcePortRanges,omitempty" tf:"source_port_ranges,omitempty"`
}

type NetworkSecurityGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Location          string `json:"location" tf:"location"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityRule []NetworkSecurityGroupSpecSecurityRule `json:"securityRule,omitempty" tf:"security_rule,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NetworkSecurityGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkSecurityGroupList is a list of NetworkSecurityGroups
type NetworkSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkSecurityGroup CRD objects
	Items []NetworkSecurityGroup `json:"items,omitempty"`
}
