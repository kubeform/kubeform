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

type AzurermApplicationGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApplicationGatewaySpec   `json:"spec,omitempty"`
	Status            AzurermApplicationGatewayStatus `json:"status,omitempty"`
}

type AzurermApplicationGatewaySpecFrontendIpConfiguration struct {
	PrivateIpAddressAllocation string `json:"private_ip_address_allocation"`
	Id                         string `json:"id"`
	Name                       string `json:"name"`
	SubnetId                   string `json:"subnet_id"`
	PrivateIpAddress           string `json:"private_ip_address"`
	PublicIpAddressId          string `json:"public_ip_address_id"`
}

type AzurermApplicationGatewaySpecFrontendPort struct {
	Name string `json:"name"`
	Port int    `json:"port"`
	Id   string `json:"id"`
}

type AzurermApplicationGatewaySpecHttpListenerCustomErrorConfiguration struct {
	StatusCode         string `json:"status_code"`
	CustomErrorPageUrl string `json:"custom_error_page_url"`
	Id                 string `json:"id"`
}

type AzurermApplicationGatewaySpecHttpListener struct {
	FrontendPortName            string                                      `json:"frontend_port_name"`
	Protocol                    string                                      `json:"protocol"`
	HostName                    string                                      `json:"host_name"`
	SslCertificateName          string                                      `json:"ssl_certificate_name"`
	FrontendIpConfigurationId   string                                      `json:"frontend_ip_configuration_id"`
	SslCertificateId            string                                      `json:"ssl_certificate_id"`
	CustomErrorConfiguration    []AzurermApplicationGatewaySpecHttpListener `json:"custom_error_configuration"`
	Name                        string                                      `json:"name"`
	RequireSni                  bool                                        `json:"require_sni"`
	FrontendPortId              string                                      `json:"frontend_port_id"`
	Id                          string                                      `json:"id"`
	FrontendIpConfigurationName string                                      `json:"frontend_ip_configuration_name"`
}

type AzurermApplicationGatewaySpecRequestRoutingRule struct {
	BackendHttpSettingsId     string `json:"backend_http_settings_id"`
	UrlPathMapId              string `json:"url_path_map_id"`
	BackendAddressPoolName    string `json:"backend_address_pool_name"`
	BackendHttpSettingsName   string `json:"backend_http_settings_name"`
	Id                        string `json:"id"`
	UrlPathMapName            string `json:"url_path_map_name"`
	BackendAddressPoolId      string `json:"backend_address_pool_id"`
	HttpListenerId            string `json:"http_listener_id"`
	Name                      string `json:"name"`
	RuleType                  string `json:"rule_type"`
	HttpListenerName          string `json:"http_listener_name"`
	RedirectConfigurationName string `json:"redirect_configuration_name"`
	RewriteRuleSetName        string `json:"rewrite_rule_set_name"`
	RedirectConfigurationId   string `json:"redirect_configuration_id"`
	RewriteRuleSetId          string `json:"rewrite_rule_set_id"`
}

type AzurermApplicationGatewaySpecRedirectConfiguration struct {
	IncludePath        bool   `json:"include_path"`
	IncludeQueryString bool   `json:"include_query_string"`
	Id                 string `json:"id"`
	TargetListenerId   string `json:"target_listener_id"`
	Name               string `json:"name"`
	RedirectType       string `json:"redirect_type"`
	TargetListenerName string `json:"target_listener_name"`
	TargetUrl          string `json:"target_url"`
}

type AzurermApplicationGatewaySpecAutoscaleConfiguration struct {
	MaxCapacity int `json:"max_capacity"`
	MinCapacity int `json:"min_capacity"`
}

type AzurermApplicationGatewaySpecSslPolicy struct {
	DisabledProtocols  []string `json:"disabled_protocols"`
	PolicyType         string   `json:"policy_type"`
	PolicyName         string   `json:"policy_name"`
	CipherSuites       []string `json:"cipher_suites"`
	MinProtocolVersion string   `json:"min_protocol_version"`
}

type AzurermApplicationGatewaySpecUrlPathMapPathRule struct {
	Name                      string   `json:"name"`
	BackendAddressPoolName    string   `json:"backend_address_pool_name"`
	BackendAddressPoolId      string   `json:"backend_address_pool_id"`
	RewriteRuleSetId          string   `json:"rewrite_rule_set_id"`
	Id                        string   `json:"id"`
	Paths                     []string `json:"paths"`
	BackendHttpSettingsName   string   `json:"backend_http_settings_name"`
	RedirectConfigurationName string   `json:"redirect_configuration_name"`
	RewriteRuleSetName        string   `json:"rewrite_rule_set_name"`
	BackendHttpSettingsId     string   `json:"backend_http_settings_id"`
	RedirectConfigurationId   string   `json:"redirect_configuration_id"`
}

type AzurermApplicationGatewaySpecUrlPathMap struct {
	DefaultRedirectConfigurationName string                                    `json:"default_redirect_configuration_name"`
	PathRule                         []AzurermApplicationGatewaySpecUrlPathMap `json:"path_rule"`
	DefaultBackendHttpSettingsId     string                                    `json:"default_backend_http_settings_id"`
	DefaultRewriteRuleSetId          string                                    `json:"default_rewrite_rule_set_id"`
	DefaultBackendAddressPoolName    string                                    `json:"default_backend_address_pool_name"`
	DefaultBackendHttpSettingsName   string                                    `json:"default_backend_http_settings_name"`
	DefaultRewriteRuleSetName        string                                    `json:"default_rewrite_rule_set_name"`
	DefaultBackendAddressPoolId      string                                    `json:"default_backend_address_pool_id"`
	DefaultRedirectConfigurationId   string                                    `json:"default_redirect_configuration_id"`
	Id                               string                                    `json:"id"`
	Name                             string                                    `json:"name"`
}

type AzurermApplicationGatewaySpecWafConfigurationDisabledRuleGroup struct {
	RuleGroupName string  `json:"rule_group_name"`
	Rules         []int64 `json:"rules"`
}

type AzurermApplicationGatewaySpecWafConfigurationExclusion struct {
	MatchVariable         string `json:"match_variable"`
	SelectorMatchOperator string `json:"selector_match_operator"`
	Selector              string `json:"selector"`
}

type AzurermApplicationGatewaySpecWafConfiguration struct {
	RequestBodyCheck     bool                                            `json:"request_body_check"`
	DisabledRuleGroup    []AzurermApplicationGatewaySpecWafConfiguration `json:"disabled_rule_group"`
	RuleSetType          string                                          `json:"rule_set_type"`
	RuleSetVersion       string                                          `json:"rule_set_version"`
	FileUploadLimitMb    int                                             `json:"file_upload_limit_mb"`
	MaxRequestBodySizeKb int                                             `json:"max_request_body_size_kb"`
	Exclusion            []AzurermApplicationGatewaySpecWafConfiguration `json:"exclusion"`
	Enabled              bool                                            `json:"enabled"`
	FirewallMode         string                                          `json:"firewall_mode"`
}

type AzurermApplicationGatewaySpecBackendAddressPool struct {
	Name          string   `json:"name"`
	Fqdns         []string `json:"fqdns"`
	IpAddresses   []string `json:"ip_addresses"`
	FqdnList      []string `json:"fqdn_list"`
	IpAddressList []string `json:"ip_address_list"`
	Id            string   `json:"id"`
}

type AzurermApplicationGatewaySpecSku struct {
	Name     string `json:"name"`
	Tier     string `json:"tier"`
	Capacity int    `json:"capacity"`
}

type AzurermApplicationGatewaySpecProbeMatch struct {
	Body       string   `json:"body"`
	StatusCode []string `json:"status_code"`
}

type AzurermApplicationGatewaySpecProbe struct {
	Path                                string                               `json:"path"`
	Interval                            int                                  `json:"interval"`
	PickHostNameFromBackendHttpSettings bool                                 `json:"pick_host_name_from_backend_http_settings"`
	MinimumServers                      int                                  `json:"minimum_servers"`
	Id                                  string                               `json:"id"`
	Name                                string                               `json:"name"`
	Protocol                            string                               `json:"protocol"`
	UnhealthyThreshold                  int                                  `json:"unhealthy_threshold"`
	Match                               []AzurermApplicationGatewaySpecProbe `json:"match"`
	Host                                string                               `json:"host"`
	Timeout                             int                                  `json:"timeout"`
}

type AzurermApplicationGatewaySpecBackendHttpSettingsConnectionDraining struct {
	Enabled         bool `json:"enabled"`
	DrainTimeoutSec int  `json:"drain_timeout_sec"`
}

type AzurermApplicationGatewaySpecBackendHttpSettingsAuthenticationCertificate struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type AzurermApplicationGatewaySpecBackendHttpSettings struct {
	CookieBasedAffinity            string                                             `json:"cookie_based_affinity"`
	AffinityCookieName             string                                             `json:"affinity_cookie_name"`
	ConnectionDraining             []AzurermApplicationGatewaySpecBackendHttpSettings `json:"connection_draining"`
	ProbeName                      string                                             `json:"probe_name"`
	ProbeId                        string                                             `json:"probe_id"`
	Path                           string                                             `json:"path"`
	Port                           int                                                `json:"port"`
	Protocol                       string                                             `json:"protocol"`
	AuthenticationCertificate      []AzurermApplicationGatewaySpecBackendHttpSettings `json:"authentication_certificate"`
	Id                             string                                             `json:"id"`
	HostName                       string                                             `json:"host_name"`
	Name                           string                                             `json:"name"`
	PickHostNameFromBackendAddress bool                                               `json:"pick_host_name_from_backend_address"`
	RequestTimeout                 int                                                `json:"request_timeout"`
}

type AzurermApplicationGatewaySpecGatewayIpConfiguration struct {
	Name     string `json:"name"`
	SubnetId string `json:"subnet_id"`
	Id       string `json:"id"`
}

type AzurermApplicationGatewaySpecAuthenticationCertificate struct {
	Name string `json:"name"`
	Data string `json:"data"`
	Id   string `json:"id"`
}

type AzurermApplicationGatewaySpecRewriteRuleSetRewriteRuleCondition struct {
	Variable   string `json:"variable"`
	Pattern    string `json:"pattern"`
	IgnoreCase bool   `json:"ignore_case"`
	Negate     bool   `json:"negate"`
}

type AzurermApplicationGatewaySpecRewriteRuleSetRewriteRuleRequestHeaderConfiguration struct {
	HeaderName  string `json:"header_name"`
	HeaderValue string `json:"header_value"`
}

type AzurermApplicationGatewaySpecRewriteRuleSetRewriteRuleResponseHeaderConfiguration struct {
	HeaderName  string `json:"header_name"`
	HeaderValue string `json:"header_value"`
}

type AzurermApplicationGatewaySpecRewriteRuleSetRewriteRule struct {
	RuleSequence                int                                                      `json:"rule_sequence"`
	Condition                   []AzurermApplicationGatewaySpecRewriteRuleSetRewriteRule `json:"condition"`
	RequestHeaderConfiguration  []AzurermApplicationGatewaySpecRewriteRuleSetRewriteRule `json:"request_header_configuration"`
	ResponseHeaderConfiguration []AzurermApplicationGatewaySpecRewriteRuleSetRewriteRule `json:"response_header_configuration"`
	Name                        string                                                   `json:"name"`
}

type AzurermApplicationGatewaySpecRewriteRuleSet struct {
	Name        string                                        `json:"name"`
	RewriteRule []AzurermApplicationGatewaySpecRewriteRuleSet `json:"rewrite_rule"`
	Id          string                                        `json:"id"`
}

type AzurermApplicationGatewaySpecSslCertificate struct {
	Data           string `json:"data"`
	Password       string `json:"password"`
	Id             string `json:"id"`
	PublicCertData string `json:"public_cert_data"`
	Name           string `json:"name"`
}

type AzurermApplicationGatewaySpecCustomErrorConfiguration struct {
	StatusCode         string `json:"status_code"`
	CustomErrorPageUrl string `json:"custom_error_page_url"`
	Id                 string `json:"id"`
}

type AzurermApplicationGatewaySpec struct {
	DisabledSslProtocols      []string                        `json:"disabled_ssl_protocols"`
	Location                  string                          `json:"location"`
	FrontendIpConfiguration   []AzurermApplicationGatewaySpec `json:"frontend_ip_configuration"`
	FrontendPort              []AzurermApplicationGatewaySpec `json:"frontend_port"`
	HttpListener              []AzurermApplicationGatewaySpec `json:"http_listener"`
	RequestRoutingRule        []AzurermApplicationGatewaySpec `json:"request_routing_rule"`
	RedirectConfiguration     []AzurermApplicationGatewaySpec `json:"redirect_configuration"`
	AutoscaleConfiguration    []AzurermApplicationGatewaySpec `json:"autoscale_configuration"`
	SslPolicy                 []AzurermApplicationGatewaySpec `json:"ssl_policy"`
	UrlPathMap                []AzurermApplicationGatewaySpec `json:"url_path_map"`
	WafConfiguration          []AzurermApplicationGatewaySpec `json:"waf_configuration"`
	Zones                     []string                        `json:"zones"`
	BackendAddressPool        []AzurermApplicationGatewaySpec `json:"backend_address_pool"`
	Sku                       []AzurermApplicationGatewaySpec `json:"sku"`
	Probe                     []AzurermApplicationGatewaySpec `json:"probe"`
	ResourceGroupName         string                          `json:"resource_group_name"`
	BackendHttpSettings       []AzurermApplicationGatewaySpec `json:"backend_http_settings"`
	GatewayIpConfiguration    []AzurermApplicationGatewaySpec `json:"gateway_ip_configuration"`
	Name                      string                          `json:"name"`
	AuthenticationCertificate []AzurermApplicationGatewaySpec `json:"authentication_certificate"`
	EnableHttp2               bool                            `json:"enable_http2"`
	RewriteRuleSet            []AzurermApplicationGatewaySpec `json:"rewrite_rule_set"`
	SslCertificate            []AzurermApplicationGatewaySpec `json:"ssl_certificate"`
	CustomErrorConfiguration  []AzurermApplicationGatewaySpec `json:"custom_error_configuration"`
	Tags                      map[string]string               `json:"tags"`
}



type AzurermApplicationGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApplicationGatewayList is a list of AzurermApplicationGateways
type AzurermApplicationGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApplicationGateway CRD objects
	Items []AzurermApplicationGateway `json:"items,omitempty"`
}