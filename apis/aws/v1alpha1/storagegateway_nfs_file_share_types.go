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

type StoragegatewayNfsFileShare struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoragegatewayNfsFileShareSpec   `json:"spec,omitempty"`
	Status            StoragegatewayNfsFileShareStatus `json:"status,omitempty"`
}

type StoragegatewayNfsFileShareSpecNfsFileShareDefaults struct {
	// +optional
	DirectoryMode string `json:"directoryMode,omitempty" tf:"directory_mode,omitempty"`
	// +optional
	FileMode string `json:"fileMode,omitempty" tf:"file_mode,omitempty"`
	// +optional
	GroupID int `json:"groupID,omitempty" tf:"group_id,omitempty"`
	// +optional
	OwnerID int `json:"ownerID,omitempty" tf:"owner_id,omitempty"`
}

type StoragegatewayNfsFileShareSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:MinItems=1
	ClientList []string `json:"clientList" tf:"client_list"`
	// +optional
	DefaultStorageClass string `json:"defaultStorageClass,omitempty" tf:"default_storage_class,omitempty"`
	// +optional
	FileshareID string `json:"fileshareID,omitempty" tf:"fileshare_id,omitempty"`
	GatewayArn  string `json:"gatewayArn" tf:"gateway_arn"`
	// +optional
	GuessMimeTypeEnabled bool `json:"guessMimeTypeEnabled,omitempty" tf:"guess_mime_type_enabled,omitempty"`
	// +optional
	KmsEncrypted bool `json:"kmsEncrypted,omitempty" tf:"kms_encrypted,omitempty"`
	// +optional
	KmsKeyArn   string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
	LocationArn string `json:"locationArn" tf:"location_arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NfsFileShareDefaults []StoragegatewayNfsFileShareSpecNfsFileShareDefaults `json:"nfsFileShareDefaults,omitempty" tf:"nfs_file_share_defaults,omitempty"`
	// +optional
	ObjectACL string `json:"objectACL,omitempty" tf:"object_acl,omitempty"`
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`
	// +optional
	RequesterPays bool   `json:"requesterPays,omitempty" tf:"requester_pays,omitempty"`
	RoleArn       string `json:"roleArn" tf:"role_arn"`
	// +optional
	Squash string `json:"squash,omitempty" tf:"squash,omitempty"`
}

type StoragegatewayNfsFileShareStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StoragegatewayNfsFileShareSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StoragegatewayNfsFileShareList is a list of StoragegatewayNfsFileShares
type StoragegatewayNfsFileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StoragegatewayNfsFileShare CRD objects
	Items []StoragegatewayNfsFileShare `json:"items,omitempty"`
}
