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

type AwsEbsDefaultKmsKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEbsDefaultKmsKeySpec   `json:"spec,omitempty"`
	Status            AwsEbsDefaultKmsKeyStatus `json:"status,omitempty"`
}

type AwsEbsDefaultKmsKeySpec struct {
	KeyArn string `json:"key_arn"`
}

type AwsEbsDefaultKmsKeyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEbsDefaultKmsKeyList is a list of AwsEbsDefaultKmsKeys
type AwsEbsDefaultKmsKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEbsDefaultKmsKey CRD objects
	Items []AwsEbsDefaultKmsKey `json:"items,omitempty"`
}
