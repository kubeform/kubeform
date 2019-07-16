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

type KmsKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KmsKeySpec   `json:"spec,omitempty"`
	Status            KmsKeyStatus `json:"status,omitempty"`
}

type KmsKeySpec struct {
	// +optional
	DeletionWindowInDays int `json:"deletion_window_in_days,omitempty"`
	// +optional
	EnableKeyRotation bool `json:"enable_key_rotation,omitempty"`
	// +optional
	IsEnabled bool `json:"is_enabled,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type KmsKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KmsKeyList is a list of KmsKeys
type KmsKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KmsKey CRD objects
	Items []KmsKey `json:"items,omitempty"`
}
