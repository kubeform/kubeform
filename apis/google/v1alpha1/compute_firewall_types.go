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

type ComputeFirewall struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeFirewallSpec   `json:"spec,omitempty"`
	Status            ComputeFirewallStatus `json:"status,omitempty"`
}

type ComputeFirewallSpecAllow struct {
	// +optional
	Ports    []string `json:"ports,omitempty" tf:"ports,omitempty"`
	Protocol string   `json:"protocol" tf:"protocol"`
}

type ComputeFirewallSpecDeny struct {
	// +optional
	Ports    []string `json:"ports,omitempty" tf:"ports,omitempty"`
	Protocol string   `json:"protocol" tf:"protocol"`
}

type ComputeFirewallSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Allow []ComputeFirewallSpecAllow `json:"allow,omitempty" tf:"allow,omitempty"`
	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
	// +optional
	Deny []ComputeFirewallSpecDeny `json:"deny,omitempty" tf:"deny,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	DestinationRanges []string `json:"destinationRanges,omitempty" tf:"destination_ranges,omitempty"`
	// +optional
	Direction string `json:"direction,omitempty" tf:"direction,omitempty"`
	// +optional
	Disabled bool `json:"disabled,omitempty" tf:"disabled,omitempty"`
	// +optional
	// Deprecated
	EnableLogging bool   `json:"enableLogging,omitempty" tf:"enable_logging,omitempty"`
	Name          string `json:"name" tf:"name"`
	Network       string `json:"network" tf:"network"`
	// +optional
	Priority int `json:"priority,omitempty" tf:"priority,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	// +optional
	SourceRanges []string `json:"sourceRanges,omitempty" tf:"source_ranges,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceServiceAccounts []string `json:"sourceServiceAccounts,omitempty" tf:"source_service_accounts,omitempty"`
	// +optional
	SourceTags []string `json:"sourceTags,omitempty" tf:"source_tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TargetServiceAccounts []string `json:"targetServiceAccounts,omitempty" tf:"target_service_accounts,omitempty"`
	// +optional
	TargetTags []string `json:"targetTags,omitempty" tf:"target_tags,omitempty"`
}

type ComputeFirewallStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeFirewallSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeFirewallList is a list of ComputeFirewalls
type ComputeFirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeFirewall CRD objects
	Items []ComputeFirewall `json:"items,omitempty"`
}
