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

type StorageBucketIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageBucketIamBindingSpec   `json:"spec,omitempty"`
	Status            StorageBucketIamBindingStatus `json:"status,omitempty"`
}

type StorageBucketIamBindingSpec struct {
	Bucket string `json:"bucket"`
	// +kubebuilder:validation:UniqueItems=true
	Members []string `json:"members"`
	Role    string   `json:"role"`
}

type StorageBucketIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageBucketIamBindingList is a list of StorageBucketIamBindings
type StorageBucketIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageBucketIamBinding CRD objects
	Items []StorageBucketIamBinding `json:"items,omitempty"`
}
