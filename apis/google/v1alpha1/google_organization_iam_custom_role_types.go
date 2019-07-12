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

type GoogleOrganizationIamCustomRole struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleOrganizationIamCustomRoleSpec   `json:"spec,omitempty"`
	Status            GoogleOrganizationIamCustomRoleStatus `json:"status,omitempty"`
}

type GoogleOrganizationIamCustomRoleSpec struct {
	Permissions []string `json:"permissions"`
	Stage       string   `json:"stage"`
	Description string   `json:"description"`
	Deleted     bool     `json:"deleted"`
	RoleId      string   `json:"role_id"`
	OrgId       string   `json:"org_id"`
	Title       string   `json:"title"`
}

type GoogleOrganizationIamCustomRoleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleOrganizationIamCustomRoleList is a list of GoogleOrganizationIamCustomRoles
type GoogleOrganizationIamCustomRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleOrganizationIamCustomRole CRD objects
	Items []GoogleOrganizationIamCustomRole `json:"items,omitempty"`
}
