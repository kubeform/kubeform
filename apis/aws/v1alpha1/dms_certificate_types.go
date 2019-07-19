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

type DmsCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DmsCertificateSpec   `json:"spec,omitempty"`
	Status            DmsCertificateStatus `json:"status,omitempty"`
}

type DmsCertificateSpec struct {
	CertificateID string `json:"certificateID" tf:"certificate_id"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	CertificatePem *core.LocalObjectReference `json:"certificatePem,omitempty" tf:"certificate_pem,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	CertificateWallet *core.LocalObjectReference `json:"certificateWallet,omitempty" tf:"certificate_wallet,omitempty"`
	ProviderRef       core.LocalObjectReference  `json:"providerRef" tf:"-"`
}

type DmsCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DmsCertificateList is a list of DmsCertificates
type DmsCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DmsCertificate CRD objects
	Items []DmsCertificate `json:"items,omitempty"`
}
