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

type ComposerEnvironment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComposerEnvironmentSpec   `json:"spec,omitempty"`
	Status            ComposerEnvironmentStatus `json:"status,omitempty"`
}

type ComposerEnvironmentSpec struct {
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	Name   string            `json:"name"`
	// +optional
	Region string `json:"region,omitempty"`
}

type ComposerEnvironmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComposerEnvironmentList is a list of ComposerEnvironments
type ComposerEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComposerEnvironment CRD objects
	Items []ComposerEnvironment `json:"items,omitempty"`
}
