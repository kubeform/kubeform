package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingSchedule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAutoscalingScheduleSpec   `json:"spec,omitempty"`
	Status            AwsAutoscalingScheduleStatus `json:"status,omitempty"`
}

type AwsAutoscalingScheduleSpec struct {
	MinSize              int    `json:"min_size"`
	DesiredCapacity      int    `json:"desired_capacity"`
	StartTime            string `json:"start_time"`
	ScheduledActionName  string `json:"scheduled_action_name"`
	AutoscalingGroupName string `json:"autoscaling_group_name"`
	EndTime              string `json:"end_time"`
	Recurrence           string `json:"recurrence"`
	MaxSize              int    `json:"max_size"`
	Arn                  string `json:"arn"`
}



type AwsAutoscalingScheduleStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingScheduleList is a list of AwsAutoscalingSchedules
type AwsAutoscalingScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAutoscalingSchedule CRD objects
	Items []AwsAutoscalingSchedule `json:"items,omitempty"`
}