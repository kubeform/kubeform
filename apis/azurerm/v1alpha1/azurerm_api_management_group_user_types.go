package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermApiManagementGroupUser struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementGroupUserSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementGroupUserStatus `json:"status,omitempty"`
}

type AzurermApiManagementGroupUserSpec struct {
	UserId            string `json:"user_id"`
	GroupName         string `json:"group_name"`
	ResourceGroupName string `json:"resource_group_name"`
	ApiManagementName string `json:"api_management_name"`
}

type AzurermApiManagementGroupUserStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermApiManagementGroupUserList is a list of AzurermApiManagementGroupUsers
type AzurermApiManagementGroupUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementGroupUser CRD objects
	Items []AzurermApiManagementGroupUser `json:"items,omitempty"`
}
