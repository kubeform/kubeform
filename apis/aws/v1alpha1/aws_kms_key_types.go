package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKmsKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsKmsKeySpec   `json:"spec,omitempty"`
	Status            AwsKmsKeyStatus `json:"status,omitempty"`
}

type AwsKmsKeySpec struct {
	Description          string            `json:"description"`
	IsEnabled            bool              `json:"is_enabled"`
	DeletionWindowInDays int               `json:"deletion_window_in_days"`
	EnableKeyRotation    bool              `json:"enable_key_rotation"`
	Tags                 map[string]string `json:"tags"`
	Arn                  string            `json:"arn"`
	KeyId                string            `json:"key_id"`
	KeyUsage             string            `json:"key_usage"`
	Policy               string            `json:"policy"`
}



type AwsKmsKeyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKmsKeyList is a list of AwsKmsKeys
type AwsKmsKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsKmsKey CRD objects
	Items []AwsKmsKey `json:"items,omitempty"`
}