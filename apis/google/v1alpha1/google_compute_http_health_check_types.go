package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeHttpHealthCheck struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeHttpHealthCheckSpec   `json:"spec,omitempty"`
	Status            GoogleComputeHttpHealthCheckStatus `json:"status,omitempty"`
}

type GoogleComputeHttpHealthCheckSpec struct {
	Name               string `json:"name"`
	Port               int    `json:"port"`
	TimeoutSec         int    `json:"timeout_sec"`
	CreationTimestamp  string `json:"creation_timestamp"`
	Project            string `json:"project"`
	CheckIntervalSec   int    `json:"check_interval_sec"`
	Description        string `json:"description"`
	HealthyThreshold   int    `json:"healthy_threshold"`
	Host               string `json:"host"`
	RequestPath        string `json:"request_path"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	SelfLink           string `json:"self_link"`
}



type GoogleComputeHttpHealthCheckStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeHttpHealthCheckList is a list of GoogleComputeHttpHealthChecks
type GoogleComputeHttpHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeHttpHealthCheck CRD objects
	Items []GoogleComputeHttpHealthCheck `json:"items,omitempty"`
}