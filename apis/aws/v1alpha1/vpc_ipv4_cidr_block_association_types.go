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

type VpcIpv4CidrBlockAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcIpv4CidrBlockAssociationSpec   `json:"spec,omitempty"`
	Status            VpcIpv4CidrBlockAssociationStatus `json:"status,omitempty"`
}

type VpcIpv4CidrBlockAssociationSpec struct {
	CidrBlock string `json:"cidr_block"`
	VpcId     string `json:"vpc_id"`
}

type VpcIpv4CidrBlockAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcIpv4CidrBlockAssociationList is a list of VpcIpv4CidrBlockAssociations
type VpcIpv4CidrBlockAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcIpv4CidrBlockAssociation CRD objects
	Items []VpcIpv4CidrBlockAssociation `json:"items,omitempty"`
}
