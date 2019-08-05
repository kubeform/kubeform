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

type DlmLifecyclePolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DlmLifecyclePolicySpec   `json:"spec,omitempty"`
	Status            DlmLifecyclePolicyStatus `json:"status,omitempty"`
}

type DlmLifecyclePolicySpecPolicyDetailsScheduleCreateRule struct {
	Interval int `json:"interval" tf:"interval"`
	// +optional
	IntervalUnit string `json:"intervalUnit,omitempty" tf:"interval_unit,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Times []string `json:"times,omitempty" tf:"times,omitempty"`
}

type DlmLifecyclePolicySpecPolicyDetailsScheduleRetainRule struct {
	Count int `json:"count" tf:"count"`
}

type DlmLifecyclePolicySpecPolicyDetailsSchedule struct {
	// +optional
	CopyTags bool `json:"copyTags,omitempty" tf:"copy_tags,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	CreateRule []DlmLifecyclePolicySpecPolicyDetailsScheduleCreateRule `json:"createRule" tf:"create_rule"`
	Name       string                                                  `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	RetainRule []DlmLifecyclePolicySpecPolicyDetailsScheduleRetainRule `json:"retainRule" tf:"retain_rule"`
	// +optional
	TagsToAdd map[string]string `json:"tagsToAdd,omitempty" tf:"tags_to_add,omitempty"`
}

type DlmLifecyclePolicySpecPolicyDetails struct {
	ResourceTypes []string                                      `json:"resourceTypes" tf:"resource_types"`
	Schedule      []DlmLifecyclePolicySpecPolicyDetailsSchedule `json:"schedule" tf:"schedule"`
	TargetTags    map[string]string                             `json:"targetTags" tf:"target_tags"`
}

type DlmLifecyclePolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Description      string `json:"description" tf:"description"`
	ExecutionRoleArn string `json:"executionRoleArn" tf:"execution_role_arn"`
	// +kubebuilder:validation:MaxItems=1
	PolicyDetails []DlmLifecyclePolicySpecPolicyDetails `json:"policyDetails" tf:"policy_details"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
}

type DlmLifecyclePolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DlmLifecyclePolicySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DlmLifecyclePolicyList is a list of DlmLifecyclePolicys
type DlmLifecyclePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DlmLifecyclePolicy CRD objects
	Items []DlmLifecyclePolicy `json:"items,omitempty"`
}
