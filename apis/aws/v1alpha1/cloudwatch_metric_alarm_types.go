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

type CloudwatchMetricAlarm struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchMetricAlarmSpec   `json:"spec,omitempty"`
	Status            CloudwatchMetricAlarmStatus `json:"status,omitempty"`
}

type CloudwatchMetricAlarmSpecMetricQueryMetric struct {
	// +optional
	Dimensions map[string]string `json:"dimensions,omitempty"`
	MetricName string            `json:"metric_name"`
	// +optional
	Namespace string `json:"namespace,omitempty"`
	Period    int    `json:"period"`
	Stat      string `json:"stat"`
	// +optional
	Unit string `json:"unit,omitempty"`
}

type CloudwatchMetricAlarmSpecMetricQuery struct {
	// +optional
	Expression string `json:"expression,omitempty"`
	Id         string `json:"id"`
	// +optional
	Label string `json:"label,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Metric *[]CloudwatchMetricAlarmSpecMetricQuery `json:"metric,omitempty"`
	// +optional
	ReturnData bool `json:"return_data,omitempty"`
}

type CloudwatchMetricAlarmSpec struct {
	// +optional
	ActionsEnabled bool `json:"actions_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AlarmActions []string `json:"alarm_actions,omitempty"`
	// +optional
	AlarmDescription   string `json:"alarm_description,omitempty"`
	AlarmName          string `json:"alarm_name"`
	ComparisonOperator string `json:"comparison_operator"`
	// +optional
	DatapointsToAlarm int `json:"datapoints_to_alarm,omitempty"`
	// +optional
	Dimensions        map[string]string `json:"dimensions,omitempty"`
	EvaluationPeriods int               `json:"evaluation_periods"`
	// +optional
	ExtendedStatistic string `json:"extended_statistic,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	InsufficientDataActions []string `json:"insufficient_data_actions,omitempty"`
	// +optional
	MetricName string `json:"metric_name,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	MetricQuery *[]CloudwatchMetricAlarmSpec `json:"metric_query,omitempty"`
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	OkActions []string `json:"ok_actions,omitempty"`
	// +optional
	Period int `json:"period,omitempty"`
	// +optional
	Statistic string `json:"statistic,omitempty"`
	// +optional
	Tags      map[string]string `json:"tags,omitempty"`
	Threshold json.Number       `json:"threshold"`
	// +optional
	TreatMissingData string `json:"treat_missing_data,omitempty"`
	// +optional
	Unit string `json:"unit,omitempty"`
}

type CloudwatchMetricAlarmStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
