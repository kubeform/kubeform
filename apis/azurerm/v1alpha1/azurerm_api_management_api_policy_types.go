package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermApiManagementApiPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementApiPolicySpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementApiPolicyStatus `json:"status,omitempty"`
}

type AzurermApiManagementApiPolicySpec struct {
	XmlLink           string `json:"xml_link"`
	ResourceGroupName string `json:"resource_group_name"`
	ApiManagementName string `json:"api_management_name"`
	ApiName           string `json:"api_name"`
	XmlContent        string `json:"xml_content"`
}



type AzurermApiManagementApiPolicyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermApiManagementApiPolicyList is a list of AzurermApiManagementApiPolicys
type AzurermApiManagementApiPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementApiPolicy CRD objects
	Items []AzurermApiManagementApiPolicy `json:"items,omitempty"`
}