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

type ApiManagementApiOperation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementApiOperationSpec   `json:"spec,omitempty"`
	Status            ApiManagementApiOperationStatus `json:"status,omitempty"`
}

type ApiManagementApiOperationSpecResponseHeader struct {
	// +optional
	DefaultValue string `json:"default_value,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	Required    bool   `json:"required"`
	Type        string `json:"type"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty"`
}

type ApiManagementApiOperationSpecResponseRepresentationFormParameter struct {
	// +optional
	DefaultValue string `json:"default_value,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	Required    bool   `json:"required"`
	Type        string `json:"type"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty"`
}

type ApiManagementApiOperationSpecResponseRepresentation struct {
	ContentType string `json:"content_type"`
	// +optional
	FormParameter *[]ApiManagementApiOperationSpecResponseRepresentation `json:"form_parameter,omitempty"`
	// +optional
	Sample string `json:"sample,omitempty"`
	// +optional
	SchemaId string `json:"schema_id,omitempty"`
	// +optional
	TypeName string `json:"type_name,omitempty"`
}

type ApiManagementApiOperationSpecResponse struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Header *[]ApiManagementApiOperationSpecResponse `json:"header,omitempty"`
	// +optional
	Representation *[]ApiManagementApiOperationSpecResponse `json:"representation,omitempty"`
	StatusCode     int                                      `json:"status_code"`
}

type ApiManagementApiOperationSpecTemplateParameter struct {
	// +optional
	DefaultValue string `json:"default_value,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	Required    bool   `json:"required"`
	Type        string `json:"type"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty"`
}

type ApiManagementApiOperationSpec struct {
	ApiManagementName string `json:"api_management_name"`
	ApiName           string `json:"api_name"`
	// +optional
	Description       string `json:"description,omitempty"`
	DisplayName       string `json:"display_name"`
	Method            string `json:"method"`
	OperationId       string `json:"operation_id"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	Response *[]ApiManagementApiOperationSpec `json:"response,omitempty"`
	// +optional
	TemplateParameter *[]ApiManagementApiOperationSpec `json:"template_parameter,omitempty"`
	UrlTemplate       string                           `json:"url_template"`
}

type ApiManagementApiOperationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementApiOperationList is a list of ApiManagementApiOperations
type ApiManagementApiOperationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementApiOperation CRD objects
	Items []ApiManagementApiOperation `json:"items,omitempty"`
}
