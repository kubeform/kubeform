package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type ApiGatewayClientCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayClientCertificateSpec   `json:"spec,omitempty"`
	Status            ApiGatewayClientCertificateStatus `json:"status,omitempty"`
}

type ApiGatewayClientCertificateSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CreatedDate string `json:"createdDate,omitempty" tf:"created_date,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	ExpirationDate string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`
	// +optional
	PemEncodedCertificate string `json:"pemEncodedCertificate,omitempty" tf:"pem_encoded_certificate,omitempty"`
}

type ApiGatewayClientCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiGatewayClientCertificateSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayClientCertificateList is a list of ApiGatewayClientCertificates
type ApiGatewayClientCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayClientCertificate CRD objects
	Items []ApiGatewayClientCertificate `json:"items,omitempty"`
}
