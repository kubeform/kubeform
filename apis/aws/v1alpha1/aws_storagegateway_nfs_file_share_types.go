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

type AwsStoragegatewayNfsFileShare struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsStoragegatewayNfsFileShareSpec   `json:"spec,omitempty"`
	Status            AwsStoragegatewayNfsFileShareStatus `json:"status,omitempty"`
}

type AwsStoragegatewayNfsFileShareSpecNfsFileShareDefaults struct {
	DirectoryMode string `json:"directory_mode"`
	FileMode      string `json:"file_mode"`
	GroupId       int    `json:"group_id"`
	OwnerId       int    `json:"owner_id"`
}

type AwsStoragegatewayNfsFileShareSpec struct {
	DefaultStorageClass  string                              `json:"default_storage_class"`
	Squash               string                              `json:"squash"`
	Arn                  string                              `json:"arn"`
	ClientList           []string                            `json:"client_list"`
	NfsFileShareDefaults []AwsStoragegatewayNfsFileShareSpec `json:"nfs_file_share_defaults"`
	ReadOnly             bool                                `json:"read_only"`
	GuessMimeTypeEnabled bool                                `json:"guess_mime_type_enabled"`
	LocationArn          string                              `json:"location_arn"`
	KmsEncrypted         bool                                `json:"kms_encrypted"`
	KmsKeyArn            string                              `json:"kms_key_arn"`
	ObjectAcl            string                              `json:"object_acl"`
	RoleArn              string                              `json:"role_arn"`
	FileshareId          string                              `json:"fileshare_id"`
	GatewayArn           string                              `json:"gateway_arn"`
	RequesterPays        bool                                `json:"requester_pays"`
}



type AwsStoragegatewayNfsFileShareStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsStoragegatewayNfsFileShareList is a list of AwsStoragegatewayNfsFileShares
type AwsStoragegatewayNfsFileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsStoragegatewayNfsFileShare CRD objects
	Items []AwsStoragegatewayNfsFileShare `json:"items,omitempty"`
}