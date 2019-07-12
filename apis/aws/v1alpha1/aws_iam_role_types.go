package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamRole struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamRoleSpec   `json:"spec,omitempty"`
	Status            AwsIamRoleStatus `json:"status,omitempty"`
}

type AwsIamRoleSpec struct {
	UniqueId            string            `json:"unique_id"`
	Path                string            `json:"path"`
	PermissionsBoundary string            `json:"permissions_boundary"`
	Tags                map[string]string `json:"tags"`
	ForceDetachPolicies bool              `json:"force_detach_policies"`
	CreateDate          string            `json:"create_date"`
	MaxSessionDuration  int               `json:"max_session_duration"`
	Arn                 string            `json:"arn"`
	Name                string            `json:"name"`
	NamePrefix          string            `json:"name_prefix"`
	Description         string            `json:"description"`
	AssumeRolePolicy    string            `json:"assume_role_policy"`
}

type AwsIamRoleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamRoleList is a list of AwsIamRoles
type AwsIamRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamRole CRD objects
	Items []AwsIamRole `json:"items,omitempty"`
}
