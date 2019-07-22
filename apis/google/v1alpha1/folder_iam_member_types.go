package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type FolderIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FolderIamMemberSpec   `json:"spec,omitempty"`
	Status            FolderIamMemberStatus `json:"status,omitempty"`
}

type FolderIamMemberSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	Folder      string                    `json:"folder" tf:"folder"`
	Member      string                    `json:"member" tf:"member"`
	Role        string                    `json:"role" tf:"role"`
}

type FolderIamMemberStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FolderIamMemberList is a list of FolderIamMembers
type FolderIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FolderIamMember CRD objects
	Items []FolderIamMember `json:"items,omitempty"`
}
