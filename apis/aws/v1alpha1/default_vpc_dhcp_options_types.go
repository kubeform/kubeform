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

type DefaultVpcDhcpOptions struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultVpcDhcpOptionsSpec   `json:"spec,omitempty"`
	Status            DefaultVpcDhcpOptionsStatus `json:"status,omitempty"`
}

type DefaultVpcDhcpOptionsSpec struct {
	// +optional
	NetbiosNameServers []string `json:"netbios_name_servers,omitempty"`
	// +optional
	NetbiosNodeType string `json:"netbios_node_type,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DefaultVpcDhcpOptionsStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DefaultVpcDhcpOptionsList is a list of DefaultVpcDhcpOptionss
type DefaultVpcDhcpOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DefaultVpcDhcpOptions CRD objects
	Items []DefaultVpcDhcpOptions `json:"items,omitempty"`
}
