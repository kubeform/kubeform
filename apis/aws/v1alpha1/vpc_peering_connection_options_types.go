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

type VpcPeeringConnectionOptions struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcPeeringConnectionOptionsSpec   `json:"spec,omitempty"`
	Status            VpcPeeringConnectionOptionsStatus `json:"status,omitempty"`
}

type VpcPeeringConnectionOptionsSpec struct {
	VpcPeeringConnectionId string `json:"vpc_peering_connection_id"`
}

type VpcPeeringConnectionOptionsStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcPeeringConnectionOptionsList is a list of VpcPeeringConnectionOptionss
type VpcPeeringConnectionOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcPeeringConnectionOptions CRD objects
	Items []VpcPeeringConnectionOptions `json:"items,omitempty"`
}
