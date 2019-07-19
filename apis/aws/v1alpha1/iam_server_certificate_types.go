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

type IamServerCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamServerCertificateSpec   `json:"spec,omitempty"`
	Status            IamServerCertificateStatus `json:"status,omitempty"`
}

type IamServerCertificateSpec struct {
	// +optional
	Arn             string `json:"arn,omitempty" tf:"arn,omitempty"`
	CertificateBody string `json:"certificateBody" tf:"certificate_body"`
	// +optional
	CertificateChain string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	// Sensitive Data. Provide secret name which contains one value only
	PrivateKey  *core.LocalObjectReference `json:"privateKey" tf:"private_key"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
}

type IamServerCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamServerCertificateList is a list of IamServerCertificates
type IamServerCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamServerCertificate CRD objects
	Items []IamServerCertificate `json:"items,omitempty"`
}
