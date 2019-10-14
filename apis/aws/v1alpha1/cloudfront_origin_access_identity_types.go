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

type CloudfrontOriginAccessIdentity struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudfrontOriginAccessIdentitySpec   `json:"spec,omitempty"`
	Status            CloudfrontOriginAccessIdentityStatus `json:"status,omitempty"`
}

type CloudfrontOriginAccessIdentitySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CallerReference string `json:"callerReference,omitempty" tf:"caller_reference,omitempty"`
	// +optional
	CloudfrontAccessIdentityPath string `json:"cloudfrontAccessIdentityPath,omitempty" tf:"cloudfront_access_identity_path,omitempty"`
	// +optional
	Comment string `json:"comment,omitempty" tf:"comment,omitempty"`
	// +optional
	Etag string `json:"etag,omitempty" tf:"etag,omitempty"`
	// +optional
	IamArn string `json:"iamArn,omitempty" tf:"iam_arn,omitempty"`
	// +optional
	S3CanonicalUserID string `json:"s3CanonicalUserID,omitempty" tf:"s3_canonical_user_id,omitempty"`
}

type CloudfrontOriginAccessIdentityStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CloudfrontOriginAccessIdentitySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudfrontOriginAccessIdentityList is a list of CloudfrontOriginAccessIdentitys
type CloudfrontOriginAccessIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudfrontOriginAccessIdentity CRD objects
	Items []CloudfrontOriginAccessIdentity `json:"items,omitempty"`
}
