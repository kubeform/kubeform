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

type GoogleStorageBucketIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleStorageBucketIamMemberSpec   `json:"spec,omitempty"`
	Status            GoogleStorageBucketIamMemberStatus `json:"status,omitempty"`
}

type GoogleStorageBucketIamMemberSpec struct {
	Role   string `json:"role"`
	Member string `json:"member"`
	Etag   string `json:"etag"`
	Bucket string `json:"bucket"`
}

type GoogleStorageBucketIamMemberStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleStorageBucketIamMemberList is a list of GoogleStorageBucketIamMembers
type GoogleStorageBucketIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleStorageBucketIamMember CRD objects
	Items []GoogleStorageBucketIamMember `json:"items,omitempty"`
}
