package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermApiManagementApiSchema struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementApiSchemaSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementApiSchemaStatus `json:"status,omitempty"`
}

type AzurermApiManagementApiSchemaSpec struct {
	SchemaId          string `json:"schema_id"`
	ApiName           string `json:"api_name"`
	ResourceGroupName string `json:"resource_group_name"`
	ApiManagementName string `json:"api_management_name"`
	ContentType       string `json:"content_type"`
	Value             string `json:"value"`
}

type AzurermApiManagementApiSchemaStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermApiManagementApiSchemaList is a list of AzurermApiManagementApiSchemas
type AzurermApiManagementApiSchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementApiSchema CRD objects
	Items []AzurermApiManagementApiSchema `json:"items,omitempty"`
}
