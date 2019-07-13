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

type GoogleComputeHttpsHealthCheck struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeHttpsHealthCheckSpec   `json:"spec,omitempty"`
	Status            GoogleComputeHttpsHealthCheckStatus `json:"status,omitempty"`
}

type GoogleComputeHttpsHealthCheckSpec struct {
	CheckIntervalSec   int    `json:"check_interval_sec"`
	Description        string `json:"description"`
	HealthyThreshold   int    `json:"healthy_threshold"`
	TimeoutSec         int    `json:"timeout_sec"`
	SelfLink           string `json:"self_link"`
	Name               string `json:"name"`
	Host               string `json:"host"`
	Port               int    `json:"port"`
	RequestPath        string `json:"request_path"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	CreationTimestamp  string `json:"creation_timestamp"`
	Project            string `json:"project"`
}



type GoogleComputeHttpsHealthCheckStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeHttpsHealthCheckList is a list of GoogleComputeHttpsHealthChecks
type GoogleComputeHttpsHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeHttpsHealthCheck CRD objects
	Items []GoogleComputeHttpsHealthCheck `json:"items,omitempty"`
}