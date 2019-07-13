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

type AwsKmsCiphertext struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsKmsCiphertextSpec   `json:"spec,omitempty"`
	Status            AwsKmsCiphertextStatus `json:"status,omitempty"`
}

type AwsKmsCiphertextSpec struct {
	Context        map[string]string `json:"context"`
	CiphertextBlob string            `json:"ciphertext_blob"`
	Plaintext      string            `json:"plaintext"`
	KeyId          string            `json:"key_id"`
}



type AwsKmsCiphertextStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsKmsCiphertextList is a list of AwsKmsCiphertexts
type AwsKmsCiphertextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsKmsCiphertext CRD objects
	Items []AwsKmsCiphertext `json:"items,omitempty"`
}