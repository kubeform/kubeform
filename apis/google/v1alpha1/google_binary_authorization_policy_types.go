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

type GoogleBinaryAuthorizationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleBinaryAuthorizationPolicySpec   `json:"spec,omitempty"`
	Status            GoogleBinaryAuthorizationPolicyStatus `json:"status,omitempty"`
}

type GoogleBinaryAuthorizationPolicySpecAdmissionWhitelistPatterns struct {
	NamePattern string `json:"name_pattern"`
}

type GoogleBinaryAuthorizationPolicySpecClusterAdmissionRules struct {
	EvaluationMode        string   `json:"evaluation_mode"`
	RequireAttestationsBy []string `json:"require_attestations_by"`
	Cluster               string   `json:"cluster"`
	EnforcementMode       string   `json:"enforcement_mode"`
}

type GoogleBinaryAuthorizationPolicySpecDefaultAdmissionRule struct {
	EnforcementMode       string   `json:"enforcement_mode"`
	EvaluationMode        string   `json:"evaluation_mode"`
	RequireAttestationsBy []string `json:"require_attestations_by"`
}

type GoogleBinaryAuthorizationPolicySpec struct {
	AdmissionWhitelistPatterns []GoogleBinaryAuthorizationPolicySpec `json:"admission_whitelist_patterns"`
	ClusterAdmissionRules      []GoogleBinaryAuthorizationPolicySpec `json:"cluster_admission_rules"`
	Description                string                                `json:"description"`
	Project                    string                                `json:"project"`
	DefaultAdmissionRule       []GoogleBinaryAuthorizationPolicySpec `json:"default_admission_rule"`
}



type GoogleBinaryAuthorizationPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleBinaryAuthorizationPolicyList is a list of GoogleBinaryAuthorizationPolicys
type GoogleBinaryAuthorizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleBinaryAuthorizationPolicy CRD objects
	Items []GoogleBinaryAuthorizationPolicy `json:"items,omitempty"`
}