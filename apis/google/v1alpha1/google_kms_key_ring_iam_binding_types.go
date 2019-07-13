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

type GoogleKmsKeyRingIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleKmsKeyRingIamBindingSpec   `json:"spec,omitempty"`
	Status            GoogleKmsKeyRingIamBindingStatus `json:"status,omitempty"`
}

type GoogleKmsKeyRingIamBindingSpec struct {
	Members   []string `json:"members"`
	Etag      string   `json:"etag"`
	KeyRingId string   `json:"key_ring_id"`
	Role      string   `json:"role"`
}



type GoogleKmsKeyRingIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleKmsKeyRingIamBindingList is a list of GoogleKmsKeyRingIamBindings
type GoogleKmsKeyRingIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleKmsKeyRingIamBinding CRD objects
	Items []GoogleKmsKeyRingIamBinding `json:"items,omitempty"`
}