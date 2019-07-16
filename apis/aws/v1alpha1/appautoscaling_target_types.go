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

type AppautoscalingTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppautoscalingTargetSpec   `json:"spec,omitempty"`
	Status            AppautoscalingTargetStatus `json:"status,omitempty"`
}

type AppautoscalingTargetSpec struct {
	MaxCapacity       int    `json:"max_capacity"`
	MinCapacity       int    `json:"min_capacity"`
	ResourceId        string `json:"resource_id"`
	ScalableDimension string `json:"scalable_dimension"`
	ServiceNamespace  string `json:"service_namespace"`
}

type AppautoscalingTargetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
