package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type LoggingProjectSink struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoggingProjectSinkSpec   `json:"spec,omitempty"`
	Status            LoggingProjectSinkStatus `json:"status,omitempty"`
}

type LoggingProjectSinkSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Destination string `json:"destination" tf:"destination"`
	// +optional
	Filter string `json:"filter,omitempty" tf:"filter,omitempty"`
	Name   string `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	UniqueWriterIdentity bool `json:"uniqueWriterIdentity,omitempty" tf:"unique_writer_identity,omitempty"`
	// +optional
	WriterIdentity string `json:"writerIdentity,omitempty" tf:"writer_identity,omitempty"`
}

type LoggingProjectSinkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LoggingProjectSinkSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
