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

type AwsIamAccessKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamAccessKeySpec   `json:"spec,omitempty"`
	Status            AwsIamAccessKeyStatus `json:"status,omitempty"`
}

type AwsIamAccessKeySpec struct {
	PgpKey          string `json:"pgp_key"`
	KeyFingerprint  string `json:"key_fingerprint"`
	EncryptedSecret string `json:"encrypted_secret"`
	User            string `json:"user"`
	Status          string `json:"status"`
	Secret          string `json:"secret"`
	SesSmtpPassword string `json:"ses_smtp_password"`
}



type AwsIamAccessKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIamAccessKeyList is a list of AwsIamAccessKeys
type AwsIamAccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamAccessKey CRD objects
	Items []AwsIamAccessKey `json:"items,omitempty"`
}