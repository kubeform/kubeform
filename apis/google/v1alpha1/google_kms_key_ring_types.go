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

type GoogleKmsKeyRing struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleKmsKeyRingSpec   `json:"spec,omitempty"`
	Status            GoogleKmsKeyRingStatus `json:"status,omitempty"`
}

type GoogleKmsKeyRingSpec struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Project  string `json:"project"`
	SelfLink string `json:"self_link"`
}



type GoogleKmsKeyRingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleKmsKeyRingList is a list of GoogleKmsKeyRings
type GoogleKmsKeyRingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleKmsKeyRing CRD objects
	Items []GoogleKmsKeyRing `json:"items,omitempty"`
}