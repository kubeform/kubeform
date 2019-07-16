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

type KeyVaultCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyVaultCertificateSpec   `json:"spec,omitempty"`
	Status            KeyVaultCertificateStatus `json:"status,omitempty"`
}

type KeyVaultCertificateSpecCertificate struct {
	Contents string `json:"contents"`
	// +optional
	Password string `json:"password,omitempty"`
}

type KeyVaultCertificateSpecCertificatePolicyIssuerParameters struct {
	Name string `json:"name"`
}

type KeyVaultCertificateSpecCertificatePolicyKeyProperties struct {
	Exportable bool   `json:"exportable"`
	KeySize    int    `json:"key_size"`
	KeyType    string `json:"key_type"`
	ReuseKey   bool   `json:"reuse_key"`
}

type KeyVaultCertificateSpecCertificatePolicyLifetimeActionAction struct {
	ActionType string `json:"action_type"`
}

type KeyVaultCertificateSpecCertificatePolicyLifetimeActionTrigger struct {
	// +optional
	DaysBeforeExpiry int `json:"days_before_expiry,omitempty"`
	// +optional
	LifetimePercentage int `json:"lifetime_percentage,omitempty"`
}

type KeyVaultCertificateSpecCertificatePolicyLifetimeAction struct {
	// +kubebuilder:validation:MaxItems=1
	Action []KeyVaultCertificateSpecCertificatePolicyLifetimeAction `json:"action"`
	// +kubebuilder:validation:MaxItems=1
	Trigger []KeyVaultCertificateSpecCertificatePolicyLifetimeAction `json:"trigger"`
}

type KeyVaultCertificateSpecCertificatePolicySecretProperties struct {
	ContentType string `json:"content_type"`
}

type KeyVaultCertificateSpecCertificatePolicy struct {
	// +kubebuilder:validation:MaxItems=1
	IssuerParameters []KeyVaultCertificateSpecCertificatePolicy `json:"issuer_parameters"`
	// +kubebuilder:validation:MaxItems=1
	KeyProperties []KeyVaultCertificateSpecCertificatePolicy `json:"key_properties"`
	// +optional
	LifetimeAction *[]KeyVaultCertificateSpecCertificatePolicy `json:"lifetime_action,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	SecretProperties []KeyVaultCertificateSpecCertificatePolicy `json:"secret_properties"`
}

type KeyVaultCertificateSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Certificate *[]KeyVaultCertificateSpec `json:"certificate,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	CertificatePolicy []KeyVaultCertificateSpec `json:"certificate_policy"`
	Name              string                    `json:"name"`
}

type KeyVaultCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
