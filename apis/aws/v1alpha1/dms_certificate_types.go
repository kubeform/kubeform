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

type DmsCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DmsCertificateSpec   `json:"spec,omitempty"`
	Status            DmsCertificateStatus `json:"status,omitempty"`
}

type DmsCertificateSpec struct {
	CertificateId string `json:"certificate_id"`
	// +optional
	CertificatePem string `json:"certificate_pem,omitempty"`
	// +optional
	CertificateWallet string `json:"certificate_wallet,omitempty"`
}

type DmsCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
