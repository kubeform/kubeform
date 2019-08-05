package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type CloudfrontOriginAccessIdentity struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudfrontOriginAccessIdentitySpec   `json:"spec,omitempty"`
	Status            CloudfrontOriginAccessIdentityStatus `json:"status,omitempty"`
}

type CloudfrontOriginAccessIdentitySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	CallerReference              string `json:"callerReference" tf:"caller_reference"`
	CloudfrontAccessIdentityPath string `json:"cloudfrontAccessIdentityPath" tf:"cloudfront_access_identity_path"`
	// +optional
	Comment           string `json:"comment,omitempty" tf:"comment,omitempty"`
	Etag              string `json:"etag" tf:"etag"`
	IamArn            string `json:"iamArn" tf:"iam_arn"`
	S3CanonicalUserID string `json:"s3CanonicalUserID" tf:"s3_canonical_user_id"`
}

type CloudfrontOriginAccessIdentityStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CloudfrontOriginAccessIdentitySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
