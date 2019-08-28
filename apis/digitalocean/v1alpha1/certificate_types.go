package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Certificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateSpec   `json:"spec,omitempty"`
	Status            CertificateStatus `json:"status,omitempty"`
}

type CertificateSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	CertificateChain string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Domains []string `json:"domains,omitempty" tf:"domains,omitempty"`
	// +optional
	LeafCertificate string `json:"leafCertificate,omitempty" tf:"leaf_certificate,omitempty"`
	Name            string `json:"name" tf:"name"`
	// +optional
	NotAfter string `json:"notAfter,omitempty" tf:"not_after,omitempty"`
	// +optional
	PrivateKey string `json:"-" sensitive:"true" tf:"private_key,omitempty"`
	// +optional
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty" tf:"sha1_fingerprint,omitempty"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}



type CertificateStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *CertificateSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CertificateList is a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Certificate CRD objects
	Items []Certificate `json:"items,omitempty"`
}