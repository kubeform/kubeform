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

type DlmLifecyclePolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DlmLifecyclePolicySpec   `json:"spec,omitempty"`
	Status            DlmLifecyclePolicyStatus `json:"status,omitempty"`
}

type DlmLifecyclePolicySpecPolicyDetailsScheduleCreateRule struct {
	Interval int `json:"interval"`
	// +optional
	IntervalUnit string `json:"interval_unit,omitempty"`
}

type DlmLifecyclePolicySpecPolicyDetailsScheduleRetainRule struct {
	Count int `json:"count"`
}

type DlmLifecyclePolicySpecPolicyDetailsSchedule struct {
	// +kubebuilder:validation:MaxItems=1
	CreateRule []DlmLifecyclePolicySpecPolicyDetailsSchedule `json:"create_rule"`
	Name       string                                        `json:"name"`
	// +kubebuilder:validation:MaxItems=1
	RetainRule []DlmLifecyclePolicySpecPolicyDetailsSchedule `json:"retain_rule"`
	// +optional
	TagsToAdd map[string]string `json:"tags_to_add,omitempty"`
}

type DlmLifecyclePolicySpecPolicyDetails struct {
	ResourceTypes []string                              `json:"resource_types"`
	Schedule      []DlmLifecyclePolicySpecPolicyDetails `json:"schedule"`
	TargetTags    map[string]string                     `json:"target_tags"`
}

type DlmLifecyclePolicySpec struct {
	Description      string `json:"description"`
	ExecutionRoleArn string `json:"execution_role_arn"`
	// +kubebuilder:validation:MaxItems=1
	PolicyDetails []DlmLifecyclePolicySpec `json:"policy_details"`
	// +optional
	State string `json:"state,omitempty"`
}

type DlmLifecyclePolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
