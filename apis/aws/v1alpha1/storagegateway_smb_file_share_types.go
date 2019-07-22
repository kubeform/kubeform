package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	Authentication string `json:"authentication,omitempty" tf:"authentication,omitempty"`
	// +optional
	DefaultStorageClass string `json:"defaultStorageClass,omitempty" tf:"default_storage_class,omitempty"`
	GatewayArn          string `json:"gatewayArn" tf:"gateway_arn"`
	// +optional
	GuessMimeTypeEnabled bool `json:"guessMimeTypeEnabled,omitempty" tf:"guess_mime_type_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	InvalidUserList []string `json:"invalidUserList,omitempty" tf:"invalid_user_list,omitempty"`
	// +optional
	KmsEncrypted bool `json:"kmsEncrypted,omitempty" tf:"kms_encrypted,omitempty"`
	// +optional
	KmsKeyArn   string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
	LocationArn string `json:"locationArn" tf:"location_arn"`
	// +optional
	ObjectACL string `json:"objectACL,omitempty" tf:"object_acl,omitempty"`
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`
	// +optional
	RequesterPays bool   `json:"requesterPays,omitempty" tf:"requester_pays,omitempty"`
	RoleArn       string `json:"roleArn" tf:"role_arn"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ValidUserList []string `json:"validUserList,omitempty" tf:"valid_user_list,omitempty"`
}

type StoragegatewaySmbFileShareStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
