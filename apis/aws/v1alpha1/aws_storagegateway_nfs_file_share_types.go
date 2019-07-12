package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsStoragegatewayNfsFileShare struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsStoragegatewayNfsFileShareSpec   `json:"spec,omitempty"`
	Status            AwsStoragegatewayNfsFileShareStatus `json:"status,omitempty"`
}

type AwsStoragegatewayNfsFileShareSpecNfsFileShareDefaults struct {
	FileMode      string `json:"file_mode"`
	GroupId       int    `json:"group_id"`
	OwnerId       int    `json:"owner_id"`
	DirectoryMode string `json:"directory_mode"`
}

type AwsStoragegatewayNfsFileShareSpec struct {
	KmsKeyArn            string                              `json:"kms_key_arn"`
	ObjectAcl            string                              `json:"object_acl"`
	GuessMimeTypeEnabled bool                                `json:"guess_mime_type_enabled"`
	NfsFileShareDefaults []AwsStoragegatewayNfsFileShareSpec `json:"nfs_file_share_defaults"`
	RequesterPays        bool                                `json:"requester_pays"`
	RoleArn              string                              `json:"role_arn"`
	KmsEncrypted         bool                                `json:"kms_encrypted"`
	DefaultStorageClass  string                              `json:"default_storage_class"`
	LocationArn          string                              `json:"location_arn"`
	ReadOnly             bool                                `json:"read_only"`
	Arn                  string                              `json:"arn"`
	FileshareId          string                              `json:"fileshare_id"`
	GatewayArn           string                              `json:"gateway_arn"`
	Squash               string                              `json:"squash"`
	ClientList           []string                            `json:"client_list"`
}

type AwsStoragegatewayNfsFileShareStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsStoragegatewayNfsFileShareList is a list of AwsStoragegatewayNfsFileShares
type AwsStoragegatewayNfsFileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsStoragegatewayNfsFileShare CRD objects
	Items []AwsStoragegatewayNfsFileShare `json:"items,omitempty"`
}
