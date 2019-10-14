package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type IotRoleAlias struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IotRoleAliasSpec   `json:"spec,omitempty"`
	Status            IotRoleAliasStatus `json:"status,omitempty"`
}

type IotRoleAliasSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Alias string `json:"alias" tf:"alias"`
	// +optional
	CredentialDuration int    `json:"credentialDuration,omitempty" tf:"credential_duration,omitempty"`
	RoleArn            string `json:"roleArn" tf:"role_arn"`
}

type IotRoleAliasStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *IotRoleAliasSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IotRoleAliasList is a list of IotRoleAliass
type IotRoleAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IotRoleAlias CRD objects
	Items []IotRoleAlias `json:"items,omitempty"`
}
