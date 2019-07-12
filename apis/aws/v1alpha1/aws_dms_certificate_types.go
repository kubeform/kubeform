package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDmsCertificateSpec   `json:"spec,omitempty"`
	Status            AwsDmsCertificateStatus `json:"status,omitempty"`
}

type AwsDmsCertificateSpec struct {
	CertificateId     string `json:"certificate_id"`
	CertificatePem    string `json:"certificate_pem"`
	CertificateWallet string `json:"certificate_wallet"`
	CertificateArn    string `json:"certificate_arn"`
}



type AwsDmsCertificateStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDmsCertificateList is a list of AwsDmsCertificates
type AwsDmsCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDmsCertificate CRD objects
	Items []AwsDmsCertificate `json:"items,omitempty"`
}