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

type VpcDhcpOptionsAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcDhcpOptionsAssociationSpec   `json:"spec,omitempty"`
	Status            VpcDhcpOptionsAssociationStatus `json:"status,omitempty"`
}

type VpcDhcpOptionsAssociationSpec struct {
	DhcpOptionsId string `json:"dhcp_options_id"`
	VpcId         string `json:"vpc_id"`
}

type VpcDhcpOptionsAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcDhcpOptionsAssociationList is a list of VpcDhcpOptionsAssociations
type VpcDhcpOptionsAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcDhcpOptionsAssociation CRD objects
	Items []VpcDhcpOptionsAssociation `json:"items,omitempty"`
}
