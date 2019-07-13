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

type AzurermApiManagement struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementStatus `json:"status,omitempty"`
}

type AzurermApiManagementSpecCertificate struct {
	CertificatePassword string `json:"certificate_password"`
	StoreName           string `json:"store_name"`
	EncodedCertificate  string `json:"encoded_certificate"`
}

type AzurermApiManagementSpecSecurity struct {
	DisableFrontendSsl30    bool `json:"disable_frontend_ssl30"`
	DisableFrontendTls10    bool `json:"disable_frontend_tls10"`
	DisableFrontendTls11    bool `json:"disable_frontend_tls11"`
	DisableBackendSsl30     bool `json:"disable_backend_ssl30"`
	DisableBackendTls10     bool `json:"disable_backend_tls10"`
	DisableBackendTls11     bool `json:"disable_backend_tls11"`
	DisableTripleDesChipers bool `json:"disable_triple_des_chipers"`
	DisableTripleDesCiphers bool `json:"disable_triple_des_ciphers"`
}

type AzurermApiManagementSpecSignIn struct {
	Enabled bool `json:"enabled"`
}

type AzurermApiManagementSpecSku struct {
	Capacity int    `json:"capacity"`
	Name     string `json:"name"`
}

type AzurermApiManagementSpecPolicy struct {
	XmlContent string `json:"xml_content"`
	XmlLink    string `json:"xml_link"`
}

type AzurermApiManagementSpecSignUpTermsOfService struct {
	ConsentRequired bool   `json:"consent_required"`
	Text            string `json:"text"`
	Enabled         bool   `json:"enabled"`
}

type AzurermApiManagementSpecSignUp struct {
	Enabled        bool                             `json:"enabled"`
	TermsOfService []AzurermApiManagementSpecSignUp `json:"terms_of_service"`
}

type AzurermApiManagementSpecIdentity struct {
	Type        string `json:"type"`
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
}

type AzurermApiManagementSpecAdditionalLocation struct {
	PublicIpAddresses  []string `json:"public_ip_addresses"`
	Location           string   `json:"location"`
	GatewayRegionalUrl string   `json:"gateway_regional_url"`
}

type AzurermApiManagementSpecHostnameConfigurationProxy struct {
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
	Certificate                string `json:"certificate"`
	CertificatePassword        string `json:"certificate_password"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
	DefaultSslBinding          bool   `json:"default_ssl_binding"`
}

type AzurermApiManagementSpecHostnameConfigurationScm struct {
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
	Certificate                string `json:"certificate"`
	CertificatePassword        string `json:"certificate_password"`
}

type AzurermApiManagementSpecHostnameConfigurationManagement struct {
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
	Certificate                string `json:"certificate"`
	CertificatePassword        string `json:"certificate_password"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
}

type AzurermApiManagementSpecHostnameConfigurationPortal struct {
	Certificate                string `json:"certificate"`
	CertificatePassword        string `json:"certificate_password"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
}

type AzurermApiManagementSpecHostnameConfiguration struct {
	Proxy      []AzurermApiManagementSpecHostnameConfiguration `json:"proxy"`
	Scm        []AzurermApiManagementSpecHostnameConfiguration `json:"scm"`
	Management []AzurermApiManagementSpecHostnameConfiguration `json:"management"`
	Portal     []AzurermApiManagementSpecHostnameConfiguration `json:"portal"`
}

type AzurermApiManagementSpec struct {
	Name                    string                     `json:"name"`
	Certificate             []AzurermApiManagementSpec `json:"certificate"`
	Security                []AzurermApiManagementSpec `json:"security"`
	SignIn                  []AzurermApiManagementSpec `json:"sign_in"`
	Tags                    map[string]string          `json:"tags"`
	PortalUrl               string                     `json:"portal_url"`
	ManagementApiUrl        string                     `json:"management_api_url"`
	Location                string                     `json:"location"`
	PublisherEmail          string                     `json:"publisher_email"`
	Sku                     []AzurermApiManagementSpec `json:"sku"`
	Policy                  []AzurermApiManagementSpec `json:"policy"`
	SignUp                  []AzurermApiManagementSpec `json:"sign_up"`
	GatewayRegionalUrl      string                     `json:"gateway_regional_url"`
	ScmUrl                  string                     `json:"scm_url"`
	ResourceGroupName       string                     `json:"resource_group_name"`
	PublicIpAddresses       []string                   `json:"public_ip_addresses"`
	PublisherName           string                     `json:"publisher_name"`
	NotificationSenderEmail string                     `json:"notification_sender_email"`
	Identity                []AzurermApiManagementSpec `json:"identity"`
	AdditionalLocation      []AzurermApiManagementSpec `json:"additional_location"`
	HostnameConfiguration   []AzurermApiManagementSpec `json:"hostname_configuration"`
	GatewayUrl              string                     `json:"gateway_url"`
}



type AzurermApiManagementStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApiManagementList is a list of AzurermApiManagements
type AzurermApiManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagement CRD objects
	Items []AzurermApiManagement `json:"items,omitempty"`
}