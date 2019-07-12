package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermApiManagementProductGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementProductGroupSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementProductGroupStatus `json:"status,omitempty"`
}

type AzurermApiManagementProductGroupSpec struct {
	GroupName         string `json:"group_name"`
	ResourceGroupName string `json:"resource_group_name"`
	ApiManagementName string `json:"api_management_name"`
	ProductId         string `json:"product_id"`
}



type AzurermApiManagementProductGroupStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermApiManagementProductGroupList is a list of AzurermApiManagementProductGroups
type AzurermApiManagementProductGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementProductGroup CRD objects
	Items []AzurermApiManagementProductGroup `json:"items,omitempty"`
}