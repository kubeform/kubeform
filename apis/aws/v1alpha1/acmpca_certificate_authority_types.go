package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type AcmpcaCertificateAuthority struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AcmpcaCertificateAuthoritySpec   `json:"spec,omitempty"`
	Status            AcmpcaCertificateAuthorityStatus `json:"status,omitempty"`
}

type AcmpcaCertificateAuthoritySpecCertificateAuthorityConfigurationSubject struct {
	// +optional
	CommonName string `json:"commonName,omitempty" tf:"common_name,omitempty"`
	// +optional
	Country string `json:"country,omitempty" tf:"country,omitempty"`
	// +optional
	DistinguishedNameQualifier string `json:"distinguishedNameQualifier,omitempty" tf:"distinguished_name_qualifier,omitempty"`
	// +optional
	GenerationQualifier string `json:"generationQualifier,omitempty" tf:"generation_qualifier,omitempty"`
	// +optional
	GivenName string `json:"givenName,omitempty" tf:"given_name,omitempty"`
	// +optional
	Initials string `json:"initials,omitempty" tf:"initials,omitempty"`
	// +optional
	Locality string `json:"locality,omitempty" tf:"locality,omitempty"`
	// +optional
	Organization string `json:"organization,omitempty" tf:"organization,omitempty"`
	// +optional
	OrganizationalUnit string `json:"organizationalUnit,omitempty" tf:"organizational_unit,omitempty"`
	// +optional
	Pseudonym string `json:"pseudonym,omitempty" tf:"pseudonym,omitempty"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
	// +optional
	Surname string `json:"surname,omitempty" tf:"surname,omitempty"`
	// +optional
	Title string `json:"title,omitempty" tf:"title,omitempty"`
}

type AcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration struct {
	KeyAlgorithm     string `json:"keyAlgorithm" tf:"key_algorithm"`
	SigningAlgorithm string `json:"signingAlgorithm" tf:"signing_algorithm"`
	// +kubebuilder:validation:MaxItems=1
	Subject []AcmpcaCertificateAuthoritySpecCertificateAuthorityConfigurationSubject `json:"subject" tf:"subject"`
}

type AcmpcaCertificateAuthoritySpecRevocationConfigurationCrlConfiguration struct {
	// +optional
	CustomCname string `json:"customCname,omitempty" tf:"custom_cname,omitempty"`
	// +optional
	Enabled          bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	ExpirationInDays int  `json:"expirationInDays" tf:"expiration_in_days"`
	// +optional
	S3BucketName string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`
}

type AcmpcaCertificateAuthoritySpecRevocationConfiguration struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CrlConfiguration []AcmpcaCertificateAuthoritySpecRevocationConfigurationCrlConfiguration `json:"crlConfiguration,omitempty" tf:"crl_configuration,omitempty"`
}

type AcmpcaCertificateAuthoritySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	Certificate string `json:"certificate,omitempty" tf:"certificate,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	CertificateAuthorityConfiguration []AcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration `json:"certificateAuthorityConfiguration" tf:"certificate_authority_configuration"`
	// +optional
	CertificateChain string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`
	// +optional
	CertificateSigningRequest string `json:"certificateSigningRequest,omitempty" tf:"certificate_signing_request,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	NotAfter string `json:"notAfter,omitempty" tf:"not_after,omitempty"`
	// +optional
	NotBefore string `json:"notBefore,omitempty" tf:"not_before,omitempty"`
	// +optional
	PermanentDeletionTimeInDays int `json:"permanentDeletionTimeInDays,omitempty" tf:"permanent_deletion_time_in_days,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RevocationConfiguration []AcmpcaCertificateAuthoritySpecRevocationConfiguration `json:"revocationConfiguration,omitempty" tf:"revocation_configuration,omitempty"`
	// +optional
	Serial string `json:"serial,omitempty" tf:"serial,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type AcmpcaCertificateAuthorityStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AcmpcaCertificateAuthoritySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
