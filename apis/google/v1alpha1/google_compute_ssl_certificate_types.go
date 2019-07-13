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

type GoogleComputeSslCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeSslCertificateSpec   `json:"spec,omitempty"`
	Status            GoogleComputeSslCertificateStatus `json:"status,omitempty"`
}

type GoogleComputeSslCertificateSpec struct {
	SelfLink          string `json:"self_link"`
	Certificate       string `json:"certificate"`
	PrivateKey        string `json:"private_key"`
	CertificateId     int    `json:"certificate_id"`
	CreationTimestamp string `json:"creation_timestamp"`
	NamePrefix        string `json:"name_prefix"`
	Description       string `json:"description"`
	Name              string `json:"name"`
	Project           string `json:"project"`
}



type GoogleComputeSslCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeSslCertificateList is a list of GoogleComputeSslCertificates
type GoogleComputeSslCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeSslCertificate CRD objects
	Items []GoogleComputeSslCertificate `json:"items,omitempty"`
}