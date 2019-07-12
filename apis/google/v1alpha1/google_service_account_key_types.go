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

type GoogleServiceAccountKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleServiceAccountKeySpec   `json:"spec,omitempty"`
	Status            GoogleServiceAccountKeyStatus `json:"status,omitempty"`
}

type GoogleServiceAccountKeySpec struct {
	ServiceAccountId      string `json:"service_account_id"`
	PgpKey                string `json:"pgp_key"`
	PrivateKey            string `json:"private_key"`
	PrivateKeyFingerprint string `json:"private_key_fingerprint"`
	KeyAlgorithm          string `json:"key_algorithm"`
	PrivateKeyType        string `json:"private_key_type"`
	PublicKeyType         string `json:"public_key_type"`
	Name                  string `json:"name"`
	PublicKey             string `json:"public_key"`
	ValidAfter            string `json:"valid_after"`
	ValidBefore           string `json:"valid_before"`
	PrivateKeyEncrypted   string `json:"private_key_encrypted"`
}

type GoogleServiceAccountKeyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleServiceAccountKeyList is a list of GoogleServiceAccountKeys
type GoogleServiceAccountKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleServiceAccountKey CRD objects
	Items []GoogleServiceAccountKey `json:"items,omitempty"`
}
