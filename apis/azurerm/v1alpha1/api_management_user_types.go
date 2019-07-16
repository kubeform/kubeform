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

type ApiManagementUser struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementUserSpec   `json:"spec,omitempty"`
	Status            ApiManagementUserStatus `json:"status,omitempty"`
}

type ApiManagementUserSpec struct {
	ApiManagementName string `json:"api_management_name"`
	// +optional
	Confirmation string `json:"confirmation,omitempty"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	// +optional
	Note string `json:"note,omitempty"`
	// +optional
	Password          string `json:"password,omitempty"`
	ResourceGroupName string `json:"resource_group_name"`
	UserId            string `json:"user_id"`
}

type ApiManagementUserStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementUserList is a list of ApiManagementUsers
type ApiManagementUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementUser CRD objects
	Items []ApiManagementUser `json:"items,omitempty"`
}
