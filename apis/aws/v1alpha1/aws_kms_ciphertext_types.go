package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKmsCiphertext struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsKmsCiphertextSpec   `json:"spec,omitempty"`
	Status            AwsKmsCiphertextStatus `json:"status,omitempty"`
}

type AwsKmsCiphertextSpec struct {
	Plaintext      string            `json:"plaintext"`
	KeyId          string            `json:"key_id"`
	Context        map[string]string `json:"context"`
	CiphertextBlob string            `json:"ciphertext_blob"`
}

type AwsKmsCiphertextStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKmsCiphertextList is a list of AwsKmsCiphertexts
type AwsKmsCiphertextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsKmsCiphertext CRD objects
	Items []AwsKmsCiphertext `json:"items,omitempty"`
}
