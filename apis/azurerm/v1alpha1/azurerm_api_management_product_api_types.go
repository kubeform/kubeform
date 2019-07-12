package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermApiManagementProductApi struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementProductApiSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementProductApiStatus `json:"status,omitempty"`
}

type AzurermApiManagementProductApiSpec struct {
	ResourceGroupName string `json:"resource_group_name"`
	ApiManagementName string `json:"api_management_name"`
	ApiName           string `json:"api_name"`
	ProductId         string `json:"product_id"`
}

type AzurermApiManagementProductApiStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermApiManagementProductApiList is a list of AzurermApiManagementProductApis
type AzurermApiManagementProductApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementProductApi CRD objects
	Items []AzurermApiManagementProductApi `json:"items,omitempty"`
}
