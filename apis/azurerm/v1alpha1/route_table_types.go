package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type RouteTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTableSpec   `json:"spec,omitempty"`
	Status            RouteTableStatus `json:"status,omitempty"`
}

type RouteTableSpecRoute struct {
	AddressPrefix string `json:"addressPrefix" tf:"address_prefix"`
	Name          string `json:"name" tf:"name"`
	// +optional
	NextHopInIPAddress string `json:"nextHopInIPAddress,omitempty" tf:"next_hop_in_ip_address,omitempty"`
	NextHopType        string `json:"nextHopType" tf:"next_hop_type"`
}

type RouteTableSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	DisableBGPRoutePropagation bool   `json:"disableBGPRoutePropagation,omitempty" tf:"disable_bgp_route_propagation,omitempty"`
	Location                   string `json:"location" tf:"location"`
	Name                       string `json:"name" tf:"name"`
	ResourceGroupName          string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Route []RouteTableSpecRoute `json:"route,omitempty" tf:"route,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RouteTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *RouteTableSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RouteTableList is a list of RouteTables
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RouteTable CRD objects
	Items []RouteTable `json:"items,omitempty"`
}
