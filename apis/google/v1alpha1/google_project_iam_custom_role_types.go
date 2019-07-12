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

type GoogleProjectIamCustomRole struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleProjectIamCustomRoleSpec   `json:"spec,omitempty"`
	Status            GoogleProjectIamCustomRoleStatus `json:"status,omitempty"`
}

type GoogleProjectIamCustomRoleSpec struct {
	Permissions []string `json:"permissions"`
	Project     string   `json:"project"`
	Stage       string   `json:"stage"`
	Description string   `json:"description"`
	Deleted     bool     `json:"deleted"`
	RoleId      string   `json:"role_id"`
	Title       string   `json:"title"`
}

type GoogleProjectIamCustomRoleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleProjectIamCustomRoleList is a list of GoogleProjectIamCustomRoles
type GoogleProjectIamCustomRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleProjectIamCustomRole CRD objects
	Items []GoogleProjectIamCustomRole `json:"items,omitempty"`
}
