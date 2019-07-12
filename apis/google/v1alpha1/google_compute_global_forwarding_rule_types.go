package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeGlobalForwardingRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeGlobalForwardingRuleSpec   `json:"spec,omitempty"`
	Status            GoogleComputeGlobalForwardingRuleStatus `json:"status,omitempty"`
}

type GoogleComputeGlobalForwardingRuleSpec struct {
	IpVersion        string            `json:"ip_version"`
	SelfLink         string            `json:"self_link"`
	Name             string            `json:"name"`
	Target           string            `json:"target"`
	IpAddress        string            `json:"ip_address"`
	IpProtocol       string            `json:"ip_protocol"`
	LabelFingerprint string            `json:"label_fingerprint"`
	PortRange        string            `json:"port_range"`
	Description      string            `json:"description"`
	Labels           map[string]string `json:"labels"`
	Project          string            `json:"project"`
	Region           string            `json:"region"`
}

type GoogleComputeGlobalForwardingRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeGlobalForwardingRuleList is a list of GoogleComputeGlobalForwardingRules
type GoogleComputeGlobalForwardingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeGlobalForwardingRule CRD objects
	Items []GoogleComputeGlobalForwardingRule `json:"items,omitempty"`
}
