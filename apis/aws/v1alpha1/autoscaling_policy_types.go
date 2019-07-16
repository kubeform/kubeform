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

type AutoscalingPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingPolicySpec   `json:"spec,omitempty"`
	Status            AutoscalingPolicyStatus `json:"status,omitempty"`
}

type AutoscalingPolicySpecStepAdjustment struct {
	// +optional
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound,omitempty"`
	// +optional
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound,omitempty"`
	ScalingAdjustment        int    `json:"scaling_adjustment"`
}

type AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification struct {
	// +optional
	MetricDimension *[]AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification `json:"metric_dimension,omitempty"`
	MetricName      string                                                                           `json:"metric_name"`
	Namespace       string                                                                           `json:"namespace"`
	Statistic       string                                                                           `json:"statistic"`
	// +optional
	Unit string `json:"unit,omitempty"`
}

type AutoscalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification struct {
	PredefinedMetricType string `json:"predefined_metric_type"`
	// +optional
	ResourceLabel string `json:"resource_label,omitempty"`
}

type AutoscalingPolicySpecTargetTrackingConfiguration struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CustomizedMetricSpecification *[]AutoscalingPolicySpecTargetTrackingConfiguration `json:"customized_metric_specification,omitempty"`
	// +optional
	DisableScaleIn bool `json:"disable_scale_in,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PredefinedMetricSpecification *[]AutoscalingPolicySpecTargetTrackingConfiguration `json:"predefined_metric_specification,omitempty"`
	TargetValue                   json.Number                                         `json:"target_value"`
}

type AutoscalingPolicySpec struct {
	// +optional
	AdjustmentType       string `json:"adjustment_type,omitempty"`
	AutoscalingGroupName string `json:"autoscaling_group_name"`
	// +optional
	Cooldown int `json:"cooldown,omitempty"`
	// +optional
	EstimatedInstanceWarmup int `json:"estimated_instance_warmup,omitempty"`
	// +optional
	MinAdjustmentMagnitude int    `json:"min_adjustment_magnitude,omitempty"`
	Name                   string `json:"name"`
	// +optional
	PolicyType string `json:"policy_type,omitempty"`
	// +optional
	ScalingAdjustment int `json:"scaling_adjustment,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	StepAdjustment *[]AutoscalingPolicySpec `json:"step_adjustment,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TargetTrackingConfiguration *[]AutoscalingPolicySpec `json:"target_tracking_configuration,omitempty"`
}

type AutoscalingPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutoscalingPolicyList is a list of AutoscalingPolicys
type AutoscalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutoscalingPolicy CRD objects
	Items []AutoscalingPolicy `json:"items,omitempty"`
}
