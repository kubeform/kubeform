package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	MaxCapacity int `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`
	// +optional
	MinCapacity int `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`
}

type AppautoscalingScheduledActionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	EndTime    string `json:"endTime,omitempty" tf:"end_time,omitempty"`
	Name       string `json:"name" tf:"name"`
	ResourceID string `json:"resourceID" tf:"resource_id"`
	// +optional
	ScalableDimension string `json:"scalableDimension,omitempty" tf:"scalable_dimension,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ScalableTargetAction []AppautoscalingScheduledActionSpecScalableTargetAction `json:"scalableTargetAction,omitempty" tf:"scalable_target_action,omitempty"`
	// +optional
	Schedule         string `json:"schedule,omitempty" tf:"schedule,omitempty"`
	ServiceNamespace string `json:"serviceNamespace" tf:"service_namespace"`
	// +optional
	StartTime string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type AppautoscalingScheduledActionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *AppautoscalingScheduledActionSpec `json:"state,omitempty"`
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
