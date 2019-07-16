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

type Loadbalancer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadbalancerSpec   `json:"spec,omitempty"`
	Status            LoadbalancerStatus `json:"status,omitempty"`
}

type LoadbalancerSpecForwardingRule struct {
	// +optional
	CertificateId  string `json:"certificate_id,omitempty"`
	EntryPort      int    `json:"entry_port"`
	EntryProtocol  string `json:"entry_protocol"`
	TargetPort     int    `json:"target_port"`
	TargetProtocol string `json:"target_protocol"`
	// +optional
	TlsPassthrough bool `json:"tls_passthrough,omitempty"`
}

type LoadbalancerSpec struct {
	// +optional
	Algorithm string `json:"algorithm,omitempty"`
	// +optional
	DropletTag string `json:"droplet_tag,omitempty"`
	// +optional
	EnableProxyProtocol bool `json:"enable_proxy_protocol,omitempty"`
	// +kubebuilder:validation:MinItems=1
	ForwardingRule []LoadbalancerSpec `json:"forwarding_rule"`
	Name           string             `json:"name"`
	// +optional
	RedirectHttpToHttps bool   `json:"redirect_http_to_https,omitempty"`
	Region              string `json:"region"`
}

type LoadbalancerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
