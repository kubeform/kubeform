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

type ComputeHTTPSHealthCheck struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeHTTPSHealthCheckSpec   `json:"spec,omitempty"`
	Status            ComputeHTTPSHealthCheckStatus `json:"status,omitempty"`
}

type ComputeHTTPSHealthCheckSpec struct {
	// +optional
	CheckIntervalSec int `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	HealthyThreshold int `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`
	// +optional
	Host string `json:"host,omitempty" tf:"host,omitempty"`
	Name string `json:"name" tf:"name"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	RequestPath string `json:"requestPath,omitempty" tf:"request_path,omitempty"`
	// +optional
	TimeoutSec int `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`
	// +optional
	UnhealthyThreshold int                       `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
	ProviderRef        core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeHTTPSHealthCheckStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeHTTPSHealthCheckList is a list of ComputeHTTPSHealthChecks
type ComputeHTTPSHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeHTTPSHealthCheck CRD objects
	Items []ComputeHTTPSHealthCheck `json:"items,omitempty"`
}
