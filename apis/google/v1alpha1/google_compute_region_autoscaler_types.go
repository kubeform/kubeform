package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	MaxReplicas              int                                                  `json:"max_replicas"`
	MinReplicas              int                                                  `json:"min_replicas"`
	CooldownPeriod           int                                                  `json:"cooldown_period"`
	CpuUtilization           []GoogleComputeRegionAutoscalerSpecAutoscalingPolicy `json:"cpu_utilization"`
	LoadBalancingUtilization []GoogleComputeRegionAutoscalerSpecAutoscalingPolicy `json:"load_balancing_utilization"`
	Metric                   []GoogleComputeRegionAutoscalerSpecAutoscalingPolicy `json:"metric"`
}

type GoogleComputeRegionAutoscalerSpec struct {
	SelfLink          string                              `json:"self_link"`
	AutoscalingPolicy []GoogleComputeRegionAutoscalerSpec `json:"autoscaling_policy"`
	Name              string                              `json:"name"`
	Target            string                              `json:"target"`
	Description       string                              `json:"description"`
	Region            string                              `json:"region"`
	CreationTimestamp string                              `json:"creation_timestamp"`
	Project           string                              `json:"project"`
}



type GoogleComputeRegionAutoscalerStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeRegionAutoscalerList is a list of GoogleComputeRegionAutoscalers
type GoogleComputeRegionAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeRegionAutoscaler CRD objects
	Items []GoogleComputeRegionAutoscaler `json:"items,omitempty"`
}