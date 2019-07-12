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

type AzurermApiManagementProductPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementProductPolicySpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementProductPolicyStatus `json:"status,omitempty"`
}

type AzurermApiManagementProductPolicySpec struct {
	ResourceGroupName string `json:"resource_group_name"`
	ApiManagementName string `json:"api_management_name"`
	ProductId         string `json:"product_id"`
	XmlContent        string `json:"xml_content"`
	XmlLink           string `json:"xml_link"`
}

type AzurermApiManagementProductPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApiManagementProductPolicyList is a list of AzurermApiManagementProductPolicys
type AzurermApiManagementProductPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementProductPolicy CRD objects
	Items []AzurermApiManagementProductPolicy `json:"items,omitempty"`
}
