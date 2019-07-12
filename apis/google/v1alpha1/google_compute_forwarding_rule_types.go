package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeForwardingRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeForwardingRuleSpec   `json:"spec,omitempty"`
	Status            GoogleComputeForwardingRuleStatus `json:"status,omitempty"`
}

type GoogleComputeForwardingRuleSpec struct {
	NetworkTier         string            `json:"network_tier"`
	CreationTimestamp   string            `json:"creation_timestamp"`
	LabelFingerprint    string            `json:"label_fingerprint"`
	SelfLink            string            `json:"self_link"`
	IpAddress           string            `json:"ip_address"`
	LoadBalancingScheme string            `json:"load_balancing_scheme"`
	Description         string            `json:"description"`
	Network             string            `json:"network"`
	Subnetwork          string            `json:"subnetwork"`
	Target              string            `json:"target"`
	ServiceName         string            `json:"service_name"`
	Name                string            `json:"name"`
	BackendService      string            `json:"backend_service"`
	Project             string            `json:"project"`
	PortRange           string            `json:"port_range"`
	Ports               []string          `json:"ports"`
	Labels              map[string]string `json:"labels"`
	Region              string            `json:"region"`
	ServiceLabel        string            `json:"service_label"`
	IpProtocol          string            `json:"ip_protocol"`
	IpVersion           string            `json:"ip_version"`
}



type GoogleComputeForwardingRuleStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeForwardingRuleList is a list of GoogleComputeForwardingRules
type GoogleComputeForwardingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeForwardingRule CRD objects
	Items []GoogleComputeForwardingRule `json:"items,omitempty"`
}