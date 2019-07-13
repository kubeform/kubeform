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

type AwsAutoscalingSchedule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAutoscalingScheduleSpec   `json:"spec,omitempty"`
	Status            AwsAutoscalingScheduleStatus `json:"status,omitempty"`
}

type AwsAutoscalingScheduleSpec struct {
	Arn                  string `json:"arn"`
	ScheduledActionName  string `json:"scheduled_action_name"`
	AutoscalingGroupName string `json:"autoscaling_group_name"`
	StartTime            string `json:"start_time"`
	MinSize              int    `json:"min_size"`
	DesiredCapacity      int    `json:"desired_capacity"`
	EndTime              string `json:"end_time"`
	Recurrence           string `json:"recurrence"`
	MaxSize              int    `json:"max_size"`
}



type AwsAutoscalingScheduleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAutoscalingScheduleList is a list of AwsAutoscalingSchedules
type AwsAutoscalingScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAutoscalingSchedule CRD objects
	Items []AwsAutoscalingSchedule `json:"items,omitempty"`
}