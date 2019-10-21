package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Loadbalancer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadbalancerSpec   `json:"spec,omitempty"`
	Status            LoadbalancerStatus `json:"status,omitempty"`
}

type LoadbalancerSpecForwardingRule struct {
	// +optional
	CertificateID  string `json:"certificateID,omitempty" tf:"certificate_id,omitempty"`
	EntryPort      int    `json:"entryPort" tf:"entry_port"`
	EntryProtocol  string `json:"entryProtocol" tf:"entry_protocol"`
	TargetPort     int    `json:"targetPort" tf:"target_port"`
	TargetProtocol string `json:"targetProtocol" tf:"target_protocol"`
	// +optional
	TlsPassthrough bool `json:"tlsPassthrough,omitempty" tf:"tls_passthrough,omitempty"`
}

type LoadbalancerSpecHealthcheck struct {
	// +optional
	CheckIntervalSeconds int `json:"checkIntervalSeconds,omitempty" tf:"check_interval_seconds,omitempty"`
	// +optional
	HealthyThreshold int `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`
	// +optional
	Path     string `json:"path,omitempty" tf:"path,omitempty"`
	Port     int    `json:"port" tf:"port"`
	Protocol string `json:"protocol" tf:"protocol"`
	// +optional
	ResponseTimeoutSeconds int `json:"responseTimeoutSeconds,omitempty" tf:"response_timeout_seconds,omitempty"`
	// +optional
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type LoadbalancerSpecStickySessions struct {
	// +optional
	CookieName string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`
	// +optional
	CookieTtlSeconds int `json:"cookieTtlSeconds,omitempty" tf:"cookie_ttl_seconds,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type LoadbalancerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Algorithm string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`
	// +optional
	DropletIDS []int64 `json:"dropletIDS,omitempty" tf:"droplet_ids,omitempty"`
	// +optional
	DropletTag string `json:"dropletTag,omitempty" tf:"droplet_tag,omitempty"`
	// +optional
	EnableProxyProtocol bool `json:"enableProxyProtocol,omitempty" tf:"enable_proxy_protocol,omitempty"`
	// +kubebuilder:validation:MinItems=1
	ForwardingRule []LoadbalancerSpecForwardingRule `json:"forwardingRule" tf:"forwarding_rule"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Healthcheck []LoadbalancerSpecHealthcheck `json:"healthcheck,omitempty" tf:"healthcheck,omitempty"`
	// +optional
	Ip   string `json:"ip,omitempty" tf:"ip,omitempty"`
	Name string `json:"name" tf:"name"`
	// +optional
	RedirectHTTPToHTTPS bool   `json:"redirectHTTPToHTTPS,omitempty" tf:"redirect_http_to_https,omitempty"`
	Region              string `json:"region" tf:"region"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StickySessions []LoadbalancerSpecStickySessions `json:"stickySessions,omitempty" tf:"sticky_sessions,omitempty"`
	// the uniform resource name for the load balancer
	// +optional
	Urn string `json:"urn,omitempty" tf:"urn,omitempty"`
}

type LoadbalancerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LoadbalancerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LoadbalancerList is a list of Loadbalancers
type LoadbalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Loadbalancer CRD objects
	Items []Loadbalancer `json:"items,omitempty"`
}
