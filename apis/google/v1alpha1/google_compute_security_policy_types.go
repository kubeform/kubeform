package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeSecurityPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeSecurityPolicySpec   `json:"spec,omitempty"`
	Status            GoogleComputeSecurityPolicyStatus `json:"status,omitempty"`
}

type GoogleComputeSecurityPolicySpecRuleMatchConfig struct {
	SrcIpRanges []string `json:"src_ip_ranges"`
}

type GoogleComputeSecurityPolicySpecRuleMatch struct {
	Config        []GoogleComputeSecurityPolicySpecRuleMatch `json:"config"`
	VersionedExpr string                                     `json:"versioned_expr"`
}

type GoogleComputeSecurityPolicySpecRule struct {
	Action      string                                `json:"action"`
	Priority    int                                   `json:"priority"`
	Match       []GoogleComputeSecurityPolicySpecRule `json:"match"`
	Description string                                `json:"description"`
	Preview     bool                                  `json:"preview"`
}

type GoogleComputeSecurityPolicySpec struct {
	Name        string                            `json:"name"`
	Description string                            `json:"description"`
	Project     string                            `json:"project"`
	Rule        []GoogleComputeSecurityPolicySpec `json:"rule"`
	Fingerprint string                            `json:"fingerprint"`
	SelfLink    string                            `json:"self_link"`
}



type GoogleComputeSecurityPolicyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeSecurityPolicyList is a list of GoogleComputeSecurityPolicys
type GoogleComputeSecurityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeSecurityPolicy CRD objects
	Items []GoogleComputeSecurityPolicy `json:"items,omitempty"`
}