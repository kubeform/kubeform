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
	Fqdn                         string            `json:"fqdn"`
	Type                         string            `json:"type"`
	IpAddress                    string            `json:"ip_address"`
	Tags                         map[string]string `json:"tags"`
	Port                         int               `json:"port"`
	ResourcePath                 string            `json:"resource_path"`
	ChildHealthThreshold         int               `json:"child_health_threshold"`
	ReferenceName                string            `json:"reference_name"`
	EnableSni                    bool              `json:"enable_sni"`
	FailureThreshold             int               `json:"failure_threshold"`
	RequestInterval              int               `json:"request_interval"`
	InvertHealthcheck            bool              `json:"invert_healthcheck"`
	SearchString                 string            `json:"search_string"`
	MeasureLatency               bool              `json:"measure_latency"`
	ChildHealthchecks            []string          `json:"child_healthchecks"`
	CloudwatchAlarmName          string            `json:"cloudwatch_alarm_name"`
	CloudwatchAlarmRegion        string            `json:"cloudwatch_alarm_region"`
	InsufficientDataHealthStatus string            `json:"insufficient_data_health_status"`
	Regions                      []string          `json:"regions"`
}



type AwsRoute53HealthCheckStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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