package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecretsmanagerSecretVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSecretsmanagerSecretVersionSpec   `json:"spec,omitempty"`
	Status            AwsSecretsmanagerSecretVersionStatus `json:"status,omitempty"`
}

type AwsSecretsmanagerSecretVersionSpec struct {
	Arn           string   `json:"arn"`
	SecretId      string   `json:"secret_id"`
	SecretString  string   `json:"secret_string"`
	SecretBinary  string   `json:"secret_binary"`
	VersionId     string   `json:"version_id"`
	VersionStages []string `json:"version_stages"`
}



type AwsSecretsmanagerSecretVersionStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecretsmanagerSecretVersionList is a list of AwsSecretsmanagerSecretVersions
type AwsSecretsmanagerSecretVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSecretsmanagerSecretVersion CRD objects
	Items []AwsSecretsmanagerSecretVersion `json:"items,omitempty"`
}