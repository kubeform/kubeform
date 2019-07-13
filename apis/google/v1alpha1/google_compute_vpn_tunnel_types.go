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
	SharedSecretHash      string            `json:"shared_secret_hash"`
	PeerIp                string            `json:"peer_ip"`
	DetailedStatus        string            `json:"detailed_status"`
	LabelFingerprint      string            `json:"label_fingerprint"`
	Router                string            `json:"router"`
	Description           string            `json:"description"`
	IkeVersion            int               `json:"ike_version"`
	LocalTrafficSelector  []string          `json:"local_traffic_selector"`
	Labels                map[string]string `json:"labels"`
	RemoteTrafficSelector []string          `json:"remote_traffic_selector"`
	CreationTimestamp     string            `json:"creation_timestamp"`
	SelfLink              string            `json:"self_link"`
	Name                  string            `json:"name"`
	SharedSecret          string            `json:"shared_secret"`
	TargetVpnGateway      string            `json:"target_vpn_gateway"`
	Region                string            `json:"region"`
	Project               string            `json:"project"`
}



type GoogleComputeVpnTunnelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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