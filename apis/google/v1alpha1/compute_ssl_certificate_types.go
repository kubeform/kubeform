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

type ComputeSSLCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSSLCertificateSpec   `json:"spec,omitempty"`
	Status            ComputeSSLCertificateStatus `json:"status,omitempty"`
}

type ComputeSSLCertificateSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	Certificate string `json:"-" sensitive:"true" tf:"certificate"`
	// +optional
	CertificateID int `json:"certificateID,omitempty" tf:"certificate_id,omitempty"`
	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	PrivateKey string `json:"-" sensitive:"true" tf:"private_key"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type ComputeSSLCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeSSLCertificateSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeSSLCertificateList is a list of ComputeSSLCertificates
type ComputeSSLCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeSSLCertificate CRD objects
	Items []ComputeSSLCertificate `json:"items,omitempty"`
}
