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
// +kubebuilder:subresource:status

type StorageShare struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageShareSpec   `json:"spec,omitempty"`
	Status            StorageShareStatus `json:"status,omitempty"`
}

type StorageShareSpecAclAccessPolicy struct {
	Expiry      string `json:"expiry" tf:"expiry"`
	Permissions string `json:"permissions" tf:"permissions"`
	Start       string `json:"start" tf:"start"`
}

type StorageShareSpecAcl struct {
	// +optional
	AccessPolicy []StorageShareSpecAclAccessPolicy `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`
	ID           string                            `json:"ID" tf:"id"`
}

type StorageShareSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Acl []StorageShareSpecAcl `json:"acl,omitempty" tf:"acl,omitempty"`
	// +optional
	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata,omitempty"`
	Name     string            `json:"name" tf:"name"`
	// +optional
	Quota int `json:"quota,omitempty" tf:"quota,omitempty"`
	// +optional
	// Deprecated
	ResourceGroupName  string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
	StorageAccountName string `json:"storageAccountName" tf:"storage_account_name"`
	// +optional
	Url string `json:"url,omitempty" tf:"url,omitempty"`
}

type StorageShareStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StorageShareSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageShareList is a list of StorageShares
type StorageShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageShare CRD objects
	Items []StorageShare `json:"items,omitempty"`
}
