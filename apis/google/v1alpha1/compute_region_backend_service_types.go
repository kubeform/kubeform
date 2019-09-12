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

type ComputeRegionBackendService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeRegionBackendServiceSpec   `json:"spec,omitempty"`
	Status            ComputeRegionBackendServiceStatus `json:"status,omitempty"`
}

type ComputeRegionBackendServiceSpecBackend struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Group string `json:"group,omitempty" tf:"group,omitempty"`
}

type ComputeRegionBackendServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Backend []ComputeRegionBackendServiceSpecBackend `json:"backend,omitempty" tf:"backend,omitempty"`
	// +optional
	ConnectionDrainingTimeoutSec int `json:"connectionDrainingTimeoutSec,omitempty" tf:"connection_draining_timeout_sec,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Fingerprint string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	HealthChecks []string `json:"healthChecks" tf:"health_checks"`
	Name         string   `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty" tf:"protocol,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	// +optional
	SessionAffinity string `json:"sessionAffinity,omitempty" tf:"session_affinity,omitempty"`
	// +optional
	TimeoutSec int `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`
}

type ComputeRegionBackendServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeRegionBackendServiceSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeRegionBackendServiceList is a list of ComputeRegionBackendServices
type ComputeRegionBackendServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeRegionBackendService CRD objects
	Items []ComputeRegionBackendService `json:"items,omitempty"`
}
