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

type StorageShare struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageShareSpec   `json:"spec,omitempty"`
	Status            StorageShareStatus `json:"status,omitempty"`
}

type StorageShareSpecAclAccessPolicy struct {
	Expiry      string `json:"expiry"`
	Permissions string `json:"permissions"`
	Start       string `json:"start"`
}

type StorageShareSpecAcl struct {
	// +optional
	AccessPolicy *[]StorageShareSpecAcl `json:"access_policy,omitempty"`
	Id           string                 `json:"id"`
}

type StorageShareSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Acl *[]StorageShareSpec `json:"acl,omitempty"`
	// +optional
	Metadata map[string]string `json:"metadata,omitempty"`
	Name     string            `json:"name"`
	// +optional
	Quota              int    `json:"quota,omitempty"`
	StorageAccountName string `json:"storage_account_name"`
}

type StorageShareStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
