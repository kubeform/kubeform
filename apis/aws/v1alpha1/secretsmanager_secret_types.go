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

type SecretsmanagerSecret struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretsmanagerSecretSpec   `json:"spec,omitempty"`
	Status            SecretsmanagerSecretStatus `json:"status,omitempty"`
}

type SecretsmanagerSecretSpecRotationRules struct {
	AutomaticallyAfterDays int `json:"automatically_after_days"`
}

type SecretsmanagerSecretSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	KmsKeyId string `json:"kms_key_id,omitempty"`
	// +optional
	Policy string `json:"policy,omitempty"`
	// +optional
	RecoveryWindowInDays int `json:"recovery_window_in_days,omitempty"`
	// +optional
	RotationLambdaArn string `json:"rotation_lambda_arn,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RotationRules *[]SecretsmanagerSecretSpec `json:"rotation_rules,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type SecretsmanagerSecretStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SecretsmanagerSecretList is a list of SecretsmanagerSecrets
type SecretsmanagerSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SecretsmanagerSecret CRD objects
	Items []SecretsmanagerSecret `json:"items,omitempty"`
}
