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
	UniqueId            string            `json:"unique_id"`
	Name                string            `json:"name"`
	AssumeRolePolicy    string            `json:"assume_role_policy"`
	MaxSessionDuration  int               `json:"max_session_duration"`
	Arn                 string            `json:"arn"`
	NamePrefix          string            `json:"name_prefix"`
	Path                string            `json:"path"`
	PermissionsBoundary string            `json:"permissions_boundary"`
	Description         string            `json:"description"`
	ForceDetachPolicies bool              `json:"force_detach_policies"`
	CreateDate          string            `json:"create_date"`
	Tags                map[string]string `json:"tags"`
}



type AwsIamRoleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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