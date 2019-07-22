package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type CloudwatchMetricAlarm struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchMetricAlarmSpec   `json:"spec,omitempty"`
	Status            CloudwatchMetricAlarmStatus `json:"status,omitempty"`
}

type CloudwatchMetricAlarmSpecMetricQueryMetric struct {
	// +optional
	Dimensions map[string]string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`
	MetricName string            `json:"metricName" tf:"metric_name"`
	// +optional
	Namespace string `json:"namespace,omitempty" tf:"namespace,omitempty"`
	Period    int    `json:"period" tf:"period"`
	Stat      string `json:"stat" tf:"stat"`
	// +optional
	Unit string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type CloudwatchMetricAlarmSpecMetricQuery struct {
	// +optional
	Expression string `json:"expression,omitempty" tf:"expression,omitempty"`
	ID         string `json:"ID" tf:"id"`
	// +optional
	Label string `json:"label,omitempty" tf:"label,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Metric []CloudwatchMetricAlarmSpecMetricQueryMetric `json:"metric,omitempty" tf:"metric,omitempty"`
	// +optional
	ReturnData bool `json:"returnData,omitempty" tf:"return_data,omitempty"`
}

type CloudwatchMetricAlarmSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	ActionsEnabled bool `json:"actionsEnabled,omitempty" tf:"actions_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AlarmActions []string `json:"alarmActions,omitempty" tf:"alarm_actions,omitempty"`
	// +optional
	AlarmDescription   string `json:"alarmDescription,omitempty" tf:"alarm_description,omitempty"`
	AlarmName          string `json:"alarmName" tf:"alarm_name"`
	ComparisonOperator string `json:"comparisonOperator" tf:"comparison_operator"`
	// +optional
	DatapointsToAlarm int `json:"datapointsToAlarm,omitempty" tf:"datapoints_to_alarm,omitempty"`
	// +optional
	Dimensions map[string]string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`
	// +optional
	EvaluateLowSampleCountPercentiles string `json:"evaluateLowSampleCountPercentiles,omitempty" tf:"evaluate_low_sample_count_percentiles,omitempty"`
	EvaluationPeriods                 int    `json:"evaluationPeriods" tf:"evaluation_periods"`
	// +optional
	ExtendedStatistic string `json:"extendedStatistic,omitempty" tf:"extended_statistic,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	InsufficientDataActions []string `json:"insufficientDataActions,omitempty" tf:"insufficient_data_actions,omitempty"`
	// +optional
	MetricName string `json:"metricName,omitempty" tf:"metric_name,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	MetricQuery []CloudwatchMetricAlarmSpecMetricQuery `json:"metricQuery,omitempty" tf:"metric_query,omitempty"`
	// +optional
	Namespace string `json:"namespace,omitempty" tf:"namespace,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	OkActions []string `json:"okActions,omitempty" tf:"ok_actions,omitempty"`
	// +optional
	Period int `json:"period,omitempty" tf:"period,omitempty"`
	// +optional
	Statistic string `json:"statistic,omitempty" tf:"statistic,omitempty"`
	// +optional
	Tags      map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Threshold json.Number       `json:"threshold" tf:"threshold"`
	// +optional
	TreatMissingData string `json:"treatMissingData,omitempty" tf:"treat_missing_data,omitempty"`
	// +optional
	Unit string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type CloudwatchMetricAlarmStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudwatchMetricAlarmList is a list of CloudwatchMetricAlarms
type CloudwatchMetricAlarmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudwatchMetricAlarm CRD objects
	Items []CloudwatchMetricAlarm `json:"items,omitempty"`
}
