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

type DigitaloceanLoadbalancerSpecHealthcheck struct {
	Port                   int    `json:"port"`
	Path                   string `json:"path"`
	CheckIntervalSeconds   int    `json:"check_interval_seconds"`
	ResponseTimeoutSeconds int    `json:"response_timeout_seconds"`
	UnhealthyThreshold     int    `json:"unhealthy_threshold"`
	HealthyThreshold       int    `json:"healthy_threshold"`
	Protocol               string `json:"protocol"`
}

type DigitaloceanLoadbalancerSpecForwardingRule struct {
	EntryProtocol  string `json:"entry_protocol"`
	EntryPort      int    `json:"entry_port"`
	TargetProtocol string `json:"target_protocol"`
	TargetPort     int    `json:"target_port"`
	CertificateId  string `json:"certificate_id"`
	TlsPassthrough bool   `json:"tls_passthrough"`
}

type DigitaloceanLoadbalancerSpecStickySessions struct {
	Type             string `json:"type"`
	CookieName       string `json:"cookie_name"`
	CookieTtlSeconds int    `json:"cookie_ttl_seconds"`
}

type DigitaloceanLoadbalancerSpec struct {
	DropletIds          []int64                        `json:"droplet_ids"`
	RedirectHttpToHttps bool                           `json:"redirect_http_to_https"`
	EnableProxyProtocol bool                           `json:"enable_proxy_protocol"`
	Status              string                         `json:"status"`
	Algorithm           string                         `json:"algorithm"`
	Healthcheck         []DigitaloceanLoadbalancerSpec `json:"healthcheck"`
	Urn                 string                         `json:"urn"`
	ForwardingRule      []DigitaloceanLoadbalancerSpec `json:"forwarding_rule"`
	StickySessions      []DigitaloceanLoadbalancerSpec `json:"sticky_sessions"`
	DropletTag          string                         `json:"droplet_tag"`
	Ip                  string                         `json:"ip"`
	Region              string                         `json:"region"`
	Name                string                         `json:"name"`
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