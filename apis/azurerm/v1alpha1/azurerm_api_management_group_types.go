package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermApiManagementGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementGroupSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementGroupStatus `json:"status,omitempty"`
}

type AzurermApiManagementGroupSpec struct {
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	ApiManagementName string `json:"api_management_name"`
	DisplayName       string `json:"display_name"`
	Description       string `json:"description"`
	ExternalId        string `json:"external_id"`
	Type              string `json:"type"`
}



type AzurermApiManagementGroupStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermApiManagementGroupList is a list of AzurermApiManagementGroups
type AzurermApiManagementGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementGroup CRD objects
	Items []AzurermApiManagementGroup `json:"items,omitempty"`
}