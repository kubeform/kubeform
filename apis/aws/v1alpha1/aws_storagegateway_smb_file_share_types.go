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
	GatewayArn           string   `json:"gateway_arn"`
	InvalidUserList      []string `json:"invalid_user_list"`
	Authentication       string   `json:"authentication"`
	GuessMimeTypeEnabled bool     `json:"guess_mime_type_enabled"`
	KmsKeyArn            string   `json:"kms_key_arn"`
	RequesterPays        bool     `json:"requester_pays"`
	ValidUserList        []string `json:"valid_user_list"`
	Arn                  string   `json:"arn"`
	DefaultStorageClass  string   `json:"default_storage_class"`
	FileshareId          string   `json:"fileshare_id"`
	ObjectAcl            string   `json:"object_acl"`
	KmsEncrypted         bool     `json:"kms_encrypted"`
	LocationArn          string   `json:"location_arn"`
	ReadOnly             bool     `json:"read_only"`
	RoleArn              string   `json:"role_arn"`
}



type AwsStoragegatewaySmbFileShareStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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