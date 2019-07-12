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

type AzurermApiManagementApi struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementApiSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementApiStatus `json:"status,omitempty"`
}

type AzurermApiManagementApiSpecSubscriptionKeyParameterNames struct {
	Header string `json:"header"`
	Query  string `json:"query"`
}

type AzurermApiManagementApiSpecImportWsdlSelector struct {
	ServiceName  string `json:"service_name"`
	EndpointName string `json:"endpoint_name"`
}

type AzurermApiManagementApiSpecImport struct {
	ContentValue  string                              `json:"content_value"`
	ContentFormat string                              `json:"content_format"`
	WsdlSelector  []AzurermApiManagementApiSpecImport `json:"wsdl_selector"`
}

type AzurermApiManagementApiSpec struct {
	ApiManagementName             string                        `json:"api_management_name"`
	Description                   string                        `json:"description"`
	SubscriptionKeyParameterNames []AzurermApiManagementApiSpec `json:"subscription_key_parameter_names"`
	IsCurrent                     bool                          `json:"is_current"`
	Version                       string                        `json:"version"`
	VersionSetId                  string                        `json:"version_set_id"`
	Name                          string                        `json:"name"`
	DisplayName                   string                        `json:"display_name"`
	Revision                      string                        `json:"revision"`
	SoapPassThrough               bool                          `json:"soap_pass_through"`
	ResourceGroupName             string                        `json:"resource_group_name"`
	Protocols                     []string                      `json:"protocols"`
	Import                        []AzurermApiManagementApiSpec `json:"import"`
	ServiceUrl                    string                        `json:"service_url"`
	IsOnline                      bool                          `json:"is_online"`
	Path                          string                        `json:"path"`
}

type AzurermApiManagementApiStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApiManagementApiList is a list of AzurermApiManagementApis
type AzurermApiManagementApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementApi CRD objects
	Items []AzurermApiManagementApi `json:"items,omitempty"`
}
