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

type AwsAcmpcaCertificateAuthority struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAcmpcaCertificateAuthoritySpec   `json:"spec,omitempty"`
	Status            AwsAcmpcaCertificateAuthorityStatus `json:"status,omitempty"`
}

type AwsAcmpcaCertificateAuthoritySpecCertificateAuthorityConfigurationSubject struct {
	DistinguishedNameQualifier string `json:"distinguished_name_qualifier"`
	GenerationQualifier        string `json:"generation_qualifier"`
	Locality                   string `json:"locality"`
	State                      string `json:"state"`
	Surname                    string `json:"surname"`
	CommonName                 string `json:"common_name"`
	GivenName                  string `json:"given_name"`
	Initials                   string `json:"initials"`
	Organization               string `json:"organization"`
	OrganizationalUnit         string `json:"organizational_unit"`
	Pseudonym                  string `json:"pseudonym"`
	Title                      string `json:"title"`
	Country                    string `json:"country"`
}

type AwsAcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration struct {
	Subject          []AwsAcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration `json:"subject"`
	KeyAlgorithm     string                                                               `json:"key_algorithm"`
	SigningAlgorithm string                                                               `json:"signing_algorithm"`
}

type AwsAcmpcaCertificateAuthoritySpecRevocationConfigurationCrlConfiguration struct {
	CustomCname      string `json:"custom_cname"`
	Enabled          bool   `json:"enabled"`
	ExpirationInDays int    `json:"expiration_in_days"`
	S3BucketName     string `json:"s3_bucket_name"`
}

type AwsAcmpcaCertificateAuthoritySpecRevocationConfiguration struct {
	CrlConfiguration []AwsAcmpcaCertificateAuthoritySpecRevocationConfiguration `json:"crl_configuration"`
}

type AwsAcmpcaCertificateAuthoritySpec struct {
	CertificateAuthorityConfiguration []AwsAcmpcaCertificateAuthoritySpec `json:"certificate_authority_configuration"`
	Enabled                           bool                                `json:"enabled"`
	NotAfter                          string                              `json:"not_after"`
	Serial                            string                              `json:"serial"`
	CertificateChain                  string                              `json:"certificate_chain"`
	Status                            string                              `json:"status"`
	PermanentDeletionTimeInDays       int                                 `json:"permanent_deletion_time_in_days"`
	Tags                              map[string]string                   `json:"tags"`
	Arn                               string                              `json:"arn"`
	Certificate                       string                              `json:"certificate"`
	CertificateSigningRequest         string                              `json:"certificate_signing_request"`
	NotBefore                         string                              `json:"not_before"`
	RevocationConfiguration           []AwsAcmpcaCertificateAuthoritySpec `json:"revocation_configuration"`
	Type                              string                              `json:"type"`
}



type AwsAcmpcaCertificateAuthorityStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAcmpcaCertificateAuthorityList is a list of AwsAcmpcaCertificateAuthoritys
type AwsAcmpcaCertificateAuthorityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAcmpcaCertificateAuthority CRD objects
	Items []AwsAcmpcaCertificateAuthority `json:"items,omitempty"`
}