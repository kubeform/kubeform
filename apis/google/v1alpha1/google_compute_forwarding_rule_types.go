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

type GoogleComputeForwardingRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeForwardingRuleSpec   `json:"spec,omitempty"`
	Status            GoogleComputeForwardingRuleStatus `json:"status,omitempty"`
}

type GoogleComputeForwardingRuleSpec struct {
	Description         string            `json:"description"`
	Network             string            `json:"network"`
	PortRange           string            `json:"port_range"`
	LabelFingerprint    string            `json:"label_fingerprint"`
	IpAddress           string            `json:"ip_address"`
	IpProtocol          string            `json:"ip_protocol"`
	Ports               []string          `json:"ports"`
	Region              string            `json:"region"`
	ServiceLabel        string            `json:"service_label"`
	BackendService      string            `json:"backend_service"`
	LoadBalancingScheme string            `json:"load_balancing_scheme"`
	Subnetwork          string            `json:"subnetwork"`
	Target              string            `json:"target"`
	ServiceName         string            `json:"service_name"`
	SelfLink            string            `json:"self_link"`
	Name                string            `json:"name"`
	Labels              map[string]string `json:"labels"`
	CreationTimestamp   string            `json:"creation_timestamp"`
	Project             string            `json:"project"`
	IpVersion           string            `json:"ip_version"`
	NetworkTier         string            `json:"network_tier"`
}

type GoogleComputeForwardingRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeForwardingRuleList is a list of GoogleComputeForwardingRules
type GoogleComputeForwardingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeForwardingRule CRD objects
	Items []GoogleComputeForwardingRule `json:"items,omitempty"`
}
