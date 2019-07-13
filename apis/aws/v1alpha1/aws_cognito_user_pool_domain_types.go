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

type AwsCognitoUserPoolDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCognitoUserPoolDomainSpec   `json:"spec,omitempty"`
	Status            AwsCognitoUserPoolDomainStatus `json:"status,omitempty"`
}

type AwsCognitoUserPoolDomainSpec struct {
	CertificateArn            string `json:"certificate_arn"`
	UserPoolId                string `json:"user_pool_id"`
	AwsAccountId              string `json:"aws_account_id"`
	CloudfrontDistributionArn string `json:"cloudfront_distribution_arn"`
	S3Bucket                  string `json:"s3_bucket"`
	Version                   string `json:"version"`
	Domain                    string `json:"domain"`
}



type AwsCognitoUserPoolDomainStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCognitoUserPoolDomainList is a list of AwsCognitoUserPoolDomains
type AwsCognitoUserPoolDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCognitoUserPoolDomain CRD objects
	Items []AwsCognitoUserPoolDomain `json:"items,omitempty"`
}