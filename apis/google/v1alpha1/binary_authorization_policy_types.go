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

type BinaryAuthorizationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BinaryAuthorizationPolicySpec   `json:"spec,omitempty"`
	Status            BinaryAuthorizationPolicyStatus `json:"status,omitempty"`
}

type BinaryAuthorizationPolicySpecAdmissionWhitelistPatterns struct {
	// +optional
	NamePattern string `json:"name_pattern,omitempty"`
}

type BinaryAuthorizationPolicySpecClusterAdmissionRules struct {
	Cluster string `json:"cluster"`
	// +optional
	EnforcementMode string `json:"enforcement_mode,omitempty"`
	// +optional
	EvaluationMode string `json:"evaluation_mode,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	RequireAttestationsBy []string `json:"require_attestations_by,omitempty"`
}

type BinaryAuthorizationPolicySpecDefaultAdmissionRule struct {
	EnforcementMode string `json:"enforcement_mode"`
	EvaluationMode  string `json:"evaluation_mode"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	RequireAttestationsBy []string `json:"require_attestations_by,omitempty"`
}

type BinaryAuthorizationPolicySpec struct {
	// +optional
	AdmissionWhitelistPatterns *[]BinaryAuthorizationPolicySpec `json:"admission_whitelist_patterns,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ClusterAdmissionRules *[]BinaryAuthorizationPolicySpec `json:"cluster_admission_rules,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	DefaultAdmissionRule []BinaryAuthorizationPolicySpec `json:"default_admission_rule"`
	// +optional
	Description string `json:"description,omitempty"`
}

type BinaryAuthorizationPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
