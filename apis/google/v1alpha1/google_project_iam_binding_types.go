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

type GoogleProjectIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleProjectIamBindingSpec   `json:"spec,omitempty"`
	Status            GoogleProjectIamBindingStatus `json:"status,omitempty"`
}

type GoogleProjectIamBindingSpec struct {
	Members []string `json:"members"`
	Etag    string   `json:"etag"`
	Project string   `json:"project"`
	Role    string   `json:"role"`
}

type GoogleProjectIamBindingStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleProjectIamBindingList is a list of GoogleProjectIamBindings
type GoogleProjectIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleProjectIamBinding CRD objects
	Items []GoogleProjectIamBinding `json:"items,omitempty"`
}
