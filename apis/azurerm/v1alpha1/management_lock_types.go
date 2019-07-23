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

type ManagementLock struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementLockSpec   `json:"spec,omitempty"`
	Status            ManagementLockStatus `json:"status,omitempty"`
}

type ManagementLockSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	LockLevel string `json:"lockLevel" tf:"lock_level"`
	Name      string `json:"name" tf:"name"`
	// +optional
	Notes string `json:"notes,omitempty" tf:"notes,omitempty"`
	Scope string `json:"scope" tf:"scope"`
}

type ManagementLockStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ManagementLockList is a list of ManagementLocks
type ManagementLockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagementLock CRD objects
	Items []ManagementLock `json:"items,omitempty"`
}
