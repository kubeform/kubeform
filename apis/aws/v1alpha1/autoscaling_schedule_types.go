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

type AutoscalingSchedule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingScheduleSpec   `json:"spec,omitempty"`
	Status            AutoscalingScheduleStatus `json:"status,omitempty"`
}

type AutoscalingScheduleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Arn                  string `json:"arn,omitempty" tf:"arn,omitempty"`
	AutoscalingGroupName string `json:"autoscalingGroupName" tf:"autoscaling_group_name"`
	// +optional
	DesiredCapacity int `json:"desiredCapacity,omitempty" tf:"desired_capacity,omitempty"`
	// +optional
	EndTime string `json:"endTime,omitempty" tf:"end_time,omitempty"`
	// +optional
	MaxSize int `json:"maxSize,omitempty" tf:"max_size,omitempty"`
	// +optional
	MinSize int `json:"minSize,omitempty" tf:"min_size,omitempty"`
	// +optional
	Recurrence          string `json:"recurrence,omitempty" tf:"recurrence,omitempty"`
	ScheduledActionName string `json:"scheduledActionName" tf:"scheduled_action_name"`
	// +optional
	StartTime string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type AutoscalingScheduleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AutoscalingScheduleSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutoscalingScheduleList is a list of AutoscalingSchedules
type AutoscalingScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutoscalingSchedule CRD objects
	Items []AutoscalingSchedule `json:"items,omitempty"`
}
