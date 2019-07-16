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

type S3BucketPublicAccessBlock struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3BucketPublicAccessBlockSpec   `json:"spec,omitempty"`
	Status            S3BucketPublicAccessBlockStatus `json:"status,omitempty"`
}

type S3BucketPublicAccessBlockSpec struct {
	// +optional
	BlockPublicAcls bool `json:"block_public_acls,omitempty"`
	// +optional
	BlockPublicPolicy bool   `json:"block_public_policy,omitempty"`
	Bucket            string `json:"bucket"`
	// +optional
	IgnorePublicAcls bool `json:"ignore_public_acls,omitempty"`
	// +optional
	RestrictPublicBuckets bool `json:"restrict_public_buckets,omitempty"`
}

type S3BucketPublicAccessBlockStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// S3BucketPublicAccessBlockList is a list of S3BucketPublicAccessBlocks
type S3BucketPublicAccessBlockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of S3BucketPublicAccessBlock CRD objects
	Items []S3BucketPublicAccessBlock `json:"items,omitempty"`
}
