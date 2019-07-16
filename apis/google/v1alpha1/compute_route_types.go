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

type ComputeRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeRouteSpec   `json:"spec,omitempty"`
	Status            ComputeRouteStatus `json:"status,omitempty"`
}

type ComputeRouteSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	DestRange   string `json:"dest_range"`
	Name        string `json:"name"`
	Network     string `json:"network"`
	// +optional
	NextHopGateway string `json:"next_hop_gateway,omitempty"`
	// +optional
	NextHopInstance string `json:"next_hop_instance,omitempty"`
	// +optional
	NextHopInstanceZone string `json:"next_hop_instance_zone,omitempty"`
	// +optional
	NextHopIp string `json:"next_hop_ip,omitempty"`
	// +optional
	NextHopVpnTunnel string `json:"next_hop_vpn_tunnel,omitempty"`
	// +optional
	Priority int `json:"priority,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty"`
}

type ComputeRouteStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
