package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

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
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
}

type BinaryAuthorizationAttestorSpecAttestationAuthorityNote struct {
	// +optional
	DelegationServiceAccountEmail string `json:"delegationServiceAccountEmail,omitempty" tf:"delegation_service_account_email,omitempty"`
	NoteReference                 string `json:"noteReference" tf:"note_reference"`
	// +optional
	PublicKeys []BinaryAuthorizationAttestorSpecAttestationAuthorityNotePublicKeys `json:"publicKeys,omitempty" tf:"public_keys,omitempty"`
}

type BinaryAuthorizationAttestorSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MaxItems=1
	AttestationAuthorityNote []BinaryAuthorizationAttestorSpecAttestationAuthorityNote `json:"attestationAuthorityNote" tf:"attestation_authority_note"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
}

type BinaryAuthorizationAttestorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *BinaryAuthorizationAttestorSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
