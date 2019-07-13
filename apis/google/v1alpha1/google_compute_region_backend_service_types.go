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

type GoogleComputeRegionBackendService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeRegionBackendServiceSpec   `json:"spec,omitempty"`
	Status            GoogleComputeRegionBackendServiceStatus `json:"status,omitempty"`
}

type GoogleComputeRegionBackendServiceSpecBackend struct {
	Group       string `json:"group"`
	Description string `json:"description"`
}

type GoogleComputeRegionBackendServiceSpec struct {
	Name                         string                                  `json:"name"`
	HealthChecks                 []string                                `json:"health_checks"`
	Fingerprint                  string                                  `json:"fingerprint"`
	Protocol                     string                                  `json:"protocol"`
	Region                       string                                  `json:"region"`
	SelfLink                     string                                  `json:"self_link"`
	Backend                      []GoogleComputeRegionBackendServiceSpec `json:"backend"`
	Description                  string                                  `json:"description"`
	Project                      string                                  `json:"project"`
	SessionAffinity              string                                  `json:"session_affinity"`
	TimeoutSec                   int                                     `json:"timeout_sec"`
	ConnectionDrainingTimeoutSec int                                     `json:"connection_draining_timeout_sec"`
}



type GoogleComputeRegionBackendServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeRegionBackendServiceList is a list of GoogleComputeRegionBackendServices
type GoogleComputeRegionBackendServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeRegionBackendService CRD objects
	Items []GoogleComputeRegionBackendService `json:"items,omitempty"`
}