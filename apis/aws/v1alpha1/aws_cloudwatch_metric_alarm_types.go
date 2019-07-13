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

type AwsCloudwatchMetricAlarm struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudwatchMetricAlarmSpec   `json:"spec,omitempty"`
	Status            AwsCloudwatchMetricAlarmStatus `json:"status,omitempty"`
}

type AwsCloudwatchMetricAlarmSpecMetricQueryMetric struct {
	Dimensions map[string]string `json:"dimensions"`
	MetricName string            `json:"metric_name"`
	Namespace  string            `json:"namespace"`
	Period     int               `json:"period"`
	Stat       string            `json:"stat"`
	Unit       string            `json:"unit"`
}

type AwsCloudwatchMetricAlarmSpecMetricQuery struct {
	Metric     []AwsCloudwatchMetricAlarmSpecMetricQuery `json:"metric"`
	Label      string                                    `json:"label"`
	ReturnData bool                                      `json:"return_data"`
	Id         string                                    `json:"id"`
	Expression string                                    `json:"expression"`
}

type AwsCloudwatchMetricAlarmSpec struct {
	Arn                               string                         `json:"arn"`
	EvaluationPeriods                 int                            `json:"evaluation_periods"`
	Dimensions                        map[string]string              `json:"dimensions"`
	InsufficientDataActions           []string                       `json:"insufficient_data_actions"`
	TreatMissingData                  string                         `json:"treat_missing_data"`
	ActionsEnabled                    bool                           `json:"actions_enabled"`
	AlarmDescription                  string                         `json:"alarm_description"`
	Unit                              string                         `json:"unit"`
	MetricName                        string                         `json:"metric_name"`
	MetricQuery                       []AwsCloudwatchMetricAlarmSpec `json:"metric_query"`
	Period                            int                            `json:"period"`
	Statistic                         string                         `json:"statistic"`
	Threshold                         float64                        `json:"threshold"`
	Tags                              map[string]string              `json:"tags"`
	OkActions                         []string                       `json:"ok_actions"`
	ExtendedStatistic                 string                         `json:"extended_statistic"`
	EvaluateLowSampleCountPercentiles string                         `json:"evaluate_low_sample_count_percentiles"`
	AlarmName                         string                         `json:"alarm_name"`
	ComparisonOperator                string                         `json:"comparison_operator"`
	Namespace                         string                         `json:"namespace"`
	AlarmActions                      []string                       `json:"alarm_actions"`
	DatapointsToAlarm                 int                            `json:"datapoints_to_alarm"`
}



type AwsCloudwatchMetricAlarmStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCloudwatchMetricAlarmList is a list of AwsCloudwatchMetricAlarms
type AwsCloudwatchMetricAlarmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudwatchMetricAlarm CRD objects
	Items []AwsCloudwatchMetricAlarm `json:"items,omitempty"`
}