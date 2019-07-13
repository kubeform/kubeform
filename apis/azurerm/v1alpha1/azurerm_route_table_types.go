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

type AzurermRouteTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermRouteTableSpec   `json:"spec,omitempty"`
	Status            AzurermRouteTableStatus `json:"status,omitempty"`
}

type AzurermRouteTableSpecRoute struct {
	Name               string `json:"name"`
	AddressPrefix      string `json:"address_prefix"`
	NextHopType        string `json:"next_hop_type"`
	NextHopInIpAddress string `json:"next_hop_in_ip_address"`
}

type AzurermRouteTableSpec struct {
	Location                   string                  `json:"location"`
	ResourceGroupName          string                  `json:"resource_group_name"`
	Route                      []AzurermRouteTableSpec `json:"route"`
	DisableBgpRoutePropagation bool                    `json:"disable_bgp_route_propagation"`
	Subnets                    []string                `json:"subnets"`
	Tags                       map[string]string       `json:"tags"`
	Name                       string                  `json:"name"`
}



type AzurermRouteTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermRouteTableList is a list of AzurermRouteTables
type AzurermRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermRouteTable CRD objects
	Items []AzurermRouteTable `json:"items,omitempty"`
}