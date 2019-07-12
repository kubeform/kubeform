package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeHealthCheck struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeHealthCheckSpec   `json:"spec,omitempty"`
	Status            GoogleComputeHealthCheckStatus `json:"status,omitempty"`
}

type GoogleComputeHealthCheckSpecHttpHealthCheck struct {
	RequestPath string `json:"request_path"`
	Response    string `json:"response"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	ProxyHeader string `json:"proxy_header"`
}

type GoogleComputeHealthCheckSpecHttpsHealthCheck struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	ProxyHeader string `json:"proxy_header"`
	RequestPath string `json:"request_path"`
	Response    string `json:"response"`
}

type GoogleComputeHealthCheckSpecSslHealthCheck struct {
	Request     string `json:"request"`
	Response    string `json:"response"`
	Port        int    `json:"port"`
	ProxyHeader string `json:"proxy_header"`
}

type GoogleComputeHealthCheckSpecTcpHealthCheck struct {
	Request     string `json:"request"`
	Response    string `json:"response"`
	Port        int    `json:"port"`
	ProxyHeader string `json:"proxy_header"`
}

type GoogleComputeHealthCheckSpec struct {
	Description        string                         `json:"description"`
	HealthyThreshold   int                            `json:"healthy_threshold"`
	TimeoutSec         int                            `json:"timeout_sec"`
	UnhealthyThreshold int                            `json:"unhealthy_threshold"`
	HttpHealthCheck    []GoogleComputeHealthCheckSpec `json:"http_health_check"`
	CreationTimestamp  string                         `json:"creation_timestamp"`
	HttpsHealthCheck   []GoogleComputeHealthCheckSpec `json:"https_health_check"`
	SslHealthCheck     []GoogleComputeHealthCheckSpec `json:"ssl_health_check"`
	SelfLink           string                         `json:"self_link"`
	Name               string                         `json:"name"`
	CheckIntervalSec   int                            `json:"check_interval_sec"`
	TcpHealthCheck     []GoogleComputeHealthCheckSpec `json:"tcp_health_check"`
	Type               string                         `json:"type"`
	Project            string                         `json:"project"`
}



type GoogleComputeHealthCheckStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeHealthCheckList is a list of GoogleComputeHealthChecks
type GoogleComputeHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeHealthCheck CRD objects
	Items []GoogleComputeHealthCheck `json:"items,omitempty"`
}