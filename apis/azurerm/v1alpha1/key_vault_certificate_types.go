package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type KeyVaultCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyVaultCertificateSpec   `json:"spec,omitempty"`
	Status            KeyVaultCertificateStatus `json:"status,omitempty"`
}

type KeyVaultCertificateSpecCertificate struct {
	// +optional
}

type KeyVaultCertificateSpecCertificatePolicyIssuerParameters struct {
	Name string `json:"name" tf:"name"`
}

type KeyVaultCertificateSpecCertificatePolicyKeyProperties struct {
	Exportable bool   `json:"exportable" tf:"exportable"`
	KeySize    int    `json:"keySize" tf:"key_size"`
	KeyType    string `json:"keyType" tf:"key_type"`
	ReuseKey   bool   `json:"reuseKey" tf:"reuse_key"`
}

type KeyVaultCertificateSpecCertificatePolicyLifetimeActionAction struct {
	ActionType string `json:"actionType" tf:"action_type"`
}

type KeyVaultCertificateSpecCertificatePolicyLifetimeActionTrigger struct {
	// +optional
	DaysBeforeExpiry int `json:"daysBeforeExpiry,omitempty" tf:"days_before_expiry,omitempty"`
	// +optional
	LifetimePercentage int `json:"lifetimePercentage,omitempty" tf:"lifetime_percentage,omitempty"`
}

type KeyVaultCertificateSpecCertificatePolicyLifetimeAction struct {
	// +kubebuilder:validation:MaxItems=1
	Action []KeyVaultCertificateSpecCertificatePolicyLifetimeActionAction `json:"action" tf:"action"`
	// +kubebuilder:validation:MaxItems=1
	Trigger []KeyVaultCertificateSpecCertificatePolicyLifetimeActionTrigger `json:"trigger" tf:"trigger"`
}

type KeyVaultCertificateSpecCertificatePolicySecretProperties struct {
	ContentType string `json:"contentType" tf:"content_type"`
}

type KeyVaultCertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames struct {
	// +optional
	DnsNames []string `json:"dnsNames,omitempty" tf:"dns_names,omitempty"`
	// +optional
	Emails []string `json:"emails,omitempty" tf:"emails,omitempty"`
	// +optional
	Upns []string `json:"upns,omitempty" tf:"upns,omitempty"`
}

type KeyVaultCertificateSpecCertificatePolicyX509CertificateProperties struct {
	// +optional
	ExtendedKeyUsage []string `json:"extendedKeyUsage,omitempty" tf:"extended_key_usage,omitempty"`
	KeyUsage         []string `json:"keyUsage" tf:"key_usage"`
	Subject          string   `json:"subject" tf:"subject"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SubjectAlternativeNames []KeyVaultCertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`
	ValidityInMonths        int                                                                                        `json:"validityInMonths" tf:"validity_in_months"`
}

type KeyVaultCertificateSpecCertificatePolicy struct {
	// +kubebuilder:validation:MaxItems=1
	IssuerParameters []KeyVaultCertificateSpecCertificatePolicyIssuerParameters `json:"issuerParameters" tf:"issuer_parameters"`
	// +kubebuilder:validation:MaxItems=1
	KeyProperties []KeyVaultCertificateSpecCertificatePolicyKeyProperties `json:"keyProperties" tf:"key_properties"`
	// +optional
	LifetimeAction []KeyVaultCertificateSpecCertificatePolicyLifetimeAction `json:"lifetimeAction,omitempty" tf:"lifetime_action,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	SecretProperties []KeyVaultCertificateSpecCertificatePolicySecretProperties `json:"secretProperties" tf:"secret_properties"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	X509CertificateProperties []KeyVaultCertificateSpecCertificatePolicyX509CertificateProperties `json:"x509CertificateProperties,omitempty" tf:"x509_certificate_properties,omitempty"`
}

type KeyVaultCertificateSpec struct {
	Secret      *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Certificate []KeyVaultCertificateSpecCertificate `json:"certificate,omitempty" tf:"certificate,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	CertificatePolicy []KeyVaultCertificateSpecCertificatePolicy `json:"certificatePolicy" tf:"certificate_policy"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty"`
	Name       string `json:"name" tf:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// Deprecated
	VaultURI string `json:"vaultURI,omitempty" tf:"vault_uri,omitempty"`
}

type KeyVaultCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KeyVaultCertificateList is a list of KeyVaultCertificates
type KeyVaultCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KeyVaultCertificate CRD objects
	Items []KeyVaultCertificate `json:"items,omitempty"`
}
