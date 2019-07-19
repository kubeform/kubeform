package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	ChildHealthThreshold int `json:"childHealthThreshold,omitempty" tf:"child_health_threshold,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ChildHealthchecks []string `json:"childHealthchecks,omitempty" tf:"child_healthchecks,omitempty"`
	// +optional
	CloudwatchAlarmName string `json:"cloudwatchAlarmName,omitempty" tf:"cloudwatch_alarm_name,omitempty"`
	// +optional
	CloudwatchAlarmRegion string `json:"cloudwatchAlarmRegion,omitempty" tf:"cloudwatch_alarm_region,omitempty"`
	// +optional
	EnableSni bool `json:"enableSni,omitempty" tf:"enable_sni,omitempty"`
	// +optional
	FailureThreshold int `json:"failureThreshold,omitempty" tf:"failure_threshold,omitempty"`
	// +optional
	Fqdn string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
	// +optional
	InsufficientDataHealthStatus string `json:"insufficientDataHealthStatus,omitempty" tf:"insufficient_data_health_status,omitempty"`
	// +optional
	InvertHealthcheck bool `json:"invertHealthcheck,omitempty" tf:"invert_healthcheck,omitempty"`
	// +optional
	IpAddress string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`
	// +optional
	MeasureLatency bool `json:"measureLatency,omitempty" tf:"measure_latency,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	ReferenceName string `json:"referenceName,omitempty" tf:"reference_name,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Regions []string `json:"regions,omitempty" tf:"regions,omitempty"`
	// +optional
	RequestInterval int `json:"requestInterval,omitempty" tf:"request_interval,omitempty"`
	// +optional
	ResourcePath string `json:"resourcePath,omitempty" tf:"resource_path,omitempty"`
	// +optional
	SearchString string `json:"searchString,omitempty" tf:"search_string,omitempty"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	Type        string                    `json:"type" tf:"type"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type Route53HealthCheckStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
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
