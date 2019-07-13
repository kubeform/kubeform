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

type DigitaloceanCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanCertificateSpec   `json:"spec,omitempty"`
	Status            DigitaloceanCertificateStatus `json:"status,omitempty"`
}

type DigitaloceanCertificateSpec struct {
	NotAfter         string   `json:"not_after"`
	CertificateChain string   `json:"certificate_chain"`
	Type             string   `json:"type"`
	LeafCertificate  string   `json:"leaf_certificate"`
	Domains          []string `json:"domains"`
	State            string   `json:"state"`
	Sha1Fingerprint  string   `json:"sha1_fingerprint"`
	Name             string   `json:"name"`
	PrivateKey       string   `json:"private_key"`
}



type DigitaloceanCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DigitaloceanCertificateList is a list of DigitaloceanCertificates
type DigitaloceanCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanCertificate CRD objects
	Items []DigitaloceanCertificate `json:"items,omitempty"`
}