package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	Certificate string `json:"-" sensitive:"true" tf:"certificate"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	PrivateKey string `json:"-" sensitive:"true" tf:"private_key"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
}

type ComputeSslCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
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
