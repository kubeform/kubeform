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

type CognitoUserPoolDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoUserPoolDomainSpec   `json:"spec,omitempty"`
	Status            CognitoUserPoolDomainStatus `json:"status,omitempty"`
}

type CognitoUserPoolDomainSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AwsAccountID string `json:"awsAccountID,omitempty" tf:"aws_account_id,omitempty"`
	// +optional
	CertificateArn string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`
	// +optional
	CloudfrontDistributionArn string `json:"cloudfrontDistributionArn,omitempty" tf:"cloudfront_distribution_arn,omitempty"`
	Domain                    string `json:"domain" tf:"domain"`
	// +optional
	S3Bucket   string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`
	UserPoolID string `json:"userPoolID" tf:"user_pool_id"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
}

type CognitoUserPoolDomainStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CognitoUserPoolDomainSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CognitoUserPoolDomainList is a list of CognitoUserPoolDomains
type CognitoUserPoolDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CognitoUserPoolDomain CRD objects
	Items []CognitoUserPoolDomain `json:"items,omitempty"`
}
