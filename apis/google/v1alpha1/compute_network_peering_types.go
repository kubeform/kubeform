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

type ComputeNetworkPeering struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeNetworkPeeringSpec   `json:"spec,omitempty"`
	Status            ComputeNetworkPeeringStatus `json:"status,omitempty"`
}

type ComputeNetworkPeeringSpec struct {
	// +optional
	AutoCreateRoutes bool   `json:"auto_create_routes,omitempty"`
	Name             string `json:"name"`
	Network          string `json:"network"`
	PeerNetwork      string `json:"peer_network"`
}

type ComputeNetworkPeeringStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
