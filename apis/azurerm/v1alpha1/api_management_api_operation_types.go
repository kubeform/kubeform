package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ApiManagementAPIOperation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementAPIOperationSpec   `json:"spec,omitempty"`
	Status            ApiManagementAPIOperationStatus `json:"status,omitempty"`
}

type ApiManagementAPIOperationSpecRequestHeader struct {
	// +optional
	DefaultValue string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	Required    bool   `json:"required" tf:"required"`
	Type        string `json:"type" tf:"type"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty" tf:"values,omitempty"`
}

type ApiManagementAPIOperationSpecRequestQueryParameter struct {
	// +optional
	DefaultValue string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	Required    bool   `json:"required" tf:"required"`
	Type        string `json:"type" tf:"type"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty" tf:"values,omitempty"`
}

type ApiManagementAPIOperationSpecRequestRepresentationFormParameter struct {
	// +optional
	DefaultValue string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	Required    bool   `json:"required" tf:"required"`
	Type        string `json:"type" tf:"type"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty" tf:"values,omitempty"`
}

type ApiManagementAPIOperationSpecRequestRepresentation struct {
	ContentType string `json:"contentType" tf:"content_type"`
	// +optional
	FormParameter []ApiManagementAPIOperationSpecRequestRepresentationFormParameter `json:"formParameter,omitempty" tf:"form_parameter,omitempty"`
	// +optional
	Sample string `json:"sample,omitempty" tf:"sample,omitempty"`
	// +optional
	SchemaID string `json:"schemaID,omitempty" tf:"schema_id,omitempty"`
	// +optional
	TypeName string `json:"typeName,omitempty" tf:"type_name,omitempty"`
}

type ApiManagementAPIOperationSpecRequest struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Header []ApiManagementAPIOperationSpecRequestHeader `json:"header,omitempty" tf:"header,omitempty"`
	// +optional
	QueryParameter []ApiManagementAPIOperationSpecRequestQueryParameter `json:"queryParameter,omitempty" tf:"query_parameter,omitempty"`
	// +optional
	Representation []ApiManagementAPIOperationSpecRequestRepresentation `json:"representation,omitempty" tf:"representation,omitempty"`
}

type ApiManagementAPIOperationSpecResponseHeader struct {
	// +optional
	DefaultValue string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	Required    bool   `json:"required" tf:"required"`
	Type        string `json:"type" tf:"type"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty" tf:"values,omitempty"`
}

type ApiManagementAPIOperationSpecResponseRepresentationFormParameter struct {
	// +optional
	DefaultValue string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	Required    bool   `json:"required" tf:"required"`
	Type        string `json:"type" tf:"type"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty" tf:"values,omitempty"`
}

type ApiManagementAPIOperationSpecResponseRepresentation struct {
	ContentType string `json:"contentType" tf:"content_type"`
	// +optional
	FormParameter []ApiManagementAPIOperationSpecResponseRepresentationFormParameter `json:"formParameter,omitempty" tf:"form_parameter,omitempty"`
	// +optional
	Sample string `json:"sample,omitempty" tf:"sample,omitempty"`
	// +optional
	SchemaID string `json:"schemaID,omitempty" tf:"schema_id,omitempty"`
	// +optional
	TypeName string `json:"typeName,omitempty" tf:"type_name,omitempty"`
}

type ApiManagementAPIOperationSpecResponse struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Header []ApiManagementAPIOperationSpecResponseHeader `json:"header,omitempty" tf:"header,omitempty"`
	// +optional
	Representation []ApiManagementAPIOperationSpecResponseRepresentation `json:"representation,omitempty" tf:"representation,omitempty"`
	StatusCode     int                                                   `json:"statusCode" tf:"status_code"`
}

type ApiManagementAPIOperationSpecTemplateParameter struct {
	// +optional
	DefaultValue string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	Required    bool   `json:"required" tf:"required"`
	Type        string `json:"type" tf:"type"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Values []string `json:"values,omitempty" tf:"values,omitempty"`
}

type ApiManagementAPIOperationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`
	ApiName           string `json:"apiName" tf:"api_name"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	DisplayName string `json:"displayName" tf:"display_name"`
	Method      string `json:"method" tf:"method"`
	OperationID string `json:"operationID" tf:"operation_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Request           []ApiManagementAPIOperationSpecRequest `json:"request,omitempty" tf:"request,omitempty"`
	ResourceGroupName string                                 `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Response []ApiManagementAPIOperationSpecResponse `json:"response,omitempty" tf:"response,omitempty"`
	// +optional
	TemplateParameter []ApiManagementAPIOperationSpecTemplateParameter `json:"templateParameter,omitempty" tf:"template_parameter,omitempty"`
	UrlTemplate       string                                           `json:"urlTemplate" tf:"url_template"`
}

type ApiManagementAPIOperationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementAPIOperationList is a list of ApiManagementAPIOperations
type ApiManagementAPIOperationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementAPIOperation CRD objects
	Items []ApiManagementAPIOperation `json:"items,omitempty"`
}
