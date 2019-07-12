package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermApiManagementApiOperation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementApiOperationSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementApiOperationStatus `json:"status,omitempty"`
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
	DefaultValue string   `json:"default_value"`
	Values       []string `json:"values"`
	Name         string   `json:"name"`
	Required     bool     `json:"required"`
	Description  string   `json:"description"`
	Type         string   `json:"type"`
}

type AzurermApiManagementApiOperationSpecResponseRepresentation struct {
	Sample        string                                                       `json:"sample"`
	SchemaId      string                                                       `json:"schema_id"`
	TypeName      string                                                       `json:"type_name"`
	ContentType   string                                                       `json:"content_type"`
	FormParameter []AzurermApiManagementApiOperationSpecResponseRepresentation `json:"form_parameter"`
}

type AzurermApiManagementApiOperationSpecResponse struct {
	StatusCode     int                                            `json:"status_code"`
	Description    string                                         `json:"description"`
	Header         []AzurermApiManagementApiOperationSpecResponse `json:"header"`
	Representation []AzurermApiManagementApiOperationSpecResponse `json:"representation"`
}

type AzurermApiManagementApiOperationSpecRequestRepresentationFormParameter struct {
	Values       []string `json:"values"`
	Name         string   `json:"name"`
	Required     bool     `json:"required"`
	Description  string   `json:"description"`
	Type         string   `json:"type"`
	DefaultValue string   `json:"default_value"`
}

type AzurermApiManagementApiOperationSpecRequestRepresentation struct {
	TypeName      string                                                      `json:"type_name"`
	ContentType   string                                                      `json:"content_type"`
	FormParameter []AzurermApiManagementApiOperationSpecRequestRepresentation `json:"form_parameter"`
	Sample        string                                                      `json:"sample"`
	SchemaId      string                                                      `json:"schema_id"`
}

type AzurermApiManagementApiOperationSpecRequestHeader struct {
	Name         string   `json:"name"`
	Required     bool     `json:"required"`
	Description  string   `json:"description"`
	Type         string   `json:"type"`
	DefaultValue string   `json:"default_value"`
	Values       []string `json:"values"`
}

type AzurermApiManagementApiOperationSpecRequestQueryParameter struct {
	Type         string   `json:"type"`
	DefaultValue string   `json:"default_value"`
	Values       []string `json:"values"`
	Name         string   `json:"name"`
	Required     bool     `json:"required"`
	Description  string   `json:"description"`
}

type AzurermApiManagementApiOperationSpecRequest struct {
	Representation []AzurermApiManagementApiOperationSpecRequest `json:"representation"`
	Description    string                                        `json:"description"`
	Header         []AzurermApiManagementApiOperationSpecRequest `json:"header"`
	QueryParameter []AzurermApiManagementApiOperationSpecRequest `json:"query_parameter"`
}

type AzurermApiManagementApiOperationSpecTemplateParameter struct {
	Name         string   `json:"name"`
	Required     bool     `json:"required"`
	Description  string   `json:"description"`
	Type         string   `json:"type"`
	DefaultValue string   `json:"default_value"`
	Values       []string `json:"values"`
}

type AzurermApiManagementApiOperationSpec struct {
	ApiManagementName string                                 `json:"api_management_name"`
	Description       string                                 `json:"description"`
	Response          []AzurermApiManagementApiOperationSpec `json:"response"`
	DisplayName       string                                 `json:"display_name"`
	Method            string                                 `json:"method"`
	UrlTemplate       string                                 `json:"url_template"`
	Request           []AzurermApiManagementApiOperationSpec `json:"request"`
	TemplateParameter []AzurermApiManagementApiOperationSpec `json:"template_parameter"`
	OperationId       string                                 `json:"operation_id"`
	ApiName           string                                 `json:"api_name"`
	ResourceGroupName string                                 `json:"resource_group_name"`
}

type AzurermApiManagementApiOperationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermApiManagementApiOperationList is a list of AzurermApiManagementApiOperations
type AzurermApiManagementApiOperationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementApiOperation CRD objects
	Items []AzurermApiManagementApiOperation `json:"items,omitempty"`
}
