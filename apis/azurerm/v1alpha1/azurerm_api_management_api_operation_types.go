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

type AzurermApiManagementApiOperation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementApiOperationSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementApiOperationStatus `json:"status,omitempty"`
}

type AzurermApiManagementApiOperationSpecRequestHeader struct {
	Type         string   `json:"type"`
	DefaultValue string   `json:"default_value"`
	Values       []string `json:"values"`
	Name         string   `json:"name"`
	Required     bool     `json:"required"`
	Description  string   `json:"description"`
}

type AzurermApiManagementApiOperationSpecRequestQueryParameter struct {
	Name         string   `json:"name"`
	Required     bool     `json:"required"`
	Description  string   `json:"description"`
	Type         string   `json:"type"`
	DefaultValue string   `json:"default_value"`
	Values       []string `json:"values"`
}

type AzurermApiManagementApiOperationSpecRequestRepresentationFormParameter struct {
	Description  string   `json:"description"`
	Type         string   `json:"type"`
	DefaultValue string   `json:"default_value"`
	Values       []string `json:"values"`
	Name         string   `json:"name"`
	Required     bool     `json:"required"`
}

type AzurermApiManagementApiOperationSpecRequestRepresentation struct {
	ContentType   string                                                      `json:"content_type"`
	FormParameter []AzurermApiManagementApiOperationSpecRequestRepresentation `json:"form_parameter"`
	Sample        string                                                      `json:"sample"`
	SchemaId      string                                                      `json:"schema_id"`
	TypeName      string                                                      `json:"type_name"`
}

type AzurermApiManagementApiOperationSpecRequest struct {
	Description    string                                        `json:"description"`
	Header         []AzurermApiManagementApiOperationSpecRequest `json:"header"`
	QueryParameter []AzurermApiManagementApiOperationSpecRequest `json:"query_parameter"`
	Representation []AzurermApiManagementApiOperationSpecRequest `json:"representation"`
}

type AzurermApiManagementApiOperationSpecResponseHeader struct {
	Name         string   `json:"name"`
	Required     bool     `json:"required"`
	Description  string   `json:"description"`
	Type         string   `json:"type"`
	DefaultValue string   `json:"default_value"`
	Values       []string `json:"values"`
}

type AzurermApiManagementApiOperationSpecResponseRepresentationFormParameter struct {
	Name         string   `json:"name"`
	Required     bool     `json:"required"`
	Description  string   `json:"description"`
	Type         string   `json:"type"`
	DefaultValue string   `json:"default_value"`
	Values       []string `json:"values"`
}

type AzurermApiManagementApiOperationSpecResponseRepresentation struct {
	FormParameter []AzurermApiManagementApiOperationSpecResponseRepresentation `json:"form_parameter"`
	Sample        string                                                       `json:"sample"`
	SchemaId      string                                                       `json:"schema_id"`
	TypeName      string                                                       `json:"type_name"`
	ContentType   string                                                       `json:"content_type"`
}

type AzurermApiManagementApiOperationSpecResponse struct {
	StatusCode     int                                            `json:"status_code"`
	Description    string                                         `json:"description"`
	Header         []AzurermApiManagementApiOperationSpecResponse `json:"header"`
	Representation []AzurermApiManagementApiOperationSpecResponse `json:"representation"`
}

type AzurermApiManagementApiOperationSpecTemplateParameter struct {
	Required     bool     `json:"required"`
	Description  string   `json:"description"`
	Type         string   `json:"type"`
	DefaultValue string   `json:"default_value"`
	Values       []string `json:"values"`
	Name         string   `json:"name"`
}

type AzurermApiManagementApiOperationSpec struct {
	ResourceGroupName string                                 `json:"resource_group_name"`
	UrlTemplate       string                                 `json:"url_template"`
	OperationId       string                                 `json:"operation_id"`
	ApiName           string                                 `json:"api_name"`
	ApiManagementName string                                 `json:"api_management_name"`
	Request           []AzurermApiManagementApiOperationSpec `json:"request"`
	Response          []AzurermApiManagementApiOperationSpec `json:"response"`
	TemplateParameter []AzurermApiManagementApiOperationSpec `json:"template_parameter"`
	DisplayName       string                                 `json:"display_name"`
	Method            string                                 `json:"method"`
	Description       string                                 `json:"description"`
}



type AzurermApiManagementApiOperationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApiManagementApiOperationList is a list of AzurermApiManagementApiOperations
type AzurermApiManagementApiOperationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementApiOperation CRD objects
	Items []AzurermApiManagementApiOperation `json:"items,omitempty"`
}