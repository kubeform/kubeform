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

type GoogleComputeGlobalForwardingRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeGlobalForwardingRuleSpec   `json:"spec,omitempty"`
	Status            GoogleComputeGlobalForwardingRuleStatus `json:"status,omitempty"`
}

type GoogleComputeGlobalForwardingRuleSpec struct {
	Name             string            `json:"name"`
	Target           string            `json:"target"`
	Labels           map[string]string `json:"labels"`
	PortRange        string            `json:"port_range"`
	Region           string            `json:"region"`
	SelfLink         string            `json:"self_link"`
	Description      string            `json:"description"`
	IpAddress        string            `json:"ip_address"`
	IpProtocol       string            `json:"ip_protocol"`
	LabelFingerprint string            `json:"label_fingerprint"`
	IpVersion        string            `json:"ip_version"`
	Project          string            `json:"project"`
}



type GoogleComputeGlobalForwardingRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeGlobalForwardingRuleList is a list of GoogleComputeGlobalForwardingRules
type GoogleComputeGlobalForwardingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeGlobalForwardingRule CRD objects
	Items []GoogleComputeGlobalForwardingRule `json:"items,omitempty"`
}