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

type GoogleComputeRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeRouteSpec   `json:"spec,omitempty"`
	Status            GoogleComputeRouteStatus `json:"status,omitempty"`
}

type GoogleComputeRouteSpec struct {
	DestRange           string   `json:"dest_range"`
	Tags                []string `json:"tags"`
	NextHopInstanceZone string   `json:"next_hop_instance_zone"`
	NextHopInstance     string   `json:"next_hop_instance"`
	Network             string   `json:"network"`
	NextHopIp           string   `json:"next_hop_ip"`
	NextHopVpnTunnel    string   `json:"next_hop_vpn_tunnel"`
	SelfLink            string   `json:"self_link"`
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	NextHopGateway      string   `json:"next_hop_gateway"`
	Priority            int      `json:"priority"`
	NextHopNetwork      string   `json:"next_hop_network"`
	Project             string   `json:"project"`
}



type GoogleComputeRouteStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeRouteList is a list of GoogleComputeRoutes
type GoogleComputeRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeRoute CRD objects
	Items []GoogleComputeRoute `json:"items,omitempty"`
}