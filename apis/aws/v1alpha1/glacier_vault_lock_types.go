package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type GlacierVaultLock struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlacierVaultLockSpec   `json:"spec,omitempty"`
	Status            GlacierVaultLockStatus `json:"status,omitempty"`
}

type GlacierVaultLockSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompleteLock bool `json:"completeLock" tf:"complete_lock"`
	// +optional
	IgnoreDeletionError bool   `json:"ignoreDeletionError,omitempty" tf:"ignore_deletion_error,omitempty"`
	Policy              string `json:"policy" tf:"policy"`
	VaultName           string `json:"vaultName" tf:"vault_name"`
}

type GlacierVaultLockStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GlacierVaultLockSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlacierVaultLockList is a list of GlacierVaultLocks
type GlacierVaultLockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlacierVaultLock CRD objects
	Items []GlacierVaultLock `json:"items,omitempty"`
}
