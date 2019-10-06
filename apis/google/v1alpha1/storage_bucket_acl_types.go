package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type StorageBucketACL struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageBucketACLSpec   `json:"spec,omitempty"`
	Status            StorageBucketACLStatus `json:"status,omitempty"`
}

type StorageBucketACLSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Bucket string `json:"bucket" tf:"bucket"`
	// +optional
	DefaultACL string `json:"defaultACL,omitempty" tf:"default_acl,omitempty"`
	// +optional
	PredefinedACL string `json:"predefinedACL,omitempty" tf:"predefined_acl,omitempty"`
	// +optional
	RoleEntity []string `json:"roleEntity,omitempty" tf:"role_entity,omitempty"`
}

type StorageBucketACLStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StorageBucketACLSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageBucketACLList is a list of StorageBucketACLs
type StorageBucketACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageBucketACL CRD objects
	Items []StorageBucketACL `json:"items,omitempty"`
}
