package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ApiGatewayDomainName struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayDomainNameSpec   `json:"spec,omitempty"`
	Status            ApiGatewayDomainNameStatus `json:"status,omitempty"`
}

type ApiGatewayDomainNameSpecEndpointConfiguration struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Types []string `json:"types" tf:"types"`
}

type ApiGatewayDomainNameSpec struct {
	Secret      *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	// +optional
	CertificateArn string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`
	// +optional
	CertificateBody string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`
	// +optional
	CertificateChain string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`
	// +optional
	CertificateName string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`
	// +optional
	DomainName string `json:"domainName" tf:"domain_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	EndpointConfiguration []ApiGatewayDomainNameSpecEndpointConfiguration `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration,omitempty"`
	// +optional
	RegionalCertificateArn string `json:"regionalCertificateArn,omitempty" tf:"regional_certificate_arn,omitempty"`
	// +optional
	RegionalCertificateName string `json:"regionalCertificateName,omitempty" tf:"regional_certificate_name,omitempty"`
}

type ApiGatewayDomainNameStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
