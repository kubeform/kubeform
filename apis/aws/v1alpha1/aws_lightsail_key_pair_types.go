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

type AwsLightsailKeyPair struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLightsailKeyPairSpec   `json:"spec,omitempty"`
	Status            AwsLightsailKeyPairStatus `json:"status,omitempty"`
}

type AwsLightsailKeyPairSpec struct {
	Name                 string `json:"name"`
	PublicKey            string `json:"public_key"`
	PrivateKey           string `json:"private_key"`
	EncryptedFingerprint string `json:"encrypted_fingerprint"`
	EncryptedPrivateKey  string `json:"encrypted_private_key"`
	NamePrefix           string `json:"name_prefix"`
	PgpKey               string `json:"pgp_key"`
	Arn                  string `json:"arn"`
	Fingerprint          string `json:"fingerprint"`
}



type AwsLightsailKeyPairStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsLightsailKeyPairList is a list of AwsLightsailKeyPairs
type AwsLightsailKeyPairList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLightsailKeyPair CRD objects
	Items []AwsLightsailKeyPair `json:"items,omitempty"`
}