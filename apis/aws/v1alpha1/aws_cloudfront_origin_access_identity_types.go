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

type AwsCloudfrontOriginAccessIdentity struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudfrontOriginAccessIdentitySpec   `json:"spec,omitempty"`
	Status            AwsCloudfrontOriginAccessIdentityStatus `json:"status,omitempty"`
}

type AwsCloudfrontOriginAccessIdentitySpec struct {
	IamArn                       string `json:"iam_arn"`
	S3CanonicalUserId            string `json:"s3_canonical_user_id"`
	Comment                      string `json:"comment"`
	CallerReference              string `json:"caller_reference"`
	CloudfrontAccessIdentityPath string `json:"cloudfront_access_identity_path"`
	Etag                         string `json:"etag"`
}

type AwsCloudfrontOriginAccessIdentityStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCloudfrontOriginAccessIdentityList is a list of AwsCloudfrontOriginAccessIdentitys
type AwsCloudfrontOriginAccessIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudfrontOriginAccessIdentity CRD objects
	Items []AwsCloudfrontOriginAccessIdentity `json:"items,omitempty"`
}
