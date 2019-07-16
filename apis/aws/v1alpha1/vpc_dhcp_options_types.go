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

type VpcDhcpOptions struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcDhcpOptionsSpec   `json:"spec,omitempty"`
	Status            VpcDhcpOptionsStatus `json:"status,omitempty"`
}

type VpcDhcpOptionsSpec struct {
	// +optional
	DomainName string `json:"domain_name,omitempty"`
	// +optional
	DomainNameServers []string `json:"domain_name_servers,omitempty"`
	// +optional
	NetbiosNameServers []string `json:"netbios_name_servers,omitempty"`
	// +optional
	NetbiosNodeType string `json:"netbios_node_type,omitempty"`
	// +optional
	NtpServers []string `json:"ntp_servers,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type VpcDhcpOptionsStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcDhcpOptionsList is a list of VpcDhcpOptionss
type VpcDhcpOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcDhcpOptions CRD objects
	Items []VpcDhcpOptions `json:"items,omitempty"`
}
