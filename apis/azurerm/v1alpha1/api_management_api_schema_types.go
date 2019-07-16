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

type ApiManagementApiSchema struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementApiSchemaSpec   `json:"spec,omitempty"`
	Status            ApiManagementApiSchemaStatus `json:"status,omitempty"`
}

type ApiManagementApiSchemaSpec struct {
	ApiManagementName string `json:"api_management_name"`
	ApiName           string `json:"api_name"`
	ContentType       string `json:"content_type"`
	ResourceGroupName string `json:"resource_group_name"`
	SchemaId          string `json:"schema_id"`
	Value             string `json:"value"`
}

type ApiManagementApiSchemaStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementApiSchemaList is a list of ApiManagementApiSchemas
type ApiManagementApiSchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementApiSchema CRD objects
	Items []ApiManagementApiSchema `json:"items,omitempty"`
}
