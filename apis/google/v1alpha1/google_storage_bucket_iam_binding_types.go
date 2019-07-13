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

type GoogleStorageBucketIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleStorageBucketIamBindingSpec   `json:"spec,omitempty"`
	Status            GoogleStorageBucketIamBindingStatus `json:"status,omitempty"`
}

type GoogleStorageBucketIamBindingSpec struct {
	Role    string   `json:"role"`
	Members []string `json:"members"`
	Etag    string   `json:"etag"`
	Bucket  string   `json:"bucket"`
}



type GoogleStorageBucketIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleStorageBucketIamBindingList is a list of GoogleStorageBucketIamBindings
type GoogleStorageBucketIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleStorageBucketIamBinding CRD objects
	Items []GoogleStorageBucketIamBinding `json:"items,omitempty"`
}