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

type GoogleLoggingProjectSink struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleLoggingProjectSinkSpec   `json:"spec,omitempty"`
	Status            GoogleLoggingProjectSinkStatus `json:"status,omitempty"`
}

type GoogleLoggingProjectSinkSpec struct {
	UniqueWriterIdentity bool   `json:"unique_writer_identity"`
	Name                 string `json:"name"`
	Destination          string `json:"destination"`
	Filter               string `json:"filter"`
	WriterIdentity       string `json:"writer_identity"`
	Project              string `json:"project"`
}



type GoogleLoggingProjectSinkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleLoggingProjectSinkList is a list of GoogleLoggingProjectSinks
type GoogleLoggingProjectSinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleLoggingProjectSink CRD objects
	Items []GoogleLoggingProjectSink `json:"items,omitempty"`
}