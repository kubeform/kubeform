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

type ApplicationGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationGatewaySpec   `json:"spec,omitempty"`
	Status            ApplicationGatewayStatus `json:"status,omitempty"`
}

type ApplicationGatewaySpecAuthenticationCertificate struct {
	Data string `json:"-" sensitive:"true" tf:"data"`
	ID   string `json:"ID" tf:"id"`
	Name string `json:"name" tf:"name"`
}

type ApplicationGatewaySpecAutoscaleConfiguration struct {
	// +optional
	MaxCapacity int `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`
	MinCapacity int `json:"minCapacity" tf:"min_capacity"`
}

type ApplicationGatewaySpecBackendAddressPool struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	// Deprecated
	FqdnList []string `json:"fqdnList,omitempty" tf:"fqdn_list,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	Fqdns []string `json:"fqdns,omitempty" tf:"fqdns,omitempty"`
	ID    string   `json:"ID" tf:"id"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	// Deprecated
	IpAddressList []string `json:"ipAddressList,omitempty" tf:"ip_address_list,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	IpAddresses []string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`
	Name        string   `json:"name" tf:"name"`
}

type ApplicationGatewaySpecBackendHTTPSettingsAuthenticationCertificate struct {
	ID   string `json:"ID" tf:"id"`
	Name string `json:"name" tf:"name"`
}

type ApplicationGatewaySpecBackendHTTPSettingsConnectionDraining struct {
	DrainTimeoutSec int  `json:"drainTimeoutSec" tf:"drain_timeout_sec"`
	Enabled         bool `json:"enabled" tf:"enabled"`
}

type ApplicationGatewaySpecBackendHTTPSettings struct {
	// +optional
	AffinityCookieName string `json:"affinityCookieName,omitempty" tf:"affinity_cookie_name,omitempty"`
	// +optional
	AuthenticationCertificate []ApplicationGatewaySpecBackendHTTPSettingsAuthenticationCertificate `json:"authenticationCertificate,omitempty" tf:"authentication_certificate,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ConnectionDraining  []ApplicationGatewaySpecBackendHTTPSettingsConnectionDraining `json:"connectionDraining,omitempty" tf:"connection_draining,omitempty"`
	CookieBasedAffinity string                                                        `json:"cookieBasedAffinity" tf:"cookie_based_affinity"`
	// +optional
	HostName string `json:"hostName,omitempty" tf:"host_name,omitempty"`
	ID       string `json:"ID" tf:"id"`
	Name     string `json:"name" tf:"name"`
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	// +optional
	PickHostNameFromBackendAddress bool   `json:"pickHostNameFromBackendAddress,omitempty" tf:"pick_host_name_from_backend_address,omitempty"`
	Port                           int    `json:"port" tf:"port"`
	ProbeID                        string `json:"probeID" tf:"probe_id"`
	// +optional
	ProbeName string `json:"probeName,omitempty" tf:"probe_name,omitempty"`
	Protocol  string `json:"protocol" tf:"protocol"`
	// +optional
	RequestTimeout int `json:"requestTimeout,omitempty" tf:"request_timeout,omitempty"`
}

type ApplicationGatewaySpecCustomErrorConfiguration struct {
	CustomErrorPageURL string `json:"customErrorPageURL" tf:"custom_error_page_url"`
	ID                 string `json:"ID" tf:"id"`
	StatusCode         string `json:"statusCode" tf:"status_code"`
}

type ApplicationGatewaySpecFrontendIPConfiguration struct {
	ID   string `json:"ID" tf:"id"`
	Name string `json:"name" tf:"name"`
	// +optional
	PrivateIPAddress string `json:"privateIPAddress,omitempty" tf:"private_ip_address,omitempty"`
	// +optional
	PrivateIPAddressAllocation string `json:"privateIPAddressAllocation,omitempty" tf:"private_ip_address_allocation,omitempty"`
	// +optional
	PublicIPAddressID string `json:"publicIPAddressID,omitempty" tf:"public_ip_address_id,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
}

type ApplicationGatewaySpecFrontendPort struct {
	ID   string `json:"ID" tf:"id"`
	Name string `json:"name" tf:"name"`
	Port int    `json:"port" tf:"port"`
}

type ApplicationGatewaySpecGatewayIPConfiguration struct {
	ID       string `json:"ID" tf:"id"`
	Name     string `json:"name" tf:"name"`
	SubnetID string `json:"subnetID" tf:"subnet_id"`
}

type ApplicationGatewaySpecHttpListenerCustomErrorConfiguration struct {
	CustomErrorPageURL string `json:"customErrorPageURL" tf:"custom_error_page_url"`
	ID                 string `json:"ID" tf:"id"`
	StatusCode         string `json:"statusCode" tf:"status_code"`
}

type ApplicationGatewaySpecHttpListener struct {
	// +optional
	CustomErrorConfiguration    []ApplicationGatewaySpecHttpListenerCustomErrorConfiguration `json:"customErrorConfiguration,omitempty" tf:"custom_error_configuration,omitempty"`
	FrontendIPConfigurationID   string                                                       `json:"frontendIPConfigurationID" tf:"frontend_ip_configuration_id"`
	FrontendIPConfigurationName string                                                       `json:"frontendIPConfigurationName" tf:"frontend_ip_configuration_name"`
	FrontendPortID              string                                                       `json:"frontendPortID" tf:"frontend_port_id"`
	FrontendPortName            string                                                       `json:"frontendPortName" tf:"frontend_port_name"`
	// +optional
	HostName string `json:"hostName,omitempty" tf:"host_name,omitempty"`
	ID       string `json:"ID" tf:"id"`
	Name     string `json:"name" tf:"name"`
	Protocol string `json:"protocol" tf:"protocol"`
	// +optional
	RequireSni       bool   `json:"requireSni,omitempty" tf:"require_sni,omitempty"`
	SslCertificateID string `json:"sslCertificateID" tf:"ssl_certificate_id"`
	// +optional
	SslCertificateName string `json:"sslCertificateName,omitempty" tf:"ssl_certificate_name,omitempty"`
}

type ApplicationGatewaySpecProbeMatch struct {
	// +optional
	Body string `json:"body,omitempty" tf:"body,omitempty"`
	// +optional
	StatusCode []string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type ApplicationGatewaySpecProbe struct {
	// +optional
	Host     string `json:"host,omitempty" tf:"host,omitempty"`
	ID       string `json:"ID" tf:"id"`
	Interval int    `json:"interval" tf:"interval"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Match []ApplicationGatewaySpecProbeMatch `json:"match,omitempty" tf:"match,omitempty"`
	// +optional
	MinimumServers int    `json:"minimumServers,omitempty" tf:"minimum_servers,omitempty"`
	Name           string `json:"name" tf:"name"`
	Path           string `json:"path" tf:"path"`
	// +optional
	PickHostNameFromBackendHTTPSettings bool   `json:"pickHostNameFromBackendHTTPSettings,omitempty" tf:"pick_host_name_from_backend_http_settings,omitempty"`
	Protocol                            string `json:"protocol" tf:"protocol"`
	Timeout                             int    `json:"timeout" tf:"timeout"`
	UnhealthyThreshold                  int    `json:"unhealthyThreshold" tf:"unhealthy_threshold"`
}

type ApplicationGatewaySpecRedirectConfiguration struct {
	ID string `json:"ID" tf:"id"`
	// +optional
	IncludePath bool `json:"includePath,omitempty" tf:"include_path,omitempty"`
	// +optional
	IncludeQueryString bool   `json:"includeQueryString,omitempty" tf:"include_query_string,omitempty"`
	Name               string `json:"name" tf:"name"`
	RedirectType       string `json:"redirectType" tf:"redirect_type"`
	TargetListenerID   string `json:"targetListenerID" tf:"target_listener_id"`
	// +optional
	TargetListenerName string `json:"targetListenerName,omitempty" tf:"target_listener_name,omitempty"`
	// +optional
	TargetURL string `json:"targetURL,omitempty" tf:"target_url,omitempty"`
}

type ApplicationGatewaySpecRequestRoutingRule struct {
	BackendAddressPoolID string `json:"backendAddressPoolID" tf:"backend_address_pool_id"`
	// +optional
	BackendAddressPoolName string `json:"backendAddressPoolName,omitempty" tf:"backend_address_pool_name,omitempty"`
	BackendHTTPSettingsID  string `json:"backendHTTPSettingsID" tf:"backend_http_settings_id"`
	// +optional
	BackendHTTPSettingsName string `json:"backendHTTPSettingsName,omitempty" tf:"backend_http_settings_name,omitempty"`
	HttpListenerID          string `json:"httpListenerID" tf:"http_listener_id"`
	HttpListenerName        string `json:"httpListenerName" tf:"http_listener_name"`
	ID                      string `json:"ID" tf:"id"`
	Name                    string `json:"name" tf:"name"`
	RedirectConfigurationID string `json:"redirectConfigurationID" tf:"redirect_configuration_id"`
	// +optional
	RedirectConfigurationName string `json:"redirectConfigurationName,omitempty" tf:"redirect_configuration_name,omitempty"`
	RewriteRuleSetID          string `json:"rewriteRuleSetID" tf:"rewrite_rule_set_id"`
	// +optional
	RewriteRuleSetName string `json:"rewriteRuleSetName,omitempty" tf:"rewrite_rule_set_name,omitempty"`
	RuleType           string `json:"ruleType" tf:"rule_type"`
	UrlPathMapID       string `json:"urlPathMapID" tf:"url_path_map_id"`
	// +optional
	UrlPathMapName string `json:"urlPathMapName,omitempty" tf:"url_path_map_name,omitempty"`
}

type ApplicationGatewaySpecRewriteRuleSetRewriteRuleCondition struct {
	// +optional
	IgnoreCase bool `json:"ignoreCase,omitempty" tf:"ignore_case,omitempty"`
	// +optional
	Negate   bool   `json:"negate,omitempty" tf:"negate,omitempty"`
	Pattern  string `json:"pattern" tf:"pattern"`
	Variable string `json:"variable" tf:"variable"`
}

type ApplicationGatewaySpecRewriteRuleSetRewriteRuleRequestHeaderConfiguration struct {
	HeaderName  string `json:"headerName" tf:"header_name"`
	HeaderValue string `json:"headerValue" tf:"header_value"`
}

type ApplicationGatewaySpecRewriteRuleSetRewriteRuleResponseHeaderConfiguration struct {
	HeaderName  string `json:"headerName" tf:"header_name"`
	HeaderValue string `json:"headerValue" tf:"header_value"`
}

type ApplicationGatewaySpecRewriteRuleSetRewriteRule struct {
	// +optional
	Condition []ApplicationGatewaySpecRewriteRuleSetRewriteRuleCondition `json:"condition,omitempty" tf:"condition,omitempty"`
	Name      string                                                     `json:"name" tf:"name"`
	// +optional
	RequestHeaderConfiguration []ApplicationGatewaySpecRewriteRuleSetRewriteRuleRequestHeaderConfiguration `json:"requestHeaderConfiguration,omitempty" tf:"request_header_configuration,omitempty"`
	// +optional
	ResponseHeaderConfiguration []ApplicationGatewaySpecRewriteRuleSetRewriteRuleResponseHeaderConfiguration `json:"responseHeaderConfiguration,omitempty" tf:"response_header_configuration,omitempty"`
	RuleSequence                int                                                                          `json:"ruleSequence" tf:"rule_sequence"`
}

type ApplicationGatewaySpecRewriteRuleSet struct {
	ID   string `json:"ID" tf:"id"`
	Name string `json:"name" tf:"name"`
	// +optional
	RewriteRule []ApplicationGatewaySpecRewriteRuleSetRewriteRule `json:"rewriteRule,omitempty" tf:"rewrite_rule,omitempty"`
}

type ApplicationGatewaySpecSku struct {
	// +optional
	Capacity int    `json:"capacity,omitempty" tf:"capacity,omitempty"`
	Name     string `json:"name" tf:"name"`
	Tier     string `json:"tier" tf:"tier"`
}

type ApplicationGatewaySpecSslCertificate struct {
	Data           string `json:"-" sensitive:"true" tf:"data"`
	ID             string `json:"ID" tf:"id"`
	Name           string `json:"name" tf:"name"`
	Password       string `json:"-" sensitive:"true" tf:"password"`
	PublicCertData string `json:"publicCertData" tf:"public_cert_data"`
}

type ApplicationGatewaySpecSslPolicy struct {
	// +optional
	CipherSuites []string `json:"cipherSuites,omitempty" tf:"cipher_suites,omitempty"`
	// +optional
	DisabledProtocols []string `json:"disabledProtocols,omitempty" tf:"disabled_protocols,omitempty"`
	// +optional
	MinProtocolVersion string `json:"minProtocolVersion,omitempty" tf:"min_protocol_version,omitempty"`
	// +optional
	PolicyName string `json:"policyName,omitempty" tf:"policy_name,omitempty"`
	// +optional
	PolicyType string `json:"policyType,omitempty" tf:"policy_type,omitempty"`
}

type ApplicationGatewaySpecUrlPathMapPathRule struct {
	BackendAddressPoolID string `json:"backendAddressPoolID" tf:"backend_address_pool_id"`
	// +optional
	BackendAddressPoolName string `json:"backendAddressPoolName,omitempty" tf:"backend_address_pool_name,omitempty"`
	BackendHTTPSettingsID  string `json:"backendHTTPSettingsID" tf:"backend_http_settings_id"`
	// +optional
	BackendHTTPSettingsName string   `json:"backendHTTPSettingsName,omitempty" tf:"backend_http_settings_name,omitempty"`
	ID                      string   `json:"ID" tf:"id"`
	Name                    string   `json:"name" tf:"name"`
	Paths                   []string `json:"paths" tf:"paths"`
	RedirectConfigurationID string   `json:"redirectConfigurationID" tf:"redirect_configuration_id"`
	// +optional
	RedirectConfigurationName string `json:"redirectConfigurationName,omitempty" tf:"redirect_configuration_name,omitempty"`
	RewriteRuleSetID          string `json:"rewriteRuleSetID" tf:"rewrite_rule_set_id"`
	// +optional
	RewriteRuleSetName string `json:"rewriteRuleSetName,omitempty" tf:"rewrite_rule_set_name,omitempty"`
}

type ApplicationGatewaySpecUrlPathMap struct {
	DefaultBackendAddressPoolID string `json:"defaultBackendAddressPoolID" tf:"default_backend_address_pool_id"`
	// +optional
	DefaultBackendAddressPoolName string `json:"defaultBackendAddressPoolName,omitempty" tf:"default_backend_address_pool_name,omitempty"`
	DefaultBackendHTTPSettingsID  string `json:"defaultBackendHTTPSettingsID" tf:"default_backend_http_settings_id"`
	// +optional
	DefaultBackendHTTPSettingsName string `json:"defaultBackendHTTPSettingsName,omitempty" tf:"default_backend_http_settings_name,omitempty"`
	DefaultRedirectConfigurationID string `json:"defaultRedirectConfigurationID" tf:"default_redirect_configuration_id"`
	// +optional
	DefaultRedirectConfigurationName string `json:"defaultRedirectConfigurationName,omitempty" tf:"default_redirect_configuration_name,omitempty"`
	DefaultRewriteRuleSetID          string `json:"defaultRewriteRuleSetID" tf:"default_rewrite_rule_set_id"`
	// +optional
	DefaultRewriteRuleSetName string                                     `json:"defaultRewriteRuleSetName,omitempty" tf:"default_rewrite_rule_set_name,omitempty"`
	ID                        string                                     `json:"ID" tf:"id"`
	Name                      string                                     `json:"name" tf:"name"`
	PathRule                  []ApplicationGatewaySpecUrlPathMapPathRule `json:"pathRule" tf:"path_rule"`
}

type ApplicationGatewaySpecWafConfigurationDisabledRuleGroup struct {
	RuleGroupName string `json:"ruleGroupName" tf:"rule_group_name"`
	// +optional
	Rules []int64 `json:"rules,omitempty" tf:"rules,omitempty"`
}

type ApplicationGatewaySpecWafConfigurationExclusion struct {
	MatchVariable string `json:"matchVariable" tf:"match_variable"`
	// +optional
	Selector string `json:"selector,omitempty" tf:"selector,omitempty"`
	// +optional
	SelectorMatchOperator string `json:"selectorMatchOperator,omitempty" tf:"selector_match_operator,omitempty"`
}

type ApplicationGatewaySpecWafConfiguration struct {
	// +optional
	DisabledRuleGroup []ApplicationGatewaySpecWafConfigurationDisabledRuleGroup `json:"disabledRuleGroup,omitempty" tf:"disabled_rule_group,omitempty"`
	Enabled           bool                                                      `json:"enabled" tf:"enabled"`
	// +optional
	Exclusion []ApplicationGatewaySpecWafConfigurationExclusion `json:"exclusion,omitempty" tf:"exclusion,omitempty"`
	// +optional
	FileUploadLimitMb int    `json:"fileUploadLimitMb,omitempty" tf:"file_upload_limit_mb,omitempty"`
	FirewallMode      string `json:"firewallMode" tf:"firewall_mode"`
	// +optional
	MaxRequestBodySizeKb int `json:"maxRequestBodySizeKb,omitempty" tf:"max_request_body_size_kb,omitempty"`
	// +optional
	RequestBodyCheck bool `json:"requestBodyCheck,omitempty" tf:"request_body_check,omitempty"`
	// +optional
	RuleSetType    string `json:"ruleSetType,omitempty" tf:"rule_set_type,omitempty"`
	RuleSetVersion string `json:"ruleSetVersion" tf:"rule_set_version"`
}

type ApplicationGatewaySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	KubeFormSecret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	AuthenticationCertificate []ApplicationGatewaySpecAuthenticationCertificate `json:"authenticationCertificate,omitempty" tf:"authentication_certificate,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AutoscaleConfiguration []ApplicationGatewaySpecAutoscaleConfiguration `json:"autoscaleConfiguration,omitempty" tf:"autoscale_configuration,omitempty"`
	BackendAddressPool     []ApplicationGatewaySpecBackendAddressPool     `json:"backendAddressPool" tf:"backend_address_pool"`
	// +kubebuilder:validation:MinItems=1
	BackendHTTPSettings []ApplicationGatewaySpecBackendHTTPSettings `json:"backendHTTPSettings" tf:"backend_http_settings"`
	// +optional
	CustomErrorConfiguration []ApplicationGatewaySpecCustomErrorConfiguration `json:"customErrorConfiguration,omitempty" tf:"custom_error_configuration,omitempty"`
	// +optional
	// Deprecated
	DisabledSslProtocols []string `json:"disabledSslProtocols,omitempty" tf:"disabled_ssl_protocols,omitempty"`
	// +optional
	EnableHttp2 bool `json:"enableHttp2,omitempty" tf:"enable_http2,omitempty"`
	// +kubebuilder:validation:MinItems=1
	FrontendIPConfiguration []ApplicationGatewaySpecFrontendIPConfiguration `json:"frontendIPConfiguration" tf:"frontend_ip_configuration"`
	FrontendPort            []ApplicationGatewaySpecFrontendPort            `json:"frontendPort" tf:"frontend_port"`
	// +kubebuilder:validation:MaxItems=2
	GatewayIPConfiguration []ApplicationGatewaySpecGatewayIPConfiguration `json:"gatewayIPConfiguration" tf:"gateway_ip_configuration"`
	HttpListener           []ApplicationGatewaySpecHttpListener           `json:"httpListener" tf:"http_listener"`
	Location               string                                         `json:"location" tf:"location"`
	Name                   string                                         `json:"name" tf:"name"`
	// +optional
	Probe []ApplicationGatewaySpecProbe `json:"probe,omitempty" tf:"probe,omitempty"`
	// +optional
	RedirectConfiguration []ApplicationGatewaySpecRedirectConfiguration `json:"redirectConfiguration,omitempty" tf:"redirect_configuration,omitempty"`
	// +kubebuilder:validation:MinItems=1
	RequestRoutingRule []ApplicationGatewaySpecRequestRoutingRule `json:"requestRoutingRule" tf:"request_routing_rule"`
	ResourceGroupName  string                                     `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	RewriteRuleSet []ApplicationGatewaySpecRewriteRuleSet `json:"rewriteRuleSet,omitempty" tf:"rewrite_rule_set,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Sku []ApplicationGatewaySpecSku `json:"sku" tf:"sku"`
	// +optional
	SslCertificate []ApplicationGatewaySpecSslCertificate `json:"sslCertificate,omitempty" tf:"ssl_certificate,omitempty"`
	// +optional
	SslPolicy []ApplicationGatewaySpecSslPolicy `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	UrlPathMap []ApplicationGatewaySpecUrlPathMap `json:"urlPathMap,omitempty" tf:"url_path_map,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	WafConfiguration []ApplicationGatewaySpecWafConfiguration `json:"wafConfiguration,omitempty" tf:"waf_configuration,omitempty"`
	// +optional
	Zones []string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type ApplicationGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApplicationGatewaySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApplicationGatewayList is a list of ApplicationGateways
type ApplicationGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApplicationGateway CRD objects
	Items []ApplicationGateway `json:"items,omitempty"`
}
