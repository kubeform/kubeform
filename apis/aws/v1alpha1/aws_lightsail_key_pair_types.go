package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailKeyPair struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLightsailKeyPairSpec   `json:"spec,omitempty"`
	Status            AwsLightsailKeyPairStatus `json:"status,omitempty"`
}

type AwsLightsailKeyPairSpec struct {
	EncryptedFingerprint string `json:"encrypted_fingerprint"`
	Name                 string `json:"name"`
	Arn                  string `json:"arn"`
	Fingerprint          string `json:"fingerprint"`
	PublicKey            string `json:"public_key"`
	PrivateKey           string `json:"private_key"`
	EncryptedPrivateKey  string `json:"encrypted_private_key"`
	NamePrefix           string `json:"name_prefix"`
	PgpKey               string `json:"pgp_key"`
}

type AwsLightsailKeyPairStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLightsailKeyPairList is a list of AwsLightsailKeyPairs
type AwsLightsailKeyPairList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLightsailKeyPair CRD objects
	Items []AwsLightsailKeyPair `json:"items,omitempty"`
}
