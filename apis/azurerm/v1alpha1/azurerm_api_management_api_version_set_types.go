package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermApiManagementApiVersionSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementApiVersionSetSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementApiVersionSetStatus `json:"status,omitempty"`
}

type AzurermApiManagementApiVersionSetSpec struct {
	VersionHeaderName string `json:"version_header_name"`
	VersionQueryName  string `json:"version_query_name"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	ApiManagementName string `json:"api_management_name"`
	DisplayName       string `json:"display_name"`
	VersioningScheme  string `json:"versioning_scheme"`
	Description       string `json:"description"`
}

type AzurermApiManagementApiVersionSetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermApiManagementApiVersionSetList is a list of AzurermApiManagementApiVersionSets
type AzurermApiManagementApiVersionSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementApiVersionSet CRD objects
	Items []AzurermApiManagementApiVersionSet `json:"items,omitempty"`
}
