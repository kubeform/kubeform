package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	GenerationQualifier        string `json:"generation_qualifier"`
	GivenName                  string `json:"given_name"`
	OrganizationalUnit         string `json:"organizational_unit"`
	Pseudonym                  string `json:"pseudonym"`
	Surname                    string `json:"surname"`
	Title                      string `json:"title"`
	CommonName                 string `json:"common_name"`
	DistinguishedNameQualifier string `json:"distinguished_name_qualifier"`
	Initials                   string `json:"initials"`
	Locality                   string `json:"locality"`
	Organization               string `json:"organization"`
	State                      string `json:"state"`
}

type AwsAcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration struct {
	Subject          []AwsAcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration `json:"subject"`
	KeyAlgorithm     string                                                               `json:"key_algorithm"`
	SigningAlgorithm string                                                               `json:"signing_algorithm"`
}

type AwsAcmpcaCertificateAuthoritySpec struct {
	Certificate                       string                              `json:"certificate"`
	CertificateChain                  string                              `json:"certificate_chain"`
	CertificateSigningRequest         string                              `json:"certificate_signing_request"`
	Status                            string                              `json:"status"`
	Arn                               string                              `json:"arn"`
	Enabled                           bool                                `json:"enabled"`
	NotAfter                          string                              `json:"not_after"`
	RevocationConfiguration           []AwsAcmpcaCertificateAuthoritySpec `json:"revocation_configuration"`
	Tags                              map[string]string                   `json:"tags"`
	Type                              string                              `json:"type"`
	CertificateAuthorityConfiguration []AwsAcmpcaCertificateAuthoritySpec `json:"certificate_authority_configuration"`
	NotBefore                         string                              `json:"not_before"`
	Serial                            string                              `json:"serial"`
	PermanentDeletionTimeInDays       int                                 `json:"permanent_deletion_time_in_days"`
}



type AwsAcmpcaCertificateAuthorityStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAcmpcaCertificateAuthorityList is a list of AwsAcmpcaCertificateAuthoritys
type AwsAcmpcaCertificateAuthorityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAcmpcaCertificateAuthority CRD objects
	Items []AwsAcmpcaCertificateAuthority `json:"items,omitempty"`
}