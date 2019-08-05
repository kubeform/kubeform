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

type AppautoscalingPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppautoscalingPolicySpec   `json:"spec,omitempty"`
	Status            AppautoscalingPolicyStatus `json:"status,omitempty"`
}

type AppautoscalingPolicySpecStepScalingPolicyConfigurationStepAdjustment struct {
	// +optional
	MetricIntervalLowerBound string `json:"metricIntervalLowerBound,omitempty" tf:"metric_interval_lower_bound,omitempty"`
	// +optional
	MetricIntervalUpperBound string `json:"metricIntervalUpperBound,omitempty" tf:"metric_interval_upper_bound,omitempty"`
	ScalingAdjustment        int    `json:"scalingAdjustment" tf:"scaling_adjustment"`
}

type AppautoscalingPolicySpecStepScalingPolicyConfiguration struct {
	// +optional
	AdjustmentType string `json:"adjustmentType,omitempty" tf:"adjustment_type,omitempty"`
	// +optional
	Cooldown int `json:"cooldown,omitempty" tf:"cooldown,omitempty"`
	// +optional
	MetricAggregationType string `json:"metricAggregationType,omitempty" tf:"metric_aggregation_type,omitempty"`
	// +optional
	MinAdjustmentMagnitude int `json:"minAdjustmentMagnitude,omitempty" tf:"min_adjustment_magnitude,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	StepAdjustment []AppautoscalingPolicySpecStepScalingPolicyConfigurationStepAdjustment `json:"stepAdjustment,omitempty" tf:"step_adjustment,omitempty"`
}

type AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensions struct {
	Name  string `json:"name" tf:"name"`
	Value string `json:"value" tf:"value"`
}

type AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Dimensions []AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensions `json:"dimensions,omitempty" tf:"dimensions,omitempty"`
	MetricName string                                                                                                    `json:"metricName" tf:"metric_name"`
	Namespace  string                                                                                                    `json:"namespace" tf:"namespace"`
	Statistic  string                                                                                                    `json:"statistic" tf:"statistic"`
	// +optional
	Unit string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification struct {
	PredefinedMetricType string `json:"predefinedMetricType" tf:"predefined_metric_type"`
	// +optional
	ResourceLabel string `json:"resourceLabel,omitempty" tf:"resource_label,omitempty"`
}

type AppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CustomizedMetricSpecification []AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification `json:"customizedMetricSpecification,omitempty" tf:"customized_metric_specification,omitempty"`
	// +optional
	DisableScaleIn bool `json:"disableScaleIn,omitempty" tf:"disable_scale_in,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PredefinedMetricSpecification []AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification `json:"predefinedMetricSpecification,omitempty" tf:"predefined_metric_specification,omitempty"`
	// +optional
	ScaleInCooldown int `json:"scaleInCooldown,omitempty" tf:"scale_in_cooldown,omitempty"`
	// +optional
	ScaleOutCooldown int         `json:"scaleOutCooldown,omitempty" tf:"scale_out_cooldown,omitempty"`
	TargetValue      json.Number `json:"targetValue" tf:"target_value"`
}

type AppautoscalingPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Alarms []string `json:"alarms,omitempty" tf:"alarms,omitempty"`
	Arn    string   `json:"arn" tf:"arn"`
	Name   string   `json:"name" tf:"name"`
	// +optional
	PolicyType        string `json:"policyType,omitempty" tf:"policy_type,omitempty"`
	ResourceID        string `json:"resourceID" tf:"resource_id"`
	ScalableDimension string `json:"scalableDimension" tf:"scalable_dimension"`
	ServiceNamespace  string `json:"serviceNamespace" tf:"service_namespace"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StepScalingPolicyConfiguration []AppautoscalingPolicySpecStepScalingPolicyConfiguration `json:"stepScalingPolicyConfiguration,omitempty" tf:"step_scaling_policy_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TargetTrackingScalingPolicyConfiguration []AppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration `json:"targetTrackingScalingPolicyConfiguration,omitempty" tf:"target_tracking_scaling_policy_configuration,omitempty"`
}

type AppautoscalingPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppautoscalingPolicySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
