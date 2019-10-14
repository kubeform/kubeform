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

type Eip struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EipSpec   `json:"spec,omitempty"`
	Status            EipStatus `json:"status,omitempty"`
}

type EipSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllocationID string `json:"allocationID,omitempty" tf:"allocation_id,omitempty"`
	// +optional
	AssociateWithPrivateIP string `json:"associateWithPrivateIP,omitempty" tf:"associate_with_private_ip,omitempty"`
	// +optional
	AssociationID string `json:"associationID,omitempty" tf:"association_id,omitempty"`
	// +optional
	Domain string `json:"domain,omitempty" tf:"domain,omitempty"`
	// +optional
	Instance string `json:"instance,omitempty" tf:"instance,omitempty"`
	// +optional
	NetworkInterface string `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`
	// +optional
	PrivateDNS string `json:"privateDNS,omitempty" tf:"private_dns,omitempty"`
	// +optional
	PrivateIP string `json:"privateIP,omitempty" tf:"private_ip,omitempty"`
	// +optional
	PublicDNS string `json:"publicDNS,omitempty" tf:"public_dns,omitempty"`
	// +optional
	PublicIP string `json:"publicIP,omitempty" tf:"public_ip,omitempty"`
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
	State *base.State `json:"state,omitempty"`
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
