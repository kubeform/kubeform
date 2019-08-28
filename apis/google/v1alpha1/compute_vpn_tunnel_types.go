package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	DetailedStatus string `json:"detailedStatus,omitempty" tf:"detailed_status,omitempty"`
	// +optional
	IkeVersion int `json:"ikeVersion,omitempty" tf:"ike_version,omitempty"`
	// +optional
	LabelFingerprint string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`
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
	// +optional
	SelfLink     string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	SharedSecret string `json:"-" sensitive:"true" tf:"shared_secret"`
	// +optional
	SharedSecretHash string `json:"sharedSecretHash,omitempty" tf:"shared_secret_hash,omitempty"`
	TargetVPNGateway string `json:"targetVPNGateway" tf:"target_vpn_gateway"`
}



type ComputeVPNTunnelStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *ComputeVPNTunnelSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
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