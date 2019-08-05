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

type NodebalancerConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodebalancerConfigSpec   `json:"spec,omitempty"`
	Status            NodebalancerConfigStatus `json:"status,omitempty"`
}

type NodebalancerConfigSpecNodeStatus struct {
	StatusDown int `json:"statusDown" tf:"status_down"`
	StatusUp   int `json:"statusUp" tf:"status_up"`
}

type NodebalancerConfigSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	KubeFormSecret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	Algorithm string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`
	// +optional
	Check string `json:"check,omitempty" tf:"check,omitempty"`
	// +optional
	CheckAttempts int `json:"checkAttempts,omitempty" tf:"check_attempts,omitempty"`
	// +optional
	CheckBody string `json:"checkBody,omitempty" tf:"check_body,omitempty"`
	// +optional
	CheckInterval int `json:"checkInterval,omitempty" tf:"check_interval,omitempty"`
	// +optional
	CheckPassive bool `json:"checkPassive,omitempty" tf:"check_passive,omitempty"`
	// +optional
	CheckPath string `json:"checkPath,omitempty" tf:"check_path,omitempty"`
	// +optional
	CheckTimeout int `json:"checkTimeout,omitempty" tf:"check_timeout,omitempty"`
	// +optional
	CipherSuite    string                                      `json:"cipherSuite,omitempty" tf:"cipher_suite,omitempty"`
	NodeStatus     map[string]NodebalancerConfigSpecNodeStatus `json:"nodeStatus" tf:"node_status"`
	NodebalancerID int                                         `json:"nodebalancerID" tf:"nodebalancer_id"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty" tf:"protocol,omitempty"`
	// +optional
	SslCert        string `json:"sslCert,omitempty" tf:"ssl_cert,omitempty"`
	SslCommonname  string `json:"sslCommonname" tf:"ssl_commonname"`
	SslFingerprint string `json:"sslFingerprint" tf:"ssl_fingerprint"`
	// +optional
	SslKey string `json:"-" sensitive:"true" tf:"ssl_key,omitempty"`
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
	State *apis.State `json:"state,omitempty"`
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
