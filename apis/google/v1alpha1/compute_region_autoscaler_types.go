package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeRegionAutoscaler struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeRegionAutoscalerSpec   `json:"spec,omitempty"`
	Status            ComputeRegionAutoscalerStatus `json:"status,omitempty"`
}

type ComputeRegionAutoscalerSpecAutoscalingPolicyCpuUtilization struct {
	Target json.Number `json:"target" tf:"target"`
}

type ComputeRegionAutoscalerSpecAutoscalingPolicyLoadBalancingUtilization struct {
	Target json.Number `json:"target" tf:"target"`
}

type ComputeRegionAutoscalerSpecAutoscalingPolicyMetric struct {
	Name   string      `json:"name" tf:"name"`
	Target json.Number `json:"target" tf:"target"`
	Type   string      `json:"type" tf:"type"`
}

type ComputeRegionAutoscalerSpecAutoscalingPolicy struct {
	// +optional
	CooldownPeriod int `json:"cooldownPeriod,omitempty" tf:"cooldown_period,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CpuUtilization []ComputeRegionAutoscalerSpecAutoscalingPolicyCpuUtilization `json:"cpuUtilization,omitempty" tf:"cpu_utilization,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoadBalancingUtilization []ComputeRegionAutoscalerSpecAutoscalingPolicyLoadBalancingUtilization `json:"loadBalancingUtilization,omitempty" tf:"load_balancing_utilization,omitempty"`
	MaxReplicas              int                                                                    `json:"maxReplicas" tf:"max_replicas"`
	// +optional
	Metric      []ComputeRegionAutoscalerSpecAutoscalingPolicyMetric `json:"metric,omitempty" tf:"metric,omitempty"`
	MinReplicas int                                                  `json:"minReplicas" tf:"min_replicas"`
}

type ComputeRegionAutoscalerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MaxItems=1
	AutoscalingPolicy []ComputeRegionAutoscalerSpecAutoscalingPolicy `json:"autoscalingPolicy" tf:"autoscaling_policy"`
	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	Target   string `json:"target" tf:"target"`
}

type ComputeRegionAutoscalerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeRegionAutoscalerSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeRegionAutoscalerList is a list of ComputeRegionAutoscalers
type ComputeRegionAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeRegionAutoscaler CRD objects
	Items []ComputeRegionAutoscaler `json:"items,omitempty"`
}
