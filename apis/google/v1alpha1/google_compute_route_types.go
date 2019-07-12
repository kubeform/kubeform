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
	SelfLink            string   `json:"self_link"`
	Name                string   `json:"name"`
	Network             string   `json:"network"`
	NextHopInstance     string   `json:"next_hop_instance"`
	NextHopVpnTunnel    string   `json:"next_hop_vpn_tunnel"`
	NextHopNetwork      string   `json:"next_hop_network"`
	Project             string   `json:"project"`
	Description         string   `json:"description"`
	DestRange           string   `json:"dest_range"`
	NextHopGateway      string   `json:"next_hop_gateway"`
	NextHopIp           string   `json:"next_hop_ip"`
	Tags                []string `json:"tags"`
	NextHopInstanceZone string   `json:"next_hop_instance_zone"`
	Priority            int      `json:"priority"`
}

type GoogleComputeRouteStatus struct {
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
