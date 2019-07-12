package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermApiManagementApi struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementApiSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementApiStatus `json:"status,omitempty"`
}

type AzurermApiManagementApiSpecImportWsdlSelector struct {
	ServiceName  string `json:"service_name"`
	EndpointName string `json:"endpoint_name"`
}

type AzurermApiManagementApiSpecImport struct {
	WsdlSelector  []AzurermApiManagementApiSpecImport `json:"wsdl_selector"`
	ContentValue  string                              `json:"content_value"`
	ContentFormat string                              `json:"content_format"`
}

type AzurermApiManagementApiSpecSubscriptionKeyParameterNames struct {
	Query  string `json:"query"`
	Header string `json:"header"`
}

type AzurermApiManagementApiSpec struct {
	IsCurrent                     bool                          `json:"is_current"`
	VersionSetId                  string                        `json:"version_set_id"`
	ServiceUrl                    string                        `json:"service_url"`
	IsOnline                      bool                          `json:"is_online"`
	Name                          string                        `json:"name"`
	ApiManagementName             string                        `json:"api_management_name"`
	ResourceGroupName             string                        `json:"resource_group_name"`
	Path                          string                        `json:"path"`
	Protocols                     []string                      `json:"protocols"`
	Description                   string                        `json:"description"`
	Import                        []AzurermApiManagementApiSpec `json:"import"`
	SubscriptionKeyParameterNames []AzurermApiManagementApiSpec `json:"subscription_key_parameter_names"`
	SoapPassThrough               bool                          `json:"soap_pass_through"`
	Version                       string                        `json:"version"`
	DisplayName                   string                        `json:"display_name"`
	Revision                      string                        `json:"revision"`
}

type AzurermApiManagementApiStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermApiManagementApiList is a list of AzurermApiManagementApis
type AzurermApiManagementApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementApi CRD objects
	Items []AzurermApiManagementApi `json:"items,omitempty"`
}
