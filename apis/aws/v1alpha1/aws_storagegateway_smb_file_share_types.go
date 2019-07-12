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

type AwsStoragegatewaySmbFileShare struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsStoragegatewaySmbFileShareSpec   `json:"spec,omitempty"`
	Status            AwsStoragegatewaySmbFileShareStatus `json:"status,omitempty"`
}

type AwsStoragegatewaySmbFileShareSpec struct {
	KmsEncrypted         bool     `json:"kms_encrypted"`
	ObjectAcl            string   `json:"object_acl"`
	RoleArn              string   `json:"role_arn"`
	DefaultStorageClass  string   `json:"default_storage_class"`
	InvalidUserList      []string `json:"invalid_user_list"`
	Authentication       string   `json:"authentication"`
	GatewayArn           string   `json:"gateway_arn"`
	GuessMimeTypeEnabled bool     `json:"guess_mime_type_enabled"`
	ReadOnly             bool     `json:"read_only"`
	RequesterPays        bool     `json:"requester_pays"`
	ValidUserList        []string `json:"valid_user_list"`
	Arn                  string   `json:"arn"`
	FileshareId          string   `json:"fileshare_id"`
	KmsKeyArn            string   `json:"kms_key_arn"`
	LocationArn          string   `json:"location_arn"`
}

type AwsStoragegatewaySmbFileShareStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsStoragegatewaySmbFileShareList is a list of AwsStoragegatewaySmbFileShares
type AwsStoragegatewaySmbFileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsStoragegatewaySmbFileShare CRD objects
	Items []AwsStoragegatewaySmbFileShare `json:"items,omitempty"`
}
