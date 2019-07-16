package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeAutoscaler struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeAutoscalerSpec   `json:"spec,omitempty"`
	Status            ComputeAutoscalerStatus `json:"status,omitempty"`
}

type ComputeAutoscalerSpecAutoscalingPolicyLoadBalancingUtilization struct {
	Target json.Number `json:"target"`
}

type ComputeAutoscalerSpecAutoscalingPolicyMetric struct {
	Name   string      `json:"name"`
	Target json.Number `json:"target"`
	Type   string      `json:"type"`
}

type ComputeAutoscalerSpecAutoscalingPolicy struct {
	// +optional
	CooldownPeriod int `json:"cooldown_period,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoadBalancingUtilization *[]ComputeAutoscalerSpecAutoscalingPolicy `json:"load_balancing_utilization,omitempty"`
	MaxReplicas              int                                       `json:"max_replicas"`
	// +optional
	Metric      *[]ComputeAutoscalerSpecAutoscalingPolicy `json:"metric,omitempty"`
	MinReplicas int                                       `json:"min_replicas"`
}

type ComputeAutoscalerSpec struct {
	// +kubebuilder:validation:MaxItems=1
	AutoscalingPolicy []ComputeAutoscalerSpec `json:"autoscaling_policy"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	Target      string `json:"target"`
}

type ComputeAutoscalerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeAutoscalerList is a list of ComputeAutoscalers
type ComputeAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeAutoscaler CRD objects
	Items []ComputeAutoscaler `json:"items,omitempty"`
}
