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

type NodebalancerConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodebalancerConfigSpec   `json:"spec,omitempty"`
	Status            NodebalancerConfigStatus `json:"status,omitempty"`
}

type NodebalancerConfigSpec struct {
	NodebalancerId int `json:"nodebalancer_id"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty"`
	// +optional
	SslCert string `json:"ssl_cert,omitempty"`
	// +optional
	SslKey string `json:"ssl_key,omitempty"`
}

type NodebalancerConfigStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
