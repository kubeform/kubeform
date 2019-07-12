package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleSourcerepoRepository struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSourcerepoRepositorySpec   `json:"spec,omitempty"`
	Status            GoogleSourcerepoRepositoryStatus `json:"status,omitempty"`
}

type GoogleSourcerepoRepositorySpec struct {
	Size    int    `json:"size"`
	Url     string `json:"url"`
	Name    string `json:"name"`
	Project string `json:"project"`
}



type GoogleSourcerepoRepositoryStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleSourcerepoRepositoryList is a list of GoogleSourcerepoRepositorys
type GoogleSourcerepoRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSourcerepoRepository CRD objects
	Items []GoogleSourcerepoRepository `json:"items,omitempty"`
}