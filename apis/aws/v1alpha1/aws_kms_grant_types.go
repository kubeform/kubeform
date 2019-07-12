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

type AwsKmsGrant struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsKmsGrantSpec   `json:"spec,omitempty"`
	Status            AwsKmsGrantStatus `json:"status,omitempty"`
}

type AwsKmsGrantSpecConstraints struct {
	EncryptionContextEquals map[string]string `json:"encryption_context_equals"`
	EncryptionContextSubset map[string]string `json:"encryption_context_subset"`
}

type AwsKmsGrantSpec struct {
	GrantId             string            `json:"grant_id"`
	KeyId               string            `json:"key_id"`
	GranteePrincipal    string            `json:"grantee_principal"`
	RetiringPrincipal   string            `json:"retiring_principal"`
	RetireOnDelete      bool              `json:"retire_on_delete"`
	GrantToken          string            `json:"grant_token"`
	Name                string            `json:"name"`
	Operations          []string          `json:"operations"`
	Constraints         []AwsKmsGrantSpec `json:"constraints"`
	GrantCreationTokens []string          `json:"grant_creation_tokens"`
}

type AwsKmsGrantStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsKmsGrantList is a list of AwsKmsGrants
type AwsKmsGrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsKmsGrant CRD objects
	Items []AwsKmsGrant `json:"items,omitempty"`
}
