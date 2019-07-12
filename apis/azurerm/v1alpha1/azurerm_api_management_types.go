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

type AzurermApiManagementSpecIdentity struct {
	Type        string `json:"type"`
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
}

type AzurermApiManagementSpecSku struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

type AzurermApiManagementSpecCertificate struct {
	CertificatePassword string `json:"certificate_password"`
	StoreName           string `json:"store_name"`
	EncodedCertificate  string `json:"encoded_certificate"`
}

type AzurermApiManagementSpecSecurity struct {
	DisableFrontendTls11    bool `json:"disable_frontend_tls11"`
	DisableBackendSsl30     bool `json:"disable_backend_ssl30"`
	DisableBackendTls10     bool `json:"disable_backend_tls10"`
	DisableBackendTls11     bool `json:"disable_backend_tls11"`
	DisableTripleDesChipers bool `json:"disable_triple_des_chipers"`
	DisableTripleDesCiphers bool `json:"disable_triple_des_ciphers"`
	DisableFrontendSsl30    bool `json:"disable_frontend_ssl30"`
	DisableFrontendTls10    bool `json:"disable_frontend_tls10"`
}

type AzurermApiManagementSpecHostnameConfigurationManagement struct {
	KeyVaultId                 string `json:"key_vault_id"`
	Certificate                string `json:"certificate"`
	CertificatePassword        string `json:"certificate_password"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
	HostName                   string `json:"host_name"`
}

type AzurermApiManagementSpecHostnameConfigurationPortal struct {
	CertificatePassword        string `json:"certificate_password"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
	Certificate                string `json:"certificate"`
}

type AzurermApiManagementSpecHostnameConfigurationProxy struct {
	DefaultSslBinding          bool   `json:"default_ssl_binding"`
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
	Certificate                string `json:"certificate"`
	CertificatePassword        string `json:"certificate_password"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
}

type AzurermApiManagementSpecHostnameConfigurationScm struct {
	HostName                   string `json:"host_name"`
	KeyVaultId                 string `json:"key_vault_id"`
	Certificate                string `json:"certificate"`
	CertificatePassword        string `json:"certificate_password"`
	NegotiateClientCertificate bool   `json:"negotiate_client_certificate"`
}

type AzurermApiManagementSpecHostnameConfiguration struct {
	Management []AzurermApiManagementSpecHostnameConfiguration `json:"management"`
	Portal     []AzurermApiManagementSpecHostnameConfiguration `json:"portal"`
	Proxy      []AzurermApiManagementSpecHostnameConfiguration `json:"proxy"`
	Scm        []AzurermApiManagementSpecHostnameConfiguration `json:"scm"`
}

type AzurermApiManagementSpecSignIn struct {
	Enabled bool `json:"enabled"`
}

type AzurermApiManagementSpecAdditionalLocation struct {
	Location           string   `json:"location"`
	GatewayRegionalUrl string   `json:"gateway_regional_url"`
	PublicIpAddresses  []string `json:"public_ip_addresses"`
}

type AzurermApiManagementSpecPolicy struct {
	XmlContent string `json:"xml_content"`
	XmlLink    string `json:"xml_link"`
}

type AzurermApiManagementSpecSignUpTermsOfService struct {
	Enabled         bool   `json:"enabled"`
	ConsentRequired bool   `json:"consent_required"`
	Text            string `json:"text"`
}

type AzurermApiManagementSpecSignUp struct {
	Enabled        bool                             `json:"enabled"`
	TermsOfService []AzurermApiManagementSpecSignUp `json:"terms_of_service"`
}

type AzurermApiManagementSpec struct {
	Location                string                     `json:"location"`
	Identity                []AzurermApiManagementSpec `json:"identity"`
	NotificationSenderEmail string                     `json:"notification_sender_email"`
	GatewayUrl              string                     `json:"gateway_url"`
	ResourceGroupName       string                     `json:"resource_group_name"`
	ScmUrl                  string                     `json:"scm_url"`
	PublicIpAddresses       []string                   `json:"public_ip_addresses"`
	PublisherEmail          string                     `json:"publisher_email"`
	Sku                     []AzurermApiManagementSpec `json:"sku"`
	Certificate             []AzurermApiManagementSpec `json:"certificate"`
	Security                []AzurermApiManagementSpec `json:"security"`
	HostnameConfiguration   []AzurermApiManagementSpec `json:"hostname_configuration"`
	SignIn                  []AzurermApiManagementSpec `json:"sign_in"`
	ManagementApiUrl        string                     `json:"management_api_url"`
	Name                    string                     `json:"name"`
	PublisherName           string                     `json:"publisher_name"`
	AdditionalLocation      []AzurermApiManagementSpec `json:"additional_location"`
	Policy                  []AzurermApiManagementSpec `json:"policy"`
	SignUp                  []AzurermApiManagementSpec `json:"sign_up"`
	Tags                    map[string]string          `json:"tags"`
	GatewayRegionalUrl      string                     `json:"gateway_regional_url"`
	PortalUrl               string                     `json:"portal_url"`
}

type AzurermApiManagementStatus struct {
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
