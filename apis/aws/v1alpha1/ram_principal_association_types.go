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

type RamPrincipalAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RamPrincipalAssociationSpec   `json:"spec,omitempty"`
	Status            RamPrincipalAssociationStatus `json:"status,omitempty"`
}

type RamPrincipalAssociationSpec struct {
	Principal        string `json:"principal"`
	ResourceShareArn string `json:"resource_share_arn"`
}

type RamPrincipalAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RamPrincipalAssociationList is a list of RamPrincipalAssociations
type RamPrincipalAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RamPrincipalAssociation CRD objects
	Items []RamPrincipalAssociation `json:"items,omitempty"`
}
