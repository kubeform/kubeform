package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type AppautoscalingTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppautoscalingTargetSpec   `json:"spec,omitempty"`
	Status            AppautoscalingTargetStatus `json:"status,omitempty"`
}

type AppautoscalingTargetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	MaxCapacity int    `json:"maxCapacity" tf:"max_capacity"`
	MinCapacity int    `json:"minCapacity" tf:"min_capacity"`
	ResourceID  string `json:"resourceID" tf:"resource_id"`
	// +optional
	RoleArn           string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
	ScalableDimension string `json:"scalableDimension" tf:"scalable_dimension"`
	ServiceNamespace  string `json:"serviceNamespace" tf:"service_namespace"`
}

type AppautoscalingTargetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppautoscalingTargetSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppautoscalingTargetList is a list of AppautoscalingTargets
type AppautoscalingTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppautoscalingTarget CRD objects
	Items []AppautoscalingTarget `json:"items,omitempty"`
}