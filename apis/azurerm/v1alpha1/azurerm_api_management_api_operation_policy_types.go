package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermApiManagementApiOperationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementApiOperationPolicySpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementApiOperationPolicyStatus `json:"status,omitempty"`
}

type AzurermApiManagementApiOperationPolicySpec struct {
	XmlLink           string `json:"xml_link"`
	ResourceGroupName string `json:"resource_group_name"`
	ApiManagementName string `json:"api_management_name"`
	ApiName           string `json:"api_name"`
	OperationId       string `json:"operation_id"`
	XmlContent        string `json:"xml_content"`
}



type AzurermApiManagementApiOperationPolicyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermApiManagementApiOperationPolicyList is a list of AzurermApiManagementApiOperationPolicys
type AzurermApiManagementApiOperationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementApiOperationPolicy CRD objects
	Items []AzurermApiManagementApiOperationPolicy `json:"items,omitempty"`
}