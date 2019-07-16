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

type VpcPeeringConnectionAccepter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcPeeringConnectionAccepterSpec   `json:"spec,omitempty"`
	Status            VpcPeeringConnectionAccepterStatus `json:"status,omitempty"`
}

type VpcPeeringConnectionAccepterSpec struct {
	// +optional
	AutoAccept bool `json:"auto_accept,omitempty"`
	// +optional
	Tags                   map[string]string `json:"tags,omitempty"`
	VpcPeeringConnectionId string            `json:"vpc_peering_connection_id"`
}

type VpcPeeringConnectionAccepterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcPeeringConnectionAccepterList is a list of VpcPeeringConnectionAccepters
type VpcPeeringConnectionAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcPeeringConnectionAccepter CRD objects
	Items []VpcPeeringConnectionAccepter `json:"items,omitempty"`
}
