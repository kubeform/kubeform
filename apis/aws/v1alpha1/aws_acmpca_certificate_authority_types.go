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

type AwsAcmpcaCertificateAuthoritySpecRevocationConfigurationCrlConfiguration struct {
	CustomCname      string `json:"custom_cname"`
	Enabled          bool   `json:"enabled"`
	ExpirationInDays int    `json:"expiration_in_days"`
	S3BucketName     string `json:"s3_bucket_name"`
}

type AwsAcmpcaCertificateAuthoritySpecRevocationConfiguration struct {
	CrlConfiguration []AwsAcmpcaCertificateAuthoritySpecRevocationConfiguration `json:"crl_configuration"`
}

type AwsAcmpcaCertificateAuthoritySpecCertificateAuthorityConfigurationSubject struct {
	Country                    string `json:"country"`
	DistinguishedNameQualifier string `json:"distinguished_name_qualifier"`
	OrganizationalUnit         string `json:"organizational_unit"`
	Surname                    string `json:"surname"`
	CommonName                 string `json:"common_name"`
	GivenName                  string `json:"given_name"`
	Initials                   string `json:"initials"`
	Locality                   string `json:"locality"`
	Organization               string `json:"organization"`
	Pseudonym                  string `json:"pseudonym"`
	State                      string `json:"state"`
	Title                      string `json:"title"`
	GenerationQualifier        string `json:"generation_qualifier"`
}

type AwsAcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration struct {
	Subject          []AwsAcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration `json:"subject"`
	KeyAlgorithm     string                                                               `json:"key_algorithm"`
	SigningAlgorithm string                                                               `json:"signing_algorithm"`
}

type AwsAcmpcaCertificateAuthoritySpec struct {
	Arn                               string                              `json:"arn"`
	Certificate                       string                              `json:"certificate"`
	NotBefore                         string                              `json:"not_before"`
	RevocationConfiguration           []AwsAcmpcaCertificateAuthoritySpec `json:"revocation_configuration"`
	Serial                            string                              `json:"serial"`
	Tags                              map[string]string                   `json:"tags"`
	PermanentDeletionTimeInDays       int                                 `json:"permanent_deletion_time_in_days"`
	Type                              string                              `json:"type"`
	Enabled                           bool                                `json:"enabled"`
	Status                            string                              `json:"status"`
	CertificateAuthorityConfiguration []AwsAcmpcaCertificateAuthoritySpec `json:"certificate_authority_configuration"`
	CertificateChain                  string                              `json:"certificate_chain"`
	CertificateSigningRequest         string                              `json:"certificate_signing_request"`
	NotAfter                          string                              `json:"not_after"`
}

type AwsAcmpcaCertificateAuthorityStatus struct {
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
