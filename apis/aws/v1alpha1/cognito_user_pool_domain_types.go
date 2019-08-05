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

type CognitoUserPoolDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoUserPoolDomainSpec   `json:"spec,omitempty"`
	Status            CognitoUserPoolDomainStatus `json:"status,omitempty"`
}

type CognitoUserPoolDomainSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	AwsAccountID string `json:"awsAccountID" tf:"aws_account_id"`
	// +optional
	CertificateArn            string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`
	CloudfrontDistributionArn string `json:"cloudfrontDistributionArn" tf:"cloudfront_distribution_arn"`
	Domain                    string `json:"domain" tf:"domain"`
	S3Bucket                  string `json:"s3Bucket" tf:"s3_bucket"`
	UserPoolID                string `json:"userPoolID" tf:"user_pool_id"`
	Version                   string `json:"version" tf:"version"`
}

type CognitoUserPoolDomainStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CognitoUserPoolDomainSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
