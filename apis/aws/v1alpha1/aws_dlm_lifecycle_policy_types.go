package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDlmLifecyclePolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDlmLifecyclePolicySpec   `json:"spec,omitempty"`
	Status            AwsDlmLifecyclePolicyStatus `json:"status,omitempty"`
}

type AwsDlmLifecyclePolicySpecPolicyDetailsScheduleCreateRule struct {
	IntervalUnit string   `json:"interval_unit"`
	Times        []string `json:"times"`
	Interval     int      `json:"interval"`
}

type AwsDlmLifecyclePolicySpecPolicyDetailsScheduleRetainRule struct {
	Count int `json:"count"`
}

type AwsDlmLifecyclePolicySpecPolicyDetailsSchedule struct {
	CreateRule []AwsDlmLifecyclePolicySpecPolicyDetailsSchedule `json:"create_rule"`
	Name       string                                           `json:"name"`
	RetainRule []AwsDlmLifecyclePolicySpecPolicyDetailsSchedule `json:"retain_rule"`
	TagsToAdd  map[string]string                                `json:"tags_to_add"`
	CopyTags   bool                                             `json:"copy_tags"`
}

type AwsDlmLifecyclePolicySpecPolicyDetails struct {
	Schedule      []AwsDlmLifecyclePolicySpecPolicyDetails `json:"schedule"`
	TargetTags    map[string]string                        `json:"target_tags"`
	ResourceTypes []string                                 `json:"resource_types"`
}

type AwsDlmLifecyclePolicySpec struct {
	Description      string                      `json:"description"`
	ExecutionRoleArn string                      `json:"execution_role_arn"`
	PolicyDetails    []AwsDlmLifecyclePolicySpec `json:"policy_details"`
	State            string                      `json:"state"`
}

type AwsDlmLifecyclePolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDlmLifecyclePolicyList is a list of AwsDlmLifecyclePolicys
type AwsDlmLifecyclePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDlmLifecyclePolicy CRD objects
	Items []AwsDlmLifecyclePolicy `json:"items,omitempty"`
}
