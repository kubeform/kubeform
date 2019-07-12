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

type AwsRoute53HealthCheck struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRoute53HealthCheckSpec   `json:"spec,omitempty"`
	Status            AwsRoute53HealthCheckStatus `json:"status,omitempty"`
}

type AwsRoute53HealthCheckSpec struct {
	Port                         int               `json:"port"`
	InvertHealthcheck            bool              `json:"invert_healthcheck"`
	MeasureLatency               bool              `json:"measure_latency"`
	CloudwatchAlarmName          string            `json:"cloudwatch_alarm_name"`
	CloudwatchAlarmRegion        string            `json:"cloudwatch_alarm_region"`
	RequestInterval              int               `json:"request_interval"`
	ChildHealthchecks            []string          `json:"child_healthchecks"`
	ChildHealthThreshold         int               `json:"child_health_threshold"`
	InsufficientDataHealthStatus string            `json:"insufficient_data_health_status"`
	Type                         string            `json:"type"`
	IpAddress                    string            `json:"ip_address"`
	Fqdn                         string            `json:"fqdn"`
	SearchString                 string            `json:"search_string"`
	ReferenceName                string            `json:"reference_name"`
	EnableSni                    bool              `json:"enable_sni"`
	FailureThreshold             int               `json:"failure_threshold"`
	ResourcePath                 string            `json:"resource_path"`
	Regions                      []string          `json:"regions"`
	Tags                         map[string]string `json:"tags"`
}

type AwsRoute53HealthCheckStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRoute53HealthCheckList is a list of AwsRoute53HealthChecks
type AwsRoute53HealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRoute53HealthCheck CRD objects
	Items []AwsRoute53HealthCheck `json:"items,omitempty"`
}
