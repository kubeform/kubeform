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

type GoogleLoggingFolderSink struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleLoggingFolderSinkSpec   `json:"spec,omitempty"`
	Status            GoogleLoggingFolderSinkStatus `json:"status,omitempty"`
}

type GoogleLoggingFolderSinkSpec struct {
	WriterIdentity  string `json:"writer_identity"`
	Folder          string `json:"folder"`
	IncludeChildren bool   `json:"include_children"`
	Name            string `json:"name"`
	Destination     string `json:"destination"`
	Filter          string `json:"filter"`
}



type GoogleLoggingFolderSinkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleLoggingFolderSinkList is a list of GoogleLoggingFolderSinks
type GoogleLoggingFolderSinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleLoggingFolderSink CRD objects
	Items []GoogleLoggingFolderSink `json:"items,omitempty"`
}