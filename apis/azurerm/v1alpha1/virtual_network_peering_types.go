package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type VirtualNetworkPeering struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkPeeringSpec   `json:"spec,omitempty"`
	Status            VirtualNetworkPeeringStatus `json:"status,omitempty"`
}

type VirtualNetworkPeeringSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	AllowForwardedTraffic bool `json:"allowForwardedTraffic,omitempty" tf:"allow_forwarded_traffic,omitempty"`
	// +optional
	AllowGatewayTransit bool `json:"allowGatewayTransit,omitempty" tf:"allow_gateway_transit,omitempty"`
	// +optional
	AllowVirtualNetworkAccess bool   `json:"allowVirtualNetworkAccess,omitempty" tf:"allow_virtual_network_access,omitempty"`
	Name                      string `json:"name" tf:"name"`
	RemoteVirtualNetworkID    string `json:"remoteVirtualNetworkID" tf:"remote_virtual_network_id"`
	ResourceGroupName         string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	UseRemoteGateways  bool   `json:"useRemoteGateways,omitempty" tf:"use_remote_gateways,omitempty"`
	VirtualNetworkName string `json:"virtualNetworkName" tf:"virtual_network_name"`
}

type VirtualNetworkPeeringStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VirtualNetworkPeeringList is a list of VirtualNetworkPeerings
type VirtualNetworkPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualNetworkPeering CRD objects
	Items []VirtualNetworkPeering `json:"items,omitempty"`
}
