package v1alpha1

import (
	"encoding/json"

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

type ComputeAutoscaler struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeAutoscalerSpec   `json:"spec,omitempty"`
	Status            ComputeAutoscalerStatus `json:"status,omitempty"`
}

type ComputeAutoscalerSpecAutoscalingPolicyCpuUtilization struct {
	Target json.Number `json:"target" tf:"target"`
}

type ComputeAutoscalerSpecAutoscalingPolicyLoadBalancingUtilization struct {
	Target json.Number `json:"target" tf:"target"`
}

type ComputeAutoscalerSpecAutoscalingPolicyMetric struct {
	Name   string      `json:"name" tf:"name"`
	Target json.Number `json:"target" tf:"target"`
	Type   string      `json:"type" tf:"type"`
}

type ComputeAutoscalerSpecAutoscalingPolicy struct {
	// +optional
	CooldownPeriod int `json:"cooldownPeriod,omitempty" tf:"cooldown_period,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CpuUtilization []ComputeAutoscalerSpecAutoscalingPolicyCpuUtilization `json:"cpuUtilization,omitempty" tf:"cpu_utilization,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoadBalancingUtilization []ComputeAutoscalerSpecAutoscalingPolicyLoadBalancingUtilization `json:"loadBalancingUtilization,omitempty" tf:"load_balancing_utilization,omitempty"`
	MaxReplicas              int                                                              `json:"maxReplicas" tf:"max_replicas"`
	// +optional
	Metric      []ComputeAutoscalerSpecAutoscalingPolicyMetric `json:"metric,omitempty" tf:"metric,omitempty"`
	MinReplicas int                                            `json:"minReplicas" tf:"min_replicas"`
}

type ComputeAutoscalerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MaxItems=1
	AutoscalingPolicy []ComputeAutoscalerSpecAutoscalingPolicy `json:"autoscalingPolicy" tf:"autoscaling_policy"`
	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	Target   string `json:"target" tf:"target"`
	// +optional
	Zone string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ComputeAutoscalerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeAutoscalerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
