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

type GoogleComputeAutoscaler struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeAutoscalerSpec   `json:"spec,omitempty"`
	Status            GoogleComputeAutoscalerStatus `json:"status,omitempty"`
}

type GoogleComputeAutoscalerSpecAutoscalingPolicyCpuUtilization struct {
	Target float64 `json:"target"`
}

type GoogleComputeAutoscalerSpecAutoscalingPolicyLoadBalancingUtilization struct {
	Target float64 `json:"target"`
}

type GoogleComputeAutoscalerSpecAutoscalingPolicyMetric struct {
	Name   string  `json:"name"`
	Target float64 `json:"target"`
	Type   string  `json:"type"`
}

type GoogleComputeAutoscalerSpecAutoscalingPolicy struct {
	CooldownPeriod           int                                            `json:"cooldown_period"`
	CpuUtilization           []GoogleComputeAutoscalerSpecAutoscalingPolicy `json:"cpu_utilization"`
	LoadBalancingUtilization []GoogleComputeAutoscalerSpecAutoscalingPolicy `json:"load_balancing_utilization"`
	Metric                   []GoogleComputeAutoscalerSpecAutoscalingPolicy `json:"metric"`
	MaxReplicas              int                                            `json:"max_replicas"`
	MinReplicas              int                                            `json:"min_replicas"`
}

type GoogleComputeAutoscalerSpec struct {
	Zone              string                        `json:"zone"`
	CreationTimestamp string                        `json:"creation_timestamp"`
	Project           string                        `json:"project"`
	SelfLink          string                        `json:"self_link"`
	AutoscalingPolicy []GoogleComputeAutoscalerSpec `json:"autoscaling_policy"`
	Name              string                        `json:"name"`
	Target            string                        `json:"target"`
	Description       string                        `json:"description"`
}

type GoogleComputeAutoscalerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeAutoscalerList is a list of GoogleComputeAutoscalers
type GoogleComputeAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeAutoscaler CRD objects
	Items []GoogleComputeAutoscaler `json:"items,omitempty"`
}
