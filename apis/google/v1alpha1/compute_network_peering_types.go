package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeNetworkPeering struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeNetworkPeeringSpec   `json:"spec,omitempty"`
	Status            ComputeNetworkPeeringStatus `json:"status,omitempty"`
}

type ComputeNetworkPeeringSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	AutoCreateRoutes bool   `json:"autoCreateRoutes,omitempty" tf:"auto_create_routes,omitempty"`
	Name             string `json:"name" tf:"name"`
	Network          string `json:"network" tf:"network"`
	PeerNetwork      string `json:"peerNetwork" tf:"peer_network"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
	// +optional
	StateDetails string `json:"stateDetails,omitempty" tf:"state_details,omitempty"`
}

type ComputeNetworkPeeringStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeNetworkPeeringSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeNetworkPeeringList is a list of ComputeNetworkPeerings
type ComputeNetworkPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeNetworkPeering CRD objects
	Items []ComputeNetworkPeering `json:"items,omitempty"`
}
