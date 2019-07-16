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

type RamResourceAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RamResourceAssociationSpec   `json:"spec,omitempty"`
	Status            RamResourceAssociationStatus `json:"status,omitempty"`
}

type RamResourceAssociationSpec struct {
	ResourceArn      string `json:"resource_arn"`
	ResourceShareArn string `json:"resource_share_arn"`
}

type RamResourceAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RamResourceAssociationList is a list of RamResourceAssociations
type RamResourceAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RamResourceAssociation CRD objects
	Items []RamResourceAssociation `json:"items,omitempty"`
}
