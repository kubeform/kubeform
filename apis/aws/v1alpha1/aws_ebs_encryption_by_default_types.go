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

type AwsEbsEncryptionByDefault struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEbsEncryptionByDefaultSpec   `json:"spec,omitempty"`
	Status            AwsEbsEncryptionByDefaultStatus `json:"status,omitempty"`
}

type AwsEbsEncryptionByDefaultSpec struct {
	Enabled bool `json:"enabled"`
}

type AwsEbsEncryptionByDefaultStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEbsEncryptionByDefaultList is a list of AwsEbsEncryptionByDefaults
type AwsEbsEncryptionByDefaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEbsEncryptionByDefault CRD objects
	Items []AwsEbsEncryptionByDefault `json:"items,omitempty"`
}
