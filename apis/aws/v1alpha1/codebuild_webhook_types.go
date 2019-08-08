package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	BranchFilter string `json:"branchFilter,omitempty" tf:"branch_filter,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	FilterGroup []CodebuildWebhookSpecFilterGroup `json:"filterGroup,omitempty" tf:"filter_group,omitempty"`
	// +optional
	PayloadURL  string `json:"payloadURL,omitempty" tf:"payload_url,omitempty"`
	ProjectName string `json:"projectName" tf:"project_name"`
	// +optional
	Secret string `json:"-" sensitive:"true" tf:"secret,omitempty"`
	// +optional
	Url string `json:"url,omitempty" tf:"url,omitempty"`
}

type CodebuildWebhookStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CodebuildWebhookSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
