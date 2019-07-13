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

type GoogleKmsKeyRingIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleKmsKeyRingIamMemberSpec   `json:"spec,omitempty"`
	Status            GoogleKmsKeyRingIamMemberStatus `json:"status,omitempty"`
}

type GoogleKmsKeyRingIamMemberSpec struct {
	Role      string `json:"role"`
	Member    string `json:"member"`
	Etag      string `json:"etag"`
	KeyRingId string `json:"key_ring_id"`
}



type GoogleKmsKeyRingIamMemberStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleKmsKeyRingIamMemberList is a list of GoogleKmsKeyRingIamMembers
type GoogleKmsKeyRingIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleKmsKeyRingIamMember CRD objects
	Items []GoogleKmsKeyRingIamMember `json:"items,omitempty"`
}