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

type ComputeRegionBackendService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeRegionBackendServiceSpec   `json:"spec,omitempty"`
	Status            ComputeRegionBackendServiceStatus `json:"status,omitempty"`
}

type ComputeRegionBackendServiceSpecBackend struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Group string `json:"group,omitempty"`
}

type ComputeRegionBackendServiceSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Backend *[]ComputeRegionBackendServiceSpec `json:"backend,omitempty"`
	// +optional
	ConnectionDrainingTimeoutSec int `json:"connection_draining_timeout_sec,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	HealthChecks []string `json:"health_checks"`
	Name         string   `json:"name"`
}

type ComputeRegionBackendServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
