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

type ComputeHttpsHealthCheck struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeHttpsHealthCheckSpec   `json:"spec,omitempty"`
	Status            ComputeHttpsHealthCheckStatus `json:"status,omitempty"`
}

type ComputeHttpsHealthCheckSpec struct {
	// +optional
	CheckIntervalSec int `json:"check_interval_sec,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	HealthyThreshold int `json:"healthy_threshold,omitempty"`
	// +optional
	Host string `json:"host,omitempty"`
	Name string `json:"name"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	RequestPath string `json:"request_path,omitempty"`
	// +optional
	TimeoutSec int `json:"timeout_sec,omitempty"`
	// +optional
	UnhealthyThreshold int `json:"unhealthy_threshold,omitempty"`
}

type ComputeHttpsHealthCheckStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeHttpsHealthCheckList is a list of ComputeHttpsHealthChecks
type ComputeHttpsHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeHttpsHealthCheck CRD objects
	Items []ComputeHttpsHealthCheck `json:"items,omitempty"`
}
