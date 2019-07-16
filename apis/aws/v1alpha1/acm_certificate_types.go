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

type AcmCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AcmCertificateSpec   `json:"spec,omitempty"`
	Status            AcmCertificateStatus `json:"status,omitempty"`
}

type AcmCertificateSpec struct {
	// +optional
	CertificateBody string `json:"certificate_body,omitempty"`
	// +optional
	CertificateChain string `json:"certificate_chain,omitempty"`
	// +optional
	PrivateKey string `json:"private_key,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type AcmCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AcmCertificateList is a list of AcmCertificates
type AcmCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AcmCertificate CRD objects
	Items []AcmCertificate `json:"items,omitempty"`
}
