package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ContainerAnalysisNote struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerAnalysisNoteSpec   `json:"spec,omitempty"`
	Status            ContainerAnalysisNoteStatus `json:"status,omitempty"`
}

type ContainerAnalysisNoteSpecAttestationAuthorityHint struct {
	HumanReadableName string `json:"humanReadableName" tf:"human_readable_name"`
}

type ContainerAnalysisNoteSpecAttestationAuthority struct {
	// +kubebuilder:validation:MaxItems=1
	Hint []ContainerAnalysisNoteSpecAttestationAuthorityHint `json:"hint" tf:"hint"`
}

type ContainerAnalysisNoteSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +kubebuilder:validation:MaxItems=1
	AttestationAuthority []ContainerAnalysisNoteSpecAttestationAuthority `json:"attestationAuthority" tf:"attestation_authority"`
	Name                 string                                          `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
}

type ContainerAnalysisNoteStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *ContainerAnalysisNoteSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ContainerAnalysisNoteList is a list of ContainerAnalysisNotes
type ContainerAnalysisNoteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ContainerAnalysisNote CRD objects
	Items []ContainerAnalysisNote `json:"items,omitempty"`
}
