package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type CognitoIdentityPoolRolesAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoIdentityPoolRolesAttachmentSpec   `json:"spec,omitempty"`
	Status            CognitoIdentityPoolRolesAttachmentStatus `json:"status,omitempty"`
}

type CognitoIdentityPoolRolesAttachmentSpecRoleMappingMappingRule struct {
	Claim     string `json:"claim" tf:"claim"`
	MatchType string `json:"matchType" tf:"match_type"`
	RoleArn   string `json:"roleArn" tf:"role_arn"`
	Value     string `json:"value" tf:"value"`
}

type CognitoIdentityPoolRolesAttachmentSpecRoleMapping struct {
	// +optional
	AmbiguousRoleResolution string `json:"ambiguousRoleResolution,omitempty" tf:"ambiguous_role_resolution,omitempty"`
	IdentityProvider        string `json:"identityProvider" tf:"identity_provider"`
	// +optional
	// +kubebuilder:validation:MaxItems=25
	MappingRule []CognitoIdentityPoolRolesAttachmentSpecRoleMappingMappingRule `json:"mappingRule,omitempty" tf:"mapping_rule,omitempty"`
	Type        string                                                         `json:"type" tf:"type"`
}

type CognitoIdentityPoolRolesAttachmentSpecRoles struct {
	// +optional
	Authenticated string `json:"authenticated,omitempty" tf:"authenticated,omitempty"`
	// +optional
	Unauthenticated string `json:"unauthenticated,omitempty" tf:"unauthenticated,omitempty"`
}

type CognitoIdentityPoolRolesAttachmentSpec struct {
	ProviderRef    core.LocalObjectReference `json:"providerRef" tf:"-"`
	IdentityPoolID string                    `json:"identityPoolID" tf:"identity_pool_id"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	RoleMapping []CognitoIdentityPoolRolesAttachmentSpecRoleMapping    `json:"roleMapping,omitempty" tf:"role_mapping,omitempty"`
	Roles       map[string]CognitoIdentityPoolRolesAttachmentSpecRoles `json:"roles" tf:"roles"`
}

type CognitoIdentityPoolRolesAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CognitoIdentityPoolRolesAttachmentList is a list of CognitoIdentityPoolRolesAttachments
type CognitoIdentityPoolRolesAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CognitoIdentityPoolRolesAttachment CRD objects
	Items []CognitoIdentityPoolRolesAttachment `json:"items,omitempty"`
}
