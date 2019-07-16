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

type CodebuildWebhook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodebuildWebhookSpec   `json:"spec,omitempty"`
	Status            CodebuildWebhookStatus `json:"status,omitempty"`
}

type CodebuildWebhookSpecFilterGroupFilter struct {
	// +optional
	ExcludeMatchedPattern bool   `json:"exclude_matched_pattern,omitempty"`
	Pattern               string `json:"pattern"`
	Type                  string `json:"type"`
}

type CodebuildWebhookSpecFilterGroup struct {
	// +optional
	Filter *[]CodebuildWebhookSpecFilterGroup `json:"filter,omitempty"`
}

type CodebuildWebhookSpec struct {
	// +optional
	BranchFilter string `json:"branch_filter,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	FilterGroup *[]CodebuildWebhookSpec `json:"filter_group,omitempty"`
	ProjectName string                  `json:"project_name"`
}

type CodebuildWebhookStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
