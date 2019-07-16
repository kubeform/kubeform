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

type StoragegatewaySmbFileShare struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoragegatewaySmbFileShareSpec   `json:"spec,omitempty"`
	Status            StoragegatewaySmbFileShareStatus `json:"status,omitempty"`
}

type StoragegatewaySmbFileShareSpec struct {
	// +optional
	Authentication string `json:"authentication,omitempty"`
	// +optional
	DefaultStorageClass string `json:"default_storage_class,omitempty"`
	GatewayArn          string `json:"gateway_arn"`
	// +optional
	GuessMimeTypeEnabled bool `json:"guess_mime_type_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	InvalidUserList []string `json:"invalid_user_list,omitempty"`
	// +optional
	KmsEncrypted bool `json:"kms_encrypted,omitempty"`
	// +optional
	KmsKeyArn   string `json:"kms_key_arn,omitempty"`
	LocationArn string `json:"location_arn"`
	// +optional
	ObjectAcl string `json:"object_acl,omitempty"`
	// +optional
	ReadOnly bool `json:"read_only,omitempty"`
	// +optional
	RequesterPays bool   `json:"requester_pays,omitempty"`
	RoleArn       string `json:"role_arn"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ValidUserList []string `json:"valid_user_list,omitempty"`
}

type StoragegatewaySmbFileShareStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StoragegatewaySmbFileShareList is a list of StoragegatewaySmbFileShares
type StoragegatewaySmbFileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StoragegatewaySmbFileShare CRD objects
	Items []StoragegatewaySmbFileShare `json:"items,omitempty"`
}
