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

type StorageDefaultObjectAccessControl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageDefaultObjectAccessControlSpec   `json:"spec,omitempty"`
	Status            StorageDefaultObjectAccessControlStatus `json:"status,omitempty"`
}

type StorageDefaultObjectAccessControlSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Bucket string `json:"bucket" tf:"bucket"`
	Entity string `json:"entity" tf:"entity"`
	// +optional
	Object string `json:"object,omitempty" tf:"object,omitempty"`
	Role   string `json:"role" tf:"role"`
}

type StorageDefaultObjectAccessControlStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageDefaultObjectAccessControlList is a list of StorageDefaultObjectAccessControls
type StorageDefaultObjectAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageDefaultObjectAccessControl CRD objects
	Items []StorageDefaultObjectAccessControl `json:"items,omitempty"`
}
