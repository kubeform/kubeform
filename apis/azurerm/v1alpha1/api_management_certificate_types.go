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

type ApiManagementCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementCertificateSpec   `json:"spec,omitempty"`
	Status            ApiManagementCertificateStatus `json:"status,omitempty"`
}

type ApiManagementCertificateSpec struct {
	ApiManagementName string `json:"api_management_name"`
	Data              string `json:"data"`
	Name              string `json:"name"`
	// +optional
	Password          string `json:"password,omitempty"`
	ResourceGroupName string `json:"resource_group_name"`
}

type ApiManagementCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
