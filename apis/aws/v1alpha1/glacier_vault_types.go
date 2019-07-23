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

type GlacierVault struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlacierVaultSpec   `json:"spec,omitempty"`
	Status            GlacierVaultStatus `json:"status,omitempty"`
}

type GlacierVaultSpecNotification struct {
	// +kubebuilder:validation:UniqueItems=true
	Events   []string `json:"events" tf:"events"`
	SnsTopic string   `json:"snsTopic" tf:"sns_topic"`
}

type GlacierVaultSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	AccessPolicy string `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`
	Name         string `json:"name" tf:"name"`
	// +optional
	Notification []GlacierVaultSpecNotification `json:"notification,omitempty" tf:"notification,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type GlacierVaultStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlacierVaultList is a list of GlacierVaults
type GlacierVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlacierVault CRD objects
	Items []GlacierVault `json:"items,omitempty"`
}
