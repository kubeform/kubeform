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

type AwsApiGatewayDomainName struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayDomainNameSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayDomainNameStatus `json:"status,omitempty"`
}

type AwsApiGatewayDomainNameSpecEndpointConfiguration struct {
	Types []string `json:"types"`
}

type AwsApiGatewayDomainNameSpec struct {
	CertificateChain        string                        `json:"certificate_chain"`
	CertificatePrivateKey   string                        `json:"certificate_private_key"`
	CloudfrontDomainName    string                        `json:"cloudfront_domain_name"`
	CertificateArn          string                        `json:"certificate_arn"`
	CertificateUploadDate   string                        `json:"certificate_upload_date"`
	CloudfrontZoneId        string                        `json:"cloudfront_zone_id"`
	CertificateBody         string                        `json:"certificate_body"`
	CertificateName         string                        `json:"certificate_name"`
	EndpointConfiguration   []AwsApiGatewayDomainNameSpec `json:"endpoint_configuration"`
	RegionalCertificateName string                        `json:"regional_certificate_name"`
	RegionalZoneId          string                        `json:"regional_zone_id"`
	DomainName              string                        `json:"domain_name"`
	RegionalCertificateArn  string                        `json:"regional_certificate_arn"`
	RegionalDomainName      string                        `json:"regional_domain_name"`
}



type AwsApiGatewayDomainNameStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsApiGatewayDomainNameList is a list of AwsApiGatewayDomainNames
type AwsApiGatewayDomainNameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayDomainName CRD objects
	Items []AwsApiGatewayDomainName `json:"items,omitempty"`
}