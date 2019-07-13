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

type AwsCloudwatchEventRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudwatchEventRuleSpec   `json:"spec,omitempty"`
	Status            AwsCloudwatchEventRuleStatus `json:"status,omitempty"`
}

type AwsCloudwatchEventRuleSpec struct {
	RoleArn            string            `json:"role_arn"`
	IsEnabled          bool              `json:"is_enabled"`
	NamePrefix         string            `json:"name_prefix"`
	Description        string            `json:"description"`
	EventPattern       string            `json:"event_pattern"`
	Arn                string            `json:"arn"`
	Tags               map[string]string `json:"tags"`
	Name               string            `json:"name"`
	ScheduleExpression string            `json:"schedule_expression"`
}



type AwsCloudwatchEventRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCloudwatchEventRuleList is a list of AwsCloudwatchEventRules
type AwsCloudwatchEventRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudwatchEventRule CRD objects
	Items []AwsCloudwatchEventRule `json:"items,omitempty"`
}