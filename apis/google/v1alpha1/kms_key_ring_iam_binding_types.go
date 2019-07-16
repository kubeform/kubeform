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

type KmsKeyRingIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KmsKeyRingIamBindingSpec   `json:"spec,omitempty"`
	Status            KmsKeyRingIamBindingStatus `json:"status,omitempty"`
}

type KmsKeyRingIamBindingSpec struct {
	KeyRingId string `json:"key_ring_id"`
	// +kubebuilder:validation:UniqueItems=true
	Members []string `json:"members"`
	Role    string   `json:"role"`
}

type KmsKeyRingIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KmsKeyRingIamBindingList is a list of KmsKeyRingIamBindings
type KmsKeyRingIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KmsKeyRingIamBinding CRD objects
	Items []KmsKeyRingIamBinding `json:"items,omitempty"`
}
