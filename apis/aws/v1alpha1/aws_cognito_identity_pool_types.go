package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoIdentityPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCognitoIdentityPoolSpec   `json:"spec,omitempty"`
	Status            AwsCognitoIdentityPoolStatus `json:"status,omitempty"`
}

type AwsCognitoIdentityPoolSpecCognitoIdentityProviders struct {
	ClientId             string `json:"client_id"`
	ProviderName         string `json:"provider_name"`
	ServerSideTokenCheck bool   `json:"server_side_token_check"`
}

type AwsCognitoIdentityPoolSpec struct {
	IdentityPoolName               string                       `json:"identity_pool_name"`
	Arn                            string                       `json:"arn"`
	CognitoIdentityProviders       []AwsCognitoIdentityPoolSpec `json:"cognito_identity_providers"`
	DeveloperProviderName          string                       `json:"developer_provider_name"`
	AllowUnauthenticatedIdentities bool                         `json:"allow_unauthenticated_identities"`
	OpenidConnectProviderArns      []string                     `json:"openid_connect_provider_arns"`
	SamlProviderArns               []string                     `json:"saml_provider_arns"`
	SupportedLoginProviders        map[string]string            `json:"supported_login_providers"`
}

type AwsCognitoIdentityPoolStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoIdentityPoolList is a list of AwsCognitoIdentityPools
type AwsCognitoIdentityPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCognitoIdentityPool CRD objects
	Items []AwsCognitoIdentityPool `json:"items,omitempty"`
}
