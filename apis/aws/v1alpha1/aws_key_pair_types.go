package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKeyPair struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsKeyPairSpec   `json:"spec,omitempty"`
	Status            AwsKeyPairStatus `json:"status,omitempty"`
}

type AwsKeyPairSpec struct {
	KeyName       string `json:"key_name"`
	KeyNamePrefix string `json:"key_name_prefix"`
	PublicKey     string `json:"public_key"`
	Fingerprint   string `json:"fingerprint"`
}



type AwsKeyPairStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKeyPairList is a list of AwsKeyPairs
type AwsKeyPairList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsKeyPair CRD objects
	Items []AwsKeyPair `json:"items,omitempty"`
}