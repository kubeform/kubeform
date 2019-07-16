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

type LoggingFolderSink struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoggingFolderSinkSpec   `json:"spec,omitempty"`
	Status            LoggingFolderSinkStatus `json:"status,omitempty"`
}

type LoggingFolderSinkSpec struct {
	Destination string `json:"destination"`
	// +optional
	Filter string `json:"filter,omitempty"`
	Folder string `json:"folder"`
	// +optional
	IncludeChildren bool   `json:"include_children,omitempty"`
	Name            string `json:"name"`
}

type LoggingFolderSinkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LoggingFolderSinkList is a list of LoggingFolderSinks
type LoggingFolderSinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LoggingFolderSink CRD objects
	Items []LoggingFolderSink `json:"items,omitempty"`
}
