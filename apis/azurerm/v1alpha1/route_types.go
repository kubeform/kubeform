package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Route struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteSpec   `json:"spec,omitempty"`
	Status            RouteStatus `json:"status,omitempty"`
}

type RouteSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AddressPrefix string `json:"addressPrefix" tf:"address_prefix"`
	Name          string `json:"name" tf:"name"`
	// +optional
	NextHopInIPAddress string `json:"nextHopInIPAddress,omitempty" tf:"next_hop_in_ip_address,omitempty"`
	NextHopType        string `json:"nextHopType" tf:"next_hop_type"`
	ResourceGroupName  string `json:"resourceGroupName" tf:"resource_group_name"`
	RouteTableName     string `json:"routeTableName" tf:"route_table_name"`
}

type RouteStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RouteSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RouteList is a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Route CRD objects
	Items []Route `json:"items,omitempty"`
}
