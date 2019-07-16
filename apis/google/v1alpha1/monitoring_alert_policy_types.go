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

type MonitoringAlertPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitoringAlertPolicySpec   `json:"spec,omitempty"`
	Status            MonitoringAlertPolicyStatus `json:"status,omitempty"`
}

type MonitoringAlertPolicySpecConditionsConditionAbsentAggregations struct {
	// +optional
	AlignmentPeriod string `json:"alignment_period,omitempty"`
	// +optional
	CrossSeriesReducer string `json:"cross_series_reducer,omitempty"`
	// +optional
	GroupByFields []string `json:"group_by_fields,omitempty"`
	// +optional
	PerSeriesAligner string `json:"per_series_aligner,omitempty"`
}

type MonitoringAlertPolicySpecConditionsConditionAbsentTrigger struct {
	// +optional
	Count int `json:"count,omitempty"`
	// +optional
	Percent json.Number `json:"percent,omitempty"`
}

type MonitoringAlertPolicySpecConditionsConditionAbsent struct {
	// +optional
	Aggregations *[]MonitoringAlertPolicySpecConditionsConditionAbsent `json:"aggregations,omitempty"`
	Duration     string                                                `json:"duration"`
	// +optional
	Filter string `json:"filter,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Trigger *[]MonitoringAlertPolicySpecConditionsConditionAbsent `json:"trigger,omitempty"`
}

type MonitoringAlertPolicySpecConditionsConditionThresholdAggregations struct {
	// +optional
	AlignmentPeriod string `json:"alignment_period,omitempty"`
	// +optional
	CrossSeriesReducer string `json:"cross_series_reducer,omitempty"`
	// +optional
	GroupByFields []string `json:"group_by_fields,omitempty"`
	// +optional
	PerSeriesAligner string `json:"per_series_aligner,omitempty"`
}

type MonitoringAlertPolicySpecConditionsConditionThresholdDenominatorAggregations struct {
	// +optional
	AlignmentPeriod string `json:"alignment_period,omitempty"`
	// +optional
	CrossSeriesReducer string `json:"cross_series_reducer,omitempty"`
	// +optional
	GroupByFields []string `json:"group_by_fields,omitempty"`
	// +optional
	PerSeriesAligner string `json:"per_series_aligner,omitempty"`
}

type MonitoringAlertPolicySpecConditionsConditionThresholdTrigger struct {
	// +optional
	Count int `json:"count,omitempty"`
	// +optional
	Percent json.Number `json:"percent,omitempty"`
}

type MonitoringAlertPolicySpecConditionsConditionThreshold struct {
	// +optional
	Aggregations *[]MonitoringAlertPolicySpecConditionsConditionThreshold `json:"aggregations,omitempty"`
	Comparison   string                                                   `json:"comparison"`
	// +optional
	DenominatorAggregations *[]MonitoringAlertPolicySpecConditionsConditionThreshold `json:"denominator_aggregations,omitempty"`
	// +optional
	DenominatorFilter string `json:"denominator_filter,omitempty"`
	Duration          string `json:"duration"`
	// +optional
	Filter string `json:"filter,omitempty"`
	// +optional
	ThresholdValue json.Number `json:"threshold_value,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Trigger *[]MonitoringAlertPolicySpecConditionsConditionThreshold `json:"trigger,omitempty"`
}

type MonitoringAlertPolicySpecConditions struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ConditionAbsent *[]MonitoringAlertPolicySpecConditions `json:"condition_absent,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ConditionThreshold *[]MonitoringAlertPolicySpecConditions `json:"condition_threshold,omitempty"`
	DisplayName        string                                 `json:"display_name"`
}

type MonitoringAlertPolicySpec struct {
	Combiner    string                      `json:"combiner"`
	Conditions  []MonitoringAlertPolicySpec `json:"conditions"`
	DisplayName string                      `json:"display_name"`
	Enabled     bool                        `json:"enabled"`
	// +optional
	Labels []string `json:"labels,omitempty"`
	// +optional
	NotificationChannels []string `json:"notification_channels,omitempty"`
}

type MonitoringAlertPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitoringAlertPolicyList is a list of MonitoringAlertPolicys
type MonitoringAlertPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitoringAlertPolicy CRD objects
	Items []MonitoringAlertPolicy `json:"items,omitempty"`
}
