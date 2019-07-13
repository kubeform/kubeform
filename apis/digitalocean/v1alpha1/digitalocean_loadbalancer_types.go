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

type DigitaloceanLoadbalancer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanLoadbalancerSpec   `json:"spec,omitempty"`
	Status            DigitaloceanLoadbalancerStatus `json:"status,omitempty"`
}

type DigitaloceanLoadbalancerSpecForwardingRule struct {
	EntryPort      int    `json:"entry_port"`
	TargetProtocol string `json:"target_protocol"`
	TargetPort     int    `json:"target_port"`
	CertificateId  string `json:"certificate_id"`
	TlsPassthrough bool   `json:"tls_passthrough"`
	EntryProtocol  string `json:"entry_protocol"`
}

type DigitaloceanLoadbalancerSpecHealthcheck struct {
	HealthyThreshold       int    `json:"healthy_threshold"`
	Protocol               string `json:"protocol"`
	Port                   int    `json:"port"`
	Path                   string `json:"path"`
	CheckIntervalSeconds   int    `json:"check_interval_seconds"`
	ResponseTimeoutSeconds int    `json:"response_timeout_seconds"`
	UnhealthyThreshold     int    `json:"unhealthy_threshold"`
}

type DigitaloceanLoadbalancerSpecStickySessions struct {
	Type             string `json:"type"`
	CookieName       string `json:"cookie_name"`
	CookieTtlSeconds int    `json:"cookie_ttl_seconds"`
}

type DigitaloceanLoadbalancerSpec struct {
	Urn                 string                         `json:"urn"`
	Ip                  string                         `json:"ip"`
	Algorithm           string                         `json:"algorithm"`
	ForwardingRule      []DigitaloceanLoadbalancerSpec `json:"forwarding_rule"`
	Healthcheck         []DigitaloceanLoadbalancerSpec `json:"healthcheck"`
	StickySessions      []DigitaloceanLoadbalancerSpec `json:"sticky_sessions"`
	DropletIds          []int64                        `json:"droplet_ids"`
	DropletTag          string                         `json:"droplet_tag"`
	Region              string                         `json:"region"`
	Name                string                         `json:"name"`
	Status              string                         `json:"status"`
	RedirectHttpToHttps bool                           `json:"redirect_http_to_https"`
	EnableProxyProtocol bool                           `json:"enable_proxy_protocol"`
}



type DigitaloceanLoadbalancerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DigitaloceanLoadbalancerList is a list of DigitaloceanLoadbalancers
type DigitaloceanLoadbalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanLoadbalancer CRD objects
	Items []DigitaloceanLoadbalancer `json:"items,omitempty"`
}