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

type LoggingProjectSink struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoggingProjectSinkSpec   `json:"spec,omitempty"`
	Status            LoggingProjectSinkStatus `json:"status,omitempty"`
}

type LoggingProjectSinkSpec struct {
	Destination string `json:"destination"`
	// +optional
	Filter string `json:"filter,omitempty"`
	Name   string `json:"name"`
	// +optional
	UniqueWriterIdentity bool `json:"unique_writer_identity,omitempty"`
}

type LoggingProjectSinkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LoggingProjectSinkList is a list of LoggingProjectSinks
type LoggingProjectSinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LoggingProjectSink CRD objects
	Items []LoggingProjectSink `json:"items,omitempty"`
}
