package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type BinaryAuthorizationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BinaryAuthorizationPolicySpec   `json:"spec,omitempty"`
	Status            BinaryAuthorizationPolicyStatus `json:"status,omitempty"`
}

type BinaryAuthorizationPolicySpecAdmissionWhitelistPatterns struct {
	// +optional
	NamePattern string `json:"namePattern,omitempty" tf:"name_pattern,omitempty"`
}

type BinaryAuthorizationPolicySpecClusterAdmissionRules struct {
	Cluster string `json:"cluster" tf:"cluster"`
	// +optional
	EnforcementMode string `json:"enforcementMode,omitempty" tf:"enforcement_mode,omitempty"`
	// +optional
	EvaluationMode string `json:"evaluationMode,omitempty" tf:"evaluation_mode,omitempty"`
	// +optional
	RequireAttestationsBy []string `json:"requireAttestationsBy,omitempty" tf:"require_attestations_by,omitempty"`
}

type BinaryAuthorizationPolicySpecDefaultAdmissionRule struct {
	EnforcementMode string `json:"enforcementMode" tf:"enforcement_mode"`
	EvaluationMode  string `json:"evaluationMode" tf:"evaluation_mode"`
	// +optional
	RequireAttestationsBy []string `json:"requireAttestationsBy,omitempty" tf:"require_attestations_by,omitempty"`
}

type BinaryAuthorizationPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdmissionWhitelistPatterns []BinaryAuthorizationPolicySpecAdmissionWhitelistPatterns `json:"admissionWhitelistPatterns,omitempty" tf:"admission_whitelist_patterns,omitempty"`
	// +optional
	ClusterAdmissionRules []BinaryAuthorizationPolicySpecClusterAdmissionRules `json:"clusterAdmissionRules,omitempty" tf:"cluster_admission_rules,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	DefaultAdmissionRule []BinaryAuthorizationPolicySpecDefaultAdmissionRule `json:"defaultAdmissionRule" tf:"default_admission_rule"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
}

type BinaryAuthorizationPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *BinaryAuthorizationPolicySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BinaryAuthorizationPolicyList is a list of BinaryAuthorizationPolicys
type BinaryAuthorizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BinaryAuthorizationPolicy CRD objects
	Items []BinaryAuthorizationPolicy `json:"items,omitempty"`
}
