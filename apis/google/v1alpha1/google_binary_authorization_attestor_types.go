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

type GoogleBinaryAuthorizationAttestor struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleBinaryAuthorizationAttestorSpec   `json:"spec,omitempty"`
	Status            GoogleBinaryAuthorizationAttestorStatus `json:"status,omitempty"`
}

type GoogleBinaryAuthorizationAttestorSpecAttestationAuthorityNotePublicKeys struct {
	AsciiArmoredPgpPublicKey string `json:"ascii_armored_pgp_public_key"`
	Comment                  string `json:"comment"`
	Id                       string `json:"id"`
}

type GoogleBinaryAuthorizationAttestorSpecAttestationAuthorityNote struct {
	NoteReference                 string                                                          `json:"note_reference"`
	PublicKeys                    []GoogleBinaryAuthorizationAttestorSpecAttestationAuthorityNote `json:"public_keys"`
	DelegationServiceAccountEmail string                                                          `json:"delegation_service_account_email"`
}

type GoogleBinaryAuthorizationAttestorSpec struct {
	AttestationAuthorityNote []GoogleBinaryAuthorizationAttestorSpec `json:"attestation_authority_note"`
	Name                     string                                  `json:"name"`
	Description              string                                  `json:"description"`
	Project                  string                                  `json:"project"`
}

type GoogleBinaryAuthorizationAttestorStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleBinaryAuthorizationAttestorList is a list of GoogleBinaryAuthorizationAttestors
type GoogleBinaryAuthorizationAttestorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleBinaryAuthorizationAttestor CRD objects
	Items []GoogleBinaryAuthorizationAttestor `json:"items,omitempty"`
}
