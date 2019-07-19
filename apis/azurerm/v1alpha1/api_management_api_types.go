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

type ApiManagementAPI struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementAPISpec   `json:"spec,omitempty"`
	Status            ApiManagementAPIStatus `json:"status,omitempty"`
}

type ApiManagementAPISpecImportWsdlSelector struct {
	EndpointName string `json:"endpointName" tf:"endpoint_name"`
	ServiceName  string `json:"serviceName" tf:"service_name"`
}

type ApiManagementAPISpecImport struct {
	ContentFormat string `json:"contentFormat" tf:"content_format"`
	ContentValue  string `json:"contentValue" tf:"content_value"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	WsdlSelector []ApiManagementAPISpecImportWsdlSelector `json:"wsdlSelector,omitempty" tf:"wsdl_selector,omitempty"`
}

type ApiManagementAPISpecSubscriptionKeyParameterNames struct {
	Header string `json:"header" tf:"header"`
	Query  string `json:"query" tf:"query"`
}

type ApiManagementAPISpec struct {
	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	DisplayName string `json:"displayName" tf:"display_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Import []ApiManagementAPISpecImport `json:"import,omitempty" tf:"import,omitempty"`
	Name   string                       `json:"name" tf:"name"`
	Path   string                       `json:"path" tf:"path"`
	// +kubebuilder:validation:UniqueItems=true
	Protocols         []string `json:"protocols" tf:"protocols"`
	ResourceGroupName string   `json:"resourceGroupName" tf:"resource_group_name"`
	Revision          string   `json:"revision" tf:"revision"`
	// +optional
	ServiceURL string `json:"serviceURL,omitempty" tf:"service_url,omitempty"`
	// +optional
	SoapPassThrough bool `json:"soapPassThrough,omitempty" tf:"soap_pass_through,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SubscriptionKeyParameterNames []ApiManagementAPISpecSubscriptionKeyParameterNames `json:"subscriptionKeyParameterNames,omitempty" tf:"subscription_key_parameter_names,omitempty"`
	ProviderRef                   core.LocalObjectReference                           `json:"providerRef" tf:"-"`
}

type ApiManagementAPIStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementAPIList is a list of ApiManagementAPIs
type ApiManagementAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementAPI CRD objects
	Items []ApiManagementAPI `json:"items,omitempty"`
}
