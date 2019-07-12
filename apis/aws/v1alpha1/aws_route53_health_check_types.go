package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53HealthCheck struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRoute53HealthCheckSpec   `json:"spec,omitempty"`
	Status            AwsRoute53HealthCheckStatus `json:"status,omitempty"`
}

type AwsRoute53HealthCheckSpec struct {
	RequestInterval              int               `json:"request_interval"`
	InsufficientDataHealthStatus string            `json:"insufficient_data_health_status"`
	CloudwatchAlarmRegion        string            `json:"cloudwatch_alarm_region"`
	ResourcePath                 string            `json:"resource_path"`
	SearchString                 string            `json:"search_string"`
	ChildHealthchecks            []string          `json:"child_healthchecks"`
	CloudwatchAlarmName          string            `json:"cloudwatch_alarm_name"`
	Tags                         map[string]string `json:"tags"`
	IpAddress                    string            `json:"ip_address"`
	Port                         int               `json:"port"`
	MeasureLatency               bool              `json:"measure_latency"`
	ReferenceName                string            `json:"reference_name"`
	ChildHealthThreshold         int               `json:"child_health_threshold"`
	EnableSni                    bool              `json:"enable_sni"`
	Regions                      []string          `json:"regions"`
	Type                         string            `json:"type"`
	FailureThreshold             int               `json:"failure_threshold"`
	Fqdn                         string            `json:"fqdn"`
	InvertHealthcheck            bool              `json:"invert_healthcheck"`
}



type AwsRoute53HealthCheckStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53HealthCheckList is a list of AwsRoute53HealthChecks
type AwsRoute53HealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRoute53HealthCheck CRD objects
	Items []AwsRoute53HealthCheck `json:"items,omitempty"`
}