package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeVPNTunnel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeVPNTunnelSpec   `json:"spec,omitempty"`
	Status            ComputeVPNTunnelStatus `json:"status,omitempty"`
}

type ComputeVPNTunnelSpec struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	IkeVersion int `json:"ikeVersion,omitempty" tf:"ike_version,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	LocalTrafficSelector []string `json:"localTrafficSelector,omitempty" tf:"local_traffic_selector,omitempty"`
	Name                 string   `json:"name" tf:"name"`
	PeerIP               string   `json:"peerIP" tf:"peer_ip"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	RemoteTrafficSelector []string `json:"remoteTrafficSelector,omitempty" tf:"remote_traffic_selector,omitempty"`
	// +optional
	Router string `json:"router,omitempty" tf:"router,omitempty"`
	// Sensitive Data. Provide secret name which contains one value only
	SharedSecret     core.LocalObjectReference `json:"sharedSecret" tf:"shared_secret"`
	TargetVPNGateway string                    `json:"targetVPNGateway" tf:"target_vpn_gateway"`
	ProviderRef      core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeVPNTunnelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeVPNTunnelList is a list of ComputeVPNTunnels
type ComputeVPNTunnelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeVPNTunnel CRD objects
	Items []ComputeVPNTunnel `json:"items,omitempty"`
}
