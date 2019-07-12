package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermApplicationGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApplicationGatewaySpec   `json:"spec,omitempty"`
	Status            AzurermApplicationGatewayStatus `json:"status,omitempty"`
}

type AzurermApplicationGatewaySpecSku struct {
	Name     string `json:"name"`
	Tier     string `json:"tier"`
	Capacity int    `json:"capacity"`
}

type AzurermApplicationGatewaySpecBackendAddressPool struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Fqdns         []string `json:"fqdns"`
	IpAddresses   []string `json:"ip_addresses"`
	FqdnList      []string `json:"fqdn_list"`
	IpAddressList []string `json:"ip_address_list"`
}

type AzurermApplicationGatewaySpecHttpListenerCustomErrorConfiguration struct {
	StatusCode         string `json:"status_code"`
	CustomErrorPageUrl string `json:"custom_error_page_url"`
	Id                 string `json:"id"`
}

type AzurermApplicationGatewaySpecHttpListener struct {
	SslCertificateId            string                                      `json:"ssl_certificate_id"`
	FrontendIpConfigurationName string                                      `json:"frontend_ip_configuration_name"`
	RequireSni                  bool                                        `json:"require_sni"`
	FrontendPortId              string                                      `json:"frontend_port_id"`
	HostName                    string                                      `json:"host_name"`
	SslCertificateName          string                                      `json:"ssl_certificate_name"`
	FrontendIpConfigurationId   string                                      `json:"frontend_ip_configuration_id"`
	Id                          string                                      `json:"id"`
	CustomErrorConfiguration    []AzurermApplicationGatewaySpecHttpListener `json:"custom_error_configuration"`
	Name                        string                                      `json:"name"`
	FrontendPortName            string                                      `json:"frontend_port_name"`
	Protocol                    string                                      `json:"protocol"`
}

type AzurermApplicationGatewaySpecFrontendPort struct {
	Name string `json:"name"`
	Port int    `json:"port"`
	Id   string `json:"id"`
}

type AzurermApplicationGatewaySpecSslPolicy struct {
	PolicyName         string   `json:"policy_name"`
	CipherSuites       []string `json:"cipher_suites"`
	MinProtocolVersion string   `json:"min_protocol_version"`
	DisabledProtocols  []string `json:"disabled_protocols"`
	PolicyType         string   `json:"policy_type"`
}

type AzurermApplicationGatewaySpecProbeMatch struct {
	Body       string   `json:"body"`
	StatusCode []string `json:"status_code"`
}

type AzurermApplicationGatewaySpecProbe struct {
	Interval                            int                                  `json:"interval"`
	MinimumServers                      int                                  `json:"minimum_servers"`
	Match                               []AzurermApplicationGatewaySpecProbe `json:"match"`
	Name                                string                               `json:"name"`
	Path                                string                               `json:"path"`
	Host                                string                               `json:"host"`
	PickHostNameFromBackendHttpSettings bool                                 `json:"pick_host_name_from_backend_http_settings"`
	Id                                  string                               `json:"id"`
	Protocol                            string                               `json:"protocol"`
	Timeout                             int                                  `json:"timeout"`
	UnhealthyThreshold                  int                                  `json:"unhealthy_threshold"`
}

type AzurermApplicationGatewaySpecSslCertificate struct {
	Name           string `json:"name"`
	Data           string `json:"data"`
	Password       string `json:"password"`
	Id             string `json:"id"`
	PublicCertData string `json:"public_cert_data"`
}

type AzurermApplicationGatewaySpecBackendHttpSettingsAuthenticationCertificate struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type AzurermApplicationGatewaySpecBackendHttpSettingsConnectionDraining struct {
	Enabled         bool `json:"enabled"`
	DrainTimeoutSec int  `json:"drain_timeout_sec"`
}

type AzurermApplicationGatewaySpecBackendHttpSettings struct {
	HostName                       string                                             `json:"host_name"`
	Path                           string                                             `json:"path"`
	RequestTimeout                 int                                                `json:"request_timeout"`
	AuthenticationCertificate      []AzurermApplicationGatewaySpecBackendHttpSettings `json:"authentication_certificate"`
	Id                             string                                             `json:"id"`
	Name                           string                                             `json:"name"`
	ProbeId                        string                                             `json:"probe_id"`
	CookieBasedAffinity            string                                             `json:"cookie_based_affinity"`
	Protocol                       string                                             `json:"protocol"`
	AffinityCookieName             string                                             `json:"affinity_cookie_name"`
	PickHostNameFromBackendAddress bool                                               `json:"pick_host_name_from_backend_address"`
	ConnectionDraining             []AzurermApplicationGatewaySpecBackendHttpSettings `json:"connection_draining"`
	ProbeName                      string                                             `json:"probe_name"`
	Port                           int                                                `json:"port"`
}

type AzurermApplicationGatewaySpecFrontendIpConfiguration struct {
	Name                       string `json:"name"`
	SubnetId                   string `json:"subnet_id"`
	PrivateIpAddress           string `json:"private_ip_address"`
	PublicIpAddressId          string `json:"public_ip_address_id"`
	PrivateIpAddressAllocation string `json:"private_ip_address_allocation"`
	Id                         string `json:"id"`
}

type AzurermApplicationGatewaySpecRewriteRuleSetRewriteRuleResponseHeaderConfiguration struct {
	HeaderName  string `json:"header_name"`
	HeaderValue string `json:"header_value"`
}

type AzurermApplicationGatewaySpecRewriteRuleSetRewriteRuleCondition struct {
	Pattern    string `json:"pattern"`
	IgnoreCase bool   `json:"ignore_case"`
	Negate     bool   `json:"negate"`
	Variable   string `json:"variable"`
}

type AzurermApplicationGatewaySpecRewriteRuleSetRewriteRuleRequestHeaderConfiguration struct {
	HeaderName  string `json:"header_name"`
	HeaderValue string `json:"header_value"`
}

type AzurermApplicationGatewaySpecRewriteRuleSetRewriteRule struct {
	ResponseHeaderConfiguration []AzurermApplicationGatewaySpecRewriteRuleSetRewriteRule `json:"response_header_configuration"`
	Name                        string                                                   `json:"name"`
	RuleSequence                int                                                      `json:"rule_sequence"`
	Condition                   []AzurermApplicationGatewaySpecRewriteRuleSetRewriteRule `json:"condition"`
	RequestHeaderConfiguration  []AzurermApplicationGatewaySpecRewriteRuleSetRewriteRule `json:"request_header_configuration"`
}

type AzurermApplicationGatewaySpecRewriteRuleSet struct {
	Name        string                                        `json:"name"`
	RewriteRule []AzurermApplicationGatewaySpecRewriteRuleSet `json:"rewrite_rule"`
	Id          string                                        `json:"id"`
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
	FileUploadLimitMb    int                                             `json:"file_upload_limit_mb"`
	DisabledRuleGroup    []AzurermApplicationGatewaySpecWafConfiguration `json:"disabled_rule_group"`
	Enabled              bool                                            `json:"enabled"`
	RuleSetType          string                                          `json:"rule_set_type"`
	RequestBodyCheck     bool                                            `json:"request_body_check"`
	MaxRequestBodySizeKb int                                             `json:"max_request_body_size_kb"`
	Exclusion            []AzurermApplicationGatewaySpecWafConfiguration `json:"exclusion"`
	FirewallMode         string                                          `json:"firewall_mode"`
	RuleSetVersion       string                                          `json:"rule_set_version"`
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

type AzurermApplicationGatewaySpecAuthenticationCertificate struct {
	Name string `json:"name"`
	Data string `json:"data"`
	Id   string `json:"id"`
}

type AzurermApplicationGatewaySpecUrlPathMapPathRule struct {
	Id                        string   `json:"id"`
	Paths                     []string `json:"paths"`
	BackendHttpSettingsName   string   `json:"backend_http_settings_name"`
	BackendAddressPoolId      string   `json:"backend_address_pool_id"`
	RewriteRuleSetId          string   `json:"rewrite_rule_set_id"`
	BackendHttpSettingsId     string   `json:"backend_http_settings_id"`
	RedirectConfigurationId   string   `json:"redirect_configuration_id"`
	Name                      string   `json:"name"`
	BackendAddressPoolName    string   `json:"backend_address_pool_name"`
	RedirectConfigurationName string   `json:"redirect_configuration_name"`
	RewriteRuleSetName        string   `json:"rewrite_rule_set_name"`
}

type AzurermApplicationGatewaySpecUrlPathMap struct {
	PathRule                         []AzurermApplicationGatewaySpecUrlPathMap `json:"path_rule"`
	DefaultBackendAddressPoolId      string                                    `json:"default_backend_address_pool_id"`
	DefaultBackendHttpSettingsId     string                                    `json:"default_backend_http_settings_id"`
	DefaultRedirectConfigurationId   string                                    `json:"default_redirect_configuration_id"`
	DefaultRewriteRuleSetId          string                                    `json:"default_rewrite_rule_set_id"`
	Id                               string                                    `json:"id"`
	Name                             string                                    `json:"name"`
	DefaultBackendAddressPoolName    string                                    `json:"default_backend_address_pool_name"`
	DefaultBackendHttpSettingsName   string                                    `json:"default_backend_http_settings_name"`
	DefaultRedirectConfigurationName string                                    `json:"default_redirect_configuration_name"`
	DefaultRewriteRuleSetName        string                                    `json:"default_rewrite_rule_set_name"`
}

type AzurermApplicationGatewaySpecCustomErrorConfiguration struct {
	Id                 string `json:"id"`
	StatusCode         string `json:"status_code"`
	CustomErrorPageUrl string `json:"custom_error_page_url"`
}

type AzurermApplicationGatewaySpecGatewayIpConfiguration struct {
	Name     string `json:"name"`
	SubnetId string `json:"subnet_id"`
	Id       string `json:"id"`
}

type AzurermApplicationGatewaySpecRequestRoutingRule struct {
	RewriteRuleSetName        string `json:"rewrite_rule_set_name"`
	BackendAddressPoolId      string `json:"backend_address_pool_id"`
	BackendHttpSettingsId     string `json:"backend_http_settings_id"`
	Id                        string `json:"id"`
	RewriteRuleSetId          string `json:"rewrite_rule_set_id"`
	RedirectConfigurationName string `json:"redirect_configuration_name"`
	BackendAddressPoolName    string `json:"backend_address_pool_name"`
	UrlPathMapId              string `json:"url_path_map_id"`
	RedirectConfigurationId   string `json:"redirect_configuration_id"`
	Name                      string `json:"name"`
	HttpListenerName          string `json:"http_listener_name"`
	BackendHttpSettingsName   string `json:"backend_http_settings_name"`
	UrlPathMapName            string `json:"url_path_map_name"`
	HttpListenerId            string `json:"http_listener_id"`
	RuleType                  string `json:"rule_type"`
}

type AzurermApplicationGatewaySpec struct {
	Sku                       []AzurermApplicationGatewaySpec `json:"sku"`
	BackendAddressPool        []AzurermApplicationGatewaySpec `json:"backend_address_pool"`
	HttpListener              []AzurermApplicationGatewaySpec `json:"http_listener"`
	FrontendPort              []AzurermApplicationGatewaySpec `json:"frontend_port"`
	SslPolicy                 []AzurermApplicationGatewaySpec `json:"ssl_policy"`
	EnableHttp2               bool                            `json:"enable_http2"`
	Probe                     []AzurermApplicationGatewaySpec `json:"probe"`
	SslCertificate            []AzurermApplicationGatewaySpec `json:"ssl_certificate"`
	Zones                     []string                        `json:"zones"`
	ResourceGroupName         string                          `json:"resource_group_name"`
	BackendHttpSettings       []AzurermApplicationGatewaySpec `json:"backend_http_settings"`
	FrontendIpConfiguration   []AzurermApplicationGatewaySpec `json:"frontend_ip_configuration"`
	RewriteRuleSet            []AzurermApplicationGatewaySpec `json:"rewrite_rule_set"`
	WafConfiguration          []AzurermApplicationGatewaySpec `json:"waf_configuration"`
	Tags                      map[string]string               `json:"tags"`
	Name                      string                          `json:"name"`
	Location                  string                          `json:"location"`
	RedirectConfiguration     []AzurermApplicationGatewaySpec `json:"redirect_configuration"`
	AutoscaleConfiguration    []AzurermApplicationGatewaySpec `json:"autoscale_configuration"`
	AuthenticationCertificate []AzurermApplicationGatewaySpec `json:"authentication_certificate"`
	DisabledSslProtocols      []string                        `json:"disabled_ssl_protocols"`
	UrlPathMap                []AzurermApplicationGatewaySpec `json:"url_path_map"`
	CustomErrorConfiguration  []AzurermApplicationGatewaySpec `json:"custom_error_configuration"`
	GatewayIpConfiguration    []AzurermApplicationGatewaySpec `json:"gateway_ip_configuration"`
	RequestRoutingRule        []AzurermApplicationGatewaySpec `json:"request_routing_rule"`
}

type AzurermApplicationGatewayStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermApplicationGatewayList is a list of AzurermApplicationGateways
type AzurermApplicationGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApplicationGateway CRD objects
	Items []AzurermApplicationGateway `json:"items,omitempty"`
}
