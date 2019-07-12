package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleKmsKeyRingIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleKmsKeyRingIamMemberSpec   `json:"spec,omitempty"`
	Status            GoogleKmsKeyRingIamMemberStatus `json:"status,omitempty"`
}

type GoogleKmsKeyRingIamMemberSpec struct {
	KeyRingId string `json:"key_ring_id"`
	Role      string `json:"role"`
	Member    string `json:"member"`
	Etag      string `json:"etag"`
}

type GoogleKmsKeyRingIamMemberStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleKmsKeyRingIamMemberList is a list of GoogleKmsKeyRingIamMembers
type GoogleKmsKeyRingIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleKmsKeyRingIamMember CRD objects
	Items []GoogleKmsKeyRingIamMember `json:"items,omitempty"`
}
