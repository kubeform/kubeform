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

type Eip struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EipSpec   `json:"spec,omitempty"`
	Status            EipStatus `json:"status,omitempty"`
}

type EipSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	AllocationID string `json:"allocationID" tf:"allocation_id"`
	// +optional
	AssociateWithPrivateIP string `json:"associateWithPrivateIP,omitempty" tf:"associate_with_private_ip,omitempty"`
	AssociationID          string `json:"associationID" tf:"association_id"`
	Domain                 string `json:"domain" tf:"domain"`
	// +optional
	Instance string `json:"instance,omitempty" tf:"instance,omitempty"`
	// +optional
	NetworkInterface string `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`
	PrivateDNS       string `json:"privateDNS" tf:"private_dns"`
	PrivateIP        string `json:"privateIP" tf:"private_ip"`
	PublicDNS        string `json:"publicDNS" tf:"public_dns"`
	PublicIP         string `json:"publicIP" tf:"public_ip"`
	// +optional
	PublicIpv4Pool string `json:"publicIpv4Pool,omitempty" tf:"public_ipv4_pool,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Vpc bool `json:"vpc,omitempty" tf:"vpc,omitempty"`
}

type EipStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EipSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EipList is a list of Eips
type EipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Eip CRD objects
	Items []Eip `json:"items,omitempty"`
}
