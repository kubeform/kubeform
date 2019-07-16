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

type ApiManagementApi struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementApiSpec   `json:"spec,omitempty"`
	Status            ApiManagementApiStatus `json:"status,omitempty"`
}

type ApiManagementApiSpecImportWsdlSelector struct {
	EndpointName string `json:"endpoint_name"`
	ServiceName  string `json:"service_name"`
}

type ApiManagementApiSpecImport struct {
	ContentFormat string `json:"content_format"`
	ContentValue  string `json:"content_value"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	WsdlSelector *[]ApiManagementApiSpecImport `json:"wsdl_selector,omitempty"`
}

type ApiManagementApiSpec struct {
	ApiManagementName string `json:"api_management_name"`
	// +optional
	Description string `json:"description,omitempty"`
	DisplayName string `json:"display_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Import *[]ApiManagementApiSpec `json:"import,omitempty"`
	Name   string                  `json:"name"`
	Path   string                  `json:"path"`
	// +kubebuilder:validation:UniqueItems=true
	Protocols         []string `json:"protocols"`
	ResourceGroupName string   `json:"resource_group_name"`
	Revision          string   `json:"revision"`
	// +optional
	SoapPassThrough bool `json:"soap_pass_through,omitempty"`
}

type ApiManagementApiStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementApiList is a list of ApiManagementApis
type ApiManagementApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementApi CRD objects
	Items []ApiManagementApi `json:"items,omitempty"`
}
