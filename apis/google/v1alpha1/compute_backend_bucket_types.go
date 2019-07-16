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

type ComputeBackendBucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeBackendBucketSpec   `json:"spec,omitempty"`
	Status            ComputeBackendBucketStatus `json:"status,omitempty"`
}

type ComputeBackendBucketSpec struct {
	BucketName string `json:"bucket_name"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	EnableCdn bool   `json:"enable_cdn,omitempty"`
	Name      string `json:"name"`
}

type ComputeBackendBucketStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeBackendBucketList is a list of ComputeBackendBuckets
type ComputeBackendBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeBackendBucket CRD objects
	Items []ComputeBackendBucket `json:"items,omitempty"`
}
