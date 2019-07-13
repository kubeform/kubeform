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

type AwsAcmCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAcmCertificateSpec   `json:"spec,omitempty"`
	Status            AwsAcmCertificateStatus `json:"status,omitempty"`
}

type AwsAcmCertificateSpecDomainValidationOptions struct {
	ResourceRecordValue string `json:"resource_record_value"`
	DomainName          string `json:"domain_name"`
	ResourceRecordName  string `json:"resource_record_name"`
	ResourceRecordType  string `json:"resource_record_type"`
}

type AwsAcmCertificateSpec struct {
	CertificateChain        string                  `json:"certificate_chain"`
	PrivateKey              string                  `json:"private_key"`
	SubjectAlternativeNames []string                `json:"subject_alternative_names"`
	ValidationMethod        string                  `json:"validation_method"`
	Arn                     string                  `json:"arn"`
	CertificateBody         string                  `json:"certificate_body"`
	DomainName              string                  `json:"domain_name"`
	DomainValidationOptions []AwsAcmCertificateSpec `json:"domain_validation_options"`
	ValidationEmails        []string                `json:"validation_emails"`
	Tags                    map[string]string       `json:"tags"`
}



type AwsAcmCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAcmCertificateList is a list of AwsAcmCertificates
type AwsAcmCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAcmCertificate CRD objects
	Items []AwsAcmCertificate `json:"items,omitempty"`
}