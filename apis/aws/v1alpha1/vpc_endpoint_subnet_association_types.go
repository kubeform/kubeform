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

type VpcEndpointSubnetAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcEndpointSubnetAssociationSpec   `json:"spec,omitempty"`
	Status            VpcEndpointSubnetAssociationStatus `json:"status,omitempty"`
}

type VpcEndpointSubnetAssociationSpec struct {
	SubnetId      string `json:"subnet_id"`
	VpcEndpointId string `json:"vpc_endpoint_id"`
}

type VpcEndpointSubnetAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcEndpointSubnetAssociationList is a list of VpcEndpointSubnetAssociations
type VpcEndpointSubnetAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcEndpointSubnetAssociation CRD objects
	Items []VpcEndpointSubnetAssociation `json:"items,omitempty"`
}
