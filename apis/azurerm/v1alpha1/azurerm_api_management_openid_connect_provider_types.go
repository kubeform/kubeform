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

type AzurermApiManagementOpenidConnectProvider struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementOpenidConnectProviderSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementOpenidConnectProviderStatus `json:"status,omitempty"`
}

type AzurermApiManagementOpenidConnectProviderSpec struct {
	MetadataEndpoint  string `json:"metadata_endpoint"`
	Description       string `json:"description"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	ApiManagementName string `json:"api_management_name"`
	ClientId          string `json:"client_id"`
	ClientSecret      string `json:"client_secret"`
	DisplayName       string `json:"display_name"`
}



type AzurermApiManagementOpenidConnectProviderStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApiManagementOpenidConnectProviderList is a list of AzurermApiManagementOpenidConnectProviders
type AzurermApiManagementOpenidConnectProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementOpenidConnectProvider CRD objects
	Items []AzurermApiManagementOpenidConnectProvider `json:"items,omitempty"`
}