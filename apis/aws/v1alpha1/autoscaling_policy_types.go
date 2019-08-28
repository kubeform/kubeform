package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	MetricIntervalLowerBound string `json:"metricIntervalLowerBound,omitempty" tf:"metric_interval_lower_bound,omitempty"`
	// +optional
	MetricIntervalUpperBound string `json:"metricIntervalUpperBound,omitempty" tf:"metric_interval_upper_bound,omitempty"`
	ScalingAdjustment        int    `json:"scalingAdjustment" tf:"scaling_adjustment"`
}

type AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension struct {
	Name  string `json:"name" tf:"name"`
	Value string `json:"value" tf:"value"`
}

type AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification struct {
	// +optional
	MetricDimension []AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension `json:"metricDimension,omitempty" tf:"metric_dimension,omitempty"`
	MetricName      string                                                                                         `json:"metricName" tf:"metric_name"`
	Namespace       string                                                                                         `json:"namespace" tf:"namespace"`
	Statistic       string                                                                                         `json:"statistic" tf:"statistic"`
	// +optional
	Unit string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type AutoscalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification struct {
	PredefinedMetricType string `json:"predefinedMetricType" tf:"predefined_metric_type"`
	// +optional
	ResourceLabel string `json:"resourceLabel,omitempty" tf:"resource_label,omitempty"`
}

type AutoscalingPolicySpecTargetTrackingConfiguration struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CustomizedMetricSpecification []AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification `json:"customizedMetricSpecification,omitempty" tf:"customized_metric_specification,omitempty"`
	// +optional
	DisableScaleIn bool `json:"disableScaleIn,omitempty" tf:"disable_scale_in,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PredefinedMetricSpecification []AutoscalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification `json:"predefinedMetricSpecification,omitempty" tf:"predefined_metric_specification,omitempty"`
	TargetValue                   json.Number                                                                     `json:"targetValue" tf:"target_value"`
}

type AutoscalingPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdjustmentType string `json:"adjustmentType,omitempty" tf:"adjustment_type,omitempty"`
	// +optional
	Arn                  string `json:"arn,omitempty" tf:"arn,omitempty"`
	AutoscalingGroupName string `json:"autoscalingGroupName" tf:"autoscaling_group_name"`
	// +optional
	Cooldown int `json:"cooldown,omitempty" tf:"cooldown,omitempty"`
	// +optional
	EstimatedInstanceWarmup int `json:"estimatedInstanceWarmup,omitempty" tf:"estimated_instance_warmup,omitempty"`
	// +optional
	MetricAggregationType string `json:"metricAggregationType,omitempty" tf:"metric_aggregation_type,omitempty"`
	// +optional
	MinAdjustmentMagnitude int    `json:"minAdjustmentMagnitude,omitempty" tf:"min_adjustment_magnitude,omitempty"`
	Name                   string `json:"name" tf:"name"`
	// +optional
	PolicyType string `json:"policyType,omitempty" tf:"policy_type,omitempty"`
	// +optional
	ScalingAdjustment int `json:"scalingAdjustment,omitempty" tf:"scaling_adjustment,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	StepAdjustment []AutoscalingPolicySpecStepAdjustment `json:"stepAdjustment,omitempty" tf:"step_adjustment,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TargetTrackingConfiguration []AutoscalingPolicySpecTargetTrackingConfiguration `json:"targetTrackingConfiguration,omitempty" tf:"target_tracking_configuration,omitempty"`
}



type AutoscalingPolicyStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *AutoscalingPolicySpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
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