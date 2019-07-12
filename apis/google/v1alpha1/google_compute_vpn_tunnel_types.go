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

type GoogleComputeVpnTunnel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeVpnTunnelSpec   `json:"spec,omitempty"`
	Status            GoogleComputeVpnTunnelStatus `json:"status,omitempty"`
}

type GoogleComputeVpnTunnelSpec struct {
	DetailedStatus        string            `json:"detailed_status"`
	LabelFingerprint      string            `json:"label_fingerprint"`
	Project               string            `json:"project"`
	PeerIp                string            `json:"peer_ip"`
	RemoteTrafficSelector []string          `json:"remote_traffic_selector"`
	Region                string            `json:"region"`
	CreationTimestamp     string            `json:"creation_timestamp"`
	Labels                map[string]string `json:"labels"`
	LocalTrafficSelector  []string          `json:"local_traffic_selector"`
	SharedSecret          string            `json:"shared_secret"`
	Router                string            `json:"router"`
	Description           string            `json:"description"`
	IkeVersion            int               `json:"ike_version"`
	SharedSecretHash      string            `json:"shared_secret_hash"`
	SelfLink              string            `json:"self_link"`
	Name                  string            `json:"name"`
	TargetVpnGateway      string            `json:"target_vpn_gateway"`
}

type GoogleComputeVpnTunnelStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeVpnTunnelList is a list of GoogleComputeVpnTunnels
type GoogleComputeVpnTunnelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeVpnTunnel CRD objects
	Items []GoogleComputeVpnTunnel `json:"items,omitempty"`
}
