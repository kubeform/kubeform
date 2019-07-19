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

type CodebuildWebhook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodebuildWebhookSpec   `json:"spec,omitempty"`
	Status            CodebuildWebhookStatus `json:"status,omitempty"`
}

type CodebuildWebhookSpecFilterGroupFilter struct {
	// +optional
	ExcludeMatchedPattern bool   `json:"excludeMatchedPattern,omitempty" tf:"exclude_matched_pattern,omitempty"`
	Pattern               string `json:"pattern" tf:"pattern"`
	Type                  string `json:"type" tf:"type"`
}

type CodebuildWebhookSpecFilterGroup struct {
	// +optional
	Filter []CodebuildWebhookSpecFilterGroupFilter `json:"filter,omitempty" tf:"filter,omitempty"`
}

type CodebuildWebhookSpec struct {
	// +optional
	BranchFilter string `json:"branchFilter,omitempty" tf:"branch_filter,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	FilterGroup []CodebuildWebhookSpecFilterGroup `json:"filterGroup,omitempty" tf:"filter_group,omitempty"`
	ProjectName string                            `json:"projectName" tf:"project_name"`
	ProviderRef core.LocalObjectReference         `json:"providerRef" tf:"-"`
}

type CodebuildWebhookStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CodebuildWebhookList is a list of CodebuildWebhooks
type CodebuildWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CodebuildWebhook CRD objects
	Items []CodebuildWebhook `json:"items,omitempty"`
}
