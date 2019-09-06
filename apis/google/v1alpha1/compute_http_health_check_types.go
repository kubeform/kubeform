package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeHTTPHealthCheck struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeHTTPHealthCheckSpec   `json:"spec,omitempty"`
	Status            ComputeHTTPHealthCheckStatus `json:"status,omitempty"`
}

type ComputeHTTPHealthCheckSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CheckIntervalSec int `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`
	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
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
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	// +optional
	TimeoutSec int `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`
	// +optional
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type ComputeHTTPHealthCheckStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeHTTPHealthCheckSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeHTTPHealthCheckList is a list of ComputeHTTPHealthChecks
type ComputeHTTPHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeHTTPHealthCheck CRD objects
	Items []ComputeHTTPHealthCheck `json:"items,omitempty"`
}
