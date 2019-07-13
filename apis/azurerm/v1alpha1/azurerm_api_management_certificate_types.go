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

type AzurermApiManagementCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementCertificateSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementCertificateStatus `json:"status,omitempty"`
}

type AzurermApiManagementCertificateSpec struct {
	ApiManagementName string `json:"api_management_name"`
	Data              string `json:"data"`
	Password          string `json:"password"`
	Expiration        string `json:"expiration"`
	Subject           string `json:"subject"`
	Thumbprint        string `json:"thumbprint"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
}



type AzurermApiManagementCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApiManagementCertificateList is a list of AzurermApiManagementCertificates
type AzurermApiManagementCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementCertificate CRD objects
	Items []AzurermApiManagementCertificate `json:"items,omitempty"`
}