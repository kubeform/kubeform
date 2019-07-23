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

type ComputeHealthCheck struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeHealthCheckSpec   `json:"spec,omitempty"`
	Status            ComputeHealthCheckStatus `json:"status,omitempty"`
}

type ComputeHealthCheckSpecHttpHealthCheck struct {
	// +optional
	Host string `json:"host,omitempty" tf:"host,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	ProxyHeader string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`
	// +optional
	RequestPath string `json:"requestPath,omitempty" tf:"request_path,omitempty"`
	// +optional
	Response string `json:"response,omitempty" tf:"response,omitempty"`
}

type ComputeHealthCheckSpecHttpsHealthCheck struct {
	// +optional
	Host string `json:"host,omitempty" tf:"host,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	ProxyHeader string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`
	// +optional
	RequestPath string `json:"requestPath,omitempty" tf:"request_path,omitempty"`
	// +optional
	Response string `json:"response,omitempty" tf:"response,omitempty"`
}

type ComputeHealthCheckSpecSslHealthCheck struct {
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	ProxyHeader string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`
	// +optional
	Request string `json:"request,omitempty" tf:"request,omitempty"`
	// +optional
	Response string `json:"response,omitempty" tf:"response,omitempty"`
}

type ComputeHealthCheckSpecTcpHealthCheck struct {
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	ProxyHeader string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`
	// +optional
	Request string `json:"request,omitempty" tf:"request,omitempty"`
	// +optional
	Response string `json:"response,omitempty" tf:"response,omitempty"`
}

type ComputeHealthCheckSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	CheckIntervalSec int `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	HealthyThreshold int `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpHealthCheck []ComputeHealthCheckSpecHttpHealthCheck `json:"httpHealthCheck,omitempty" tf:"http_health_check,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpsHealthCheck []ComputeHealthCheckSpecHttpsHealthCheck `json:"httpsHealthCheck,omitempty" tf:"https_health_check,omitempty"`
	Name             string                                   `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SslHealthCheck []ComputeHealthCheckSpecSslHealthCheck `json:"sslHealthCheck,omitempty" tf:"ssl_health_check,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TcpHealthCheck []ComputeHealthCheckSpecTcpHealthCheck `json:"tcpHealthCheck,omitempty" tf:"tcp_health_check,omitempty"`
	// +optional
	TimeoutSec int `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`
	// +optional
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type ComputeHealthCheckStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
