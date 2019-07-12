package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAcmCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAcmCertificateSpec   `json:"spec,omitempty"`
	Status            AwsAcmCertificateStatus `json:"status,omitempty"`
}

type AwsAcmCertificateSpecDomainValidationOptions struct {
	ResourceRecordType  string `json:"resource_record_type"`
	ResourceRecordValue string `json:"resource_record_value"`
	DomainName          string `json:"domain_name"`
	ResourceRecordName  string `json:"resource_record_name"`
}

type AwsAcmCertificateSpec struct {
	CertificateChain        string                  `json:"certificate_chain"`
	PrivateKey              string                  `json:"private_key"`
	Arn                     string                  `json:"arn"`
	DomainValidationOptions []AwsAcmCertificateSpec `json:"domain_validation_options"`
	Tags                    map[string]string       `json:"tags"`
	CertificateBody         string                  `json:"certificate_body"`
	DomainName              string                  `json:"domain_name"`
	SubjectAlternativeNames []string                `json:"subject_alternative_names"`
	ValidationMethod        string                  `json:"validation_method"`
	ValidationEmails        []string                `json:"validation_emails"`
}

type AwsAcmCertificateStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAcmCertificateList is a list of AwsAcmCertificates
type AwsAcmCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAcmCertificate CRD objects
	Items []AwsAcmCertificate `json:"items,omitempty"`
}
