package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleMonitoringAlertPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleMonitoringAlertPolicySpec   `json:"spec,omitempty"`
	Status            GoogleMonitoringAlertPolicyStatus `json:"status,omitempty"`
}

type GoogleMonitoringAlertPolicySpecCreationRecord struct {
	MutateTime string `json:"mutate_time"`
	MutatedBy  string `json:"mutated_by"`
}

type GoogleMonitoringAlertPolicySpecConditionsConditionAbsentAggregations struct {
	AlignmentPeriod    string   `json:"alignment_period"`
	CrossSeriesReducer string   `json:"cross_series_reducer"`
	GroupByFields      []string `json:"group_by_fields"`
	PerSeriesAligner   string   `json:"per_series_aligner"`
}

type GoogleMonitoringAlertPolicySpecConditionsConditionAbsentTrigger struct {
	Count   int     `json:"count"`
	Percent float64 `json:"percent"`
}

type GoogleMonitoringAlertPolicySpecConditionsConditionAbsent struct {
	Duration     string                                                     `json:"duration"`
	Aggregations []GoogleMonitoringAlertPolicySpecConditionsConditionAbsent `json:"aggregations"`
	Filter       string                                                     `json:"filter"`
	Trigger      []GoogleMonitoringAlertPolicySpecConditionsConditionAbsent `json:"trigger"`
}

type GoogleMonitoringAlertPolicySpecConditionsConditionThresholdDenominatorAggregations struct {
	AlignmentPeriod    string   `json:"alignment_period"`
	CrossSeriesReducer string   `json:"cross_series_reducer"`
	GroupByFields      []string `json:"group_by_fields"`
	PerSeriesAligner   string   `json:"per_series_aligner"`
}

type GoogleMonitoringAlertPolicySpecConditionsConditionThresholdTrigger struct {
	Count   int     `json:"count"`
	Percent float64 `json:"percent"`
}

type GoogleMonitoringAlertPolicySpecConditionsConditionThresholdAggregations struct {
	AlignmentPeriod    string   `json:"alignment_period"`
	CrossSeriesReducer string   `json:"cross_series_reducer"`
	GroupByFields      []string `json:"group_by_fields"`
	PerSeriesAligner   string   `json:"per_series_aligner"`
}

type GoogleMonitoringAlertPolicySpecConditionsConditionThreshold struct {
	DenominatorAggregations []GoogleMonitoringAlertPolicySpecConditionsConditionThreshold `json:"denominator_aggregations"`
	DenominatorFilter       string                                                        `json:"denominator_filter"`
	Filter                  string                                                        `json:"filter"`
	ThresholdValue          float64                                                       `json:"threshold_value"`
	Trigger                 []GoogleMonitoringAlertPolicySpecConditionsConditionThreshold `json:"trigger"`
	Comparison              string                                                        `json:"comparison"`
	Duration                string                                                        `json:"duration"`
	Aggregations            []GoogleMonitoringAlertPolicySpecConditionsConditionThreshold `json:"aggregations"`
}

type GoogleMonitoringAlertPolicySpecConditions struct {
	DisplayName        string                                      `json:"display_name"`
	ConditionAbsent    []GoogleMonitoringAlertPolicySpecConditions `json:"condition_absent"`
	ConditionThreshold []GoogleMonitoringAlertPolicySpecConditions `json:"condition_threshold"`
	Name               string                                      `json:"name"`
}

type GoogleMonitoringAlertPolicySpec struct {
	CreationRecord       []GoogleMonitoringAlertPolicySpec `json:"creation_record"`
	Name                 string                            `json:"name"`
	Project              string                            `json:"project"`
	Labels               []string                          `json:"labels"`
	Conditions           []GoogleMonitoringAlertPolicySpec `json:"conditions"`
	DisplayName          string                            `json:"display_name"`
	Enabled              bool                              `json:"enabled"`
	NotificationChannels []string                          `json:"notification_channels"`
	Combiner             string                            `json:"combiner"`
}

type GoogleMonitoringAlertPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleMonitoringAlertPolicyList is a list of GoogleMonitoringAlertPolicys
type GoogleMonitoringAlertPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleMonitoringAlertPolicy CRD objects
	Items []GoogleMonitoringAlertPolicy `json:"items,omitempty"`
}
