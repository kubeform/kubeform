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

type GoogleSourcerepoRepository struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSourcerepoRepositorySpec   `json:"spec,omitempty"`
	Status            GoogleSourcerepoRepositoryStatus `json:"status,omitempty"`
}

type GoogleSourcerepoRepositorySpec struct {
	Name    string `json:"name"`
	Project string `json:"project"`
	Size    int    `json:"size"`
	Url     string `json:"url"`
}



type GoogleSourcerepoRepositoryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleSourcerepoRepositoryList is a list of GoogleSourcerepoRepositorys
type GoogleSourcerepoRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSourcerepoRepository CRD objects
	Items []GoogleSourcerepoRepository `json:"items,omitempty"`
}