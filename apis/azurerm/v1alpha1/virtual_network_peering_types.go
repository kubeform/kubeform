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

type VirtualNetworkPeering struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkPeeringSpec   `json:"spec,omitempty"`
	Status            VirtualNetworkPeeringStatus `json:"status,omitempty"`
}

type VirtualNetworkPeeringSpec struct {
	Name                   string `json:"name"`
	RemoteVirtualNetworkId string `json:"remote_virtual_network_id"`
	ResourceGroupName      string `json:"resource_group_name"`
	VirtualNetworkName     string `json:"virtual_network_name"`
}

type VirtualNetworkPeeringStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VirtualNetworkPeeringList is a list of VirtualNetworkPeerings
type VirtualNetworkPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualNetworkPeering CRD objects
	Items []VirtualNetworkPeering `json:"items,omitempty"`
}
