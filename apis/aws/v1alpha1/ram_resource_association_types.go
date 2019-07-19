package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	ResourceArn      string                    `json:"resourceArn" tf:"resource_arn"`
	ResourceShareArn string                    `json:"resourceShareArn" tf:"resource_share_arn"`
	ProviderRef      core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type RamResourceAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
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
