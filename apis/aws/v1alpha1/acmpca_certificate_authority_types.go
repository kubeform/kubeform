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

type AcmpcaCertificateAuthority struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AcmpcaCertificateAuthoritySpec   `json:"spec,omitempty"`
	Status            AcmpcaCertificateAuthorityStatus `json:"status,omitempty"`
}

type AcmpcaCertificateAuthoritySpecCertificateAuthorityConfigurationSubject struct {
	// +optional
	CommonName string `json:"common_name,omitempty"`
	// +optional
	Country string `json:"country,omitempty"`
	// +optional
	DistinguishedNameQualifier string `json:"distinguished_name_qualifier,omitempty"`
	// +optional
	GenerationQualifier string `json:"generation_qualifier,omitempty"`
	// +optional
	GivenName string `json:"given_name,omitempty"`
	// +optional
	Initials string `json:"initials,omitempty"`
	// +optional
	Locality string `json:"locality,omitempty"`
	// +optional
	Organization string `json:"organization,omitempty"`
	// +optional
	OrganizationalUnit string `json:"organizational_unit,omitempty"`
	// +optional
	Pseudonym string `json:"pseudonym,omitempty"`
	// +optional
	State string `json:"state,omitempty"`
	// +optional
	Surname string `json:"surname,omitempty"`
	// +optional
	Title string `json:"title,omitempty"`
}

type AcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration struct {
	KeyAlgorithm     string `json:"key_algorithm"`
	SigningAlgorithm string `json:"signing_algorithm"`
	// +kubebuilder:validation:MaxItems=1
	Subject []AcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration `json:"subject"`
}

type AcmpcaCertificateAuthoritySpecRevocationConfigurationCrlConfiguration struct {
	// +optional
	CustomCname string `json:"custom_cname,omitempty"`
	// +optional
	Enabled          bool `json:"enabled,omitempty"`
	ExpirationInDays int  `json:"expiration_in_days"`
	// +optional
	S3BucketName string `json:"s3_bucket_name,omitempty"`
}

type AcmpcaCertificateAuthoritySpecRevocationConfiguration struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CrlConfiguration *[]AcmpcaCertificateAuthoritySpecRevocationConfiguration `json:"crl_configuration,omitempty"`
}

type AcmpcaCertificateAuthoritySpec struct {
	// +kubebuilder:validation:MaxItems=1
	CertificateAuthorityConfiguration []AcmpcaCertificateAuthoritySpec `json:"certificate_authority_configuration"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	PermanentDeletionTimeInDays int `json:"permanent_deletion_time_in_days,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RevocationConfiguration *[]AcmpcaCertificateAuthoritySpec `json:"revocation_configuration,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	Type string `json:"type,omitempty"`
}

type AcmpcaCertificateAuthorityStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AcmpcaCertificateAuthorityList is a list of AcmpcaCertificateAuthoritys
type AcmpcaCertificateAuthorityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AcmpcaCertificateAuthority CRD objects
	Items []AcmpcaCertificateAuthority `json:"items,omitempty"`
}
