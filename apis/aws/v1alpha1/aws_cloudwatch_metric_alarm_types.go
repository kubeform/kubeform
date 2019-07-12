package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchMetricAlarm struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudwatchMetricAlarmSpec   `json:"spec,omitempty"`
	Status            AwsCloudwatchMetricAlarmStatus `json:"status,omitempty"`
}

type AwsCloudwatchMetricAlarmSpecMetricQueryMetric struct {
	Unit       string            `json:"unit"`
	Dimensions map[string]string `json:"dimensions"`
	MetricName string            `json:"metric_name"`
	Namespace  string            `json:"namespace"`
	Period     int               `json:"period"`
	Stat       string            `json:"stat"`
}

type AwsCloudwatchMetricAlarmSpecMetricQuery struct {
	Metric     []AwsCloudwatchMetricAlarmSpecMetricQuery `json:"metric"`
	Label      string                                    `json:"label"`
	ReturnData bool                                      `json:"return_data"`
	Id         string                                    `json:"id"`
	Expression string                                    `json:"expression"`
}

type AwsCloudwatchMetricAlarmSpec struct {
	Statistic                         string                         `json:"statistic"`
	Unit                              string                         `json:"unit"`
	Tags                              map[string]string              `json:"tags"`
	Threshold                         float64                        `json:"threshold"`
	AlarmDescription                  string                         `json:"alarm_description"`
	Arn                               string                         `json:"arn"`
	ComparisonOperator                string                         `json:"comparison_operator"`
	MetricName                        string                         `json:"metric_name"`
	AlarmActions                      []string                       `json:"alarm_actions"`
	Dimensions                        map[string]string              `json:"dimensions"`
	InsufficientDataActions           []string                       `json:"insufficient_data_actions"`
	ExtendedStatistic                 string                         `json:"extended_statistic"`
	EvaluateLowSampleCountPercentiles string                         `json:"evaluate_low_sample_count_percentiles"`
	EvaluationPeriods                 int                            `json:"evaluation_periods"`
	Namespace                         string                         `json:"namespace"`
	Period                            int                            `json:"period"`
	DatapointsToAlarm                 int                            `json:"datapoints_to_alarm"`
	OkActions                         []string                       `json:"ok_actions"`
	TreatMissingData                  string                         `json:"treat_missing_data"`
	AlarmName                         string                         `json:"alarm_name"`
	MetricQuery                       []AwsCloudwatchMetricAlarmSpec `json:"metric_query"`
	ActionsEnabled                    bool                           `json:"actions_enabled"`
}



type AwsCloudwatchMetricAlarmStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchMetricAlarmList is a list of AwsCloudwatchMetricAlarms
type AwsCloudwatchMetricAlarmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudwatchMetricAlarm CRD objects
	Items []AwsCloudwatchMetricAlarm `json:"items,omitempty"`
}