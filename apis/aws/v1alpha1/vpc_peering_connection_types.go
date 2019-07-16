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

type VpcPeeringConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcPeeringConnectionSpec   `json:"spec,omitempty"`
	Status            VpcPeeringConnectionStatus `json:"status,omitempty"`
}

type VpcPeeringConnectionSpec struct {
	// +optional
	AutoAccept bool   `json:"auto_accept,omitempty"`
	PeerVpcId  string `json:"peer_vpc_id"`
	// +optional
	Tags  map[string]string `json:"tags,omitempty"`
	VpcId string            `json:"vpc_id"`
}

type VpcPeeringConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcPeeringConnectionList is a list of VpcPeeringConnections
type VpcPeeringConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcPeeringConnection CRD objects
	Items []VpcPeeringConnection `json:"items,omitempty"`
}
