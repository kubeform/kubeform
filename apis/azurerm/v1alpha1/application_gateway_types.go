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

type ApplicationGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationGatewaySpec   `json:"spec,omitempty"`
	Status            ApplicationGatewayStatus `json:"status,omitempty"`
}

type ApplicationGatewaySpecAuthenticationCertificate struct {
	Data string `json:"data"`
	Name string `json:"name"`
}

type ApplicationGatewaySpecAutoscaleConfiguration struct {
	// +optional
	MaxCapacity int `json:"max_capacity,omitempty"`
	MinCapacity int `json:"min_capacity"`
}

type ApplicationGatewaySpecBackendAddressPool struct {
	Name string `json:"name"`
}

type ApplicationGatewaySpecBackendHttpSettingsAuthenticationCertificate struct {
	Name string `json:"name"`
}

type ApplicationGatewaySpecBackendHttpSettingsConnectionDraining struct {
	DrainTimeoutSec int  `json:"drain_timeout_sec"`
	Enabled         bool `json:"enabled"`
}

type ApplicationGatewaySpecBackendHttpSettings struct {
	// +optional
	AffinityCookieName string `json:"affinity_cookie_name,omitempty"`
	// +optional
	AuthenticationCertificate *[]ApplicationGatewaySpecBackendHttpSettings `json:"authentication_certificate,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ConnectionDraining  *[]ApplicationGatewaySpecBackendHttpSettings `json:"connection_draining,omitempty"`
	CookieBasedAffinity string                                       `json:"cookie_based_affinity"`
	// +optional
	HostName string `json:"host_name,omitempty"`
	Name     string `json:"name"`
	// +optional
	Path string `json:"path,omitempty"`
	// +optional
	PickHostNameFromBackendAddress bool `json:"pick_host_name_from_backend_address,omitempty"`
	Port                           int  `json:"port"`
	// +optional
	ProbeName string `json:"probe_name,omitempty"`
	Protocol  string `json:"protocol"`
	// +optional
	RequestTimeout int `json:"request_timeout,omitempty"`
}

type ApplicationGatewaySpecCustomErrorConfiguration struct {
	CustomErrorPageUrl string `json:"custom_error_page_url"`
	StatusCode         string `json:"status_code"`
}

type ApplicationGatewaySpecFrontendIpConfiguration struct {
	Name string `json:"name"`
}

type ApplicationGatewaySpecFrontendPort struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

type ApplicationGatewaySpecGatewayIpConfiguration struct {
	Name     string `json:"name"`
	SubnetId string `json:"subnet_id"`
}

type ApplicationGatewaySpecHttpListenerCustomErrorConfiguration struct {
	CustomErrorPageUrl string `json:"custom_error_page_url"`
	StatusCode         string `json:"status_code"`
}

type ApplicationGatewaySpecHttpListener struct {
	// +optional
	CustomErrorConfiguration    *[]ApplicationGatewaySpecHttpListener `json:"custom_error_configuration,omitempty"`
	FrontendIpConfigurationName string                                `json:"frontend_ip_configuration_name"`
	FrontendPortName            string                                `json:"frontend_port_name"`
	// +optional
	HostName string `json:"host_name,omitempty"`
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
	// +optional
	RequireSni bool `json:"require_sni,omitempty"`
	// +optional
	SslCertificateName string `json:"ssl_certificate_name,omitempty"`
}

type ApplicationGatewaySpecProbe struct {
	// +optional
	Host     string `json:"host,omitempty"`
	Interval int    `json:"interval"`
	// +optional
	MinimumServers int    `json:"minimum_servers,omitempty"`
	Name           string `json:"name"`
	Path           string `json:"path"`
	// +optional
	PickHostNameFromBackendHttpSettings bool   `json:"pick_host_name_from_backend_http_settings,omitempty"`
	Protocol                            string `json:"protocol"`
	Timeout                             int    `json:"timeout"`
	UnhealthyThreshold                  int    `json:"unhealthy_threshold"`
}

type ApplicationGatewaySpecRedirectConfiguration struct {
	// +optional
	IncludePath bool `json:"include_path,omitempty"`
	// +optional
	IncludeQueryString bool   `json:"include_query_string,omitempty"`
	Name               string `json:"name"`
	RedirectType       string `json:"redirect_type"`
	// +optional
	TargetListenerName string `json:"target_listener_name,omitempty"`
	// +optional
	TargetUrl string `json:"target_url,omitempty"`
}

type ApplicationGatewaySpecRequestRoutingRule struct {
	// +optional
	BackendAddressPoolName string `json:"backend_address_pool_name,omitempty"`
	// +optional
	BackendHttpSettingsName string `json:"backend_http_settings_name,omitempty"`
	HttpListenerName        string `json:"http_listener_name"`
	Name                    string `json:"name"`
	// +optional
	RedirectConfigurationName string `json:"redirect_configuration_name,omitempty"`
	// +optional
	RewriteRuleSetName string `json:"rewrite_rule_set_name,omitempty"`
	RuleType           string `json:"rule_type"`
	// +optional
	UrlPathMapName string `json:"url_path_map_name,omitempty"`
}

type ApplicationGatewaySpecRewriteRuleSetRewriteRuleCondition struct {
	// +optional
	IgnoreCase bool `json:"ignore_case,omitempty"`
	// +optional
	Negate   bool   `json:"negate,omitempty"`
	Pattern  string `json:"pattern"`
	Variable string `json:"variable"`
}

type ApplicationGatewaySpecRewriteRuleSetRewriteRuleRequestHeaderConfiguration struct {
	HeaderName  string `json:"header_name"`
	HeaderValue string `json:"header_value"`
}

type ApplicationGatewaySpecRewriteRuleSetRewriteRuleResponseHeaderConfiguration struct {
	HeaderName  string `json:"header_name"`
	HeaderValue string `json:"header_value"`
}

type ApplicationGatewaySpecRewriteRuleSetRewriteRule struct {
	// +optional
	Condition *[]ApplicationGatewaySpecRewriteRuleSetRewriteRule `json:"condition,omitempty"`
	Name      string                                             `json:"name"`
	// +optional
	RequestHeaderConfiguration *[]ApplicationGatewaySpecRewriteRuleSetRewriteRule `json:"request_header_configuration,omitempty"`
	// +optional
	ResponseHeaderConfiguration *[]ApplicationGatewaySpecRewriteRuleSetRewriteRule `json:"response_header_configuration,omitempty"`
	RuleSequence                int                                                `json:"rule_sequence"`
}

type ApplicationGatewaySpecRewriteRuleSet struct {
	Name string `json:"name"`
	// +optional
	RewriteRule *[]ApplicationGatewaySpecRewriteRuleSet `json:"rewrite_rule,omitempty"`
}

type ApplicationGatewaySpecSku struct {
	// +optional
	Capacity int    `json:"capacity,omitempty"`
	Name     string `json:"name"`
	Tier     string `json:"tier"`
}

type ApplicationGatewaySpecSslCertificate struct {
	Data     string `json:"data"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ApplicationGatewaySpecUrlPathMapPathRule struct {
	// +optional
	BackendAddressPoolName string `json:"backend_address_pool_name,omitempty"`
	// +optional
	BackendHttpSettingsName string   `json:"backend_http_settings_name,omitempty"`
	Name                    string   `json:"name"`
	Paths                   []string `json:"paths"`
	// +optional
	RedirectConfigurationName string `json:"redirect_configuration_name,omitempty"`
	// +optional
	RewriteRuleSetName string `json:"rewrite_rule_set_name,omitempty"`
}

type ApplicationGatewaySpecUrlPathMap struct {
	// +optional
	DefaultBackendAddressPoolName string `json:"default_backend_address_pool_name,omitempty"`
	// +optional
	DefaultBackendHttpSettingsName string `json:"default_backend_http_settings_name,omitempty"`
	// +optional
	DefaultRedirectConfigurationName string `json:"default_redirect_configuration_name,omitempty"`
	// +optional
	DefaultRewriteRuleSetName string                             `json:"default_rewrite_rule_set_name,omitempty"`
	Name                      string                             `json:"name"`
	PathRule                  []ApplicationGatewaySpecUrlPathMap `json:"path_rule"`
}

type ApplicationGatewaySpecWafConfigurationDisabledRuleGroup struct {
	RuleGroupName string `json:"rule_group_name"`
	// +optional
	Rules []int64 `json:"rules,omitempty"`
}

type ApplicationGatewaySpecWafConfigurationExclusion struct {
	MatchVariable string `json:"match_variable"`
	// +optional
	Selector string `json:"selector,omitempty"`
	// +optional
	SelectorMatchOperator string `json:"selector_match_operator,omitempty"`
}

type ApplicationGatewaySpecWafConfiguration struct {
	// +optional
	DisabledRuleGroup *[]ApplicationGatewaySpecWafConfiguration `json:"disabled_rule_group,omitempty"`
	Enabled           bool                                      `json:"enabled"`
	// +optional
	Exclusion *[]ApplicationGatewaySpecWafConfiguration `json:"exclusion,omitempty"`
	// +optional
	FileUploadLimitMb int    `json:"file_upload_limit_mb,omitempty"`
	FirewallMode      string `json:"firewall_mode"`
	// +optional
	MaxRequestBodySizeKb int `json:"max_request_body_size_kb,omitempty"`
	// +optional
	RequestBodyCheck bool `json:"request_body_check,omitempty"`
	// +optional
	RuleSetType    string `json:"rule_set_type,omitempty"`
	RuleSetVersion string `json:"rule_set_version"`
}

type ApplicationGatewaySpec struct {
	// +optional
	AuthenticationCertificate *[]ApplicationGatewaySpec `json:"authentication_certificate,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AutoscaleConfiguration *[]ApplicationGatewaySpec `json:"autoscale_configuration,omitempty"`
	BackendAddressPool     []ApplicationGatewaySpec  `json:"backend_address_pool"`
	// +kubebuilder:validation:MinItems=1
	BackendHttpSettings []ApplicationGatewaySpec `json:"backend_http_settings"`
	// +optional
	CustomErrorConfiguration *[]ApplicationGatewaySpec `json:"custom_error_configuration,omitempty"`
	// +optional
	EnableHttp2 bool `json:"enable_http2,omitempty"`
	// +kubebuilder:validation:MinItems=1
	FrontendIpConfiguration []ApplicationGatewaySpec `json:"frontend_ip_configuration"`
	FrontendPort            []ApplicationGatewaySpec `json:"frontend_port"`
	// +kubebuilder:validation:MaxItems=2
	GatewayIpConfiguration []ApplicationGatewaySpec `json:"gateway_ip_configuration"`
	HttpListener           []ApplicationGatewaySpec `json:"http_listener"`
	Location               string                   `json:"location"`
	Name                   string                   `json:"name"`
	// +optional
	Probe *[]ApplicationGatewaySpec `json:"probe,omitempty"`
	// +optional
	RedirectConfiguration *[]ApplicationGatewaySpec `json:"redirect_configuration,omitempty"`
	// +kubebuilder:validation:MinItems=1
	RequestRoutingRule []ApplicationGatewaySpec `json:"request_routing_rule"`
	ResourceGroupName  string                   `json:"resource_group_name"`
	// +optional
	RewriteRuleSet *[]ApplicationGatewaySpec `json:"rewrite_rule_set,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Sku []ApplicationGatewaySpec `json:"sku"`
	// +optional
	SslCertificate *[]ApplicationGatewaySpec `json:"ssl_certificate,omitempty"`
	// +optional
	UrlPathMap *[]ApplicationGatewaySpec `json:"url_path_map,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	WafConfiguration *[]ApplicationGatewaySpec `json:"waf_configuration,omitempty"`
	// +optional
	Zones []string `json:"zones,omitempty"`
}

type ApplicationGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
