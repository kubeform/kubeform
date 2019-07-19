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

type BinaryAuthorizationAttestor struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BinaryAuthorizationAttestorSpec   `json:"spec,omitempty"`
	Status            BinaryAuthorizationAttestorStatus `json:"status,omitempty"`
}

type BinaryAuthorizationAttestorSpecAttestationAuthorityNotePublicKeys struct {
	AsciiArmoredPgpPublicKey string `json:"asciiArmoredPgpPublicKey" tf:"ascii_armored_pgp_public_key"`
	// +optional
	Comment string `json:"comment,omitempty" tf:"comment,omitempty"`
}

type BinaryAuthorizationAttestorSpecAttestationAuthorityNote struct {
	NoteReference string `json:"noteReference" tf:"note_reference"`
	// +optional
	PublicKeys []BinaryAuthorizationAttestorSpecAttestationAuthorityNotePublicKeys `json:"publicKeys,omitempty" tf:"public_keys,omitempty"`
}

type BinaryAuthorizationAttestorSpec struct {
	// +kubebuilder:validation:MaxItems=1
	AttestationAuthorityNote []BinaryAuthorizationAttestorSpecAttestationAuthorityNote `json:"attestationAuthorityNote" tf:"attestation_authority_note"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Project     string                    `json:"project,omitempty" tf:"project,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BinaryAuthorizationAttestorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BinaryAuthorizationAttestorList is a list of BinaryAuthorizationAttestors
type BinaryAuthorizationAttestorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BinaryAuthorizationAttestor CRD objects
	Items []BinaryAuthorizationAttestor `json:"items,omitempty"`
}
