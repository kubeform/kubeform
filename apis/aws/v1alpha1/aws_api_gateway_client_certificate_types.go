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

type AwsApiGatewayClientCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayClientCertificateSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayClientCertificateStatus `json:"status,omitempty"`
}

type AwsApiGatewayClientCertificateSpec struct {
	PemEncodedCertificate string `json:"pem_encoded_certificate"`
	Description           string `json:"description"`
	CreatedDate           string `json:"created_date"`
	ExpirationDate        string `json:"expiration_date"`
}

type AwsApiGatewayClientCertificateStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsApiGatewayClientCertificateList is a list of AwsApiGatewayClientCertificates
type AwsApiGatewayClientCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayClientCertificate CRD objects
	Items []AwsApiGatewayClientCertificate `json:"items,omitempty"`
}
