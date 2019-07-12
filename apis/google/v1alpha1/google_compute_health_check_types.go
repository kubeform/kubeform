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

type GoogleComputeHealthCheck struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeHealthCheckSpec   `json:"spec,omitempty"`
	Status            GoogleComputeHealthCheckStatus `json:"status,omitempty"`
}

type GoogleComputeHealthCheckSpecHttpsHealthCheck struct {
	Response    string `json:"response"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	ProxyHeader string `json:"proxy_header"`
	RequestPath string `json:"request_path"`
}

type GoogleComputeHealthCheckSpecTcpHealthCheck struct {
	Port        int    `json:"port"`
	ProxyHeader string `json:"proxy_header"`
	Request     string `json:"request"`
	Response    string `json:"response"`
}

type GoogleComputeHealthCheckSpecHttpHealthCheck struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	ProxyHeader string `json:"proxy_header"`
	RequestPath string `json:"request_path"`
	Response    string `json:"response"`
}

type GoogleComputeHealthCheckSpecSslHealthCheck struct {
	Response    string `json:"response"`
	Port        int    `json:"port"`
	ProxyHeader string `json:"proxy_header"`
	Request     string `json:"request"`
}

type GoogleComputeHealthCheckSpec struct {
	Description        string                         `json:"description"`
	CreationTimestamp  string                         `json:"creation_timestamp"`
	Type               string                         `json:"type"`
	HttpsHealthCheck   []GoogleComputeHealthCheckSpec `json:"https_health_check"`
	TcpHealthCheck     []GoogleComputeHealthCheckSpec `json:"tcp_health_check"`
	SelfLink           string                         `json:"self_link"`
	Name               string                         `json:"name"`
	HttpHealthCheck    []GoogleComputeHealthCheckSpec `json:"http_health_check"`
	SslHealthCheck     []GoogleComputeHealthCheckSpec `json:"ssl_health_check"`
	TimeoutSec         int                            `json:"timeout_sec"`
	UnhealthyThreshold int                            `json:"unhealthy_threshold"`
	Project            string                         `json:"project"`
	CheckIntervalSec   int                            `json:"check_interval_sec"`
	HealthyThreshold   int                            `json:"healthy_threshold"`
}

type GoogleComputeHealthCheckStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeHealthCheckList is a list of GoogleComputeHealthChecks
type GoogleComputeHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeHealthCheck CRD objects
	Items []GoogleComputeHealthCheck `json:"items,omitempty"`
}
