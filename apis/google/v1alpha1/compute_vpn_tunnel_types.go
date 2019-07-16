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

type ComputeVpnTunnel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeVpnTunnelSpec   `json:"spec,omitempty"`
	Status            ComputeVpnTunnelStatus `json:"status,omitempty"`
}

type ComputeVpnTunnelSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	IkeVersion int `json:"ike_version,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	Name   string            `json:"name"`
	PeerIp string            `json:"peer_ip"`
	// +optional
	Router           string `json:"router,omitempty"`
	SharedSecret     string `json:"shared_secret"`
	TargetVpnGateway string `json:"target_vpn_gateway"`
}

type ComputeVpnTunnelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeVpnTunnelList is a list of ComputeVpnTunnels
type ComputeVpnTunnelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeVpnTunnel CRD objects
	Items []ComputeVpnTunnel `json:"items,omitempty"`
}
