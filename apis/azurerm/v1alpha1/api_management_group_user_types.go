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

type ApiManagementGroupUser struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementGroupUserSpec   `json:"spec,omitempty"`
	Status            ApiManagementGroupUserStatus `json:"status,omitempty"`
}

type ApiManagementGroupUserSpec struct {
	ApiManagementName string `json:"api_management_name"`
	GroupName         string `json:"group_name"`
	ResourceGroupName string `json:"resource_group_name"`
	UserId            string `json:"user_id"`
}

type ApiManagementGroupUserStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementGroupUserList is a list of ApiManagementGroupUsers
type ApiManagementGroupUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementGroupUser CRD objects
	Items []ApiManagementGroupUser `json:"items,omitempty"`
}
