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

type AppautoscalingPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppautoscalingPolicySpec   `json:"spec,omitempty"`
	Status            AppautoscalingPolicyStatus `json:"status,omitempty"`
}

type AppautoscalingPolicySpecStepScalingPolicyConfigurationStepAdjustment struct {
	// +optional
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound,omitempty"`
	// +optional
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound,omitempty"`
	ScalingAdjustment        int    `json:"scaling_adjustment"`
}

type AppautoscalingPolicySpecStepScalingPolicyConfiguration struct {
	// +optional
	AdjustmentType string `json:"adjustment_type,omitempty"`
	// +optional
	Cooldown int `json:"cooldown,omitempty"`
	// +optional
	MetricAggregationType string `json:"metric_aggregation_type,omitempty"`
	// +optional
	MinAdjustmentMagnitude int `json:"min_adjustment_magnitude,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	StepAdjustment *[]AppautoscalingPolicySpecStepScalingPolicyConfiguration `json:"step_adjustment,omitempty"`
}

type AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensions struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Dimensions *[]AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification `json:"dimensions,omitempty"`
	MetricName string                                                                                           `json:"metric_name"`
	Namespace  string                                                                                           `json:"namespace"`
	Statistic  string                                                                                           `json:"statistic"`
	// +optional
	Unit string `json:"unit,omitempty"`
}

type AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification struct {
	PredefinedMetricType string `json:"predefined_metric_type"`
	// +optional
	ResourceLabel string `json:"resource_label,omitempty"`
}

type AppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CustomizedMetricSpecification *[]AppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration `json:"customized_metric_specification,omitempty"`
	// +optional
	DisableScaleIn bool `json:"disable_scale_in,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PredefinedMetricSpecification *[]AppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration `json:"predefined_metric_specification,omitempty"`
	// +optional
	ScaleInCooldown int `json:"scale_in_cooldown,omitempty"`
	// +optional
	ScaleOutCooldown int         `json:"scale_out_cooldown,omitempty"`
	TargetValue      json.Number `json:"target_value"`
}

type AppautoscalingPolicySpec struct {
	// +optional
	Alarms []string `json:"alarms,omitempty"`
	Name   string   `json:"name"`
	// +optional
	PolicyType        string `json:"policy_type,omitempty"`
	ResourceId        string `json:"resource_id"`
	ScalableDimension string `json:"scalable_dimension"`
	ServiceNamespace  string `json:"service_namespace"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StepScalingPolicyConfiguration *[]AppautoscalingPolicySpec `json:"step_scaling_policy_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TargetTrackingScalingPolicyConfiguration *[]AppautoscalingPolicySpec `json:"target_tracking_scaling_policy_configuration,omitempty"`
}

type AppautoscalingPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppautoscalingPolicyList is a list of AppautoscalingPolicys
type AppautoscalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppautoscalingPolicy CRD objects
	Items []AppautoscalingPolicy `json:"items,omitempty"`
}
