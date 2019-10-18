package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type MonitoringAlertPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitoringAlertPolicySpec   `json:"spec,omitempty"`
	Status            MonitoringAlertPolicyStatus `json:"status,omitempty"`
}

type MonitoringAlertPolicySpecConditionsConditionAbsentAggregations struct {
	// +optional
	AlignmentPeriod string `json:"alignmentPeriod,omitempty" tf:"alignment_period,omitempty"`
	// +optional
	CrossSeriesReducer string `json:"crossSeriesReducer,omitempty" tf:"cross_series_reducer,omitempty"`
	// +optional
	GroupByFields []string `json:"groupByFields,omitempty" tf:"group_by_fields,omitempty"`
	// +optional
	PerSeriesAligner string `json:"perSeriesAligner,omitempty" tf:"per_series_aligner,omitempty"`
}

type MonitoringAlertPolicySpecConditionsConditionAbsentTrigger struct {
	// +optional
	Count int `json:"count,omitempty" tf:"count,omitempty"`
	// +optional
	Percent json.Number `json:"percent,omitempty" tf:"percent,omitempty"`
}

type MonitoringAlertPolicySpecConditionsConditionAbsent struct {
	// +optional
	Aggregations []MonitoringAlertPolicySpecConditionsConditionAbsentAggregations `json:"aggregations,omitempty" tf:"aggregations,omitempty"`
	Duration     string                                                           `json:"duration" tf:"duration"`
	// +optional
	Filter string `json:"filter,omitempty" tf:"filter,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Trigger []MonitoringAlertPolicySpecConditionsConditionAbsentTrigger `json:"trigger,omitempty" tf:"trigger,omitempty"`
}

type MonitoringAlertPolicySpecConditionsConditionThresholdAggregations struct {
	// +optional
	AlignmentPeriod string `json:"alignmentPeriod,omitempty" tf:"alignment_period,omitempty"`
	// +optional
	CrossSeriesReducer string `json:"crossSeriesReducer,omitempty" tf:"cross_series_reducer,omitempty"`
	// +optional
	GroupByFields []string `json:"groupByFields,omitempty" tf:"group_by_fields,omitempty"`
	// +optional
	PerSeriesAligner string `json:"perSeriesAligner,omitempty" tf:"per_series_aligner,omitempty"`
}

type MonitoringAlertPolicySpecConditionsConditionThresholdDenominatorAggregations struct {
	// +optional
	AlignmentPeriod string `json:"alignmentPeriod,omitempty" tf:"alignment_period,omitempty"`
	// +optional
	CrossSeriesReducer string `json:"crossSeriesReducer,omitempty" tf:"cross_series_reducer,omitempty"`
	// +optional
	GroupByFields []string `json:"groupByFields,omitempty" tf:"group_by_fields,omitempty"`
	// +optional
	PerSeriesAligner string `json:"perSeriesAligner,omitempty" tf:"per_series_aligner,omitempty"`
}

type MonitoringAlertPolicySpecConditionsConditionThresholdTrigger struct {
	// +optional
	Count int `json:"count,omitempty" tf:"count,omitempty"`
	// +optional
	Percent json.Number `json:"percent,omitempty" tf:"percent,omitempty"`
}

type MonitoringAlertPolicySpecConditionsConditionThreshold struct {
	// +optional
	Aggregations []MonitoringAlertPolicySpecConditionsConditionThresholdAggregations `json:"aggregations,omitempty" tf:"aggregations,omitempty"`
	Comparison   string                                                              `json:"comparison" tf:"comparison"`
	// +optional
	DenominatorAggregations []MonitoringAlertPolicySpecConditionsConditionThresholdDenominatorAggregations `json:"denominatorAggregations,omitempty" tf:"denominator_aggregations,omitempty"`
	// +optional
	DenominatorFilter string `json:"denominatorFilter,omitempty" tf:"denominator_filter,omitempty"`
	Duration          string `json:"duration" tf:"duration"`
	// +optional
	Filter string `json:"filter,omitempty" tf:"filter,omitempty"`
	// +optional
	ThresholdValue json.Number `json:"thresholdValue,omitempty" tf:"threshold_value,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Trigger []MonitoringAlertPolicySpecConditionsConditionThresholdTrigger `json:"trigger,omitempty" tf:"trigger,omitempty"`
}

type MonitoringAlertPolicySpecConditions struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ConditionAbsent []MonitoringAlertPolicySpecConditionsConditionAbsent `json:"conditionAbsent,omitempty" tf:"condition_absent,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ConditionThreshold []MonitoringAlertPolicySpecConditionsConditionThreshold `json:"conditionThreshold,omitempty" tf:"condition_threshold,omitempty"`
	DisplayName        string                                                  `json:"displayName" tf:"display_name"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
}

type MonitoringAlertPolicySpecCreationRecord struct {
	// +optional
	MutateTime string `json:"mutateTime,omitempty" tf:"mutate_time,omitempty"`
	// +optional
	MutatedBy string `json:"mutatedBy,omitempty" tf:"mutated_by,omitempty"`
}

type MonitoringAlertPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Combiner   string                                `json:"combiner" tf:"combiner"`
	Conditions []MonitoringAlertPolicySpecConditions `json:"conditions" tf:"conditions"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CreationRecord []MonitoringAlertPolicySpecCreationRecord `json:"creationRecord,omitempty" tf:"creation_record,omitempty"`
	DisplayName    string                                    `json:"displayName" tf:"display_name"`
	Enabled        bool                                      `json:"enabled" tf:"enabled"`
	// +optional
	Labels []string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NotificationChannels []string `json:"notificationChannels,omitempty" tf:"notification_channels,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
}

type MonitoringAlertPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MonitoringAlertPolicySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
