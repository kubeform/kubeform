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

type Route53HealthCheck struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53HealthCheckSpec   `json:"spec,omitempty"`
	Status            Route53HealthCheckStatus `json:"status,omitempty"`
}

type Route53HealthCheckSpec struct {
	// +optional
	ChildHealthThreshold int `json:"child_health_threshold,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ChildHealthchecks []string `json:"child_healthchecks,omitempty"`
	// +optional
	CloudwatchAlarmName string `json:"cloudwatch_alarm_name,omitempty"`
	// +optional
	CloudwatchAlarmRegion string `json:"cloudwatch_alarm_region,omitempty"`
	// +optional
	FailureThreshold int `json:"failure_threshold,omitempty"`
	// +optional
	Fqdn string `json:"fqdn,omitempty"`
	// +optional
	InsufficientDataHealthStatus string `json:"insufficient_data_health_status,omitempty"`
	// +optional
	InvertHealthcheck bool `json:"invert_healthcheck,omitempty"`
	// +optional
	IpAddress string `json:"ip_address,omitempty"`
	// +optional
	MeasureLatency bool `json:"measure_latency,omitempty"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	ReferenceName string `json:"reference_name,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Regions []string `json:"regions,omitempty"`
	// +optional
	RequestInterval int `json:"request_interval,omitempty"`
	// +optional
	ResourcePath string `json:"resource_path,omitempty"`
	// +optional
	SearchString string `json:"search_string,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	Type string            `json:"type"`
}

type Route53HealthCheckStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Route53HealthCheckList is a list of Route53HealthChecks
type Route53HealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Route53HealthCheck CRD objects
	Items []Route53HealthCheck `json:"items,omitempty"`
}
