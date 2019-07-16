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

type ContainerAnalysisNote struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerAnalysisNoteSpec   `json:"spec,omitempty"`
	Status            ContainerAnalysisNoteStatus `json:"status,omitempty"`
}

type ContainerAnalysisNoteSpecAttestationAuthorityHint struct {
	HumanReadableName string `json:"human_readable_name"`
}

type ContainerAnalysisNoteSpecAttestationAuthority struct {
	// +kubebuilder:validation:MaxItems=1
	Hint []ContainerAnalysisNoteSpecAttestationAuthority `json:"hint"`
}

type ContainerAnalysisNoteSpec struct {
	// +kubebuilder:validation:MaxItems=1
	AttestationAuthority []ContainerAnalysisNoteSpec `json:"attestation_authority"`
	Name                 string                      `json:"name"`
}

type ContainerAnalysisNoteStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
