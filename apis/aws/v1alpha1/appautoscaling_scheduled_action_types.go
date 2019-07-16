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

type AppautoscalingScheduledAction struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppautoscalingScheduledActionSpec   `json:"spec,omitempty"`
	Status            AppautoscalingScheduledActionStatus `json:"status,omitempty"`
}

type AppautoscalingScheduledActionSpecScalableTargetAction struct {
	// +optional
	MaxCapacity int `json:"max_capacity,omitempty"`
	// +optional
	MinCapacity int `json:"min_capacity,omitempty"`
}

type AppautoscalingScheduledActionSpec struct {
	// +optional
	EndTime    string `json:"end_time,omitempty"`
	Name       string `json:"name"`
	ResourceId string `json:"resource_id"`
	// +optional
	ScalableDimension string `json:"scalable_dimension,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ScalableTargetAction *[]AppautoscalingScheduledActionSpec `json:"scalable_target_action,omitempty"`
	// +optional
	Schedule         string `json:"schedule,omitempty"`
	ServiceNamespace string `json:"service_namespace"`
	// +optional
	StartTime string `json:"start_time,omitempty"`
}

type AppautoscalingScheduledActionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppautoscalingScheduledActionList is a list of AppautoscalingScheduledActions
type AppautoscalingScheduledActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppautoscalingScheduledAction CRD objects
	Items []AppautoscalingScheduledAction `json:"items,omitempty"`
}
