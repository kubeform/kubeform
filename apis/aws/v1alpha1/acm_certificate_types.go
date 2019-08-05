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

type AcmCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AcmCertificateSpec   `json:"spec,omitempty"`
	Status            AcmCertificateStatus `json:"status,omitempty"`
}

type AcmCertificateSpecDomainValidationOptions struct {
	DomainName          string `json:"domainName" tf:"domain_name"`
	ResourceRecordName  string `json:"resourceRecordName" tf:"resource_record_name"`
	ResourceRecordType  string `json:"resourceRecordType" tf:"resource_record_type"`
	ResourceRecordValue string `json:"resourceRecordValue" tf:"resource_record_value"`
}

type AcmCertificateSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	Arn string `json:"arn" tf:"arn"`
	// +optional
	CertificateBody string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`
	// +optional
	CertificateChain string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`
	// +optional
	DomainName              string                                      `json:"domainName,omitempty" tf:"domain_name,omitempty"`
	DomainValidationOptions []AcmCertificateSpecDomainValidationOptions `json:"domainValidationOptions" tf:"domain_validation_options"`
	// +optional
	PrivateKey string `json:"-" sensitive:"true" tf:"private_key,omitempty"`
	// +optional
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`
	// +optional
	Tags             map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	ValidationEmails []string          `json:"validationEmails" tf:"validation_emails"`
	// +optional
	ValidationMethod string `json:"validationMethod,omitempty" tf:"validation_method,omitempty"`
}

type AcmCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AcmCertificateSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AcmCertificateList is a list of AcmCertificates
type AcmCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AcmCertificate CRD objects
	Items []AcmCertificate `json:"items,omitempty"`
}
