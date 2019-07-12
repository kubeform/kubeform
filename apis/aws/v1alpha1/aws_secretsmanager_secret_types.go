package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecretsmanagerSecret struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSecretsmanagerSecretSpec   `json:"spec,omitempty"`
	Status            AwsSecretsmanagerSecretStatus `json:"status,omitempty"`
}

type AwsSecretsmanagerSecretSpecRotationRules struct {
	AutomaticallyAfterDays int `json:"automatically_after_days"`
}

type AwsSecretsmanagerSecretSpec struct {
	Name                 string                        `json:"name"`
	NamePrefix           string                        `json:"name_prefix"`
	Arn                  string                        `json:"arn"`
	Description          string                        `json:"description"`
	KmsKeyId             string                        `json:"kms_key_id"`
	RotationLambdaArn    string                        `json:"rotation_lambda_arn"`
	RotationRules        []AwsSecretsmanagerSecretSpec `json:"rotation_rules"`
	Tags                 map[string]string             `json:"tags"`
	Policy               string                        `json:"policy"`
	RecoveryWindowInDays int                           `json:"recovery_window_in_days"`
	RotationEnabled      bool                          `json:"rotation_enabled"`
}

type AwsSecretsmanagerSecretStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecretsmanagerSecretList is a list of AwsSecretsmanagerSecrets
type AwsSecretsmanagerSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSecretsmanagerSecret CRD objects
	Items []AwsSecretsmanagerSecret `json:"items,omitempty"`
}
