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

type AwsAppautoscalingScheduledAction struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppautoscalingScheduledActionSpec   `json:"spec,omitempty"`
	Status            AwsAppautoscalingScheduledActionStatus `json:"status,omitempty"`
}

type AwsAppautoscalingScheduledActionSpecScalableTargetAction struct {
	MaxCapacity int `json:"max_capacity"`
	MinCapacity int `json:"min_capacity"`
}

type AwsAppautoscalingScheduledActionSpec struct {
	ScalableTargetAction []AwsAppautoscalingScheduledActionSpec `json:"scalable_target_action"`
	Schedule             string                                 `json:"schedule"`
	ScalableDimension    string                                 `json:"scalable_dimension"`
	ServiceNamespace     string                                 `json:"service_namespace"`
	ResourceId           string                                 `json:"resource_id"`
	StartTime            string                                 `json:"start_time"`
	EndTime              string                                 `json:"end_time"`
	Arn                  string                                 `json:"arn"`
	Name                 string                                 `json:"name"`
}



type AwsAppautoscalingScheduledActionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAppautoscalingScheduledActionList is a list of AwsAppautoscalingScheduledActions
type AwsAppautoscalingScheduledActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppautoscalingScheduledAction CRD objects
	Items []AwsAppautoscalingScheduledAction `json:"items,omitempty"`
}