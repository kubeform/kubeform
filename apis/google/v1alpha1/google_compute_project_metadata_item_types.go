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

type GoogleComputeProjectMetadataItem struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeProjectMetadataItemSpec   `json:"spec,omitempty"`
	Status            GoogleComputeProjectMetadataItemStatus `json:"status,omitempty"`
}

type GoogleComputeProjectMetadataItemSpec struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Project string `json:"project"`
}

type GoogleComputeProjectMetadataItemStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeProjectMetadataItemList is a list of GoogleComputeProjectMetadataItems
type GoogleComputeProjectMetadataItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeProjectMetadataItem CRD objects
	Items []GoogleComputeProjectMetadataItem `json:"items,omitempty"`
}
