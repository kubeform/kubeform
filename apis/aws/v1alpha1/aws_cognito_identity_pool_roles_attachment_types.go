package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoIdentityPoolRolesAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCognitoIdentityPoolRolesAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsCognitoIdentityPoolRolesAttachmentStatus `json:"status,omitempty"`
}

type AwsCognitoIdentityPoolRolesAttachmentSpecRoleMappingMappingRule struct {
	Claim     string `json:"claim"`
	MatchType string `json:"match_type"`
	RoleArn   string `json:"role_arn"`
	Value     string `json:"value"`
}

type AwsCognitoIdentityPoolRolesAttachmentSpecRoleMapping struct {
	AmbiguousRoleResolution string                                                 `json:"ambiguous_role_resolution"`
	MappingRule             []AwsCognitoIdentityPoolRolesAttachmentSpecRoleMapping `json:"mapping_rule"`
	Type                    string                                                 `json:"type"`
	IdentityProvider        string                                                 `json:"identity_provider"`
}

type AwsCognitoIdentityPoolRolesAttachmentSpecRoles struct {
	Authenticated   string `json:"authenticated"`
	Unauthenticated string `json:"unauthenticated"`
}

type AwsCognitoIdentityPoolRolesAttachmentSpec struct {
	IdentityPoolId string                                      `json:"identity_pool_id"`
	RoleMapping    []AwsCognitoIdentityPoolRolesAttachmentSpec `json:"role_mapping"`
	Roles          map[string]string                           `json:"roles"`
}



type AwsCognitoIdentityPoolRolesAttachmentStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoIdentityPoolRolesAttachmentList is a list of AwsCognitoIdentityPoolRolesAttachments
type AwsCognitoIdentityPoolRolesAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCognitoIdentityPoolRolesAttachment CRD objects
	Items []AwsCognitoIdentityPoolRolesAttachment `json:"items,omitempty"`
}