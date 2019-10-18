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
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ComputeSecurityPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSecurityPolicySpec   `json:"spec,omitempty"`
	Status            ComputeSecurityPolicyStatus `json:"status,omitempty"`
}

type ComputeSecurityPolicySpecRuleMatchConfig struct {
	// +kubebuilder:validation:MaxItems=5
	// +kubebuilder:validation:MinItems=1
	SrcIPRanges []string `json:"srcIPRanges" tf:"src_ip_ranges"`
}

type ComputeSecurityPolicySpecRuleMatch struct {
	// +kubebuilder:validation:MaxItems=1
	Config        []ComputeSecurityPolicySpecRuleMatchConfig `json:"config" tf:"config"`
	VersionedExpr string                                     `json:"versionedExpr" tf:"versioned_expr"`
}

type ComputeSecurityPolicySpecRule struct {
	Action string `json:"action" tf:"action"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Match []ComputeSecurityPolicySpecRuleMatch `json:"match" tf:"match"`
	// +optional
	Preview  bool `json:"preview,omitempty" tf:"preview,omitempty"`
	Priority int  `json:"priority" tf:"priority"`
}

type ComputeSecurityPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Fingerprint string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Rule []ComputeSecurityPolicySpecRule `json:"rule,omitempty" tf:"rule,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type ComputeSecurityPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeSecurityPolicySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeSecurityPolicyList is a list of ComputeSecurityPolicys
type ComputeSecurityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeSecurityPolicy CRD objects
	Items []ComputeSecurityPolicy `json:"items,omitempty"`
}
