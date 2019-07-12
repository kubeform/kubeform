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

type AzurermApplicationGatewaySpecBackendHttpSettingsAuthenticationCertificate struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type AzurermApplicationGatewaySpecBackendHttpSettingsConnectionDraining struct {
	DrainTimeoutSec int  `json:"drain_timeout_sec"`
	Enabled         bool `json:"enabled"`
}

type AzurermApplicationGatewaySpecBackendHttpSettings struct {
	Path                           string                                             `json:"path"`
	PickHostNameFromBackendAddress bool                                               `json:"pick_host_name_from_backend_address"`
	RequestTimeout                 int                                                `json:"request_timeout"`
	Name                           string                                             `json:"name"`
	Port                           int                                                `json:"port"`
	AuthenticationCertificate      []AzurermApplicationGatewaySpecBackendHttpSettings `json:"authentication_certificate"`
	Id                             string                                             `json:"id"`
	ProbeId                        string                                             `json:"probe_id"`
	Protocol                       string                                             `json:"protocol"`
	AffinityCookieName             string                                             `json:"affinity_cookie_name"`
	HostName                       string                                             `json:"host_name"`
	CookieBasedAffinity            string                                             `json:"cookie_based_affinity"`
	ConnectionDraining             []AzurermApplicationGatewaySpecBackendHttpSettings `json:"connection_draining"`
	ProbeName                      string                                             `json:"probe_name"`
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
	Id                          string                                      `json:"id"`
	Name                        string                                      `json:"name"`
	FrontendPortName            string                                      `json:"frontend_port_name"`
	Protocol                    string                                      `json:"protocol"`
	HostName                    string                                      `json:"host_name"`
	FrontendPortId              string                                      `json:"frontend_port_id"`
	SslCertificateId            string                                      `json:"ssl_certificate_id"`
	CustomErrorConfiguration    []AzurermApplicationGatewaySpecHttpListener `json:"custom_error_configuration"`
	FrontendIpConfigurationName string                                      `json:"frontend_ip_configuration_name"`
	SslCertificateName          string                                      `json:"ssl_certificate_name"`
	RequireSni                  bool                                        `json:"require_sni"`
	FrontendIpConfigurationId   string                                      `json:"frontend_ip_configuration_id"`
}

type AzurermApplicationGatewaySpecAutoscaleConfiguration struct {
	MinCapacity int `json:"min_capacity"`
	MaxCapacity int `json:"max_capacity"`
}

type AzurermApplicationGatewaySpecUrlPathMapPathRule struct {
	BackendAddressPoolName    string   `json:"backend_address_pool_name"`
	BackendAddressPoolId      string   `json:"backend_address_pool_id"`
	BackendHttpSettingsId     string   `json:"backend_http_settings_id"`
	RedirectConfigurationId   string   `json:"redirect_configuration_id"`
	Name                      string   `json:"name"`
	Paths                     []string `json:"paths"`
	RewriteRuleSetName        string   `json:"rewrite_rule_set_name"`
	RewriteRuleSetId          string   `json:"rewrite_rule_set_id"`
	Id                        string   `json:"id"`
	BackendHttpSettingsName   string   `json:"backend_http_settings_name"`
	RedirectConfigurationName string   `json:"redirect_configuration_name"`
}

type AzurermApplicationGatewaySpecUrlPathMap struct {
	DefaultRedirectConfigurationName string                                    `json:"default_redirect_configuration_name"`
	DefaultBackendAddressPoolId      string                                    `json:"default_backend_address_pool_id"`
	DefaultBackendHttpSettingsId     string                                    `json:"default_backend_http_settings_id"`
	DefaultRedirectConfigurationId   string                                    `json:"default_redirect_configuration_id"`
	DefaultRewriteRuleSetId          string                                    `json:"default_rewrite_rule_set_id"`
	Name                             string                                    `json:"name"`
	DefaultBackendAddressPoolName    string                                    `json:"default_backend_address_pool_name"`
	DefaultBackendHttpSettingsName   string                                    `json:"default_backend_http_settings_name"`
	DefaultRewriteRuleSetName        string                                    `json:"default_rewrite_rule_set_name"`
	PathRule                         []AzurermApplicationGatewaySpecUrlPathMap `json:"path_rule"`
	Id                               string                                    `json:"id"`
}

type AzurermApplicationGatewaySpecRequestRoutingRule struct {
	Name                      string `json:"name"`
	RuleType                  string `json:"rule_type"`
	BackendHttpSettingsId     string `json:"backend_http_settings_id"`
	HttpListenerId            string `json:"http_listener_id"`
	UrlPathMapId              string `json:"url_path_map_id"`
	RedirectConfigurationId   string `json:"redirect_configuration_id"`
	RedirectConfigurationName string `json:"redirect_configuration_name"`
	RewriteRuleSetName        string `json:"rewrite_rule_set_name"`
	BackendAddressPoolId      string `json:"backend_address_pool_id"`
	Id                        string `json:"id"`
	HttpListenerName          string `json:"http_listener_name"`
	BackendAddressPoolName    string `json:"backend_address_pool_name"`
	BackendHttpSettingsName   string `json:"backend_http_settings_name"`
	UrlPathMapName            string `json:"url_path_map_name"`
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

type AzurermApplicationGatewaySpecSslPolicy struct {
	DisabledProtocols  []string `json:"disabled_protocols"`
	PolicyType         string   `json:"policy_type"`
	PolicyName         string   `json:"policy_name"`
	CipherSuites       []string `json:"cipher_suites"`
	MinProtocolVersion string   `json:"min_protocol_version"`
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
	PublicCertData string `json:"public_cert_data"`
	Name           string `json:"name"`
	Data           string `json:"data"`
	Password       string `json:"password"`
	Id             string `json:"id"`
}

type AzurermApplicationGatewaySpecBackendAddressPool struct {
	Fqdns         []string `json:"fqdns"`
	IpAddresses   []string `json:"ip_addresses"`
	FqdnList      []string `json:"fqdn_list"`
	IpAddressList []string `json:"ip_address_list"`
	Id            string   `json:"id"`
	Name          string   `json:"name"`
}

type AzurermApplicationGatewaySpecAuthenticationCertificate struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Data string `json:"data"`
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
	RuleSetType          string                                          `json:"rule_set_type"`
	FileUploadLimitMb    int                                             `json:"file_upload_limit_mb"`
	RequestBodyCheck     bool                                            `json:"request_body_check"`
	DisabledRuleGroup    []AzurermApplicationGatewaySpecWafConfiguration `json:"disabled_rule_group"`
	Exclusion            []AzurermApplicationGatewaySpecWafConfiguration `json:"exclusion"`
	Enabled              bool                                            `json:"enabled"`
	FirewallMode         string                                          `json:"firewall_mode"`
	RuleSetVersion       string                                          `json:"rule_set_version"`
	MaxRequestBodySizeKb int                                             `json:"max_request_body_size_kb"`
}

type AzurermApplicationGatewaySpecCustomErrorConfiguration struct {
	StatusCode         string `json:"status_code"`
	CustomErrorPageUrl string `json:"custom_error_page_url"`
	Id                 string `json:"id"`
}

type AzurermApplicationGatewaySpecFrontendIpConfiguration struct {
	PrivateIpAddress           string `json:"private_ip_address"`
	PublicIpAddressId          string `json:"public_ip_address_id"`
	PrivateIpAddressAllocation string `json:"private_ip_address_allocation"`
	Id                         string `json:"id"`
	Name                       string `json:"name"`
	SubnetId                   string `json:"subnet_id"`
}

type AzurermApplicationGatewaySpecGatewayIpConfiguration struct {
	Name     string `json:"name"`
	SubnetId string `json:"subnet_id"`
	Id       string `json:"id"`
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
	Name                                string                               `json:"name"`
	Host                                string                               `json:"host"`
	Interval                            int                                  `json:"interval"`
	UnhealthyThreshold                  int                                  `json:"unhealthy_threshold"`
	PickHostNameFromBackendHttpSettings bool                                 `json:"pick_host_name_from_backend_http_settings"`
	Id                                  string                               `json:"id"`
	Protocol                            string                               `json:"protocol"`
	Path                                string                               `json:"path"`
	Timeout                             int                                  `json:"timeout"`
	MinimumServers                      int                                  `json:"minimum_servers"`
	Match                               []AzurermApplicationGatewaySpecProbe `json:"match"`
}

type AzurermApplicationGatewaySpec struct {
	BackendHttpSettings       []AzurermApplicationGatewaySpec `json:"backend_http_settings"`
	FrontendPort              []AzurermApplicationGatewaySpec `json:"frontend_port"`
	HttpListener              []AzurermApplicationGatewaySpec `json:"http_listener"`
	AutoscaleConfiguration    []AzurermApplicationGatewaySpec `json:"autoscale_configuration"`
	UrlPathMap                []AzurermApplicationGatewaySpec `json:"url_path_map"`
	Zones                     []string                        `json:"zones"`
	RequestRoutingRule        []AzurermApplicationGatewaySpec `json:"request_routing_rule"`
	RedirectConfiguration     []AzurermApplicationGatewaySpec `json:"redirect_configuration"`
	SslPolicy                 []AzurermApplicationGatewaySpec `json:"ssl_policy"`
	EnableHttp2               bool                            `json:"enable_http2"`
	RewriteRuleSet            []AzurermApplicationGatewaySpec `json:"rewrite_rule_set"`
	SslCertificate            []AzurermApplicationGatewaySpec `json:"ssl_certificate"`
	ResourceGroupName         string                          `json:"resource_group_name"`
	BackendAddressPool        []AzurermApplicationGatewaySpec `json:"backend_address_pool"`
	AuthenticationCertificate []AzurermApplicationGatewaySpec `json:"authentication_certificate"`
	DisabledSslProtocols      []string                        `json:"disabled_ssl_protocols"`
	WafConfiguration          []AzurermApplicationGatewaySpec `json:"waf_configuration"`
	CustomErrorConfiguration  []AzurermApplicationGatewaySpec `json:"custom_error_configuration"`
	Name                      string                          `json:"name"`
	Location                  string                          `json:"location"`
	FrontendIpConfiguration   []AzurermApplicationGatewaySpec `json:"frontend_ip_configuration"`
	GatewayIpConfiguration    []AzurermApplicationGatewaySpec `json:"gateway_ip_configuration"`
	Sku                       []AzurermApplicationGatewaySpec `json:"sku"`
	Probe                     []AzurermApplicationGatewaySpec `json:"probe"`
	Tags                      map[string]string               `json:"tags"`
}

type AzurermApplicationGatewayStatus struct {
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
