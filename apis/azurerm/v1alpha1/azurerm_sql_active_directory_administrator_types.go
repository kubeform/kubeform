package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermSqlActiveDirectoryAdministrator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSqlActiveDirectoryAdministratorSpec   `json:"spec,omitempty"`
	Status            AzurermSqlActiveDirectoryAdministratorStatus `json:"status,omitempty"`
}

type AzurermSqlActiveDirectoryAdministratorSpec struct {
	TenantId          string `json:"tenant_id"`
	ServerName        string `json:"server_name"`
	ResourceGroupName string `json:"resource_group_name"`
	Login             string `json:"login"`
	ObjectId          string `json:"object_id"`
}



type AzurermSqlActiveDirectoryAdministratorStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermSqlActiveDirectoryAdministratorList is a list of AzurermSqlActiveDirectoryAdministrators
type AzurermSqlActiveDirectoryAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSqlActiveDirectoryAdministrator CRD objects
	Items []AzurermSqlActiveDirectoryAdministrator `json:"items,omitempty"`
}