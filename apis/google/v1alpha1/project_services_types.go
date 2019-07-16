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

type ProjectServices struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectServicesSpec   `json:"spec,omitempty"`
	Status            ProjectServicesStatus `json:"status,omitempty"`
}

type ProjectServicesSpec struct {
	// +optional
	DisableOnDestroy bool `json:"disable_on_destroy,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	Services []string `json:"services"`
}

type ProjectServicesStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ProjectServicesList is a list of ProjectServicess
type ProjectServicesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ProjectServices CRD objects
	Items []ProjectServices `json:"items,omitempty"`
}
