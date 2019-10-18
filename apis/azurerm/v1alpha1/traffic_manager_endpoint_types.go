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
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type TrafficManagerEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrafficManagerEndpointSpec   `json:"spec,omitempty"`
	Status            TrafficManagerEndpointStatus `json:"status,omitempty"`
}

type TrafficManagerEndpointSpecCustomHeader struct {
	Name  string `json:"name" tf:"name"`
	Value string `json:"value" tf:"value"`
}

type TrafficManagerEndpointSpecSubnet struct {
	First string `json:"first" tf:"first"`
	// +optional
	Last string `json:"last,omitempty" tf:"last,omitempty"`
	// +optional
	Scope int `json:"scope,omitempty" tf:"scope,omitempty"`
}

type TrafficManagerEndpointSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CustomHeader []TrafficManagerEndpointSpecCustomHeader `json:"customHeader,omitempty" tf:"custom_header,omitempty"`
	// +optional
	EndpointLocation string `json:"endpointLocation,omitempty" tf:"endpoint_location,omitempty"`
	// +optional
	EndpointMonitorStatus string `json:"endpointMonitorStatus,omitempty" tf:"endpoint_monitor_status,omitempty"`
	// +optional
	EndpointStatus string `json:"endpointStatus,omitempty" tf:"endpoint_status,omitempty"`
	// +optional
	GeoMappings []string `json:"geoMappings,omitempty" tf:"geo_mappings,omitempty"`
	// +optional
	MinChildEndpoints int    `json:"minChildEndpoints,omitempty" tf:"min_child_endpoints,omitempty"`
	Name              string `json:"name" tf:"name"`
	// +optional
	Priority          int    `json:"priority,omitempty" tf:"priority,omitempty"`
	ProfileName       string `json:"profileName" tf:"profile_name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Subnet []TrafficManagerEndpointSpecSubnet `json:"subnet,omitempty" tf:"subnet,omitempty"`
	// +optional
	Target string `json:"target,omitempty" tf:"target,omitempty"`
	// +optional
	TargetResourceID string `json:"targetResourceID,omitempty" tf:"target_resource_id,omitempty"`
	Type             string `json:"type" tf:"type"`
	// +optional
	Weight int `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TrafficManagerEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *TrafficManagerEndpointSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// TrafficManagerEndpointList is a list of TrafficManagerEndpoints
type TrafficManagerEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TrafficManagerEndpoint CRD objects
	Items []TrafficManagerEndpoint `json:"items,omitempty"`
}
