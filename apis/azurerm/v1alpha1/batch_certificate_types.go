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

type BatchCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BatchCertificateSpec   `json:"spec,omitempty"`
	Status            BatchCertificateStatus `json:"status,omitempty"`
}

type BatchCertificateSpec struct {
	AccountName string `json:"accountName" tf:"account_name"`
	// Sensitive Data. Provide secret name which contains one value only
	Certificate core.LocalObjectReference `json:"certificate" tf:"certificate"`
	Format      string                    `json:"format" tf:"format"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	Password            core.LocalObjectReference `json:"password,omitempty" tf:"password,omitempty"`
	ResourceGroupName   string                    `json:"resourceGroupName" tf:"resource_group_name"`
	Thumbprint          string                    `json:"thumbprint" tf:"thumbprint"`
	ThumbprintAlgorithm string                    `json:"thumbprintAlgorithm" tf:"thumbprint_algorithm"`
	ProviderRef         core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BatchCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BatchCertificateList is a list of BatchCertificates
type BatchCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BatchCertificate CRD objects
	Items []BatchCertificate `json:"items,omitempty"`
}
