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

type AzurermKeyVaultCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermKeyVaultCertificateSpec   `json:"spec,omitempty"`
	Status            AzurermKeyVaultCertificateStatus `json:"status,omitempty"`
}

type AzurermKeyVaultCertificateSpecCertificate struct {
	Contents string `json:"contents"`
	Password string `json:"password"`
}

type AzurermKeyVaultCertificateSpecCertificatePolicyIssuerParameters struct {
	Name string `json:"name"`
}

type AzurermKeyVaultCertificateSpecCertificatePolicyKeyProperties struct {
	KeySize    int    `json:"key_size"`
	KeyType    string `json:"key_type"`
	ReuseKey   bool   `json:"reuse_key"`
	Exportable bool   `json:"exportable"`
}

type AzurermKeyVaultCertificateSpecCertificatePolicyLifetimeActionAction struct {
	ActionType string `json:"action_type"`
}

type AzurermKeyVaultCertificateSpecCertificatePolicyLifetimeActionTrigger struct {
	DaysBeforeExpiry   int `json:"days_before_expiry"`
	LifetimePercentage int `json:"lifetime_percentage"`
}

type AzurermKeyVaultCertificateSpecCertificatePolicyLifetimeAction struct {
	Action  []AzurermKeyVaultCertificateSpecCertificatePolicyLifetimeAction `json:"action"`
	Trigger []AzurermKeyVaultCertificateSpecCertificatePolicyLifetimeAction `json:"trigger"`
}

type AzurermKeyVaultCertificateSpecCertificatePolicySecretProperties struct {
	ContentType string `json:"content_type"`
}

type AzurermKeyVaultCertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames struct {
	Emails   []string `json:"emails"`
	DnsNames []string `json:"dns_names"`
	Upns     []string `json:"upns"`
}

type AzurermKeyVaultCertificateSpecCertificatePolicyX509CertificateProperties struct {
	SubjectAlternativeNames []AzurermKeyVaultCertificateSpecCertificatePolicyX509CertificateProperties `json:"subject_alternative_names"`
	ValidityInMonths        int                                                                        `json:"validity_in_months"`
	ExtendedKeyUsage        []string                                                                   `json:"extended_key_usage"`
	KeyUsage                []string                                                                   `json:"key_usage"`
	Subject                 string                                                                     `json:"subject"`
}

type AzurermKeyVaultCertificateSpecCertificatePolicy struct {
	IssuerParameters          []AzurermKeyVaultCertificateSpecCertificatePolicy `json:"issuer_parameters"`
	KeyProperties             []AzurermKeyVaultCertificateSpecCertificatePolicy `json:"key_properties"`
	LifetimeAction            []AzurermKeyVaultCertificateSpecCertificatePolicy `json:"lifetime_action"`
	SecretProperties          []AzurermKeyVaultCertificateSpecCertificatePolicy `json:"secret_properties"`
	X509CertificateProperties []AzurermKeyVaultCertificateSpecCertificatePolicy `json:"x509_certificate_properties"`
}

type AzurermKeyVaultCertificateSpec struct {
	Name              string                           `json:"name"`
	KeyVaultId        string                           `json:"key_vault_id"`
	VaultUri          string                           `json:"vault_uri"`
	Certificate       []AzurermKeyVaultCertificateSpec `json:"certificate"`
	CertificatePolicy []AzurermKeyVaultCertificateSpec `json:"certificate_policy"`
	CertificateData   string                           `json:"certificate_data"`
	Thumbprint        string                           `json:"thumbprint"`
	Version           string                           `json:"version"`
	SecretId          string                           `json:"secret_id"`
	Tags              map[string]string                `json:"tags"`
}



type AzurermKeyVaultCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermKeyVaultCertificateList is a list of AzurermKeyVaultCertificates
type AzurermKeyVaultCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermKeyVaultCertificate CRD objects
	Items []AzurermKeyVaultCertificate `json:"items,omitempty"`
}