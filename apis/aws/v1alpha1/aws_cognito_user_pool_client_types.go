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

type AwsCognitoUserPoolClient struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCognitoUserPoolClientSpec   `json:"spec,omitempty"`
	Status            AwsCognitoUserPoolClientStatus `json:"status,omitempty"`
}

type AwsCognitoUserPoolClientSpec struct {
	ExplicitAuthFlows               []string `json:"explicit_auth_flows"`
	ReadAttributes                  []string `json:"read_attributes"`
	WriteAttributes                 []string `json:"write_attributes"`
	RefreshTokenValidity            int      `json:"refresh_token_validity"`
	AllowedOauthScopes              []string `json:"allowed_oauth_scopes"`
	CallbackUrls                    []string `json:"callback_urls"`
	ClientSecret                    string   `json:"client_secret"`
	GenerateSecret                  bool     `json:"generate_secret"`
	DefaultRedirectUri              string   `json:"default_redirect_uri"`
	AllowedOauthFlows               []string `json:"allowed_oauth_flows"`
	AllowedOauthFlowsUserPoolClient bool     `json:"allowed_oauth_flows_user_pool_client"`
	UserPoolId                      string   `json:"user_pool_id"`
	SupportedIdentityProviders      []string `json:"supported_identity_providers"`
	Name                            string   `json:"name"`
	LogoutUrls                      []string `json:"logout_urls"`
}

type AwsCognitoUserPoolClientStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCognitoUserPoolClientList is a list of AwsCognitoUserPoolClients
type AwsCognitoUserPoolClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCognitoUserPoolClient CRD objects
	Items []AwsCognitoUserPoolClient `json:"items,omitempty"`
}
