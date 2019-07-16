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

type IamRole struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamRoleSpec   `json:"spec,omitempty"`
	Status            IamRoleStatus `json:"status,omitempty"`
}

type IamRoleSpec struct {
	AssumeRolePolicy string `json:"assume_role_policy"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	ForceDetachPolicies bool `json:"force_detach_policies,omitempty"`
	// +optional
	MaxSessionDuration int `json:"max_session_duration,omitempty"`
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
	// +optional
	Path string `json:"path,omitempty"`
	// +optional
	PermissionsBoundary string `json:"permissions_boundary,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type IamRoleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamRoleList is a list of IamRoles
type IamRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamRole CRD objects
	Items []IamRole `json:"items,omitempty"`
}
