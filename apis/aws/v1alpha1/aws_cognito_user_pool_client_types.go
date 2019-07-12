package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserPoolClient struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCognitoUserPoolClientSpec   `json:"spec,omitempty"`
	Status            AwsCognitoUserPoolClientStatus `json:"status,omitempty"`
}

type AwsCognitoUserPoolClientSpec struct {
	GenerateSecret                  bool     `json:"generate_secret"`
	WriteAttributes                 []string `json:"write_attributes"`
	SupportedIdentityProviders      []string `json:"supported_identity_providers"`
	Name                            string   `json:"name"`
	UserPoolId                      string   `json:"user_pool_id"`
	RefreshTokenValidity            int      `json:"refresh_token_validity"`
	AllowedOauthScopes              []string `json:"allowed_oauth_scopes"`
	DefaultRedirectUri              string   `json:"default_redirect_uri"`
	ClientSecret                    string   `json:"client_secret"`
	ReadAttributes                  []string `json:"read_attributes"`
	AllowedOauthFlows               []string `json:"allowed_oauth_flows"`
	ExplicitAuthFlows               []string `json:"explicit_auth_flows"`
	AllowedOauthFlowsUserPoolClient bool     `json:"allowed_oauth_flows_user_pool_client"`
	CallbackUrls                    []string `json:"callback_urls"`
	LogoutUrls                      []string `json:"logout_urls"`
}

type AwsCognitoUserPoolClientStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoUserPoolClientList is a list of AwsCognitoUserPoolClients
type AwsCognitoUserPoolClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCognitoUserPoolClient CRD objects
	Items []AwsCognitoUserPoolClient `json:"items,omitempty"`
}
