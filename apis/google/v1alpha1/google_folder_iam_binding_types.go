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

type GoogleFolderIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleFolderIamBindingSpec   `json:"spec,omitempty"`
	Status            GoogleFolderIamBindingStatus `json:"status,omitempty"`
}

type GoogleFolderIamBindingSpec struct {
	Role    string   `json:"role"`
	Members []string `json:"members"`
	Etag    string   `json:"etag"`
	Folder  string   `json:"folder"`
}

type GoogleFolderIamBindingStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleFolderIamBindingList is a list of GoogleFolderIamBindings
type GoogleFolderIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleFolderIamBinding CRD objects
	Items []GoogleFolderIamBinding `json:"items,omitempty"`
}
