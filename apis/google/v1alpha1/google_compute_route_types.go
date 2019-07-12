package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeRouteSpec   `json:"spec,omitempty"`
	Status            GoogleComputeRouteStatus `json:"status,omitempty"`
}

type GoogleComputeRouteSpec struct {
	DestRange           string   `json:"dest_range"`
	Description         string   `json:"description"`
	NextHopGateway      string   `json:"next_hop_gateway"`
	Name                string   `json:"name"`
	NextHopInstance     string   `json:"next_hop_instance"`
	Priority            int      `json:"priority"`
	NextHopInstanceZone string   `json:"next_hop_instance_zone"`
	SelfLink            string   `json:"self_link"`
	Network             string   `json:"network"`
	NextHopVpnTunnel    string   `json:"next_hop_vpn_tunnel"`
	Tags                []string `json:"tags"`
	NextHopNetwork      string   `json:"next_hop_network"`
	Project             string   `json:"project"`
	NextHopIp           string   `json:"next_hop_ip"`
}

type GoogleComputeRouteStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeRouteList is a list of GoogleComputeRoutes
type GoogleComputeRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeRoute CRD objects
	Items []GoogleComputeRoute `json:"items,omitempty"`
}
