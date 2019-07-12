package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeSslCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeSslCertificateSpec   `json:"spec,omitempty"`
	Status            GoogleComputeSslCertificateStatus `json:"status,omitempty"`
}

type GoogleComputeSslCertificateSpec struct {
	Description       string `json:"description"`
	CertificateId     int    `json:"certificate_id"`
	CreationTimestamp string `json:"creation_timestamp"`
	SelfLink          string `json:"self_link"`
	Certificate       string `json:"certificate"`
	PrivateKey        string `json:"private_key"`
	Name              string `json:"name"`
	NamePrefix        string `json:"name_prefix"`
	Project           string `json:"project"`
}

type GoogleComputeSslCertificateStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeSslCertificateList is a list of GoogleComputeSslCertificates
type GoogleComputeSslCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeSslCertificate CRD objects
	Items []GoogleComputeSslCertificate `json:"items,omitempty"`
}
