package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleBinaryAuthorizationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleBinaryAuthorizationPolicySpec   `json:"spec,omitempty"`
	Status            GoogleBinaryAuthorizationPolicyStatus `json:"status,omitempty"`
}

type GoogleBinaryAuthorizationPolicySpecDefaultAdmissionRule struct {
	EnforcementMode       string   `json:"enforcement_mode"`
	EvaluationMode        string   `json:"evaluation_mode"`
	RequireAttestationsBy []string `json:"require_attestations_by"`
}

type GoogleBinaryAuthorizationPolicySpecAdmissionWhitelistPatterns struct {
	NamePattern string `json:"name_pattern"`
}

type GoogleBinaryAuthorizationPolicySpecClusterAdmissionRules struct {
	Cluster               string   `json:"cluster"`
	EnforcementMode       string   `json:"enforcement_mode"`
	EvaluationMode        string   `json:"evaluation_mode"`
	RequireAttestationsBy []string `json:"require_attestations_by"`
}

type GoogleBinaryAuthorizationPolicySpec struct {
	DefaultAdmissionRule       []GoogleBinaryAuthorizationPolicySpec `json:"default_admission_rule"`
	AdmissionWhitelistPatterns []GoogleBinaryAuthorizationPolicySpec `json:"admission_whitelist_patterns"`
	ClusterAdmissionRules      []GoogleBinaryAuthorizationPolicySpec `json:"cluster_admission_rules"`
	Description                string                                `json:"description"`
	Project                    string                                `json:"project"`
}

type GoogleBinaryAuthorizationPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleBinaryAuthorizationPolicyList is a list of GoogleBinaryAuthorizationPolicys
type GoogleBinaryAuthorizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleBinaryAuthorizationPolicy CRD objects
	Items []GoogleBinaryAuthorizationPolicy `json:"items,omitempty"`
}
