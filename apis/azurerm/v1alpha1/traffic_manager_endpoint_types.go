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

type TrafficManagerEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrafficManagerEndpointSpec   `json:"spec,omitempty"`
	Status            TrafficManagerEndpointStatus `json:"status,omitempty"`
}

type TrafficManagerEndpointSpecCustomHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TrafficManagerEndpointSpecSubnet struct {
	First string `json:"first"`
	// +optional
	Last string `json:"last,omitempty"`
	// +optional
	Scope int `json:"scope,omitempty"`
}

type TrafficManagerEndpointSpec struct {
	// +optional
	CustomHeader *[]TrafficManagerEndpointSpec `json:"custom_header,omitempty"`
	// +optional
	GeoMappings []string `json:"geo_mappings,omitempty"`
	// +optional
	MinChildEndpoints int    `json:"min_child_endpoints,omitempty"`
	Name              string `json:"name"`
	ProfileName       string `json:"profile_name"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	Subnet *[]TrafficManagerEndpointSpec `json:"subnet,omitempty"`
	// +optional
	TargetResourceId string `json:"target_resource_id,omitempty"`
	Type             string `json:"type"`
}

type TrafficManagerEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
