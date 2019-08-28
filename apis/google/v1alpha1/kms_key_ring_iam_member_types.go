package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type KmsKeyRingIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KmsKeyRingIamMemberSpec   `json:"spec,omitempty"`
	Status            KmsKeyRingIamMemberStatus `json:"status,omitempty"`
}

type KmsKeyRingIamMemberSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Etag      string `json:"etag,omitempty" tf:"etag,omitempty"`
	KeyRingID string `json:"keyRingID" tf:"key_ring_id"`
	Member    string `json:"member" tf:"member"`
	Role      string `json:"role" tf:"role"`
}



type KmsKeyRingIamMemberStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *KmsKeyRingIamMemberSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KmsKeyRingIamMemberList is a list of KmsKeyRingIamMembers
type KmsKeyRingIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KmsKeyRingIamMember CRD objects
	Items []KmsKeyRingIamMember `json:"items,omitempty"`
}