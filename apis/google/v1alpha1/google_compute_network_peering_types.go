package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeNetworkPeering struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeNetworkPeeringSpec   `json:"spec,omitempty"`
	Status            GoogleComputeNetworkPeeringStatus `json:"status,omitempty"`
}

type GoogleComputeNetworkPeeringSpec struct {
	Network          string `json:"network"`
	PeerNetwork      string `json:"peer_network"`
	AutoCreateRoutes bool   `json:"auto_create_routes"`
	State            string `json:"state"`
	StateDetails     string `json:"state_details"`
	Name             string `json:"name"`
}

type GoogleComputeNetworkPeeringStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeNetworkPeeringList is a list of GoogleComputeNetworkPeerings
type GoogleComputeNetworkPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeNetworkPeering CRD objects
	Items []GoogleComputeNetworkPeering `json:"items,omitempty"`
}
