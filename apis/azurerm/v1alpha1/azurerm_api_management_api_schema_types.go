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

type AzurermApiManagementApiSchema struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementApiSchemaSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementApiSchemaStatus `json:"status,omitempty"`
}

type AzurermApiManagementApiSchemaSpec struct {
	ResourceGroupName string `json:"resource_group_name"`
	ApiManagementName string `json:"api_management_name"`
	ContentType       string `json:"content_type"`
	Value             string `json:"value"`
	SchemaId          string `json:"schema_id"`
	ApiName           string `json:"api_name"`
}



type AzurermApiManagementApiSchemaStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApiManagementApiSchemaList is a list of AzurermApiManagementApiSchemas
type AzurermApiManagementApiSchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementApiSchema CRD objects
	Items []AzurermApiManagementApiSchema `json:"items,omitempty"`
}