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

type ComputeGlobalForwardingRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeGlobalForwardingRuleSpec   `json:"spec,omitempty"`
	Status            ComputeGlobalForwardingRuleStatus `json:"status,omitempty"`
}

type ComputeGlobalForwardingRuleSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	IpVersion string `json:"ip_version,omitempty"`
	// +optional
	// Deprecated
	Labels map[string]string `json:"labels,omitempty"`
	Name   string            `json:"name"`
	// +optional
	PortRange string `json:"port_range,omitempty"`
	Target    string `json:"target"`
}

type ComputeGlobalForwardingRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeGlobalForwardingRuleList is a list of ComputeGlobalForwardingRules
type ComputeGlobalForwardingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeGlobalForwardingRule CRD objects
	Items []ComputeGlobalForwardingRule `json:"items,omitempty"`
}
