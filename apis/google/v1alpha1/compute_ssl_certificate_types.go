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

type ComputeSslCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSslCertificateSpec   `json:"spec,omitempty"`
	Status            ComputeSslCertificateStatus `json:"status,omitempty"`
}

type ComputeSslCertificateSpec struct {
	Certificate string `json:"certificate"`
	// +optional
	Description string `json:"description,omitempty"`
	PrivateKey  string `json:"private_key"`
}

type ComputeSslCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeSslCertificateList is a list of ComputeSslCertificates
type ComputeSslCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeSslCertificate CRD objects
	Items []ComputeSslCertificate `json:"items,omitempty"`
}
