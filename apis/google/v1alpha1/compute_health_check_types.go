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

type ComputeHealthCheck struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeHealthCheckSpec   `json:"spec,omitempty"`
	Status            ComputeHealthCheckStatus `json:"status,omitempty"`
}

type ComputeHealthCheckSpecHttpHealthCheck struct {
	// +optional
	Host string `json:"host,omitempty"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	ProxyHeader string `json:"proxy_header,omitempty"`
	// +optional
	RequestPath string `json:"request_path,omitempty"`
	// +optional
	Response string `json:"response,omitempty"`
}

type ComputeHealthCheckSpecHttpsHealthCheck struct {
	// +optional
	Host string `json:"host,omitempty"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	ProxyHeader string `json:"proxy_header,omitempty"`
	// +optional
	RequestPath string `json:"request_path,omitempty"`
	// +optional
	Response string `json:"response,omitempty"`
}

type ComputeHealthCheckSpecSslHealthCheck struct {
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	ProxyHeader string `json:"proxy_header,omitempty"`
	// +optional
	Request string `json:"request,omitempty"`
	// +optional
	Response string `json:"response,omitempty"`
}

type ComputeHealthCheckSpecTcpHealthCheck struct {
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	ProxyHeader string `json:"proxy_header,omitempty"`
	// +optional
	Request string `json:"request,omitempty"`
	// +optional
	Response string `json:"response,omitempty"`
}

type ComputeHealthCheckSpec struct {
	// +optional
	CheckIntervalSec int `json:"check_interval_sec,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	HealthyThreshold int `json:"healthy_threshold,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpHealthCheck *[]ComputeHealthCheckSpec `json:"http_health_check,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpsHealthCheck *[]ComputeHealthCheckSpec `json:"https_health_check,omitempty"`
	Name             string                    `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SslHealthCheck *[]ComputeHealthCheckSpec `json:"ssl_health_check,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TcpHealthCheck *[]ComputeHealthCheckSpec `json:"tcp_health_check,omitempty"`
	// +optional
	TimeoutSec int `json:"timeout_sec,omitempty"`
	// +optional
	UnhealthyThreshold int `json:"unhealthy_threshold,omitempty"`
}

type ComputeHealthCheckStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeHealthCheckList is a list of ComputeHealthChecks
type ComputeHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeHealthCheck CRD objects
	Items []ComputeHealthCheck `json:"items,omitempty"`
}
