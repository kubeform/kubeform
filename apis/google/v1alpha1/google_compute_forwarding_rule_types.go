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
	IpProtocol          string            `json:"ip_protocol"`
	Target              string            `json:"target"`
	Project             string            `json:"project"`
	LabelFingerprint    string            `json:"label_fingerprint"`
	Description         string            `json:"description"`
	Labels              map[string]string `json:"labels"`
	LoadBalancingScheme string            `json:"load_balancing_scheme"`
	Network             string            `json:"network"`
	Region              string            `json:"region"`
	SelfLink            string            `json:"self_link"`
	BackendService      string            `json:"backend_service"`
	IpVersion           string            `json:"ip_version"`
	NetworkTier         string            `json:"network_tier"`
	PortRange           string            `json:"port_range"`
	ServiceLabel        string            `json:"service_label"`
	ServiceName         string            `json:"service_name"`
	Name                string            `json:"name"`
	IpAddress           string            `json:"ip_address"`
	Ports               []string          `json:"ports"`
	Subnetwork          string            `json:"subnetwork"`
	CreationTimestamp   string            `json:"creation_timestamp"`
}



type GoogleComputeForwardingRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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