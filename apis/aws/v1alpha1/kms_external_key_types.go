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

type KmsExternalKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KmsExternalKeySpec   `json:"spec,omitempty"`
	Status            KmsExternalKeyStatus `json:"status,omitempty"`
}

type KmsExternalKeySpec struct {
	// +optional
	DeletionWindowInDays int `json:"deletion_window_in_days,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	KeyMaterialBase64 string `json:"key_material_base64,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	ValidTo string `json:"valid_to,omitempty"`
}

type KmsExternalKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KmsExternalKeyList is a list of KmsExternalKeys
type KmsExternalKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KmsExternalKey CRD objects
	Items []KmsExternalKey `json:"items,omitempty"`
}
