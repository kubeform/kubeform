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

type ComputeProjectMetadataItem struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeProjectMetadataItemSpec   `json:"spec,omitempty"`
	Status            ComputeProjectMetadataItemStatus `json:"status,omitempty"`
}

type ComputeProjectMetadataItemSpec struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ComputeProjectMetadataItemStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeProjectMetadataItemList is a list of ComputeProjectMetadataItems
type ComputeProjectMetadataItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeProjectMetadataItem CRD objects
	Items []ComputeProjectMetadataItem `json:"items,omitempty"`
}
