package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type S3AccountPublicAccessBlock struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3AccountPublicAccessBlockSpec   `json:"spec,omitempty"`
	Status            S3AccountPublicAccessBlockStatus `json:"status,omitempty"`
}

type S3AccountPublicAccessBlockSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccountID string `json:"accountID,omitempty" tf:"account_id,omitempty"`
	// +optional
	BlockPublicAcls bool `json:"blockPublicAcls,omitempty" tf:"block_public_acls,omitempty"`
	// +optional
	BlockPublicPolicy bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`
	// +optional
	IgnorePublicAcls bool `json:"ignorePublicAcls,omitempty" tf:"ignore_public_acls,omitempty"`
	// +optional
	RestrictPublicBuckets bool `json:"restrictPublicBuckets,omitempty" tf:"restrict_public_buckets,omitempty"`
}

type S3AccountPublicAccessBlockStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *S3AccountPublicAccessBlockSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// S3AccountPublicAccessBlockList is a list of S3AccountPublicAccessBlocks
type S3AccountPublicAccessBlockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of S3AccountPublicAccessBlock CRD objects
	Items []S3AccountPublicAccessBlock `json:"items,omitempty"`
}
