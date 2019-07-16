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

type CognitoIdentityPoolRolesAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoIdentityPoolRolesAttachmentSpec   `json:"spec,omitempty"`
	Status            CognitoIdentityPoolRolesAttachmentStatus `json:"status,omitempty"`
}

type CognitoIdentityPoolRolesAttachmentSpecRoleMappingMappingRule struct {
	Claim     string `json:"claim"`
	MatchType string `json:"match_type"`
	RoleArn   string `json:"role_arn"`
	Value     string `json:"value"`
}

type CognitoIdentityPoolRolesAttachmentSpecRoleMapping struct {
	// +optional
	AmbiguousRoleResolution string `json:"ambiguous_role_resolution,omitempty"`
	IdentityProvider        string `json:"identity_provider"`
	// +optional
	// +kubebuilder:validation:MaxItems=25
	MappingRule *[]CognitoIdentityPoolRolesAttachmentSpecRoleMapping `json:"mapping_rule,omitempty"`
	Type        string                                               `json:"type"`
}

type CognitoIdentityPoolRolesAttachmentSpecRoles struct {
	// +optional
	Authenticated string `json:"authenticated,omitempty"`
	// +optional
	Unauthenticated string `json:"unauthenticated,omitempty"`
}

type CognitoIdentityPoolRolesAttachmentSpec struct {
	IdentityPoolId string `json:"identity_pool_id"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	RoleMapping *[]CognitoIdentityPoolRolesAttachmentSpec              `json:"role_mapping,omitempty"`
	Roles       map[string]CognitoIdentityPoolRolesAttachmentSpecRoles `json:"roles"`
}

type CognitoIdentityPoolRolesAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
