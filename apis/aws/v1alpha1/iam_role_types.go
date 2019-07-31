package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	AssumeRolePolicy string `json:"assumeRolePolicy" tf:"assume_role_policy"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	ForceDetachPolicies bool `json:"forceDetachPolicies,omitempty" tf:"force_detach_policies,omitempty"`
	// +optional
	MaxSessionDuration int `json:"maxSessionDuration,omitempty" tf:"max_session_duration,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	// +optional
	PermissionsBoundary string `json:"permissionsBoundary,omitempty" tf:"permissions_boundary,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IamRoleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
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
