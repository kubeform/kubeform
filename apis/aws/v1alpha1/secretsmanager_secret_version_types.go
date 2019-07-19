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

type SecretsmanagerSecretVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretsmanagerSecretVersionSpec   `json:"spec,omitempty"`
	Status            SecretsmanagerSecretVersionStatus `json:"status,omitempty"`
}

type SecretsmanagerSecretVersionSpec struct {
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	SecretBinary core.LocalObjectReference `json:"secretBinary,omitempty" tf:"secret_binary,omitempty"`
	SecretID     string                    `json:"secretID" tf:"secret_id"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	SecretString core.LocalObjectReference `json:"secretString,omitempty" tf:"secret_string,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VersionStages []string                  `json:"versionStages,omitempty" tf:"version_stages,omitempty"`
	ProviderRef   core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SecretsmanagerSecretVersionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
