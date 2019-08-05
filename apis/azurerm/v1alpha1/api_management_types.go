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

type ApiManagement struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementSpec   `json:"spec,omitempty"`
	Status            ApiManagementStatus `json:"status,omitempty"`
}

type ApiManagementSpecAdditionalLocation struct {
	Location string `json:"location" tf:"location"`
}

type ApiManagementSpecCertificate struct {
	CertificatePassword string `json:"-" sensitive:"true" tf:"certificate_password"`
	EncodedCertificate  string `json:"-" sensitive:"true" tf:"encoded_certificate"`
	StoreName           string `json:"storeName" tf:"store_name"`
}

type ApiManagementSpecHostnameConfigurationManagement struct {
	// +optional
	Certificate string `json:"-" sensitive:"true" tf:"certificate,omitempty"`
	// +optional
	CertificatePassword string `json:"-" sensitive:"true" tf:"certificate_password,omitempty"`
	HostName            string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty"`
	// +optional
	NegotiateClientCertificate bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`
}

type ApiManagementSpecHostnameConfigurationPortal struct {
	// +optional
	Certificate string `json:"-" sensitive:"true" tf:"certificate,omitempty"`
	// +optional
	CertificatePassword string `json:"-" sensitive:"true" tf:"certificate_password,omitempty"`
	HostName            string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty"`
	// +optional
	NegotiateClientCertificate bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`
}

type ApiManagementSpecHostnameConfigurationProxy struct {
	// +optional
	Certificate string `json:"-" sensitive:"true" tf:"certificate,omitempty"`
	// +optional
	CertificatePassword string `json:"-" sensitive:"true" tf:"certificate_password,omitempty"`
	// +optional
	DefaultSslBinding bool   `json:"defaultSslBinding,omitempty" tf:"default_ssl_binding,omitempty"`
	HostName          string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty"`
	// +optional
	NegotiateClientCertificate bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`
}

type ApiManagementSpecHostnameConfigurationScm struct {
	// +optional
	Certificate string `json:"-" sensitive:"true" tf:"certificate,omitempty"`
	// +optional
	CertificatePassword string `json:"-" sensitive:"true" tf:"certificate_password,omitempty"`
	HostName            string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty"`
	// +optional
	NegotiateClientCertificate bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`
}

type ApiManagementSpecHostnameConfiguration struct {
	// +optional
	Management []ApiManagementSpecHostnameConfigurationManagement `json:"management,omitempty" tf:"management,omitempty"`
	// +optional
	Portal []ApiManagementSpecHostnameConfigurationPortal `json:"portal,omitempty" tf:"portal,omitempty"`
	// +optional
	Proxy []ApiManagementSpecHostnameConfigurationProxy `json:"proxy,omitempty" tf:"proxy,omitempty"`
	// +optional
	Scm []ApiManagementSpecHostnameConfigurationScm `json:"scm,omitempty" tf:"scm,omitempty"`
}

type ApiManagementSpecIdentity struct {
	Type string `json:"type" tf:"type"`
}

type ApiManagementSpecPolicy struct {
	// +optional
	XmlContent string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`
	// +optional
	XmlLink string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

type ApiManagementSpecSecurity struct {
	// +optional
	DisableBackendSsl30 bool `json:"disableBackendSsl30,omitempty" tf:"disable_backend_ssl30,omitempty"`
	// +optional
	DisableBackendTls10 bool `json:"disableBackendTls10,omitempty" tf:"disable_backend_tls10,omitempty"`
	// +optional
	DisableBackendTls11 bool `json:"disableBackendTls11,omitempty" tf:"disable_backend_tls11,omitempty"`
	// +optional
	DisableFrontendSsl30 bool `json:"disableFrontendSsl30,omitempty" tf:"disable_frontend_ssl30,omitempty"`
	// +optional
	DisableFrontendTls10 bool `json:"disableFrontendTls10,omitempty" tf:"disable_frontend_tls10,omitempty"`
	// +optional
	DisableFrontendTls11 bool `json:"disableFrontendTls11,omitempty" tf:"disable_frontend_tls11,omitempty"`
	// +optional
	// Deprecated
	DisableTripleDESChipers bool `json:"disableTripleDESChipers,omitempty" tf:"disable_triple_des_chipers,omitempty"`
	// +optional
	DisableTripleDESCiphers bool `json:"disableTripleDESCiphers,omitempty" tf:"disable_triple_des_ciphers,omitempty"`
}

type ApiManagementSpecSignIn struct {
	Enabled bool `json:"enabled" tf:"enabled"`
}

type ApiManagementSpecSignUpTermsOfService struct {
	ConsentRequired bool `json:"consentRequired" tf:"consent_required"`
	Enabled         bool `json:"enabled" tf:"enabled"`
	// +optional
	Text string `json:"text,omitempty" tf:"text,omitempty"`
}

type ApiManagementSpecSignUp struct {
	Enabled bool `json:"enabled" tf:"enabled"`
	// +kubebuilder:validation:MaxItems=1
	TermsOfService []ApiManagementSpecSignUpTermsOfService `json:"termsOfService" tf:"terms_of_service"`
}

type ApiManagementSpecSku struct {
	Capacity int    `json:"capacity" tf:"capacity"`
	Name     string `json:"name" tf:"name"`
}

type ApiManagementSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	AdditionalLocation []ApiManagementSpecAdditionalLocation `json:"additionalLocation,omitempty" tf:"additional_location,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Certificate []ApiManagementSpecCertificate `json:"certificate,omitempty" tf:"certificate,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HostnameConfiguration []ApiManagementSpecHostnameConfiguration `json:"hostnameConfiguration,omitempty" tf:"hostname_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity []ApiManagementSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	Location string                      `json:"location" tf:"location"`
	Name     string                      `json:"name" tf:"name"`
	// +optional
	NotificationSenderEmail string `json:"notificationSenderEmail,omitempty" tf:"notification_sender_email,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Policy            []ApiManagementSpecPolicy `json:"policy,omitempty" tf:"policy,omitempty"`
	PublisherEmail    string                    `json:"publisherEmail" tf:"publisher_email"`
	PublisherName     string                    `json:"publisherName" tf:"publisher_name"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Security []ApiManagementSpecSecurity `json:"security,omitempty" tf:"security,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SignIn []ApiManagementSpecSignIn `json:"signIn,omitempty" tf:"sign_in,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SignUp []ApiManagementSpecSignUp `json:"signUp,omitempty" tf:"sign_up,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Sku []ApiManagementSpecSku `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ApiManagementStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiManagementSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementList is a list of ApiManagements
type ApiManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagement CRD objects
	Items []ApiManagement `json:"items,omitempty"`
}
