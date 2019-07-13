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

type AwsDlmLifecyclePolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDlmLifecyclePolicySpec   `json:"spec,omitempty"`
	Status            AwsDlmLifecyclePolicyStatus `json:"status,omitempty"`
}

type AwsDlmLifecyclePolicySpecPolicyDetailsScheduleRetainRule struct {
	Count int `json:"count"`
}

type AwsDlmLifecyclePolicySpecPolicyDetailsScheduleCreateRule struct {
	Interval     int      `json:"interval"`
	IntervalUnit string   `json:"interval_unit"`
	Times        []string `json:"times"`
}

type AwsDlmLifecyclePolicySpecPolicyDetailsSchedule struct {
	Name       string                                           `json:"name"`
	RetainRule []AwsDlmLifecyclePolicySpecPolicyDetailsSchedule `json:"retain_rule"`
	TagsToAdd  map[string]string                                `json:"tags_to_add"`
	CopyTags   bool                                             `json:"copy_tags"`
	CreateRule []AwsDlmLifecyclePolicySpecPolicyDetailsSchedule `json:"create_rule"`
}

type AwsDlmLifecyclePolicySpecPolicyDetails struct {
	ResourceTypes []string                                 `json:"resource_types"`
	Schedule      []AwsDlmLifecyclePolicySpecPolicyDetails `json:"schedule"`
	TargetTags    map[string]string                        `json:"target_tags"`
}

type AwsDlmLifecyclePolicySpec struct {
	Description      string                      `json:"description"`
	ExecutionRoleArn string                      `json:"execution_role_arn"`
	PolicyDetails    []AwsDlmLifecyclePolicySpec `json:"policy_details"`
	State            string                      `json:"state"`
}



type AwsDlmLifecyclePolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDlmLifecyclePolicyList is a list of AwsDlmLifecyclePolicys
type AwsDlmLifecyclePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDlmLifecyclePolicy CRD objects
	Items []AwsDlmLifecyclePolicy `json:"items,omitempty"`
}