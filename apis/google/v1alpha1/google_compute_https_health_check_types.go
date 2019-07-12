package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeHttpsHealthCheck struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeHttpsHealthCheckSpec   `json:"spec,omitempty"`
	Status            GoogleComputeHttpsHealthCheckStatus `json:"status,omitempty"`
}

type GoogleComputeHttpsHealthCheckSpec struct {
	RequestPath        string `json:"request_path"`
	TimeoutSec         int    `json:"timeout_sec"`
	CreationTimestamp  string `json:"creation_timestamp"`
	Project            string `json:"project"`
	Name               string `json:"name"`
	CheckIntervalSec   int    `json:"check_interval_sec"`
	Host               string `json:"host"`
	Port               int    `json:"port"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	SelfLink           string `json:"self_link"`
	Description        string `json:"description"`
	HealthyThreshold   int    `json:"healthy_threshold"`
}



type GoogleComputeHttpsHealthCheckStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeHttpsHealthCheckList is a list of GoogleComputeHttpsHealthChecks
type GoogleComputeHttpsHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeHttpsHealthCheck CRD objects
	Items []GoogleComputeHttpsHealthCheck `json:"items,omitempty"`
}