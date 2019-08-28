package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type VpcDHCPOptions struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcDHCPOptionsSpec   `json:"spec,omitempty"`
	Status            VpcDHCPOptionsStatus `json:"status,omitempty"`
}

type VpcDHCPOptionsSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DomainName string `json:"domainName,omitempty" tf:"domain_name,omitempty"`
	// +optional
	DomainNameServers []string `json:"domainNameServers,omitempty" tf:"domain_name_servers,omitempty"`
	// +optional
	NetbiosNameServers []string `json:"netbiosNameServers,omitempty" tf:"netbios_name_servers,omitempty"`
	// +optional
	NetbiosNodeType string `json:"netbiosNodeType,omitempty" tf:"netbios_node_type,omitempty"`
	// +optional
	NtpServers []string `json:"ntpServers,omitempty" tf:"ntp_servers,omitempty"`
	// +optional
	OwnerID string `json:"ownerID,omitempty" tf:"owner_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}



type VpcDHCPOptionsStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *VpcDHCPOptionsSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcDHCPOptionsList is a list of VpcDHCPOptionss
type VpcDHCPOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcDHCPOptions CRD objects
	Items []VpcDHCPOptions `json:"items,omitempty"`
}