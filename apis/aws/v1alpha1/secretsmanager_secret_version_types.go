package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SecretsmanagerSecretVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretsmanagerSecretVersionSpec   `json:"spec,omitempty"`
	Status            SecretsmanagerSecretVersionStatus `json:"status,omitempty"`
}

type SecretsmanagerSecretVersionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	SecretBinary string `json:"-" sensitive:"true" tf:"secret_binary,omitempty"`
	SecretID     string `json:"secretID" tf:"secret_id"`
	// +optional
	SecretString string `json:"-" sensitive:"true" tf:"secret_string,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VersionStages []string `json:"versionStages,omitempty" tf:"version_stages,omitempty"`
}

type SecretsmanagerSecretVersionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SecretsmanagerSecretVersionList is a list of SecretsmanagerSecretVersions
type SecretsmanagerSecretVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SecretsmanagerSecretVersion CRD objects
	Items []SecretsmanagerSecretVersion `json:"items,omitempty"`
}
