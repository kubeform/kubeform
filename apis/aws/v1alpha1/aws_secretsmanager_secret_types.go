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
	RecoveryWindowInDays int                           `json:"recovery_window_in_days"`
	RotationLambdaArn    string                        `json:"rotation_lambda_arn"`
	RotationRules        []AwsSecretsmanagerSecretSpec `json:"rotation_rules"`
	Tags                 map[string]string             `json:"tags"`
	Arn                  string                        `json:"arn"`
	KmsKeyId             string                        `json:"kms_key_id"`
	NamePrefix           string                        `json:"name_prefix"`
	Policy               string                        `json:"policy"`
	RotationEnabled      bool                          `json:"rotation_enabled"`
	Description          string                        `json:"description"`
}



type AwsSecretsmanagerSecretStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSecretsmanagerSecretList is a list of AwsSecretsmanagerSecrets
type AwsSecretsmanagerSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSecretsmanagerSecret CRD objects
	Items []AwsSecretsmanagerSecret `json:"items,omitempty"`
}