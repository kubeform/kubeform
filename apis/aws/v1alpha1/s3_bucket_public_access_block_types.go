package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	BlockPublicAcls bool `json:"blockPublicAcls,omitempty" tf:"block_public_acls,omitempty"`
	// +optional
	BlockPublicPolicy bool   `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`
	Bucket            string `json:"bucket" tf:"bucket"`
	// +optional
	IgnorePublicAcls bool `json:"ignorePublicAcls,omitempty" tf:"ignore_public_acls,omitempty"`
	// +optional
	RestrictPublicBuckets bool `json:"restrictPublicBuckets,omitempty" tf:"restrict_public_buckets,omitempty"`
}

type S3BucketPublicAccessBlockStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
