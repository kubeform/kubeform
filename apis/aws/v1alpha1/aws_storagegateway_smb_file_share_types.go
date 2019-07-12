package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsStoragegatewaySmbFileShare struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsStoragegatewaySmbFileShareSpec   `json:"spec,omitempty"`
	Status            AwsStoragegatewaySmbFileShareStatus `json:"status,omitempty"`
}

type AwsStoragegatewaySmbFileShareSpec struct {
	Arn                  string   `json:"arn"`
	DefaultStorageClass  string   `json:"default_storage_class"`
	KmsEncrypted         bool     `json:"kms_encrypted"`
	ReadOnly             bool     `json:"read_only"`
	RequesterPays        bool     `json:"requester_pays"`
	FileshareId          string   `json:"fileshare_id"`
	GuessMimeTypeEnabled bool     `json:"guess_mime_type_enabled"`
	InvalidUserList      []string `json:"invalid_user_list"`
	LocationArn          string   `json:"location_arn"`
	ObjectAcl            string   `json:"object_acl"`
	KmsKeyArn            string   `json:"kms_key_arn"`
	RoleArn              string   `json:"role_arn"`
	Authentication       string   `json:"authentication"`
	GatewayArn           string   `json:"gateway_arn"`
	ValidUserList        []string `json:"valid_user_list"`
}

type AwsStoragegatewaySmbFileShareStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsStoragegatewaySmbFileShareList is a list of AwsStoragegatewaySmbFileShares
type AwsStoragegatewaySmbFileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsStoragegatewaySmbFileShare CRD objects
	Items []AwsStoragegatewaySmbFileShare `json:"items,omitempty"`
}
