package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ApiManagementCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementCertificateSpec   `json:"spec,omitempty"`
	Status            ApiManagementCertificateStatus `json:"status,omitempty"`
}

type ApiManagementCertificateSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`
	Data              string `json:"-" sensitive:"true" tf:"data"`
	// +optional
	Expiration string `json:"expiration,omitempty" tf:"expiration,omitempty"`
	Name       string `json:"name" tf:"name"`
	// +optional
	Password          string `json:"-" sensitive:"true" tf:"password,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Subject string `json:"subject,omitempty" tf:"subject,omitempty"`
	// +optional
	Thumbprint string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type ApiManagementCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiManagementCertificateSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementCertificateList is a list of ApiManagementCertificates
type ApiManagementCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementCertificate CRD objects
	Items []ApiManagementCertificate `json:"items,omitempty"`
}
