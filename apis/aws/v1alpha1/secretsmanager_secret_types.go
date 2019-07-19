package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	AutomaticallyAfterDays int `json:"automaticallyAfterDays" tf:"automatically_after_days"`
}

type SecretsmanagerSecretSpec struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	Policy string `json:"policy,omitempty" tf:"policy,omitempty"`
	// +optional
	RecoveryWindowInDays int `json:"recoveryWindowInDays,omitempty" tf:"recovery_window_in_days,omitempty"`
	// +optional
	RotationLambdaArn string `json:"rotationLambdaArn,omitempty" tf:"rotation_lambda_arn,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RotationRules []SecretsmanagerSecretSpecRotationRules `json:"rotationRules,omitempty" tf:"rotation_rules,omitempty"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SecretsmanagerSecretStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
