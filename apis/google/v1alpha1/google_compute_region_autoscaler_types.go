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

type GoogleComputeRegionAutoscaler struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeRegionAutoscalerSpec   `json:"spec,omitempty"`
	Status            GoogleComputeRegionAutoscalerStatus `json:"status,omitempty"`
}

type GoogleComputeRegionAutoscalerSpecAutoscalingPolicyCpuUtilization struct {
	Target float64 `json:"target"`
}

type GoogleComputeRegionAutoscalerSpecAutoscalingPolicyLoadBalancingUtilization struct {
	Target float64 `json:"target"`
}

type GoogleComputeRegionAutoscalerSpecAutoscalingPolicyMetric struct {
	Name   string  `json:"name"`
	Target float64 `json:"target"`
	Type   string  `json:"type"`
}

type GoogleComputeRegionAutoscalerSpecAutoscalingPolicy struct {
	CpuUtilization           []GoogleComputeRegionAutoscalerSpecAutoscalingPolicy `json:"cpu_utilization"`
	LoadBalancingUtilization []GoogleComputeRegionAutoscalerSpecAutoscalingPolicy `json:"load_balancing_utilization"`
	Metric                   []GoogleComputeRegionAutoscalerSpecAutoscalingPolicy `json:"metric"`
	MaxReplicas              int                                                  `json:"max_replicas"`
	MinReplicas              int                                                  `json:"min_replicas"`
	CooldownPeriod           int                                                  `json:"cooldown_period"`
}

type GoogleComputeRegionAutoscalerSpec struct {
	Description       string                              `json:"description"`
	Region            string                              `json:"region"`
	CreationTimestamp string                              `json:"creation_timestamp"`
	Project           string                              `json:"project"`
	SelfLink          string                              `json:"self_link"`
	AutoscalingPolicy []GoogleComputeRegionAutoscalerSpec `json:"autoscaling_policy"`
	Name              string                              `json:"name"`
	Target            string                              `json:"target"`
}



type GoogleComputeRegionAutoscalerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeRegionAutoscalerList is a list of GoogleComputeRegionAutoscalers
type GoogleComputeRegionAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeRegionAutoscaler CRD objects
	Items []GoogleComputeRegionAutoscaler `json:"items,omitempty"`
}