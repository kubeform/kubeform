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

type IamServerCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamServerCertificateSpec   `json:"spec,omitempty"`
	Status            IamServerCertificateStatus `json:"status,omitempty"`
}

type IamServerCertificateSpec struct {
	CertificateBody string `json:"certificate_body"`
	// +optional
	CertificateChain string `json:"certificate_chain,omitempty"`
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
	// +optional
	Path       string `json:"path,omitempty"`
	PrivateKey string `json:"private_key"`
}

type IamServerCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamServerCertificateList is a list of IamServerCertificates
type IamServerCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamServerCertificate CRD objects
	Items []IamServerCertificate `json:"items,omitempty"`
}
