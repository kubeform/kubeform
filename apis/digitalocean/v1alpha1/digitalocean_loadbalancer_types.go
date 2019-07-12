package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DigitaloceanLoadbalancer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanLoadbalancerSpec   `json:"spec,omitempty"`
	Status            DigitaloceanLoadbalancerStatus `json:"status,omitempty"`
}

type ForwardingRuleSpec struct {
	EntryPort      int    `json:"entry_port"`
	TargetProtocol string `json:"target_protocol"`
	TargetPort     int    `json:"target_port"`
	CertificateId  string `json:"certificate_id"`
	TlsPassthrough bool   `json:"tls_passthrough"`
	EntryProtocol  string `json:"entry_protocol"`
}

type StickySessionsSpec struct {
	Type             string `json:"type"`
	CookieName       string `json:"cookie_name"`
	CookieTtlSeconds int    `json:"cookie_ttl_seconds"`
}

type HealthcheckSpec struct {
	Protocol               string `json:"protocol"`
	Port                   int    `json:"port"`
	Path                   string `json:"path"`
	CheckIntervalSeconds   int    `json:"check_interval_seconds"`
	ResponseTimeoutSeconds int    `json:"response_timeout_seconds"`
	UnhealthyThreshold     int    `json:"unhealthy_threshold"`
	HealthyThreshold       int    `json:"healthy_threshold"`
}

type DigitaloceanLoadbalancerSpec struct {
	DropletTag          string               `json:"droplet_tag"`
	RedirectHttpToHttps bool                 `json:"redirect_http_to_https"`
	Ip                  string               `json:"ip"`
	Status              string               `json:"status"`
	ForwardingRule      []ForwardingRuleSpec `json:"forwarding_rule"`
	StickySessions      []StickySessionsSpec `json:"sticky_sessions"`
	DropletIds          []int64              `json:"droplet_ids"`
	Algorithm           string               `json:"algorithm"`
	Healthcheck         []HealthcheckSpec    `json:"healthcheck"`
	EnableProxyProtocol bool                 `json:"enable_proxy_protocol"`
	Region              string               `json:"region"`
	Name                string               `json:"name"`
	Urn                 string               `json:"urn"`
}



type DigitaloceanLoadbalancerStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DigitaloceanLoadbalancerList is a list of DigitaloceanLoadbalancers
type DigitaloceanLoadbalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanLoadbalancer CRD objects
	Items []DigitaloceanLoadbalancer `json:"items,omitempty"`
}