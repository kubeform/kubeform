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
	// +optional
	DomainName string `json:"domainName,omitempty" tf:"domain_name,omitempty"`
	// +optional
	ResourceRecordName string `json:"resourceRecordName,omitempty" tf:"resource_record_name,omitempty"`
	// +optional
	ResourceRecordType string `json:"resourceRecordType,omitempty" tf:"resource_record_type,omitempty"`
	// +optional
	ResourceRecordValue string `json:"resourceRecordValue,omitempty" tf:"resource_record_value,omitempty"`
}

type AcmCertificateSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	KubeFormSecret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	CertificateBody string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`
	// +optional
	CertificateChain string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`
	// +optional
	DomainName string `json:"domainName,omitempty" tf:"domain_name,omitempty"`
	// +optional
	DomainValidationOptions []AcmCertificateSpecDomainValidationOptions `json:"domainValidationOptions,omitempty" tf:"domain_validation_options,omitempty"`
	// +optional
	PrivateKey string `json:"-" sensitive:"true" tf:"private_key,omitempty"`
	// +optional
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	ValidationEmails []string `json:"validationEmails,omitempty" tf:"validation_emails,omitempty"`
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
