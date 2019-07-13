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
	RefreshTokenValidity            int      `json:"refresh_token_validity"`
	CallbackUrls                    []string `json:"callback_urls"`
	SupportedIdentityProviders      []string `json:"supported_identity_providers"`
	GenerateSecret                  bool     `json:"generate_secret"`
	ReadAttributes                  []string `json:"read_attributes"`
	AllowedOauthFlowsUserPoolClient bool     `json:"allowed_oauth_flows_user_pool_client"`
	AllowedOauthScopes              []string `json:"allowed_oauth_scopes"`
	UserPoolId                      string   `json:"user_pool_id"`
	WriteAttributes                 []string `json:"write_attributes"`
	AllowedOauthFlows               []string `json:"allowed_oauth_flows"`
	DefaultRedirectUri              string   `json:"default_redirect_uri"`
	LogoutUrls                      []string `json:"logout_urls"`
	Name                            string   `json:"name"`
	ClientSecret                    string   `json:"client_secret"`
	ExplicitAuthFlows               []string `json:"explicit_auth_flows"`
}



type AwsCognitoUserPoolClientStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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