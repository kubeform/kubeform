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

type AzurermApiManagementUser struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementUserSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementUserStatus `json:"status,omitempty"`
}

type AzurermApiManagementUserSpec struct {
	UserId            string `json:"user_id"`
	Email             string `json:"email"`
	LastName          string `json:"last_name"`
	Confirmation      string `json:"confirmation"`
	Password          string `json:"password"`
	ApiManagementName string `json:"api_management_name"`
	ResourceGroupName string `json:"resource_group_name"`
	FirstName         string `json:"first_name"`
	Note              string `json:"note"`
	State             string `json:"state"`
}



type AzurermApiManagementUserStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApiManagementUserList is a list of AzurermApiManagementUsers
type AzurermApiManagementUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementUser CRD objects
	Items []AzurermApiManagementUser `json:"items,omitempty"`
}