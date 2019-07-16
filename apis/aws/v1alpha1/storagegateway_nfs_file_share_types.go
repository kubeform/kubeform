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

type StoragegatewayNfsFileShare struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoragegatewayNfsFileShareSpec   `json:"spec,omitempty"`
	Status            StoragegatewayNfsFileShareStatus `json:"status,omitempty"`
}

type StoragegatewayNfsFileShareSpecNfsFileShareDefaults struct {
	// +optional
	DirectoryMode string `json:"directory_mode,omitempty"`
	// +optional
	FileMode string `json:"file_mode,omitempty"`
	// +optional
	GroupId int `json:"group_id,omitempty"`
	// +optional
	OwnerId int `json:"owner_id,omitempty"`
}

type StoragegatewayNfsFileShareSpec struct {
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	ClientList []string `json:"client_list"`
	// +optional
	DefaultStorageClass string `json:"default_storage_class,omitempty"`
	GatewayArn          string `json:"gateway_arn"`
	// +optional
	GuessMimeTypeEnabled bool `json:"guess_mime_type_enabled,omitempty"`
	// +optional
	KmsEncrypted bool `json:"kms_encrypted,omitempty"`
	// +optional
	KmsKeyArn   string `json:"kms_key_arn,omitempty"`
	LocationArn string `json:"location_arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NfsFileShareDefaults *[]StoragegatewayNfsFileShareSpec `json:"nfs_file_share_defaults,omitempty"`
	// +optional
	ObjectAcl string `json:"object_acl,omitempty"`
	// +optional
	ReadOnly bool `json:"read_only,omitempty"`
	// +optional
	RequesterPays bool   `json:"requester_pays,omitempty"`
	RoleArn       string `json:"role_arn"`
	// +optional
	Squash string `json:"squash,omitempty"`
}

type StoragegatewayNfsFileShareStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
