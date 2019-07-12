package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleFolderIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleFolderIamMemberSpec   `json:"spec,omitempty"`
	Status            GoogleFolderIamMemberStatus `json:"status,omitempty"`
}

type GoogleFolderIamMemberSpec struct {
	Role   string `json:"role"`
	Folder string `json:"folder"`
	Member string `json:"member"`
	Etag   string `json:"etag"`
}

type GoogleFolderIamMemberStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleFolderIamMemberList is a list of GoogleFolderIamMembers
type GoogleFolderIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleFolderIamMember CRD objects
	Items []GoogleFolderIamMember `json:"items,omitempty"`
}
