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

type GoogleComputeProjectMetadata struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeProjectMetadataSpec   `json:"spec,omitempty"`
	Status            GoogleComputeProjectMetadataStatus `json:"status,omitempty"`
}

type GoogleComputeProjectMetadataSpec struct {
	Metadata map[string]string `json:"metadata"`
	Project  string            `json:"project"`
}



type GoogleComputeProjectMetadataStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeProjectMetadataList is a list of GoogleComputeProjectMetadatas
type GoogleComputeProjectMetadataList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeProjectMetadata CRD objects
	Items []GoogleComputeProjectMetadata `json:"items,omitempty"`
}