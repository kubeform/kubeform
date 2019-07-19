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

type ApiManagementCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementCertificateSpec   `json:"spec,omitempty"`
	Status            ApiManagementCertificateStatus `json:"status,omitempty"`
}

type ApiManagementCertificateSpec struct {
	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`
	// Sensitive Data. Provide secret name which contains one value only
	Data core.LocalObjectReference `json:"data" tf:"data"`
	Name string                    `json:"name" tf:"name"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	Password          core.LocalObjectReference `json:"password,omitempty" tf:"password,omitempty"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
	ProviderRef       core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ApiManagementCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
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
