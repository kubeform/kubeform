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

type AwsKmsExternalKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsKmsExternalKeySpec   `json:"spec,omitempty"`
	Status            AwsKmsExternalKeyStatus `json:"status,omitempty"`
}

type AwsKmsExternalKeySpec struct {
	DeletionWindowInDays int               `json:"deletion_window_in_days"`
	Description          string            `json:"description"`
	ExpirationModel      string            `json:"expiration_model"`
	KeyMaterialBase64    string            `json:"key_material_base64"`
	KeyState             string            `json:"key_state"`
	Arn                  string            `json:"arn"`
	Enabled              bool              `json:"enabled"`
	KeyUsage             string            `json:"key_usage"`
	Policy               string            `json:"policy"`
	Tags                 map[string]string `json:"tags"`
	ValidTo              string            `json:"valid_to"`
}



type AwsKmsExternalKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsKmsExternalKeyList is a list of AwsKmsExternalKeys
type AwsKmsExternalKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsKmsExternalKey CRD objects
	Items []AwsKmsExternalKey `json:"items,omitempty"`
}