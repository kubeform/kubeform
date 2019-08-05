package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type CognitoUserGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoUserGroupSpec   `json:"spec,omitempty"`
	Status            CognitoUserGroupStatus `json:"status,omitempty"`
}

type CognitoUserGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Precedence int `json:"precedence,omitempty" tf:"precedence,omitempty"`
	// +optional
	RoleArn    string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
	UserPoolID string `json:"userPoolID" tf:"user_pool_id"`
}

type CognitoUserGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CognitoUserGroupSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CognitoUserGroupList is a list of CognitoUserGroups
type CognitoUserGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CognitoUserGroup CRD objects
	Items []CognitoUserGroup `json:"items,omitempty"`
}
