package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeSecurityPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSecurityPolicySpec   `json:"spec,omitempty"`
	Status            ComputeSecurityPolicyStatus `json:"status,omitempty"`
}

type ComputeSecurityPolicySpecRuleMatchConfig struct {
	// +kubebuilder:validation:MaxItems=5
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
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

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Rule []ComputeSecurityPolicySpecRule `json:"rule,omitempty" tf:"rule,omitempty"`
}

type ComputeSecurityPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
