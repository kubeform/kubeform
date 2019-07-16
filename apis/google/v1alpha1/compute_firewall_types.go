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

type ComputeFirewall struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeFirewallSpec   `json:"spec,omitempty"`
	Status            ComputeFirewallStatus `json:"status,omitempty"`
}

type ComputeFirewallSpecAllow struct {
	// +optional
	Ports    []string `json:"ports,omitempty"`
	Protocol string   `json:"protocol"`
}

type ComputeFirewallSpecDeny struct {
	// +optional
	Ports    []string `json:"ports,omitempty"`
	Protocol string   `json:"protocol"`
}

type ComputeFirewallSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Allow *[]ComputeFirewallSpec `json:"allow,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Deny *[]ComputeFirewallSpec `json:"deny,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Disabled bool `json:"disabled,omitempty"`
	// +optional
	// Deprecated
	EnableLogging bool   `json:"enable_logging,omitempty"`
	Name          string `json:"name"`
	Network       string `json:"network"`
	// +optional
	Priority int `json:"priority,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	SourceServiceAccounts []string `json:"source_service_accounts,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SourceTags []string `json:"source_tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	TargetServiceAccounts []string `json:"target_service_accounts,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	TargetTags []string `json:"target_tags,omitempty"`
}

type ComputeFirewallStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
