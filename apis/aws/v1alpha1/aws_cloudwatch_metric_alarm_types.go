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
	Label      string                                    `json:"label"`
	ReturnData bool                                      `json:"return_data"`
	Id         string                                    `json:"id"`
	Expression string                                    `json:"expression"`
	Metric     []AwsCloudwatchMetricAlarmSpecMetricQuery `json:"metric"`
}

type AwsCloudwatchMetricAlarmSpec struct {
	AlarmName                         string                         `json:"alarm_name"`
	Namespace                         string                         `json:"namespace"`
	AlarmDescription                  string                         `json:"alarm_description"`
	DatapointsToAlarm                 int                            `json:"datapoints_to_alarm"`
	TreatMissingData                  string                         `json:"treat_missing_data"`
	EvaluationPeriods                 int                            `json:"evaluation_periods"`
	AlarmActions                      []string                       `json:"alarm_actions"`
	InsufficientDataActions           []string                       `json:"insufficient_data_actions"`
	OkActions                         []string                       `json:"ok_actions"`
	Tags                              map[string]string              `json:"tags"`
	ComparisonOperator                string                         `json:"comparison_operator"`
	MetricQuery                       []AwsCloudwatchMetricAlarmSpec `json:"metric_query"`
	Period                            int                            `json:"period"`
	Threshold                         float64                        `json:"threshold"`
	Unit                              string                         `json:"unit"`
	Arn                               string                         `json:"arn"`
	MetricName                        string                         `json:"metric_name"`
	Statistic                         string                         `json:"statistic"`
	ActionsEnabled                    bool                           `json:"actions_enabled"`
	Dimensions                        map[string]string              `json:"dimensions"`
	ExtendedStatistic                 string                         `json:"extended_statistic"`
	EvaluateLowSampleCountPercentiles string                         `json:"evaluate_low_sample_count_percentiles"`
}

type AwsCloudwatchMetricAlarmStatus struct {
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
