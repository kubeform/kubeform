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

type ApiGatewayDomainName struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayDomainNameSpec   `json:"spec,omitempty"`
	Status            ApiGatewayDomainNameStatus `json:"status,omitempty"`
}

type ApiGatewayDomainNameSpec struct {
	// +optional
	CertificateArn string `json:"certificate_arn,omitempty"`
	// +optional
	CertificateBody string `json:"certificate_body,omitempty"`
	// +optional
	CertificateChain string `json:"certificate_chain,omitempty"`
	// +optional
	CertificateName string `json:"certificate_name,omitempty"`
	// +optional
	CertificatePrivateKey string `json:"certificate_private_key,omitempty"`
	DomainName            string `json:"domain_name"`
	// +optional
	RegionalCertificateArn string `json:"regional_certificate_arn,omitempty"`
	// +optional
	RegionalCertificateName string `json:"regional_certificate_name,omitempty"`
}

type ApiGatewayDomainNameStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayDomainNameList is a list of ApiGatewayDomainNames
type ApiGatewayDomainNameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayDomainName CRD objects
	Items []ApiGatewayDomainName `json:"items,omitempty"`
}
