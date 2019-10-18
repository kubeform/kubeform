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

type NodebalancerConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodebalancerConfigSpec   `json:"spec,omitempty"`
	Status            NodebalancerConfigStatus `json:"status,omitempty"`
}

type NodebalancerConfigSpecNodeStatus struct {
	// The number of backends considered to be 'DOWN' and unhealthy. These are not in rotation, and not serving requests.
	// +optional
	StatusDown int `json:"statusDown,omitempty" tf:"status_down,omitempty"`
	// The number of backends considered to be 'UP' and healthy, and that are serving requests.
	// +optional
	StatusUp int `json:"statusUp,omitempty" tf:"status_up,omitempty"`
}

type NodebalancerConfigSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// What algorithm this NodeBalancer should use for routing traffic to backends: roundrobin, leastconn, source
	// +optional
	Algorithm string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`
	// The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. If none no check is performed. connection requires only a connection to the backend to succeed. http and http_body rely on the backend serving HTTP, and that the response returned matches what is expected.
	// +optional
	Check string `json:"check,omitempty" tf:"check,omitempty"`
	// How many times to attempt a check before considering a backend to be down. (1-30)
	// +optional
	CheckAttempts int `json:"checkAttempts,omitempty" tf:"check_attempts,omitempty"`
	// This value must be present in the response body of the check in order for it to pass. If this value is not present in the response body of a check request, the backend is considered to be down
	// +optional
	CheckBody string `json:"checkBody,omitempty" tf:"check_body,omitempty"`
	// How often, in seconds, to check that backends are up and serving requests.
	// +optional
	CheckInterval int `json:"checkInterval,omitempty" tf:"check_interval,omitempty"`
	// If true, any response from this backend with a 5xx status code will be enough for it to be considered unhealthy and taken out of rotation.
	// +optional
	CheckPassive bool `json:"checkPassive,omitempty" tf:"check_passive,omitempty"`
	// The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.
	// +optional
	CheckPath string `json:"checkPath,omitempty" tf:"check_path,omitempty"`
	// How long, in seconds, to wait for a check attempt before considering it failed. (1-30)
	// +optional
	CheckTimeout int `json:"checkTimeout,omitempty" tf:"check_timeout,omitempty"`
	// What ciphers to use for SSL connections served by this NodeBalancer. `legacy` is considered insecure and should only be used if necessary.
	// +optional
	CipherSuite string `json:"cipherSuite,omitempty" tf:"cipher_suite,omitempty"`
	// +optional
	NodeStatus map[string]NodebalancerConfigSpecNodeStatus `json:"nodeStatus,omitempty" tf:"node_status,omitempty"`
	// The ID of the NodeBalancer to access.
	NodebalancerID int `json:"nodebalancerID" tf:"nodebalancer_id"`
	// The TCP port this Config is for. These values must be unique across configs on a single NodeBalancer (you can't have two configs for port 80, for example). While some ports imply some protocols, no enforcement is done and you may configure your NodeBalancer however is useful to you. For example, while port 443 is generally used for HTTPS, you do not need SSL configured to have a NodeBalancer listening on port 443.
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// The protocol this port is configured to serve. If this is set to https you must include an ssl_cert and an ssl_key.
	// +optional
	Protocol string `json:"protocol,omitempty" tf:"protocol,omitempty"`
	// The certificate this port is serving. This is not returned. If set, this field will come back as `<REDACTED>`. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.
	// +optional
	SslCert string `json:"sslCert,omitempty" tf:"ssl_cert,omitempty"`
	// The common name for the SSL certification this port is serving if this port is not configured to use SSL.
	// +optional
	SslCommonname string `json:"sslCommonname,omitempty" tf:"ssl_commonname,omitempty"`
	// The fingerprint for the SSL certification this port is serving if this port is not configured to use SSL.
	// +optional
	SslFingerprint string `json:"sslFingerprint,omitempty" tf:"ssl_fingerprint,omitempty"`
	// The private key corresponding to this port's certificate. This is not returned. If set, this field will come back as `<REDACTED>`. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.
	// +optional
	SslKey string `json:"-" sensitive:"true" tf:"ssl_key,omitempty"`
	// Controls how session stickiness is handled on this port: 'none', 'table', 'http_cookie'
	// +optional
	Stickiness string `json:"stickiness,omitempty" tf:"stickiness,omitempty"`
}

type NodebalancerConfigStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NodebalancerConfigSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NodebalancerConfigList is a list of NodebalancerConfigs
type NodebalancerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NodebalancerConfig CRD objects
	Items []NodebalancerConfig `json:"items,omitempty"`
}
