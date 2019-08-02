package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeRouteSpec   `json:"spec,omitempty"`
	Status            ComputeRouteStatus `json:"status,omitempty"`
}

type ComputeRouteSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	DestRange   string `json:"destRange" tf:"dest_range"`
	Name        string `json:"name" tf:"name"`
	Network     string `json:"network" tf:"network"`
	// +optional
	NextHopGateway string `json:"nextHopGateway,omitempty" tf:"next_hop_gateway,omitempty"`
	// +optional
	NextHopInstance string `json:"nextHopInstance,omitempty" tf:"next_hop_instance,omitempty"`
	// +optional
	NextHopInstanceZone string `json:"nextHopInstanceZone,omitempty" tf:"next_hop_instance_zone,omitempty"`
	// +optional
	NextHopIP string `json:"nextHopIP,omitempty" tf:"next_hop_ip,omitempty"`
	// +optional
	NextHopVPNTunnel string `json:"nextHopVPNTunnel,omitempty" tf:"next_hop_vpn_tunnel,omitempty"`
	// +optional
	Priority int `json:"priority,omitempty" tf:"priority,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ComputeRouteStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *ComputeRouteSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeRouteList is a list of ComputeRoutes
type ComputeRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeRoute CRD objects
	Items []ComputeRoute `json:"items,omitempty"`
}
