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

type AwsIamRole struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamRoleSpec   `json:"spec,omitempty"`
	Status            AwsIamRoleStatus `json:"status,omitempty"`
}

type AwsIamRoleSpec struct {
	Arn                 string            `json:"arn"`
	Description         string            `json:"description"`
	AssumeRolePolicy    string            `json:"assume_role_policy"`
	CreateDate          string            `json:"create_date"`
	MaxSessionDuration  int               `json:"max_session_duration"`
	UniqueId            string            `json:"unique_id"`
	Name                string            `json:"name"`
	NamePrefix          string            `json:"name_prefix"`
	Path                string            `json:"path"`
	PermissionsBoundary string            `json:"permissions_boundary"`
	ForceDetachPolicies bool              `json:"force_detach_policies"`
	Tags                map[string]string `json:"tags"`
}

type AwsIamRoleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIamRoleList is a list of AwsIamRoles
type AwsIamRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamRole CRD objects
	Items []AwsIamRole `json:"items,omitempty"`
}
