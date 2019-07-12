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

type GoogleProjectIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleProjectIamMemberSpec   `json:"spec,omitempty"`
	Status            GoogleProjectIamMemberStatus `json:"status,omitempty"`
}

type GoogleProjectIamMemberSpec struct {
	Role    string `json:"role"`
	Member  string `json:"member"`
	Etag    string `json:"etag"`
	Project string `json:"project"`
}

type GoogleProjectIamMemberStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleProjectIamMemberList is a list of GoogleProjectIamMembers
type GoogleProjectIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleProjectIamMember CRD objects
	Items []GoogleProjectIamMember `json:"items,omitempty"`
}
